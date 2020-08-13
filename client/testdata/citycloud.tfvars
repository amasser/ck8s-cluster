prefix_sc                      = ""
prefix_wc                      = ""
master_names_sc                = ["master-0"]
master_name_flavor_map_sc      = { master-0 = "dc67a9eb-0685-4bb6-9383-a01c717e02e8" }
worker_names_sc                = ["worker-0", "worker-1"]
worker_name_flavor_map_sc      = { worker-0 = "ea0dbe3b-f93a-47e0-84e4-b09ec5873bdf", worker-1 = "dc67a9eb-0685-4bb6-9383-a01c717e02e8" }
master_names_wc                = ["master-0"]
master_name_flavor_map_wc      = { master-0 = "dc67a9eb-0685-4bb6-9383-a01c717e02e8" }
worker_names_wc                = ["worker-0"]
worker_name_flavor_map_wc      = { worker-0 = "dc67a9eb-0685-4bb6-9383-a01c717e02e8" }
master_anti_affinity_policy_sc = "anti-affinity"
worker_anti_affinity_policy_sc = "soft-anti-affinity"
master_anti_affinity_policy_wc = "anti-affinity"
worker_anti_affinity_policy_wc = "soft-anti-affinity"
public_ingress_cidr_whitelist  = ["1.2.3.4/32", "4.3.2.1/32"]
api_server_whitelist           = ["1.2.3.4/32", "4.3.2.1/32"]
nodeport_whitelist             = ["1.2.3.4/32", "4.3.2.1/32"]
aws_dns_zone_id                = "testAWSDNSZoneID"
aws_dns_role_arn               = "testAWSDNSRoleARN"
external_network_id            = "2aec7a99-3783-4e2a-bd2b-bbe4fef97d1c"
external_network_name          = "ext-net"
