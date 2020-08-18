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
# 2d1a9178-b9f0-4c01-b007-60c8ce65ec99 1 2gb 50gb
# 96c7903e-32f0-421d-b6a2-a45c97b15665 2 4gb 50gb
# 572a3b2e-6329-4053-b872-aecb1e70d8a6 4 8gb 50gb
# 73e99a76-a55c-402f-83e6-72dee465c675 2 8gb 50gb
# d430b3cd-0216-43ff-878c-c08689c0001b 4 16gb 50gb
variable worker_names_sc {
  description = "List of names for worker instances to create."
  type        = list(string)

  default = ["worker-0", "worker-1"]
}

variable worker_name_flavor_map_sc {
  description = "Map of instance name to openstack flavor."
  type        = map
  default = {
    "worker-0" : "d430b3cd-0216-43ff-878c-c08689c0001b",
    "worker-1" : "572a3b2e-6329-4053-b872-aecb1e70d8a6"
  }
}

variable "worker_anti_affinity_policy_sc" {
  description = "This can be set to 'anti-affinity' or 'soft-anti-affinity' to spread out workers on different physical machines, otherwise leave it empty"
  type        = string
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
    "worker-0" : "572a3b2e-6329-4053-b872-aecb1e70d8a6",
    "worker-1" : "572a3b2e-6329-4053-b872-aecb1e70d8a6"
  }
}

variable "worker_anti_affinity_policy_wc" {
  description = "This can be set to 'anti-affinity' or 'soft-anti-affinity' to spread out workers on different physical machines, otherwise leave it empty"
  type        = string
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
    "master-0" : "96c7903e-32f0-421d-b6a2-a45c97b15665"
  }
}

variable "master_anti_affinity_policy_sc" {
  description = "This can be set to 'anti-affinity' or 'soft-anti-affinity' to spread out masters on different physical machines, otherwise leave it empty"
  type        = string
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
    "master-0" : "96c7903e-32f0-421d-b6a2-a45c97b15665"
  }
}

variable "master_anti_affinity_policy_wc" {
  description = "This can be set to 'anti-affinity' or 'soft-anti-affinity' to spread out masters on different physical machines, otherwise leave it empty"
  type        = string
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
