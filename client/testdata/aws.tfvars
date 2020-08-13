region                        = "us-west-1"
prefix_sc                     = ""
prefix_wc                     = ""
master_nodes_sc               = { master-0 = "t3.small" }
worker_nodes_sc               = { worker-0 = "t3.xlarge", worker-1 = "t3.large" }
master_nodes_wc               = { master-0 = "t3.small" }
worker_nodes_wc               = { worker-0 = "t3.large", worker-1 = "t3.large" }
public_ingress_cidr_whitelist = ["1.2.3.4/32", "4.3.2.1/32"]
api_server_whitelist          = ["1.2.3.4/32", "4.3.2.1/32"]
nodeport_whitelist            = ["1.2.3.4/32", "4.3.2.1/32"]
