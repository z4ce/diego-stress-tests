#!/usr/bin/env bash
OM_PASS=''
o() {
  om -k -t https://127.0.0.1 -u admin -p "${OM_PASS}" "$@"
}

o curl --path /api/v0/deployed/director/credentials/bbr_ssh_credentials | jq -r '.credential.value.private_key_pem' -r > /tmp/bbr_key.key
BBR_PW=$(o curl --path /api/v0/deployed/director/credentials/uaa_bbr_client_credentials | jq -r '.credential.value.password')
eval "$(o bosh-env)"
chmod 600 /tmp/bbr_key.key

bbr director \
--private-key-path /tmp/bbr_key.key \
--username bbr \
--host "$BOSH_ENVIRONMENT" \
pre-backup-check

export BOSH_DEPLOYMENT="$(bosh deployments --column=name -n | grep "^cf-" | xargs)"
if [[ "$BOSH_DEPLOYMENT" == "" ]]; then
  echo "Couldn't find bosh deployment"
  exit 1
fi
echo Found cf deployment "$BOSH_DEPLOYMENT"

echo bbr deployment \
--target "$BOSH_ENVIRONMENT" \
--username bbr_client \
--password "$BBR_PW" \
--deployment "$BOSH_DEPLOYMENT" \
pre-backup-check

echo begin backing up ops mgr
o export-installation -o /tmp/bbr_backup/om.zip

echo begin backing up director
bbr director \
--private-key-path /tmp/bbr_key.key \
--username bbr \
--host "$BOSH_ENVIRONMENT" \
backup

bbr deployment \
--target "$BOSH_ENVIRONMENT" \
--username bbr_client \
--password "$BBR_PW" \
--deployment "$BOSH_DEPLOYMENT" \
backup
