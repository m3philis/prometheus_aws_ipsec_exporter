# Prometheus exporter for AWS IPSec tunnels

This prometheus exporter checks the state of IPSec tunnels and provides the metrics on port 9080


## Installation

Either build the binary with go build (or go install) or build a docker image with the Dockerfile.


## Usage

Currently this exporter needs to run inside an AWS VPC with an IPSec tunnel running.
The instance the exporter runs on needs `ec2:DescribeVpnConnections` as an IAM policy attached to access the VPN data from the cli.
