package main

import (
  "fmt"
  "time"
  "strconv"
  "strings"

  "github.com/aws/aws-sdk-go/aws"
  "github.com/aws/aws-sdk-go/aws/session"
  "github.com/aws/aws-sdk-go/service/ec2"
  "github.com/aws/aws-sdk-go/service/cloudformation"
  "github.com/prometheus/client_golang/prometheus"
)

func ipsecMetrics() {
  // create a session to AWS with a set region
  sess := session.Must(session.NewSession(&aws.Config{
    Region: aws.String("eu-west-1"),
  }))

  // create client object to access cloudformation
  svcCfn := cloudformation.New(sess)
  stacks, err := svcCfn.DescribeStacks(nil)
  if err != nil {
    fmt.Println(err.Error())
  }

  // loop over all stacks to find the base-region stack and get the AccountName
  var accountName string
  for _, stack := range stacks.Stacks {
    if strings.HasPrefix(*stack.StackName, "base-region") {
      for _, tag := range stack.Parameters {
        if *tag.ParameterKey == "AccountName" {
          accountName = *tag.ParameterKey
        }
      }
    }
  }

  svcEc2 := ec2.New(sess)
  // inner function as go function to run endlessly but don't block the exporter itself
  go func() {
    for {
      result, err := svcEc2.DescribeVpnConnections(nil)
      if err != nil {
        fmt.Println(err.Error())
        return
      }

      // loop over all tunnel to get metrics
      for _, connection := range result.VpnConnections {
        // loop over tags to find the name of the tunnel
        var name string
        for _, tag := range connection.Tags {
          if *tag.Key == "Name" {
            name = *tag.Value
          }
        }

        // each tunnel is really two tunnels so we need to check if both tunnels are working
        for id, tunnel := range connection.VgwTelemetry {
          if *tunnel.Status == "UP" {
            tunnelMetric.With(prometheus.Labels{"name": name, "id": strconv.Itoa(id+1), "account": accountName}).Set(1)
          } else {
            tunnelMetric.With(prometheus.Labels{"name": name, "id": strconv.Itoa(id+1), "account": accountName}).Set(0)
          }
        }
      }

      // sleep for 10s and restart the loop
      time.Sleep(10 * time.Second)

    }
  }()
}
