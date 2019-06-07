package main

import (
	"github.com/prometheus/client_golang/prometheus"
)

//Define the metrics we wish to expose
var tunnelMetric = prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: "aws_ipsec_tunnel_status",
	Help: "display the state of found IPSec tunnel",
}, []string{"name"})

func init() {
	//Register metrics with prometheus
	prometheus.MustRegister(tunnelMetric)
}
