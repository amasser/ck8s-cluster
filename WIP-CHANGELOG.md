### Breaking changes

- The old Terraform variables configuration (config.tfvars) is no longer
  supported and needs to be updated to the new format (tfvars.json). However,
  once the configuration has been updated it can be applied without any change.
  See: https://github.com/elastisys/ck8s-cluster/pull/61

### Release notes

### Fixed

- Bug where tfe provider do not read the configured value in backend_config.hcl
- Bug where folders where not created before uploading crds

### Changed

- Removed some default values to make this ready for open sourcing
- tfvars machines configuration so that name, node type, size and more are all
  grouped into a single object.
- tfvars configuration language from HCL to JSON.
- all commands that previously required node type as an argument no longer do,
  only node name is now needed (e.g. `drain master foo` -> `drain foo`).
- Citycloud: Api server white list is applied to api server loadbalancer now as well as the VMs.
- Increases timeout for api server loadbalancer to 10m
