#!/bin/bash

set -e

: "${ECK_SC_DOMAIN:?Missing ECK_SC_DOMAIN}"
: "${ENABLE_PSP:?Missing ENABLE_PSP}"

if [[ "$#" -ne 1 ]]
then 
  echo "Usage: gen-rke-conf-wc.sh <path-to-infra-file>"
  exit 1
fi

infra="$1"

master_ip_addresses=($(cat $infra | jq -r '.service_cluster.master_ip_addresses[]'))
worker_ip_addresses=($(cat $infra | jq -r '.service_cluster.worker_ip_addresses[]'))

cat <<EOF
cluster_name: eck-service_cluster

ssh_agent_auth: true

nodes:
EOF

for i in $(seq 0 $((${#master_ip_addresses[@]} - 1)))
do
cat <<EOF
  - address: ${master_ip_addresses[${i}]}
    user: rancher
    role: [controlplane,etcd]
EOF
done

for i in $(seq 0 $((${#worker_ip_addresses[@]} - 1)))
do
cat <<EOF
  - address: ${worker_ip_addresses[${i}]}
    user: rancher
    role: [worker]
EOF
done

cat <<EOF

services:
  kube-api:
EOF

if [[ $ENABLE_PSP == "true" ]]
then
cat <<EOF
    pod_security_policy: true
EOF
else
cat <<EOF
    pod_security_policy: false
EOF
fi

cat <<EOF
    # Add additional arguments to the kubernetes API server
    # This WILL OVERRIDE any existing defaults
    extra_binds:
    # Adds file from node into docker container running api-server
      - "/etc/kubernetes/conf:/etc/kubernetes/conf"
      - "/var/log/kube-audit:/var/log/kube-audit"
    extra_args:
      oidc-issuer-url: https://dex.${ECK_SC_DOMAIN}
      oidc-client-id: kubernetes
      oidc-username-claim: email
      oidc-groups-claim: groups
      audit-policy-file: "/etc/kubernetes/conf/audit-policy.yaml"
      audit-log-path: "/var/log/kube-audit/kube-apiserver.log"
      # Increase number of delete workers
      delete-collection-workers: 3
      # Set the level of log output to debug-level
      v: 4
EOF

if [[ $ENABLE_PSP == "true" ]]
then
cat <<EOF
      enable-admission-plugins: "NamespaceLifecycle,LimitRanger,ServiceAccount,DefaultStorageClass,DefaultTolerationSeconds,MutatingAdmissionWebhook,ValidatingAdmissionWebhook,ResourceQuota,NodeRestriction,PodSecurityPolicy"
EOF
else
cat <<EOF
      enable-admission-plugins: "NamespaceLifecycle,LimitRanger,ServiceAccount,DefaultStorageClass,DefaultTolerationSeconds,MutatingAdmissionWebhook,ValidatingAdmissionWebhook,ResourceQuota,NodeRestriction"
EOF
fi

cat <<EOF

  etcd:
    snapshot: true
    creation: 6h
    retention: 24h

ingress:
  provider: "nginx"
  extra_args:
    enable-ssl-passthrough: ""
EOF
