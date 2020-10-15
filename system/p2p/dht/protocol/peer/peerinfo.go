package peer

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/33cn/chain33/common/version"
	"github.com/33cn/chain33/system/p2p/dht/protocol"
	"github.com/33cn/chain33/types"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/multiformats/go-multiaddr"
)

func (p *Protocol) getLocalPeerInfo() *types.Peer {
	msg := p.QueueClient.NewMessage(mempool, types.EventGetMempoolSize, nil)
	err := p.QueueClient.Send(msg, true)
	if err != nil {
		return nil
	}
	resp, err := p.QueueClient.WaitTimeout(msg, time.Second*10)
	if err != nil {
		log.Error("getLocalPeerInfo", "mempool WaitTimeout", err)
		return nil
	}
	var localPeer types.Peer
	localPeer.MempoolSize = int32(resp.Data.(*types.MempoolSize).GetSize())

	msg = p.QueueClient.NewMessage(blockchain, types.EventGetLastHeader, nil)
	err = p.QueueClient.Send(msg, true)
	if err != nil {
		return nil
	}
	resp, err = p.QueueClient.WaitTimeout(msg, time.Second*10)
	if err != nil {
		log.Error("getLocalPeerInfo", "blockchain WaitTimeout", err)
		return nil
	}
	localPeer.Header = resp.Data.(*types.Header)
	localPeer.Name = p.Host.ID().Pretty()
	ip, port := parseIPAndPort(p.getExternalAddr())
	localPeer.Addr = ip
	localPeer.Port = int32(port)
	localPeer.Version = version.GetVersion() + "@" + version.GetAppVersion()
	localPeer.StoreDBVersion = version.GetStoreDBVersion()
	localPeer.LocalDBVersion = version.GetLocalDBVersion()
	return &localPeer
}

func (p *Protocol) refreshPeerInfo() {
	var wg sync.WaitGroup
	for _, remoteID := range p.RoutingTable.ListPeers() {
		if p.checkDone() {
			log.Warn("getPeerInfo", "process", "done+++++++")
			return
		}
		if remoteID == p.Host.ID() {
			continue
		}
		//修改为并发获取peerinfo信息
		wg.Add(1)
		go func(pid peer.ID) {
			defer wg.Done()
			ctx, cancel := context.WithTimeout(p.Ctx, time.Second*3)
			defer cancel()
			stream, err := p.Host.NewStream(ctx, pid, peerInfo)
			if err != nil {
				log.Error("refreshPeerInfo", "new stream error", err, "peer id", pid)
				return
			}
			defer protocol.CloseStream(stream)
			var resp types.Peer
			err = protocol.ReadStream(&resp, stream)
			if err != nil {
				log.Error("refreshPeerInfo", "read stream error", err, "peer id", pid)
				return
			}
			p.PeerInfoManager.Refresh(&resp)
		}(remoteID)
	}
	selfPeer := p.getLocalPeerInfo()
	selfPeer.Self = true
	p.PeerInfoManager.Refresh(selfPeer)
	wg.Wait()
}

func (p *Protocol) setExternalAddr(addr string) {
	ip, _ := parseIPAndPort(addr)
	if isPublicIP(ip) {
		p.mutex.Lock()
		p.externalAddr = addr
		p.mutex.Unlock()
		ma, _ := multiaddr.NewMultiaddr(addr)
		p.Host.Peerstore().AddAddr(p.Host.ID(), ma, time.Hour*24)
	}
}

func (p *Protocol) getExternalAddr() string {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.externalAddr
}

func (p *Protocol) getPublicIP() string {
	ip, _ := parseIPAndPort(p.getExternalAddr())
	return ip
}

func (p *Protocol) detectNodeAddr() {

	// 通过bootstrap获取本节点公网ip
	for _, bootstrap := range p.SubConfig.BootStraps {
		if p.checkDone() {
			break
		}
		addr, _ := multiaddr.NewMultiaddr(bootstrap)
		peerInfo, err := peer.AddrInfoFromP2pAddr(addr)
		if err != nil {
			continue
		}
		err = p.queryVersion(peerInfo.ID)
		if err != nil {
			continue
		}
	}
	var rangeCount int
	for {
		if p.checkDone() {
			break
		}
		if p.RoutingTable.Size() == 0 {
			time.Sleep(time.Second)
			continue
		}
		//启动后间隔1分钟，以充分获得节点外网地址
		rangeCount++
		if rangeCount > 2 {
			time.Sleep(time.Minute)
		}
		for _, pid := range p.RoutingTable.ListPeers() {
			if p.containsPublicIP(pid) {
				continue
			}
			err := p.queryVersion(pid)
			if err != nil {
				log.Error("detectNodeAddr", "queryVersion error", err, "pid", pid)
				continue
			}
		}
	}
}

func (p *Protocol) queryVersion(pid peer.ID) error {
	stream, err := p.Host.NewStream(p.Ctx, pid, peerVersion)
	if err != nil {
		log.Error("NewStream", "err", err, "remoteID", pid)
		return err
	}
	defer protocol.CloseStream(stream)

	req := &types.P2PVersion{
		Version:  p.SubConfig.Channel,
		AddrFrom: fmt.Sprintf("/ip4/%v/tcp/%d", p.getExternalAddr(), p.SubConfig.Port),
		AddrRecv: stream.Conn().RemoteMultiaddr().String(),
	}
	err = protocol.WriteStream(req, stream)
	if err != nil {
		log.Error("queryVersion", "WriteStream err", err)
		return err
	}
	var resp types.P2PVersion
	err = protocol.ReadStream(&resp, stream)
	if err != nil {
		log.Error("queryVersion", "ReadStream err", err)
		return err
	}
	addr := resp.GetAddrRecv()
	p.setExternalAddr(addr)

	if ip, _ := parseIPAndPort(resp.GetAddrFrom()); isPublicIP(ip) {
		remoteMAddr, err := multiaddr.NewMultiaddr(resp.GetAddrFrom())
		if err != nil {
			return err
		}
		p.Host.Peerstore().AddAddr(pid, remoteMAddr, time.Hour*12)
	}
	return nil
}

func (p *Protocol) containsPublicIP(pid peer.ID) bool {
	for _, maddr := range p.Host.Peerstore().Addrs(pid) {
		if ip, _ := parseIPAndPort(maddr.String()); isPublicIP(ip) {
			return true
		}
	}
	return false
}

func parseIPAndPort(multiAddr string) (ip string, port int) {
	split := strings.Split(multiAddr, "/")
	if len(split) < 5 {
		return
	}
	port, err := strconv.Atoi(split[4])
	if err != nil {
		return
	}
	ip = split[2]
	return
}
