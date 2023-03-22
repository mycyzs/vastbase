package collector

import (
	"database/sql"

	"github.com/prometheus/client_golang/prometheus"
)

func dbConnectState(db *sql.DB, ch chan<- prometheus.Metric) error {
	err := db.Ping()
	const UP = 1
	const DOWN = 0
	if err != nil {
		ch <- prometheus.MustNewConstMetric(NewDesc("dbConnectState", "数据库连接状态", []string{}), prometheus.GaugeValue, DOWN)
		logging.Debug("connect db error.")
		return err
	} else {
		ch <- prometheus.MustNewConstMetric(NewDesc("dbConnectState", "数据库连接状态", []string{}), prometheus.GaugeValue, UP)
		logging.Debug("connect db success.")
		return nil
	}
}
