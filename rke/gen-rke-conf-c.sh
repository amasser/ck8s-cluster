#!/bin/sh

set -e

SCRIPTS_PATH="$(dirname "$(readlink -f "$0")")"
cd ${SCRIPTS_PATH}/../terraform/customer/

w1_ip=$(terraform output c-worker1-ip)
w2_ip=$(terraform output c-worker2-ip)
m_ip=$(terraform output c-master-ip)

cat <<EOF
cluster_name: rke-customer

# Change this path later
ssh_key_path: ../.ssh/id_rsa

nodes:
  - address: $m_ip
    user: rancher
    role: [controlplane,etcd]
  - address: $w1_ip
    user: rancher
    role: [worker]
  - address: $w2_ip
    user: rancher
    role: [worker]

services:
  kube-api:
    pod_security_policy: true
    # Add additional arguments to the kubernetes API server
    # This WILL OVERRIDE any existing defaults
    extra_args:
      # Enable audit log to stdout
      audit-log-path: "-"
      # Increase number of delete workers
      delete-collection-workers: 3
      # Set the level of log output to debug-level
      v: 4

  etcd:
    snapshot: true
    creation: 6h
    retention: 24h

ingress:
  provider: "nginx"
  extra_args:
    default-ssl-certificate: "ingress-nginx/ingress-default-cert"
    enable-ssl-passthrough: ""
EOF
