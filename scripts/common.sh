SCRIPTS_PATH="$(dirname "$(readlink -f "$BASH_SOURCE")")"
pushd "${SCRIPTS_PATH}/../" > /dev/null

export ECK_CUSTOMER_DOMAIN=$(cat hosts.json | jq -r '.customer_dns_name.value' | sed 's/[^.]*[.]//')
export ECK_SYSTEM_DOMAIN=$(cat hosts.json | jq -r '.system_services_dns_name.value' | sed 's/[^.]*[.]//')

popd > /dev/null