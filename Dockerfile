FROM scratch
COPY huawei-lte-exporter-linux-amd64 /go/huawei-lte-exporter
ENTRYPOINT ["/go/huawei-lte-exporter"] 


