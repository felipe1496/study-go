#!/usr/bin/env bash
set -euo pipefail
set -x

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/../.." && pwd -P)"
export GOMOD="${ROOT}/go.mod"
export GO111MODULE=on

# Tudo que o Go precisa, primeiro:
export GOBIN="${PWD}/gobin"
export GOCACHE="${PWD}/gocache"
export GOMODCACHE="${PWD}/gomodcache"
mkdir -p "${GOBIN}" "${GOCACHE}" "${GOMODCACHE}"

export KUBE_VERBOSE=9

echo "ROOT=$ROOT"
go env GOMOD GOPATH GOMODCACHE GOCACHE GOBIN

echo "go list in v1alpha1:"
( cd "${ROOT}/pkg/apis/sample/v1alpha1" && go list -find . )

source "${ROOT}/tools/k8s-codegen/kube_codegen.sh"

kube::codegen::gen_helpers \
  --boilerplate "${ROOT}/tools/k8s-codegen/hack/boilerplate.go.txt" \
  "${ROOT}/pkg/apis"

kube::codegen::gen_client \
  --boilerplate "${ROOT}/tools/k8s-codegen/hack/boilerplate.go.txt" \
  --output-dir "${ROOT}/pkg/generated" \
  --output-pkg "github.com/felipe1496/hello_go/pkg/generated" \
  "${ROOT}/pkg/apis"