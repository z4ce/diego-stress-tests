#!/usr/bin/env bash
OM_PASS=''
o() {
  om -k -t https://127.0.0.1 -u admin -p "${OM_PASS}" "$@"
}

echo Applying changes
o apply-changes
