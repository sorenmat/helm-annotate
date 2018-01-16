#!/usr/bin/env bash

set -euf -o pipefail

HELM_ANNOTATE_VERSION=${HELM_ANNOTATE_VERSION:-"0.2.0"}

file="${HELM_PLUGIN_DIR:-"$(helm home)/plugins/helm-annotate"}/helm-annotate"
os=$(uname -s | tr '[:upper:]' '[:lower:]')
url="https://github.com/Tradeshift/helm-annotate/releases/download/v${HELM_ANNOTATE_VERSION}/helm-annotate_${os}_amd64"

if command -v wget; then
  wget -O "${file}"  "${url}"
elif command -v curl; then
  curl -o "${file}" "${url}"
fi

chmod +x "${file}"
