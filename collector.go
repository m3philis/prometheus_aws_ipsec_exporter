package main

import (
	"github.com/prometheus/client_golang/prometheus"
)

//Define the metrics we wish to expose
var tunnelMetric1 = prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: "aws_ipsec_tunnel1_status",
	Help: "display the state of the first IPSec tunnel",
}, []string{"name"})

var tunnelMetric2 = prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: "aws_ipsec_tunnel2_status",
	Help: "display the state of the second IPSec tunnel",
}, []string{"name"})

func init() {
	//Register metrics with prometheus
	prometheus.MustRegister(tunnelMetric1)
	prometheus.MustRegister(tunnelMetric2)
}
