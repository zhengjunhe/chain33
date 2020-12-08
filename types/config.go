// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package types

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"strconv"
	"strings"
	"sync"

	"fmt"

	"github.com/33cn/dplatform/types/chaincfg"
	tml "github.com/BurntSushi/toml"
)

//Create ...
type Create func(cfg *DplatformConfig)

//区块链共识相关的参数，重要参数不要随便修改
var (
	AllowUserExec = [][]byte{ExecerNone}
	EmptyValue    = []byte("FFFFFFFFemptyBVBiCj5jvE15pEiwro8TQRGnJSNsJF") //这字符串表示数据库中的空值
	cliSysParam   = make(map[string]*DplatformConfig)                     // map key is title
	regModuleInit = make(map[string]Create)
	regExecInit   = make(map[string]Create)
	runonce       = sync.Once{}
)

// coin conversation
const (
	Coin            int64 = 1e8
	MaxCoin         int64 = 1e17
	MaxTxSize             = 100000 //100K
	MaxTxGroupSize  int32 = 20
	MaxBlockSize          = 20000000 //20M
	MaxTxsPerBlock        = 100000
	TokenPrecision  int64 = 1e8
	MaxTokenBalance int64 = 900 * 1e8 * TokenPrecision //900亿
	DefaultMinFee   int64 = 1e5
)

//DplatformConfig ...
type DplatformConfig struct {
	mcfg            *Config
	scfg            *ConfigSubModule
	minerExecs      []string
	title           string
	mu              sync.Mutex
	chainConfig     map[string]interface{}
	mver            *mversion
	coinSymbol      string
	forks           *Forks
	enableCheckFork bool
}

//ChainParam 结构体
type ChainParam struct {
	MaxTxNumber  int64
	PowLimitBits uint32
}

//RegFork Reg 注册每个模块的自动初始化函数
func RegFork(name string, create Create) {
	if create == nil {
		panic("config: Register Module Init is nil")
	}
	if _, dup := regModuleInit[name]; dup {
		panic("config: Register Init called twice for driver " + name)
	}
	regModuleInit[name] = create
}

//RegForkInit ...
func RegForkInit(cfg *DplatformConfig) {
	for _, item := range regModuleInit {
		item(cfg)
	}
}

//RegExec ...
func RegExec(name string, create Create) {
	if create == nil {
		panic("config: Register Exec Init is nil")
	}
	if _, dup := regExecInit[name]; dup {
		panic("config: Register Exec called twice for driver " + name)
	}
	regExecInit[name] = create
}

//RegExecInit ...
func RegExecInit(cfg *DplatformConfig) {
	runonce.Do(func() {
		for _, item := range regExecInit {
			item(cfg)
		}
	})
}

//NewDplatformConfig ...
func NewDplatformConfig(cfgstring string) *DplatformConfig {
	dplatformCfg := NewDplatformConfigNoInit(cfgstring)
	dplatformCfg.dplatformCfgInit(dplatformCfg.mcfg)
	return dplatformCfg
}

//NewDplatformConfigNoInit ...
func NewDplatformConfigNoInit(cfgstring string) *DplatformConfig {
	cfg, sub := InitCfgString(cfgstring)
	dplatformCfg := &DplatformConfig{
		mcfg:        cfg,
		scfg:        sub,
		minerExecs:  []string{"ticket"}, //挖矿的合约名单，适配旧配置，默认ticket
		title:       cfg.Title,
		chainConfig: make(map[string]interface{}),
		coinSymbol:  "dpom",
		forks:       &Forks{make(map[string]int64)},
	}
	// 先将每个模块的fork初始化到DplatformConfig中，然后如果需要再将toml中的替换
	dplatformCfg.setDefaultConfig()
	dplatformCfg.setFlatConfig(cfgstring)
	dplatformCfg.setMver(cfgstring)
	// TODO 需要测试是否与NewDplatformConfig分开
	RegForkInit(dplatformCfg)
	RegExecInit(dplatformCfg)
	return dplatformCfg
}

//GetModuleConfig ...
func (c *DplatformConfig) GetModuleConfig() *Config {
	return c.mcfg
}

//GetSubConfig ...
func (c *DplatformConfig) GetSubConfig() *ConfigSubModule {
	return c.scfg
}

//EnableCheckFork ...
func (c *DplatformConfig) EnableCheckFork(enable bool) {
	c.enableCheckFork = false
}

//GetForks ...
func (c *DplatformConfig) GetForks() (map[string]int64, error) {
	if c.forks == nil {
		return nil, ErrNotFound
	}
	return c.forks.forks, nil
}

func (c *DplatformConfig) setDefaultConfig() {
	c.S("TestNet", false)
	c.SetMinFee(DefaultMinFee)
	for key, cfg := range chaincfg.LoadAll() {
		c.S("cfg."+key, cfg)
	}
	//防止报error 错误，不影响功能
	if !c.HasConf("cfg.dplatform") {
		c.S("cfg.dplatform", "")
	}
	if !c.HasConf("cfg.local") {
		c.S("cfg.local", "")
	}
}

func (c *DplatformConfig) setFlatConfig(cfgstring string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	cfg := make(map[string]interface{})
	if _, err := tml.Decode(cfgstring, &cfg); err != nil {
		panic(err)
	}
	flat := FlatConfig(cfg)
	for k, v := range flat {
		c.setChainConfig("config."+k, v)
	}
}

func (c *DplatformConfig) setChainConfig(key string, value interface{}) {
	c.chainConfig[key] = value
}

func (c *DplatformConfig) getChainConfig(key string) (value interface{}, err error) {
	if data, ok := c.chainConfig[key]; ok {
		return data, nil
	}
	//报错警告
	tlog.Error("chain config " + key + " not found")
	return nil, ErrNotFound
}

// Init 初始化
func (c *DplatformConfig) dplatformCfgInit(cfg *Config) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.forks == nil {
		c.forks = &Forks{}
	}
	c.forks.SetTestNetFork()

	if cfg != nil {
		if c.isLocal() {
			c.setTestNet(true)
		} else {
			c.setTestNet(cfg.TestNet)
		}

		if cfg.Wallet.MinFee < cfg.Mempool.MinTxFeeRate {
			panic("config must meet: wallet.minFee >= mempool.minTxFeeRate")
		}
		if cfg.Mempool.MaxTxFeeRate == 0 {
			cfg.Mempool.MaxTxFeeRate = 1e7 //0.1 coins
		}
		if cfg.Mempool.MaxTxFee == 0 {
			cfg.Mempool.MaxTxFee = 1e9 // 10 coins
		}
		c.setTxFeeConfig(cfg.Mempool.MinTxFeeRate, cfg.Mempool.MaxTxFeeRate, cfg.Mempool.MaxTxFee)
		if cfg.Consensus != nil {
			c.setMinerExecs(cfg.Consensus.MinerExecs)
		}
		c.setChainConfig("FixTime", cfg.FixTime)
		if cfg.CoinSymbol != "" {
			if strings.Contains(cfg.CoinSymbol, "-") {
				panic("config CoinSymbol must without '-'")
			}
			c.coinSymbol = cfg.CoinSymbol
		} else {
			if c.isPara() {
				panic("must config CoinSymbol in para chain")
			} else {
				c.coinSymbol = DefaultCoinsSymbol
			}
		}
		//TxHeight
		c.setChainConfig("TxHeight", cfg.TxHeight)
	}
	if c.needSetForkZero() { //local 只用于单元测试
		if c.isLocal() {
			c.forks.setLocalFork()
			c.setChainConfig("Debug", true)
		} else {
			c.forks.setForkForParaZero()
		}
	} else {
		if cfg != nil && cfg.Fork != nil {
			c.initForkConfig(cfg.Fork)
		}
	}
	// 更新fork配置信息
	if c.mver != nil {
		c.mver.UpdateFork(c.forks)
	}
}

func (c *DplatformConfig) needSetForkZero() bool {
	if c.isLocal() {
		return true
	} else if c.isPara() &&
		(c.mcfg == nil || c.mcfg.Fork == nil || c.mcfg.Fork.System == nil) &&
		!c.mcfg.EnableParaFork {
		//如果para 没有配置fork，那么默认所有的fork 为 0（一般只用于测试）
		return true
	}
	return false
}

func (c *DplatformConfig) setTestNet(isTestNet bool) {
	if !isTestNet {
		c.setChainConfig("TestNet", false)
		return
	}
	c.setChainConfig("TestNet", true)
	//const 初始化TestNet 的初始化参数
}

// GetP 获取ChainParam
func (c *DplatformConfig) GetP(height int64) *ChainParam {
	conf := Conf(c, "mver.consensus")
	chain := &ChainParam{}
	chain.MaxTxNumber = conf.MGInt("maxTxNumber", height)
	chain.PowLimitBits = uint32(conf.MGInt("powLimitBits", height))
	return chain
}

// GetMinerExecs 获取挖矿的合约名单
func (c *DplatformConfig) GetMinerExecs() []string {
	return c.minerExecs
}

func (c *DplatformConfig) setMinerExecs(execs []string) {
	if len(execs) > 0 {
		c.minerExecs = execs
	}
}

// GetFundAddr 获取基金账户地址
func (c *DplatformConfig) GetFundAddr() string {
	return c.MGStr("mver.consensus.fundKeyAddr", 0)
}

// G 获取ChainConfig中的配置
func (c *DplatformConfig) G(key string) (value interface{}, err error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	value, err = c.getChainConfig(key)
	return
}

// MG 获取mver config中的配置
func (c *DplatformConfig) MG(key string, height int64) (value interface{}, err error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.mver == nil {
		panic("mver is nil")
	}
	return c.mver.Get(key, height)
}

// GStr 获取ChainConfig中的字符串格式
func (c *DplatformConfig) GStr(name string) string {
	value, err := c.G(name)
	if err != nil {
		return ""
	}
	if i, ok := value.(string); ok {
		return i
	}
	return ""
}

// MGStr 获取mver config 中的字符串格式
func (c *DplatformConfig) MGStr(name string, height int64) string {
	value, err := c.MG(name, height)
	if err != nil {
		return ""
	}
	if i, ok := value.(string); ok {
		return i
	}
	return ""
}

func parseInt(value interface{}) int64 {
	if i, ok := value.(int64); ok {
		return i
	}
	if s, ok := value.(string); ok {
		if strings.HasPrefix(s, "0x") {
			i, err := strconv.ParseUint(s, 0, 64)
			if err == nil {
				return int64(i)
			}
		}
	}
	return 0
}

// GInt 解析ChainConfig配置
func (c *DplatformConfig) GInt(name string) int64 {
	value, err := c.G(name)
	if err != nil {
		return 0
	}
	return parseInt(value)
}

// MGInt 解析mver config 配置
func (c *DplatformConfig) MGInt(name string, height int64) int64 {
	value, err := c.MG(name, height)
	if err != nil {
		return 0
	}
	return parseInt(value)
}

// IsEnable 解析ChainConfig配置
func (c *DplatformConfig) IsEnable(name string) bool {
	isenable, err := c.G(name)
	if err == nil && isenable.(bool) {
		return true
	}
	return false
}

// MIsEnable 解析mver config 配置
func (c *DplatformConfig) MIsEnable(name string, height int64) bool {
	isenable, err := c.MG(name, height)
	if err == nil && isenable.(bool) {
		return true
	}
	return false
}

// HasConf 解析chainConfig配置
func (c *DplatformConfig) HasConf(key string) bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	_, ok := c.chainConfig[key]
	return ok
}

// S 设置chainConfig配置
func (c *DplatformConfig) S(key string, value interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if strings.HasPrefix(key, "config.") {
		if !c.isLocal() && !c.isTestPara() { //only local and test para can modify for test
			panic("prefix config. is readonly")
		} else {
			tlog.Error("modify " + key + " is only for test")
		}
	}
	c.setChainConfig(key, value)
}

//SetTitleOnlyForTest set title only for test use
func (c *DplatformConfig) SetTitleOnlyForTest(ti string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.title = ti

}

// GetTitle 获取title
func (c *DplatformConfig) GetTitle() string {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.title
}

// GetCoinSymbol 获取 coin symbol
func (c *DplatformConfig) GetCoinSymbol() string {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.coinSymbol
}

func (c *DplatformConfig) isLocal() bool {
	return c.title == "local"
}

// IsLocal 是否locak title
func (c *DplatformConfig) IsLocal() bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.isLocal()
}

// GetMinTxFeeRate get min transaction fee rate
func (c *DplatformConfig) GetMinTxFeeRate() int64 {
	return c.GInt("MinTxFeeRate")
}

// GetMaxTxFeeRate get max transaction fee rate
func (c *DplatformConfig) GetMaxTxFeeRate() int64 {
	return c.GInt("MaxTxFeeRate")
}

// GetMaxTxFee get max transaction fee
func (c *DplatformConfig) GetMaxTxFee() int64 {
	return c.GInt("MaxTxFee")
}

// SetTxFeeConfig 设置交易费相关配置
func (c *DplatformConfig) SetTxFeeConfig(minTxFeeRate, maxTxFeeRate, maxTxFee int64) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.setTxFeeConfig(minTxFeeRate, maxTxFeeRate, maxTxFee)
}

func (c *DplatformConfig) setTxFeeConfig(minTxFeeRate, maxTxFeeRate, maxTxFee int64) {
	if minTxFeeRate < 0 {
		panic("minTxFeeRate less than zero")
	}

	if minTxFeeRate > maxTxFeeRate || maxTxFeeRate > maxTxFee {
		panic("SetTxFee, tx fee must meet, minTxFeeRate <= maxTxFeeRate <= maxTxFee")
	}
	c.setChainConfig("MinTxFeeRate", minTxFeeRate)
	c.setChainConfig("MaxTxFeeRate", maxTxFeeRate)
	c.setChainConfig("MaxTxFee", maxTxFee)
	c.setChainConfig("MinBalanceTransfer", minTxFeeRate*10)
}

// SetMinFee 设置最小费用
func (c *DplatformConfig) SetMinFee(fee int64) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.setTxFeeConfig(fee, fee*100, fee*10000)
}

func (c *DplatformConfig) isPara() bool {
	return strings.Count(c.title, ".") == 3 && strings.HasPrefix(c.title, ParaKeyX)
}

func (c *DplatformConfig) isTestPara() bool {
	return strings.Count(c.title, ".") == 3 && strings.HasPrefix(c.title, ParaKeyX) && strings.HasSuffix(c.title, "test.")
}

// IsPara 是否平行链
func (c *DplatformConfig) IsPara() bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.isPara()
}

// IsParaExecName 是否平行链执行器
func IsParaExecName(exec string) bool {
	return strings.HasPrefix(exec, ParaKeyX)
}

//IsMyParaExecName 是否是我的para链的执行器
func (c *DplatformConfig) IsMyParaExecName(exec string) bool {
	return IsParaExecName(exec) && strings.HasPrefix(exec, c.GetTitle())
}

//IsSpecificParaExecName 是否是某一个平行链的执行器
func IsSpecificParaExecName(title, exec string) bool {
	return IsParaExecName(exec) && strings.HasPrefix(exec, title)
}

//GetParaExecTitleName 如果是平行链执行器，获取对应title
func GetParaExecTitleName(exec string) (string, bool) {
	if IsParaExecName(exec) {
		for i := len(ParaKey); i < len(exec); i++ {
			if exec[i] == '.' {
				return exec[:i+1], true
			}
		}
	}
	return "", false
}

// IsTestNet 是否测试链
func (c *DplatformConfig) IsTestNet() bool {
	return c.IsEnable("TestNet")
}

// GetParaName 获取平行链name
func (c *DplatformConfig) GetParaName() string {
	if c.IsPara() {
		return c.GetTitle()
	}
	return ""
}

// FlagKV 获取kv对
func FlagKV(key []byte, value int64) *KeyValue {
	return &KeyValue{Key: key, Value: Encode(&Int64{Data: value})}
}

// MergeConfig Merge配置
func MergeConfig(conf map[string]interface{}, def map[string]interface{}) string {
	errstr := checkConfig("", conf, def)
	if errstr != "" {
		return errstr
	}
	mergeConfig(conf, def)
	return ""
}

//检查默认配置文件
func checkConfig(key string, conf map[string]interface{}, def map[string]interface{}) string {
	errstr := ""
	for key1, value1 := range conf {
		if vdef, ok := def[key1]; ok {
			conf1, ok1 := value1.(map[string]interface{})
			def1, ok2 := vdef.(map[string]interface{})
			if ok1 && ok2 {
				errstr += checkConfig(getkey(key, key1), conf1, def1)
			} else {
				errstr += "rewrite defalut key " + getkey(key, key1) + "\n"
			}
		}
	}
	return errstr
}

func mergeConfig(conf map[string]interface{}, def map[string]interface{}) {
	for key1, value1 := range def {
		if vdef, ok := conf[key1]; ok {
			conf1, ok1 := value1.(map[string]interface{})
			def1, ok2 := vdef.(map[string]interface{})
			if ok1 && ok2 {
				mergeConfig(conf1, def1)
				conf[key1] = conf1
			}
		} else {
			conf[key1] = value1
		}
	}
}

func getkey(key, key1 string) string {
	if key == "" {
		return key1
	}
	return key + "." + key1
}

//MergeCfg ...
func MergeCfg(cfgstring, cfgdefault string) string {
	if cfgdefault != "" {
		return mergeCfgString(cfgstring, cfgdefault)
	}
	return cfgstring
}

func mergeCfgString(cfgstring, cfgdefault string) string {
	//1. defconfig
	def := make(map[string]interface{})
	_, err := tml.Decode(cfgdefault, &def)
	if err != nil {
		panic(err)
	}
	//2. userconfig
	conf := make(map[string]interface{})
	_, err = tml.Decode(cfgstring, &conf)
	if err != nil {
		panic(err)
	}
	errstr := MergeConfig(conf, def)
	if errstr != "" {
		panic(errstr)
	}
	buf := new(bytes.Buffer)
	tml.NewEncoder(buf).Encode(conf)
	return buf.String()
}

func initCfgString(cfgstring string) (*Config, error) {
	var cfg Config
	if _, err := tml.Decode(cfgstring, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

// InitCfg 初始化配置
func InitCfg(path string) (*Config, *ConfigSubModule) {
	return InitCfgString(readFile(path))
}

func flatConfig(key string, conf map[string]interface{}, flat map[string]interface{}) {
	for key1, value1 := range conf {
		conf1, ok := value1.(map[string]interface{})
		if ok {
			flatConfig(getkey(key, key1), conf1, flat)
		} else {
			flat[getkey(key, key1)] = value1
		}
	}
}

// FlatConfig Flat配置
func FlatConfig(conf map[string]interface{}) map[string]interface{} {
	flat := make(map[string]interface{})
	flatConfig("", conf, flat)
	return flat
}

func (c *DplatformConfig) setMver(cfgstring string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.mver = newMversion(cfgstring)
}

// InitCfgString 初始化配置
func InitCfgString(cfgstring string) (*Config, *ConfigSubModule) {
	//cfgstring = c.mergeCfg(cfgstring) // TODO 是否可以去除
	//setFlatConfig(cfgstring)         // 将set的全部去除
	cfg, err := initCfgString(cfgstring)
	if err != nil {
		panic(err)
	}
	//setMver(cfg.Title, cfgstring)    // 将set的全部去除
	sub, err := initSubModuleString(cfgstring)
	if err != nil {
		panic(err)
	}
	return cfg, sub
}

// subModule 子模块结构体
type subModule struct {
	Store     map[string]interface{}
	Exec      map[string]interface{}
	Consensus map[string]interface{}
	Wallet    map[string]interface{}
	Mempool   map[string]interface{}
	Metrics   map[string]interface{}
	P2P       map[string]interface{}
}

//ReadFile ...
func ReadFile(path string) string {
	return readFile(path)
}

func readFile(path string) string {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(data)
}

func initSubModuleString(cfgstring string) (*ConfigSubModule, error) {
	var cfg subModule
	if _, err := tml.Decode(cfgstring, &cfg); err != nil {
		return nil, err
	}
	return parseSubModule(&cfg)
}

func parseSubModule(cfg *subModule) (*ConfigSubModule, error) {
	var subcfg ConfigSubModule
	subcfg.Store = parseItem(cfg.Store)
	subcfg.Exec = parseItem(cfg.Exec)
	subcfg.Consensus = parseItem(cfg.Consensus)
	subcfg.Wallet = parseItem(cfg.Wallet)
	subcfg.Mempool = parseItem(cfg.Mempool)
	subcfg.Metrics = parseItem(cfg.Metrics)
	subcfg.P2P = parseItem(cfg.P2P)
	return &subcfg, nil
}

//ModifySubConfig json data modify
func ModifySubConfig(sub []byte, key string, value interface{}) ([]byte, error) {
	var data map[string]interface{}
	err := json.Unmarshal(sub, &data)
	if err != nil {
		return nil, err
	}
	data[key] = value
	return json.Marshal(data)
}

func parseItem(data map[string]interface{}) map[string][]byte {
	subconfig := make(map[string][]byte)
	if len(data) == 0 {
		return subconfig
	}
	for key := range data {
		if key == "sub" {
			subcfg := data[key].(map[string]interface{})
			for k := range subcfg {
				subconfig[k], _ = json.Marshal(subcfg[k])
			}
		}
	}
	return subconfig
}

// ConfQuery 结构体
type ConfQuery struct {
	cfg    *DplatformConfig
	prefix string
}

// Conf 配置
func Conf(cfg *DplatformConfig, prefix string) *ConfQuery {
	if prefix == "" || (!strings.HasPrefix(prefix, "config.") && !strings.HasPrefix(prefix, "mver.")) {
		panic("ConfQuery must init buy prefix config. or mver.")
	}
	return &ConfQuery{cfg: cfg, prefix: prefix}
}

// ConfSub 子模块配置
func ConfSub(cfg *DplatformConfig, name string) *ConfQuery {
	return Conf(cfg, "config.exec.sub."+name)
}

// G 获取指定key的配置信息
func (query *ConfQuery) G(key string) (interface{}, error) {
	return query.cfg.G(getkey(query.prefix, key))
}

func parseStrList(data interface{}) []string {
	var list []string
	if item, ok := data.([]interface{}); ok {
		for i := 0; i < len(item); i++ {
			one, ok := item[i].(string)
			if ok {
				list = append(list, one)
			}
		}
	}
	return list
}

// GStrList 解析字符串列表
func (query *ConfQuery) GStrList(key string) []string {
	data, err := query.G(key)
	if err == nil {
		return parseStrList(data)
	}
	return []string{}
}

// GInt 解析int类型
func (query *ConfQuery) GInt(key string) int64 {
	return query.cfg.GInt(getkey(query.prefix, key))
}

// GStr 解析string类型
func (query *ConfQuery) GStr(key string) string {
	return query.cfg.GStr(getkey(query.prefix, key))
}

// IsEnable 解析bool类型
func (query *ConfQuery) IsEnable(key string) bool {
	return query.cfg.IsEnable(getkey(query.prefix, key))
}

// MG 解析mversion
func (query *ConfQuery) MG(key string, height int64) (interface{}, error) {
	return query.cfg.MG(getkey(query.prefix, key), height)
}

// MGInt 解析mversion int类型配置
func (query *ConfQuery) MGInt(key string, height int64) int64 {
	return query.cfg.MGInt(getkey(query.prefix, key), height)
}

// MGStr 解析mversion string类型配置
func (query *ConfQuery) MGStr(key string, height int64) string {
	return query.cfg.MGStr(getkey(query.prefix, key), height)
}

// MGStrList 解析mversion string list类型配置
func (query *ConfQuery) MGStrList(key string, height int64) []string {
	data, err := query.MG(key, height)
	if err == nil {
		return parseStrList(data)
	}
	return []string{}
}

// MIsEnable 解析mversion bool类型配置
func (query *ConfQuery) MIsEnable(key string, height int64) bool {
	return query.cfg.MIsEnable(getkey(query.prefix, key), height)
}

//SetCliSysParam ...
func SetCliSysParam(title string, cfg *DplatformConfig) {
	if cfg == nil {
		panic("set cli system DplatformConfig param is nil")
	}
	cliSysParam[title] = cfg
}

//GetCliSysParam ...
func GetCliSysParam(title string) *DplatformConfig {
	if v, ok := cliSysParam[title]; ok {
		return v
	}
	panic(fmt.Sprintln("can not find CliSysParam title", title))
}

//AssertConfig ...
func AssertConfig(check interface{}) {
	if check == nil {
		panic("check object is nil (DplatformConfig)")
	}
}
