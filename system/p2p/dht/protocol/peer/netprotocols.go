package peer

import (
	"encoding/json"

	"github.com/33cn/dplatform/queue"
	"github.com/33cn/dplatform/types"
)

func (p *peerInfoProtol) netprotocolsHandleEvent(msg *queue.Message) {
	//allproto netinfo
	bandprotocols := p.ConnManager.BandTrackerByProtocol()
	allprotonetinfo, _ := json.MarshalIndent(bandprotocols, "", "\t")
	log.Debug("netinfoHandleEvent", string(allprotonetinfo))
	msg.Reply(p.GetQueueClient().NewMessage("rpc", types.EventNetProtocols, bandprotocols))
}
