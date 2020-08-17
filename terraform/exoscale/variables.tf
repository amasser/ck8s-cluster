# Exoscale credentials.
variable exoscale_api_key {
  description = "Either use .cloudstack.ini or this to set the API key."
  type        = string
}

variable exoscale_secret_key {
  description = "Either use .cloudstack.ini or this to set the API secret."
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

variable "compute_instance_image" {
  default = "CK8S BaseOS v0.0.6"
}

variable machines_sc {
  description = "Service cluster machines"
  type = map(object({
    node_type                 = string
    size                      = string
    es_local_storage_capacity = string
  }))
}

variable machines_wc {
  description = "Workload cluster machines"
  type = map(object({
    node_type                 = string
    size                      = string
    es_local_storage_capacity = string
  }))
}

variable nfs_size {
  description = "The size of the nfs machine"
  type        = string
  default     = "Small"
}

variable ssh_pub_key_sc {
  description = "Path to public SSH key file which is injected into the VMs."
  type        = string
}

variable ssh_pub_key_wc {
  description = "Path to public SSH key file which is injected into the VMs."
  type        = string
}

variable public_ingress_cidr_whitelist {
  type = list(string)
}

variable api_server_whitelist {
  type = list(string)
}

variable nodeport_whitelist {
  type = list(string)
}

variable dns_prefix {
  description = "Prefix name for dns"
  type        = string
}
