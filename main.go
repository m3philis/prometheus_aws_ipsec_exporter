package main

import (
  "net/http"

  "github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
  // executing the metrics collector which runs in an endless loop
  ipsecMetrics()

  //This section will start the HTTP server and expose
  //any metrics on the /metrics endpoint.
  http.Handle("/metrics", promhttp.Handler())
  http.ListenAndServe(":9080", nil)
}
