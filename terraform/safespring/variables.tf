variable ssh_pub_key_sc {
  description = "Path to public SSH key file which is injected into the VMs."
  type        = string
}

variable ssh_pub_key_wc {
  description = "Path to public SSH key file which is injected into the VMs."
  type        = string
}

variable dns_prefix {
  description = "Prefix name for dns"
  type        = string
}

variable aws_dns_zone_id {
  description = "Id for the AWS DNS zone"
  type        = string
}

variable aws_dns_role_arn {
  description = "AWS role to asume while creating DNS entries"
  type        = string
}

variable prefix_sc {
  description = "Prefix for resource names"
  default     = ""
}

variable prefix_wc {
  description = "Prefix for resource names"
  default     = ""
}

# For workers
# Common flavors
# b.small  : 1493be98-d150-4f69-8154-4d59ea49681c
# b.medium : 9d82d1ee-ca29-4928-a868-d56e224b92a1
# b.large  : 16d11558-62fe-4bce-b8de-f49a077dc881
# m.medium : 2c1708d1-3974-4ab8-97cc-cbf58aa27ad9
# b.xlarge : fce2b54d-c0ef-4ad4-aa81-bcdcaa54f7cb
# AMD flavors, preferred!
# lb.tiny     : 51d480b8-2517-4ba8-bfe0-c649ac93eb61
# lb.large.1d : dc67a9eb-0685-4bb6-9383-a01c717e02e8
variable worker_names_sc {
  description = "List of names for worker instances to create."
  type        = list(string)

  default = ["worker-0", "worker-1"]
}

variable worker_name_flavor_map_sc {
  description = "Map of instance name to openstack flavor."
  type        = map
  default = {
    "worker-0" : "fce2b54d-c0ef-4ad4-aa81-bcdcaa54f7cb",
    "worker-1" : "16d11558-62fe-4bce-b8de-f49a077dc881"
  }
}

variable "worker_anti_affinity_policy_sc" {
  description = "This can be set to 'anti-affinity' to spread out workers on different physical machines, otherwise leave it empty"
  type = string
}

variable worker_names_wc {
  description = "List of names for worker instances to create."
  type        = list(string)
  default     = ["worker-0", "worker-1"]
}

variable worker_name_flavor_map_wc {
  description = "Map of instance name to openstack flavor."
  type        = map
  default = {
    "worker-0" : "16d11558-62fe-4bce-b8de-f49a077dc881",
    "worker-1" : "16d11558-62fe-4bce-b8de-f49a077dc881"
  }
}

variable "worker_anti_affinity_policy_wc" {
  description = "This can be set to 'anti-affinity' to spread out workers on different physical machines, otherwise leave it empty"
  type = string
}

# For masters
variable master_names_sc {
  description = "List of names for master instances to create."
  type        = list(string)
  default     = ["master-0"]
}

variable master_name_flavor_map_sc {
  description = "Map of instance name to openstack flavor."
  type        = map
  default = {
    "master-0" : "9d82d1ee-ca29-4928-a868-d56e224b92a1"
  }
}

variable "master_anti_affinity_policy_sc" {
  description = "This can be set to anti-affinity to spread out masters on different physical machines, otherwise leave it empty"
  type = string
}

variable master_names_wc {
  description = "List of names for master instances to create."
  type        = list(string)
  default     = ["master-0"]
}

variable master_name_flavor_map_wc {
  description = "Map of instance name to openstack flavor."
  type        = map
  default = {
    "master-0" : "9d82d1ee-ca29-4928-a868-d56e224b92a1"
  }
}

variable "master_anti_affinity_policy_wc" {
  description = "This can be set to anti-affinity to spread out masters on different physical machines, otherwise leave it empty"
  type = string
}

variable loadbalancer_names_sc {
  description = "List of names for loadbalancer instances to create."
  type        = list(string)
  default     = ["loadbalancer-0"]
}

variable loadbalancer_name_flavor_map_sc {
  description = "Map of instance name to openstack flavor."
  type        = map
  default = {
    "loadbalancer-0" : "51d480b8-2517-4ba8-bfe0-c649ac93eb61"
  }
}

variable loadbalancer_names_wc {
  description = "List of names for loadbalancer instances to create."
  type        = list(string)
  default     = ["loadbalancer-0"]
}

variable loadbalancer_name_flavor_map_wc {
  description = "Map of instance name to openstack flavor."
  type        = map
  default = {
    "loadbalancer-0" : "51d480b8-2517-4ba8-bfe0-c649ac93eb61"
  }
}

variable public_ingress_cidr_whitelist {
  type = list
}

variable api_server_whitelist {
  type = list
}

variable nodeport_whitelist {
  type = list
}

variable external_network_id {
  description = "the id of the external network"
  type        = string
}

variable external_network_name {
  description = "the name of the external network"
  type        = string
}
