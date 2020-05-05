module "network" {
  source = "../../../modules/openstack/network"

  prefix = var.prefix

  external_network_id = var.external_network_id
}

module "secgroups" {
  source = "../../../modules/openstack/secgroups"

  prefix = var.prefix

  public_ingress_cidr_whitelist = var.public_ingress_cidr_whitelist
  api_server_whitelist          = var.api_server_whitelist
}

resource "openstack_compute_keypair_v2" "sshkey" {
  name       = "${var.prefix}-ssh-key"
  public_key = file(pathexpand(var.ssh_pub_key))
}

module "master" {
  source = "../../../modules/openstack/vm"

  prefix          = var.prefix
  names           = var.master_names
  name_flavor_map = var.master_name_flavor_map
  image_id        = var.cluster_image
  key_pair        = openstack_compute_keypair_v2.sshkey.id

  external_network_name = var.external_network_name

  network_id = module.network.network_id
  subnet_id  = module.network.subnet_id

  security_group_ids = [
    module.secgroups.cluster_secgroup,
    module.secgroups.master_secgroup,
  ]
}

module "worker" {
  source = "../../../modules/openstack/vm"

  prefix          = var.prefix
  names           = var.worker_names
  name_flavor_map = var.worker_name_flavor_map
  image_id        = var.cluster_image
  key_pair        = openstack_compute_keypair_v2.sshkey.id

  external_network_name = var.external_network_name

  network_id = module.network.network_id
  subnet_id  = module.network.subnet_id

  security_group_ids = [
    module.secgroups.cluster_secgroup,
    module.secgroups.worker_secgroup,
  ]
}

module "octavia_lb" {
  source = "../../../modules/openstack/octavia-lb"

  external_network_name = var.external_network_name
  prefix                = var.prefix

  loadbalancer_targets = {
    http = {
      port               = 80
      protocol           = "HTTP"
      target_ips         = module.worker.instance_ips
      health_path        = "/healthz"
      health_codes       = "200"
      health_delay       = 20
      health_timeout     = 10
      health_max_retries = 5
    }
    https = {
      port               = 443
      protocol           = "HTTPS"
      target_ips         = module.worker.instance_ips
      health_path        = "/healthz"
      health_codes       = "200"
      health_delay       = 20
      health_timeout     = 10
      health_max_retries = 5
    }
    kube_api = {
      port               = 6443
      protocol           = "TCP"
      target_ips         = module.master.instance_ips
      health_path        = "ignore"
      health_codes       = "ignore"
      health_delay       = 20
      health_timeout     = 10
      health_max_retries = 5
    }
  }

  subnet_id = module.network.subnet_id
}

module "dns" {
  source = "../../../modules/openstack/aws-dns"

  dns_list   = var.dns_list
  dns_prefix = var.dns_prefix

  aws_dns_zone_id  = var.aws_dns_zone_id
  aws_dns_role_arn = var.aws_dns_role_arn

  record_ips = module.octavia_lb.floating_ips
}
