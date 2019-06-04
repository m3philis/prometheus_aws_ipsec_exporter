FROM golang:alpine as base

ENV GO111MODULE on

WORKDIR /go/src/github.com/m3philis/prometheus_aws-ipsec_exporter

COPY . .

RUN go install

FROM golang:alpine as runner

COPY --from=base /go/bin/prometheus_aws-ipsec_exporter /go/bin/ipsec_exporter

ENTRYPOINT ["/go/bin/ipsec_exporter"]
