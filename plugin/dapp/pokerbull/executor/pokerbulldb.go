package executor

import (
	"gitlab.33.cn/chain33/chain33/types"
	"gitlab.33.cn/chain33/chain33/account"
	dbm "gitlab.33.cn/chain33/chain33/common/db"
	"gitlab.33.cn/chain33/chain33/common"
	pkt "gitlab.33.cn/chain33/chain33/plugin/dapp/pokerbull/types"
	"strconv"
	"sort"
	"gitlab.33.cn/chain33/chain33/system/dapp"
)

const (
    ListDESC = int32(0)
    ListASC  = int32(1)

	DefaultCount  = int32(20)  //默认一次取多少条记录
)

type Action struct {
	coinsAccount *account.DB
	db           dbm.KV
	txhash       []byte
	fromaddr     string
	blocktime    int64
	height       int64
	execaddr     string
	localDB      dbm.Lister
	index        int
}

func NewAction(pb *PokerBull, tx *types.Transaction, index int) *Action {
	hash := tx.Hash()
	fromaddr := tx.From()

	return &Action{pb.GetCoinsAccount(), pb.GetStateDB(), hash, fromaddr,
		pb.GetBlockTime(), pb.GetHeight(), dapp.ExecAddress(string(tx.Execer)), pb.GetLocalDB(), index}
}

func (action *Action) CheckExecAccountBalance(fromAddr string, ToFrozen, ToActive int64) bool {
	acc := action.coinsAccount.LoadExecAccount(fromAddr, action.execaddr)
	if acc.GetBalance() >= ToFrozen && acc.GetFrozen() >= ToActive {
		return true
	}
	return false
}

func Key(id string) (key []byte) {
	key = append(key, []byte("mavl-"+types.ExecName(pkt.PokerBullX)+"-")...)
	key = append(key, []byte(id)...)
	return key
}

func readGame(db dbm.KV, id string) (*pkt.PokerBull, error) {
	data, err := db.Get(Key(id))
	if err != nil {
		logger.Error("query data have err:", err.Error())
		return nil, err
	}
	var game pkt.PokerBull
	//decode
	err = types.Decode(data, &game)
	if err != nil {
		logger.Error("decode game have err:", err.Error())
		return nil, err
	}
	return &game, nil
}

//安全批量查询方式,防止因为脏数据导致查询接口奔溃
func GetGameList(db dbm.KV, values []string) []*pkt.PokerBull {
	var games []*pkt.PokerBull
	for _, value := range values {
		game, err := readGame(db, value)
		if err != nil {
			continue
		}
		games = append(games, game)
	}
	return games
}

func queryGameListByStatus(db dbm.Lister, stat int32) ([]string, error) {
	values, err := db.List(calcPBGameStatusPrefix(stat), nil, DefaultCount, ListDESC)
	if err != nil {
		return nil, err
	}

	var gameIds []string
	for _, value := range values {
		var record pkt.PBGameRecord
		err := types.Decode(value, &record)
		if err != nil {
			continue
		}
		gameIds = append(gameIds, record.GetGameId())
	}

	return gameIds, nil
}

func Infos(db dbm.KV, infos *pkt.QueryPBGameInfos) (types.Message, error) {
	var games []*pkt.PokerBull
	for i := 0; i < len(infos.GameIds); i++ {
		id := infos.GameIds[i]
		game, err := readGame(db, id)
		if err != nil {
			return nil, err
		}
		games = append(games, game)
	}
	return &pkt.ReplyPBGameList{Games: games}, nil
}

func queryGameListByStatusAndPlayer(db dbm.Lister, stat int32, player int32) ([]string, error) {
	values, err := db.List(calcPBGameStatusAndPlayerPrefix(stat, player), nil, DefaultCount, ListDESC)
	if err != nil {
		return nil, err
	}

	var gameIds []string
	for _, value := range values {
		var record pkt.PBGameRecord
		err := types.Decode(value, &record)
		if err != nil {
			continue
		}
		gameIds = append(gameIds, record.GetGameId())
	}

	return gameIds, nil
}

func queryGameList(db dbm.Lister, stateDB dbm.KV, param *pkt.QueryPBGameListByStatusAndPlayerNum) (types.Message, error) {
	var gameIds []string
	var err error
	if param.PlayerNum == 0 {
		gameIds,err = queryGameListByStatus(db, param.Status)
	} else {
		gameIds,err = queryGameListByStatusAndPlayer(db, param.Status, param.PlayerNum)
	}
	if err != nil {
		return nil, err
	}

	return &pkt.ReplyPBGameList{GetGameList(stateDB, gameIds)}, nil
}

func (action *Action) saveGame(game *pkt.PokerBull) (kvset []*types.KeyValue) {
	value := types.Encode(game)
	action.db.Set(Key(game.GetGameId()), value)
	kvset = append(kvset, &types.KeyValue{Key(game.GameId), value})
	return kvset
}

func (action *Action) getIndex(game *pkt.PokerBull) int64 {
	return action.height*types.MaxTxsPerBlock + int64(action.index)
}

func (action *Action) GetReceiptLog(game *pkt.PokerBull) *types.ReceiptLog {
	log := &types.ReceiptLog{}
	r := &pkt.ReceiptPBGame{}
	r.Addr = action.fromaddr
	r.GameId = game.GameId
	r.Status = game.Status
	r.Index = game.GetIndex()
	r.PrevIndex = game.GetPrevIndex()
	r.PlayerNum = game.PlayerNum
	log.Log = types.Encode(r)
	return log
}

func (action *Action) readGame(id string) (*pkt.PokerBull, error) {
	data, err := action.db.Get(Key(id))
	if err != nil {
		return nil, err
	}
	var game pkt.PokerBull
	//decode
	err = types.Decode(data, &game)
	if err != nil {
		return nil, err
	}
	return &game, nil
}

func (action *Action) calculate(game *pkt.PokerBull) *pkt.PBResult{
	var handS HandSlice = make([]*pkt.PBHand, 1)
	for _, player := range game.Players {
		hand := &pkt.PBHand{}
		hand.Cards = Deal(game.Poker, player.TxHash) //发牌
		hand.Result = Result(hand.Cards) //计算结果
		hand.Address = player.Address

		//存入玩家数组
		player.Hands = append(player.Hands, hand)

		//存入临时切片待比大小排序
		handS = append(handS, hand)

		//为下一个continue状态初始化player
		player.Ready = false
	}

	// 升序排列
	if !sort.IsSorted(handS) {
		sort.Sort(handS)
	}

	// 将有序的临时切片加入到结果数组
	result := &pkt.PBResult{}
	result.Hands = make([]*pkt.PBHand, len(handS))
	copy(result.Hands, handS)
	result.Winner = handS[len(handS)-1].Address

	game.Results = make([]*pkt.PBResult, 1)
	game.Results = append(game.Results, result)

	return result
}

func (action *Action) gameCheckOut(game *pkt.PokerBull) ([]*types.ReceiptLog, []*types.KeyValue, error) {
	var logs []*types.ReceiptLog
	var kv []*types.KeyValue
	result := action.calculate(game)
	for _,player := range game.Players {
		if player.Address == result.Winner {
			receipt, err := action.coinsAccount.ExecActive(player.GetAddress(), action.execaddr, game.GetValue())
			if err != nil {
				logger.Error("GameClose.execActive", "addr", player.GetAddress(), "execaddr", action.execaddr, "amount", game.GetValue(),
					"err", err)
				return nil, nil, err
			}
			logs = append(logs, receipt.Logs...)
			kv = append(kv, receipt.KV...)
			continue
		}

		receipt, err := action.coinsAccount.ExecTransferFrozen(player.Address, result.Winner, action.execaddr, game.GetValue())
		if err != nil {
			action.coinsAccount.ExecFrozen(result.Winner, action.execaddr, game.GetValue()) // rollback
			logger.Error("GameClose.ExecTransferFrozen", "addr", result.Winner, "execaddr", action.execaddr, "amount", game.GetValue(),
				"err", err)
			return nil, nil, err
		}
		logs = append(logs, receipt.Logs...)
		kv = append(kv, receipt.KV...)
	}

	return logs, kv, nil
}

func (action *Action) GameStart(start *pkt.PBGameStart) (*types.Receipt, error) {
	var logs []*types.ReceiptLog
	var kv []*types.KeyValue

	gameId := common.ToHex(action.txhash)
	if !action.CheckExecAccountBalance(action.fromaddr, start.GetValue(), 0) {
		logger.Error("GameStart", "addr", action.fromaddr, "execaddr", action.execaddr, "id",
			gameId, "err", types.ErrNoBalance)
		return nil, types.ErrNoBalance
	}

	//发牌随机数取txhash
	txrng,err := strconv.ParseInt(gameId, 0, 64)
	if err != nil {
		return nil, err
	}

	var game *pkt.PokerBull
	ids, err := queryGameListByStatusAndPlayer(action.localDB, pkt.PBGameActionStart, start.PlayerNum)
	if err != nil || len(ids) == 0 {
		if err != types.ErrNotFound {
			return nil, err
		}

		// 没有匹配到要求的牌局，创建一个
		game = &pkt.PokerBull{
			GameId:        gameId,
			Status:        pkt.PBGameActionStart,
			StartTime:     action.blocktime,
			StartTxHash:   gameId,
			Value:         start.GetValue(),
			Poker:         NewPoker(),
			Players:       make([]*pkt.PBPlayer, 1),
			PlayerNum:     start.PlayerNum,
			Results:       make([]*pkt.PBResult, 1),
			Index:         action.getIndex(game),
		}

		Shuffle(game.Poker, action.blocktime) //洗牌
	} else {
		id := ids[0] // 取第一个牌局加入
		game, err := action.readGame(id)
		if err != nil {
			logger.Error("Poker bull game start", "addr", action.fromaddr, "execaddr", action.execaddr, "get game failed", id, "err", err)
			return nil, err
		}
		game.PrevIndex = game.Index
		game.Index = action.getIndex(game)
	}

	//加入当前玩家信息
	game.Players = append(game.Players, &pkt.PBPlayer{
		Address:  action.fromaddr,
		TxHash:   txrng,
		Ready:    false,
		Hands:    make([]*pkt.PBHand, 1),
	})

	//冻结子账户资金
	receiptO, err := action.coinsAccount.ExecFrozen(action.fromaddr, action.execaddr, start.GetValue())
	if err != nil {
		logger.Error("GameCreate.ExecFrozen", "addr", action.fromaddr, "execaddr", action.execaddr, "amount", start.GetValue(), "err", err.Error())
		return nil, err
	}

	// 如果人数达标，则发牌计算斗牛结果
	if len(game.Players) == int(game.PlayerNum) {
		logsH, kvH, err := action.gameCheckOut(game)
		if err != nil {
			return nil, err
		}
		logs = append(logs, logsH...)
		kv = append(kv, kvH...)

		game.Status = pkt.PBGameActionContinue // 更新游戏状态
	}
	receiptLog := action.GetReceiptLog(game)
	logs = append(logs, receiptLog)
	logs = append(logs, receiptO.Logs...)

	kv = append(kv, action.saveGame(game)...)
	kv = append(kv, receiptO.KV...)
	receiptO = &types.Receipt{types.ExecOk, kv, logs}
	return receiptO, nil
}

func getReadyPlayerNum(players []*pkt.PBPlayer) int {
	var readyC = 0
	for _, player := range players {
		if player.Ready == true {
			readyC++
		}
	}
	return  readyC
}

func getPlayerFromAddress(players []*pkt.PBPlayer, addr string) *pkt.PBPlayer {
	for _, player := range players {
		if player.Address == addr {
			return player
		}
	}
	return nil
}

func (action *Action) GameContinue(pbcontinue *pkt.PBGameContinue) (*types.Receipt, error) {
	var logs []*types.ReceiptLog
	var kv []*types.KeyValue

	game, err := action.readGame(pbcontinue.GetGameId())
	if err != nil {
		logger.Error("GameContinue", "addr", action.fromaddr, "execaddr", action.execaddr, "get game failed",
			pbcontinue.GetGameId(), "err", err)
		return nil, err
	}

	if game.Status != pkt.PBGameActionContinue {
		logger.Error("GameContinue", "addr", action.fromaddr, "execaddr", action.execaddr, "Status error",
			pbcontinue.GetGameId())
		return nil, err
	}

	// 检查余额
	if !action.CheckExecAccountBalance(action.fromaddr, game.GetValue(), 0) {
		logger.Error("GameStart", "addr", action.fromaddr, "execaddr", action.execaddr, "id",
			pbcontinue.GetGameId(), "err", types.ErrNoBalance)
		return nil, types.ErrNoBalance
	}

	// 寻找对应玩家
	pbplayer := getPlayerFromAddress(game.Players, action.fromaddr)
	if pbplayer == nil {
		logger.Error("GameContinue", "addr", action.fromaddr, "execaddr", action.execaddr, "get game player failed",
			pbcontinue.GetGameId(), "err", types.ErrNotFound)
		return nil, types.ErrNotFound
	}
	pbplayer.Ready = true

	//发牌随机数取txhash
	txrng,err := strconv.ParseInt(common.ToHex(action.txhash), 0, 64)
	if err != nil {
		return nil, err
	}
	pbplayer.TxHash = txrng

	game.PrevIndex = game.Index
	game.Index = action.getIndex(game)

	//冻结子账户资金
	receipt, err := action.coinsAccount.ExecFrozen(action.fromaddr, action.execaddr, game.GetValue())
	if err != nil {
		logger.Error("GameCreate.ExecFrozen", "addr", action.fromaddr, "execaddr", action.execaddr, "amount", game.GetValue(), "err", err.Error())
		return nil, err
	}

	if getReadyPlayerNum(game.Players) == int(game.PlayerNum) {
		logsH, kvH, err := action.gameCheckOut(game)
		if err != nil {
			return nil, err
		}
		logs = append(logs, logsH...)
		kv = append(kv, kvH...)
	}

	receiptLog := action.GetReceiptLog(game)
	logs = append(logs, receiptLog)
	logs = append(logs, receipt.Logs...)

	kv = append(kv, action.saveGame(game)...)
	kv = append(kv, receipt.KV...)
	receipt = &types.Receipt{types.ExecOk, kv, logs}
	return receipt, nil
}

func (action *Action) GameQuit(pbend *pkt.PBGameQuit) (*types.Receipt, error) {
	game, err := action.readGame(pbend.GetGameId())
	if err != nil {
		logger.Error("GameEnd", "addr", action.fromaddr, "execaddr", action.execaddr, "get game failed",
			pbend.GetGameId(), "err", err)
		return nil, err
	}

	game.Status = pkt.PBGameActionQuit
	game.PrevIndex = game.Index
	game.Index = action.getIndex(game)
	game.QuitTime = action.blocktime
	game.QuitTxHash = common.ToHex(action.txhash)

	var logs []*types.ReceiptLog
	var kv []*types.KeyValue
	receiptLog := action.GetReceiptLog(game)
	logs = append(logs, receiptLog)
	kv = append(kv, action.saveGame(game)...)
	return &types.Receipt{types.ExecOk, kv, logs}, nil
}

type HandSlice []*pkt.PBHand

func (h HandSlice) Len() int {
	return len(h)
}

func (h HandSlice) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h HandSlice) Less(i, j int) bool {
	if h[i].Result < h[j].Result {
		return true
	}

	if h[i].Result == h[j].Result {
		return Compare(h[i].Cards, h[j].Cards)
	}

	return false
}