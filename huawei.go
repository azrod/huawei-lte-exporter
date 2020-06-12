package main

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"time"

	"github.com/knq/hilink"
	"github.com/prometheus/client_golang/prometheus"
)

func huaweiConnect() *huawei {
	_log.Info("Connecting to Huawei router : " + huaweiEndpoint)
	// hilink options
	opts := []hilink.Option{
		hilink.URL(huaweiEndpoint),
		hilink.Auth(huaweiUsername, huaweiPassword),
	}
	reflect.TypeOf(&hilink.Client{})
	cli, err := hilink.NewClient(opts...)
	if err != nil {
		_log.Error("Init Huawei client fail")
		panic(err)
	} else {
		_log.Info("Connected to Huawei router")
	}

	return &huawei{client: cli, connected: true}
}

func (huawei *huawei) registerTrafficInfo() {

	TrafficInfo, _ := huawei.client.TrafficInfo()
	for trafficInfometricName := range TrafficInfo {
		switch trafficInfometricName {
		case "TotalUpload":
			prometheus.MustRegister(TotalUpload)
			_log.Debug("Register metric huawei_" + trafficInfometricName)
		case "CurrentConnectTime":
			prometheus.MustRegister(CurrentConnectTime)
			_log.Debug("Register metric huawei_" + trafficInfometricName)
		case "CurrentDownloadRate":
			prometheus.MustRegister(CurrentDownloadRate)
			_log.Debug("Register metric huawei_" + trafficInfometricName)
		case "CurrentDownload":
			prometheus.MustRegister(CurrentDownload)
			_log.Debug("Register metric huawei_" + trafficInfometricName)
		case "TotalDownload":
			prometheus.MustRegister(TotalDownload)
			_log.Debug("Register metric huawei_" + trafficInfometricName)
		case "showtraffic":
			prometheus.MustRegister(showtraffic)
			_log.Debug("Register metric huawei_" + trafficInfometricName)
		case "CurrentUpload":
			prometheus.MustRegister(CurrentUpload)
			_log.Debug("Register metric huawei_" + trafficInfometricName)
		case "CurrentUploadRate":
			prometheus.MustRegister(CurrentUploadRate)
			_log.Debug("Register metric huawei_" + trafficInfometricName)
		case "TotalConnectTime":
			prometheus.MustRegister(TotalConnectTime)
			_log.Debug("Register metric huawei_" + trafficInfometricName)
		}
	}
}

func (huawei *huawei) registerSignalInfo() {

	SignalInfo, _ := huawei.client.SignalInfo()
	for signalInfometricName := range SignalInfo {

		switch signalInfometricName {
		case "rsrp":
			prometheus.MustRegister(SignalRsrp)
			_log.Debug("Register metric huawei_" + signalInfometricName)
		case "rssi":
			prometheus.MustRegister(SignalRssi)
			_log.Debug("Register metric huawei_" + signalInfometricName)
		case "rsrq":
			prometheus.MustRegister(SignalRsrq)
			_log.Debug("Register metric huawei_" + signalInfometricName)
		case "sinr":
			prometheus.MustRegister(SignalSinr)
			_log.Debug("Register metric huawei_" + signalInfometricName)
		}
	}
}

func (huawei *huawei) registerMetrics() {
	huawei.registerTrafficInfo()
	huawei.registerSignalInfo()

	// NETWORK NAME
	// ConnectionInfo, _ := huawei.client.NetworkInfo()

}

func (huawei *huawei) recordMetrics() {
	go func() {
		for {
			TrafficInfo, _ := huawei.client.TrafficInfo()

			for n, v := range TrafficInfo {
				m := &metric{}
				m.MetricName = n
				s := fmt.Sprintf("%v", v)
				_int, _ := strconv.ParseInt(s, 10, 64)
				m.MetricValue = float64(_int)
				switch m.MetricName {
				case "TotalUpload":
					TotalUpload.WithLabelValues().Set(m.MetricValue)
				case "CurrentConnectTime":
					CurrentConnectTime.WithLabelValues().Set(m.MetricValue)
				case "CurrentDownloadRate":
					CurrentDownloadRate.WithLabelValues().Set(m.MetricValue)
				case "CurrentDownload":
					CurrentDownload.WithLabelValues().Set(m.MetricValue)
				case "TotalDownload":
					TotalDownload.WithLabelValues().Set(m.MetricValue)
				case "showtraffic":
					showtraffic.WithLabelValues().Set(m.MetricValue)
				case "CurrentUpload":
					CurrentUpload.WithLabelValues().Set(m.MetricValue)
				case "CurrentUploadRate":
					CurrentUploadRate.WithLabelValues().Set(m.MetricValue)
				case "TotalConnectTime":
					TotalConnectTime.WithLabelValues().Set(m.MetricValue)
				}
			}
			SignalInfo, _ := huawei.client.SignalInfo()
			for n, v := range SignalInfo {
				m := &metric{}
				m.MetricName = n
				s := fmt.Sprintf("%v", v)
				zp := regexp.MustCompile(`[a-zA-Z]+`)
				val := zp.Split(s, 2)
				_int, _ := strconv.ParseInt(val[0], 10, 64)
				m.MetricValue = float64(_int)
				switch m.MetricName {
				case "rsrp":
					SignalRsrp.WithLabelValues().Set(m.MetricValue)
				case "rssi":
					SignalRssi.WithLabelValues().Set(m.MetricValue)
				case "rsrq":
					SignalRsrq.WithLabelValues().Set(m.MetricValue)
				case "sinr":
					SignalSinr.WithLabelValues().Set(m.MetricValue)
				}
			}
			time.Sleep(5 * time.Second)
		}
	}()
}
