# Huawei-LTE-Exporter

Huawei LTE exporter expose metrics from your LTE router. 

## Tested devices

* Huawei B525s-23a 
  
Normaly this exporter work for all device using Huawei hilink.
> You have tested on other model. Contact Me. 

## How to use ? 

```
# Minimal setup 
./huawei-lte-exporter -username <username> -password <password>
```

### Available parameters 
| Parameter Name     | Required | Default              | Description                   |
| ------------------ | -------- | -------------------- | ----------------------------- |
| username           | yes      | ' '                   | Username of your LTE router   |
| password           | yes      | ' '                   | Password of your LTE router   |
| endpoint           | no       | 'http://192.168.8.1' | Endpoint of your LTE router   |
|                    |          |                      |
| web.listen-address | no       | ':12112'             | Exposed port for your metrics |
| web.telemetry-path | no       | '/metrics'           | Path for your metrics         |

### Docker Compose 

```
version: '3.8'
services:
  huawei-lte-exporter:
    build:
      context: .
      dockerfile: huawei-lte-exporter/Dockerfile
    image: huawei-lte-exporter
    command: 
      - '--username=<username>'
      - '--password=<password>'
    restart: always
```

## Metrics 

| Metrics Name               | Metrics Description                                         |
| -------------------------- | ----------------------------------------------------------- |
| huawei_TotalUpload         | Total Upload since first boot or last counter reset         |
| huawei_TotalConnectTime    | Total Connected time since first boot or last counter reset |
| huawei_TotalDownload       | Total Download since first boot or last counter reset       |
| huawei_CurrentConnectTime  | Connect Time since last reboot                              |
| huawei_CurrentDownloadRate | Current download speed                                      |
| huawei_CurrentDownload     | Total Download since last connection established            |
| huawei_CurrentUpload       | Total Upload since last connection established              |
| huawei_CurrentUploadRate   | Current upload speed                                        |
| huawei_rsrp                | Signal RSRP                                                 |
| huawei_rssi                | Signal RSSI                                                 |
| huawei_rsrq                | Signal RSRQ                                                 |
| huawei_sinr                | Signal SINR                                                 |

## Todo 

* Add more metrics
* Publish image in DockerHub
* Publish grafana dashboard