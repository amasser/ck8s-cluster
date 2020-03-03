output "sc_master_ips" {
  value = module.service_cluster.master_ips
}

output "wc_master_ips" {
  value = module.workload_cluster.master_ips
}

output "sc_worker_ips" {
  value = module.service_cluster.worker_ips
}

output "wc_worker_ips" {
  value = module.workload_cluster.worker_ips
}

output "sc_master_external_loadbalancer_fqdn" {
  value = module.service_cluster.master_external_loadbalancer_fqdn
}

output "sc_master_internal_loadbalancer_fqdn" {
  value = module.service_cluster.master_internal_loadbalancer_fqdn
}

output "wc_master_external_loadbalancer_fqdn" {
  value = module.workload_cluster.master_external_loadbalancer_fqdn
}

output "wc_master_internal_loadbalancer_fqdn" {
  value = module.workload_cluster.master_internal_loadbalancer_fqdn
}

output "wc_dns_name" {
  value = module.workload_dns.dns_record_name
}

output "sc_dns_name" {
  value = module.service_dns.dns_record_name
}

output "domain_name" {
  value = "${var.dns_prefix}.${module.service_dns.dns_suffix}"
}

output "ansible_inventory" {
  value = data.template_file.ansible_inventory.rendered
}
