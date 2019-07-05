cat <<EOF > ../rke/cluster.yaml

cluster_name: rancher-rke-test

# Maybe change this path later.
ssh_key_path: ../.ssh/id_rsa

nodes:
  - address: $(terraform output master-ip)
    user: rancher
    role: [controlplane,etcd]
  - address: $(terraform output worker1-ip)
    user: rancher
    role: [worker]
  - address: $(terraform output worker2-ip)
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
EOF