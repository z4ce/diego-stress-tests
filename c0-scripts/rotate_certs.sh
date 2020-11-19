#!/usr/bin/env bash
OM_PASS='xx'
o() {
  om -k -t https://127.0.0.1 -u admin -p "${OM_PASS}" "$@"
}

echo creating new cert authority
o generate-certificate-authority
o certificate-authorities -f json | jq '.[]|select(.active==false).guid' -r > /tmp/inactive_cert
if [[ "$(wc -l /tmp/inactive_cert)" != "1 /tmp/inactive_cert" ]]; then
  echo "Already a cert in rotation. Delete it first"
  exit 1
fi

o apply-changes --recreate-vms

echo Activating Certificate Authority "$(cat /tmp/inactive_cert)"

o activate-certificate-authority --id "$(cat /tmp/inactive_cert)"
o regenerate-certificates
o apply-changes --recreate-vms
# If you want to also delete the old CA
# o delete-certiciate-authorities <old guid>
# o apply-changes -recreate-vms
