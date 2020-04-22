terraform {
  backend "remote" {}
}

locals {
  # Base image used to provision master and worker instances
  cluster_image = "CK8S-BaseOS-v0.0.6"
  # Base image used to provision loadbalancer instances
  loadbalancer_image = "ubuntu-18.04-server-cloudimg-amd64-20190212.1"
}

data "openstack_images_image_v2" "cluster_image" {
  name        = local.cluster_image
  most_recent = true
}

data "openstack_images_image_v2" "loadbalancer_image" {
  name        = local.loadbalancer_image
  most_recent = true
}

module "service_cluster" {
  source = "./modules/cluster"

  prefix = var.prefix_sc == "" ? "${terraform.workspace}-service-cluster" : var.prefix_sc

  ssh_pub_key = var.ssh_pub_key_sc

  cluster_image      = data.openstack_images_image_v2.cluster_image.id
  loadbalancer_image = data.openstack_images_image_v2.loadbalancer_image.id

  public_ingress_cidr_whitelist = var.public_ingress_cidr_whitelist

  external_network_id   = var.external_network_id
  external_network_name = var.external_network_name

  master_names           = var.master_names_sc
  master_name_flavor_map = var.master_name_flavor_map_sc

  worker_names           = var.worker_names_sc
  worker_name_flavor_map = var.worker_name_flavor_map_sc

  loadbalancer_names           = var.loadbalancer_names_sc
  loadbalancer_name_flavor_map = var.loadbalancer_name_flavor_map_sc

  dns_prefix = var.dns_prefix
  dns_list = [
    "*.ops",
    "grafana",
    "harbor",
    "dex",
    "kibana",
    "notary.harbor"
  ]
  aws_dns_zone_id  = var.aws_dns_zone_id
  aws_dns_role_arn = var.aws_dns_role_arn
}

module "workload_cluster" {
  source = "./modules/cluster"

  prefix = var.prefix_wc == "" ? "${terraform.workspace}-workload-cluster" : var.prefix_wc

  ssh_pub_key = var.ssh_pub_key_wc

  cluster_image      = data.openstack_images_image_v2.cluster_image.id
  loadbalancer_image = data.openstack_images_image_v2.loadbalancer_image.id

  public_ingress_cidr_whitelist = var.public_ingress_cidr_whitelist

  external_network_id   = var.external_network_id
  external_network_name = var.external_network_name

  master_names           = var.master_names_wc
  master_name_flavor_map = var.master_name_flavor_map_wc

  worker_names           = var.worker_names_wc
  worker_name_flavor_map = var.worker_name_flavor_map_wc

  loadbalancer_names           = var.loadbalancer_names_wc
  loadbalancer_name_flavor_map = var.loadbalancer_name_flavor_map_wc

  dns_prefix = var.dns_prefix
  dns_list = [
    "*",
    "prometheus.ops"
  ]
  aws_dns_zone_id  = var.aws_dns_zone_id
  aws_dns_role_arn = var.aws_dns_role_arn
}
