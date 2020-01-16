SOURCE_PATH="$(dirname "$(readlink -f "$BASH_SOURCE")")"

export CONFIG_PATH="${CONFIG_PATH:-${SOURCE_PATH}/clusters/${CLOUD_PROVIDER}/${ENVIRONMENT_NAME}}"
export VAULT_ADDR=https://vault.eck.elastisys.se
export TF_VAR_dns_prefix=${ENVIRONMENT_NAME}
export TF_VAR_ssh_pub_key_file_sc=${CONFIG_PATH}/ssh-keys/id_rsa_sc.pub
export TF_VAR_ssh_pub_key_file_wc=${CONFIG_PATH}/ssh-keys/id_rsa_wc.pub

#
# Kubernetes
#

# Rancher kubernetes image to use.
export KUBERNETES_VERSION=${KUBERNETES_VERSION:-'"v1.15.6-rancher1-2"'}

#
# Service settings
#

# Influx cronjob backup variables.
export INFLUX_ADDR=influxdb.influxdb-prometheus.svc:8088
export INFLUX_BACKUP_SCHEDULE="0 0 * * *"

# Domains that should be allowed to log in using OAuth
export OAUTH_ALLOWED_DOMAINS="${OAUTH_ALLOWED_DOMAINS:-example.com}"

# Alerting variables
export ALERT_TO=${ALERT_TO:-null}
# Default URL is for sending to the #ck8s-ops channel
export SLACK_API_URL=${SLACK_API_URL:-https://hooks.slack.com/services/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}

# If unset -> true
export ENABLE_PSP=${ENABLE_PSP:-true}
export ENABLE_FALCO=${ENABLE_FALCO:-true}
export ENABLE_HARBOR=${ENABLE_HARBOR:-true}
export ENABLE_OPA=${ENABLE_OPA:-true}
export ENABLE_CUSTOMER_PROMETHEUS=${ENABLE_CUSTOMER_PROMETHEUS:-false}
export ENABLE_CUSTOMER_GRAFANA=${ENABLE_CUSTOMER_GRAFANA:-true}
export ENABLE_CUSTOMER_ALERTMANAGER=${ENABLE_CUSTOMER_ALERTMANAGER:-false}
export ENABLE_CUSTOMER_ALERTMANAGER_INGRESS=${ENABLE_CUSTOMER_ALERTMANAGER_INGRESS:-false}
export ENABLE_POSTGRESQL=${ENABLE_POSTGRESQL:-false}

export CUSTOMER_NAMESPACES=${CUSTOMER_NAMESPACES:-"demo"}
export CUSTOMER_ADMIN_USERS=${CUSTOMER_ADMIN_USERS:-"admin@example.com"}

#retention variables

export KUBEAUDIT_RETENTION_SIZE=${KUBEAUDIT_RETENTION_SIZE:-50}
export KUBEAUDIT_RETENTION_AGE=${KUBEAUDIT_RETENTION_AGE:-30}
export KUBECOMPONENTS_RETENTION_SIZE=${KUBECOMPONENTS_RETENTION_SIZE:-10}
export KUBECOMPONENTS_RETENTION_AGE=${KUBECOMPONENTS_RETENTION_SIZE:-10}
export KUBERNETES_RETENTION_SIZE=${KUBERNETES_RETENTION_SIZE:-50}
export KUBERNETES_RETENTION_AGE=${KUBERNETES_RETENTION_SIZE:-30}
export POSTGRESQL_RETENTION_SIZE=${POSTGRESQL_RETENTION_SIZE:-30}
export POSTGRESQL_RETENTION_AGE=${POSTGRESQL_RETENTION_AGE:-30}
export OTHER_RETENTION_SIZE=${OTHER_RETENTION_SIZE:-10}
export OTHER_RETENTION_AGE=${OTHER_RETENTION_SIZE:-30}
export ROLLOVER_SIZE=${ROLLOVER_SIZE:-1}
export ROLLOVER_AGE=${ROLLOVER_AGE:-1}

export INFLUXDB_RETENTION_WC=${INFLUXDB_RETENTION_WC:-7d}
export INFLUXDB_RETENTION_SC=${INFLUXDB_RETENTION_SC:-3d}
