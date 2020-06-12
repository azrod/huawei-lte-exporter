// Copyright (C) Mickael Stanislas <contact@mickael-stanislas.com>

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"common-nighthawk/go-figure"
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/op/go-logging"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const (
	banner = "Huawei LTE Exporteur\n"
)

var _log = logging.MustGetLogger("example")
var format = logging.MustStringFormatter(
	`%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`,
)

func main() {

	flag.StringVar(&listenAddress, "web.listen-address", ":12112", "Address to listen on for web interface and telemetry.")
	flag.StringVar(&metricsPath, "web.telemetry-path", "/metrics", "Path under which to expose metrics.")
	flag.StringVar(&huaweiEndpoint, "endpoint", "http://192.168.8.1/", "Endpoint of huawei device")
	flag.StringVar(&huaweiUsername, "username", "foo", "Username of huawei device")
	flag.StringVar(&huaweiPassword, "password", "bar", "Password of huawei device")
	flag.Usage = func() {
		fmt.Fprint(os.Stderr, fmt.Sprintf(banner))
		flag.PrintDefaults()
	}
	flag.Parse()

	backend1 := logging.NewLogBackend(os.Stderr, "", 0)
	backend2 := logging.NewLogBackend(os.Stderr, "", 0)
	backend2Formatter := logging.NewBackendFormatter(backend2, format)
	backend1Leveled := logging.AddModuleLevel(backend1)
	backend1Leveled.SetLevel(logging.CRITICAL, "")
	logging.SetBackend(backend1Leveled, backend2Formatter)

	myFigure := figure.NewFigure("Huawei LTE Exporter", "", true)
	myFigure.Print()
	fmt.Println("Huawei LTE Exporter : v0.1")
	fmt.Println("---------------------------------------------------------------------------------------------")
	fmt.Println("")

	huawei := huaweiConnect()
	_log.Info("Register metrics")
	huawei.registerMetrics()
	huawei.recordMetrics()

	http.Handle(metricsPath, promhttp.Handler())
	http.ListenAndServe(listenAddress, nil)
}
