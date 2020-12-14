package mempool

import (
	"github.com/33cn/dplatform/queue"
	"github.com/33cn/dplatform/system/mempool"
	"github.com/33cn/dplatform/types"
)

// New new mempool queue module
func New(cfg *types.DplatformOSConfig) queue.Module {
	mcfg := cfg.GetModuleConfig().Mempool
	sub := cfg.GetSubConfig().Mempool
	con, err := mempool.Load(mcfg.Name)
	if err != nil {
		panic("Unsupported mempool type:" + mcfg.Name + " " + err.Error())
	}
	subcfg, ok := sub[mcfg.Name]
	if !ok {
		subcfg = nil
	}
	obj := con(mcfg, subcfg)
	return obj
}
