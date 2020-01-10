package metrics

import (
	"time"

	chain33log "github.com/33cn/chain33/common/log/log15"
	"github.com/33cn/chain33/metrics/influxdb"
	"github.com/33cn/chain33/types"
	go_metrics "github.com/rcrowley/go-metrics"
)

type influxDBPara struct {
	// 以纳秒为单位
	Duration             int64    `json:"duration,omitempty"`
	Url                  string   `json:"url,omitempty"`
	Database             string   `json:"database,omitempty"`
	Username             string   `json:"username,omitempty"`
	Password             string   `json:"password,omitempty"`
	Namespace            string   `json:"namespace,omitempty"`
}

var (
	log = chain33log.New("module", "chain33 metrics")
)

//根据配置文件相关参数启动m
func StartMetrics (cfg *types.Chain33Config) {
	metrics := cfg.GetModuleConfig().Metrics
	if !metrics.EnableMetrics {
		return
	}

	switch metrics.DataEmitMode {
	case "influxdb":
		sub := cfg.GetSubConfig().Metrics
		subcfg, ok := sub[metrics.DataEmitMode]
		if !ok {
			log.Error("nil parameter for influxdb")
		}
		var influxdbcfg influxDBPara
		types.MustDecode(subcfg, &influxdbcfg)
		log.Info("StartMetrics with influxdb", "influxdbcfg.Duration", influxdbcfg.Duration,
		"influxdbcfg.Url", influxdbcfg.Url,
		"influxdbcfg.DatabaseName,", influxdbcfg.Database,
		"influxdbcfg.Username", influxdbcfg.Username,
		"influxdbcfg.Password", influxdbcfg.Password,
		"influxdbcfg.Namespace", influxdbcfg.Namespace)
		go influxdb.InfluxDBWithTags(go_metrics.DefaultRegistry,
			time.Duration(influxdbcfg.Duration),
			influxdbcfg.Url,
			influxdbcfg.Database,
			influxdbcfg.Username,
			influxdbcfg.Password,
			influxdbcfg.Namespace,
			nil)
	default:
		log.Error("startMetrics", "The dataEmitMode set is not supported now ", metrics.DataEmitMode)
		return
	}
}