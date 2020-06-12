package main

import (
	"reflect"

	"github.com/knq/hilink"
	"github.com/prometheus/client_golang/prometheus"
)

// TrafficInfo Huawei trafficInfo
type metric struct {
	MetricName  string
	MetricValue float64
}

type huawei struct {
	client    *hilink.Client
	connected bool
}

var (
	debug          bool
	vrs            bool
	listenAddress  string
	metricsPath    string
	huaweiEndpoint string
	huaweiUsername string
	huaweiPassword string
	endpoint       string
	username       string
	password       string
	ids            string
	errorInterface = reflect.TypeOf((*error)(nil)).Elem()

	// TotalUpload bits
	TotalUpload = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "huawei_TotalUpload",
			Help: "Total upload",
		},
		nil,
	)
	// CurrentConnectTime time
	CurrentConnectTime = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "huawei_CurrentConnectTime",
			Help: "CurrentConnectTime",
		},
		nil,
	)
	// CurrentDownloadRate Download size
	CurrentDownloadRate = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "huawei_CurrentDownloadRate",
			Help: "CurrentDownloadRate",
		},
		nil,
	)
	// CurrentDownload Download size
	CurrentDownload = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "huawei_CurrentDownload",
			Help: "CurrentDownload",
		},
		nil,
	)
	// TotalDownload Download size
	TotalDownload = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "huawei_TotalDownload",
			Help: "TotalDownload",
		},
		nil,
	)
	// showtraffic Download size
	showtraffic = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "huawei_showtraffic",
			Help: "showtraffic",
		},
		nil,
	)
	// CurrentUpload Download size
	CurrentUpload = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "huawei_CurrentUpload",
			Help: "CurrentUpload",
		},
		nil,
	)
	// CurrentUploadRate Download size
	CurrentUploadRate = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "huawei_CurrentUploadRate",
			Help: "CurrentUploadRate",
		},
		nil,
	)
	// TotalConnectTime Download size
	TotalConnectTime = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "huawei_TotalConnectTime",
			Help: "TotalConnectTime",
		},
		nil,
	)

	// SignalRsrp Rsrp
	SignalRsrp = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "huawei_Signal_Rsrp",
			Help: "SignalRsrp",
		},
		nil,
	)
	// SignalRssi Rssi
	SignalRssi = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "huawei_Signal_Rssi",
			Help: "SignalRssi",
		},
		nil,
	)
	// SignalRsrq Rsrq
	SignalRsrq = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "huawei_Signal_Rsrq",
			Help: "SignalRsrq",
		},
		nil,
	)
	// SignalSinr Sinr
	SignalSinr = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "huawei_Signal_Sinr",
			Help: "SignalSinr",
		},
		nil,
	)
)
