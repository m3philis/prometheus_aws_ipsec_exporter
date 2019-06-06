package main

import (
  "fmt"
  "time"

  "github.com/aws/aws-sdk-go/aws"
  "github.com/aws/aws-sdk-go/aws/session"
  "github.com/aws/aws-sdk-go/service/ec2"
  "github.com/prometheus/client_golang/prometheus"
)

func ipsecMetrics() {
  // create a session to AWS with a set region
  sess := session.Must(session.NewSession(&aws.Config{
    Region: aws.String("eu-west-1"),
  }))
  svc := ec2.New(sess)

  // inner function as go function to run endlessly but don't block the exporter itself
  go func() {
    for {
      result, err := svc.DescribeVpnConnections(nil)
      if err != nil {
        fmt.Println("Error", err)
        return
      }

      for _, connection := range result.VpnConnections {
        var name string
        for _, tags := range connection.Tags {
          if *tags.Key == "Name" {
            name = *tags.Value
          }
        }

        // Set state of primary tunnel
        if *connection.VgwTelemetry[0].Status == "UP" {
          tunnelMetric1.With(prometheus.Labels{"name": name}).Set(1)
        } else {
          tunnelMetric1.With(prometheus.Labels{"name": name}).Set(0)
        }

        // Set state of secondary tunnel
        if *connection.VgwTelemetry[1].Status == "UP" {
          tunnelMetric2.With(prometheus.Labels{"name": name}).Set(1)
        } else {
          tunnelMetric2.With(prometheus.Labels{"name": name}).Set(0)
        }
      }

      time.Sleep(10 * time.Second)

    }
  }()
}
