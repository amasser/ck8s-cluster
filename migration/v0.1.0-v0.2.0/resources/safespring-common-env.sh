export OS_IDENTITY_API_VERSION=3
export OS_AUTH_URL=https://keystone.api.cloud.ipnett.se/v3
export OS_PROJECT_DOMAIN_NAME=${OS_PROJECT_DOMAIN_NAME:-elastisys.se}
export OS_USER_DOMAIN_NAME=${OS_USER_DOMAIN_NAME:-elastisys.se}
export OS_PROJECT_NAME=${OS_PROJECT_NAME:-infra.elastisys.se}
export OS_REGION_NAME=se-east-1
export OS_PROJECT_ID=${OS_PROJECT_ID:-9f91e56185fb4f929c36430ac4bcbe6e}
export S3_REGION=${S3_REGION:-sto1}
export S3_REGION_ADDRESS=${S3_REGION_ADDRESS:-s3.sto1.safedc.net}
export S3_REGION_ENDPOINT=https://$S3_REGION_ADDRESS

export ECK_OPS_DOMAIN=ops.${ENVIRONMENT_NAME}.elastisys.se
export ECK_BASE_DOMAIN=${ENVIRONMENT_NAME}.elastisys.se