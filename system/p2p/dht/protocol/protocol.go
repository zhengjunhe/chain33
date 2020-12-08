// Package protocol p2p protocol
package protocol

import (
	"context"

	"github.com/33cn/dplatform/p2p"
	"github.com/33cn/dplatform/queue"
	types2 "github.com/33cn/dplatform/system/p2p/dht/types"
	"github.com/33cn/dplatform/types"
	ds "github.com/ipfs/go-datastore"
	core "github.com/libp2p/go-libp2p-core"
	discovery "github.com/libp2p/go-libp2p-discovery"
	kbt "github.com/libp2p/go-libp2p-kbucket"
)

// all protocols
const (
	//p2pstore protocols
	FetchChunk        = "/dplatform/fetch-chunk/" + types2.Version
	StoreChunk        = "/dplatform/store-chunk/" + types2.Version
	GetHeader         = "/dplatform/headers/" + types2.Version
	GetChunkRecord    = "/dplatform/chunk-record/" + types2.Version
	BroadcastFullNode = "/dplatform/full-node/" + types2.Version

	//sync protocols
	IsSync        = "/dplatform/is-sync/" + types2.Version
	IsHealthy     = "/dplatform/is-healthy/" + types2.Version
	GetLastHeader = "/dplatform/last-header/" + types2.Version
)

// P2PEnv p2p全局公共变量
type P2PEnv struct {
	Ctx         context.Context
	ChainCfg    *types.DplatformConfig
	QueueClient queue.Client
	Host        core.Host
	P2PManager  *p2p.Manager
	SubConfig   *types2.P2PSubConfig
	DB          ds.Datastore
	*discovery.RoutingDiscovery

	RoutingTable *kbt.RoutingTable
}
