prefix_sc                        = ""
prefix_wc                        = ""
master_names_sc                  = ["master-0"]
master_name_size_map_sc          = { master-0 = "Small" }
worker_names_sc                  = ["worker-0", "worker-1"]
worker_name_size_map_sc          = { worker-0 = "Extra-large", worker-1 = "Large" }
es_local_storage_capacity_map_sc = { master-0 = 0, worker-0 = 1, worker-1 = 2 }
master_names_wc                  = ["master-0"]
master_name_size_map_wc          = { master-0 = "Small" }
worker_names_wc                  = ["worker-1"]
worker_name_size_map_wc          = { worker-1 = "Large" }
es_local_storage_capacity_map_wc = { master-0 = 0, worker-1 = 1 }
nfs_size                         = "Small"
public_ingress_cidr_whitelist    = ["1.2.3.4/32", "4.3.2.1/32"]
api_server_whitelist             = ["1.2.3.4/32", "4.3.2.1/32"]
nodeport_whitelist               = ["1.2.3.4/32", "4.3.2.1/32"]
