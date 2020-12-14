package protocol

import (
	"github.com/33cn/dplatformos/common/log/log15"

	"github.com/33cn/dplatformos/queue"
	"github.com/33cn/dplatformos/system/p2p/dht/protocol/types"
)

var (
	log = log15.New("module", "p2p.protocol")
)

// HandleEvent handle p2p event
func HandleEvent(msg *queue.Message) {

	if eventHander, ok := types.GetEventHandler(msg.Ty); ok {
		//log.Debug("HandleEvent", "msgTy", msg.Ty)
		eventHander(msg)
	} else if eventHandler := GetEventHandler(msg.Ty); eventHandler != nil {
		//log.Debug("HandleEvent2", "msgTy", msg.Ty)
		eventHandler(msg)
	} else {

		log.Error("HandleEvent", "unknown msgTy", msg.Ty)
	}
}
