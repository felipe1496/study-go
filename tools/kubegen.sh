#!/usr/bin/env bash
set -euo pipefail

DEEPCOPY_GEN="$1"
BOILERPLATE="$2"
PKG="$3"

"$DEEPCOPY_GEN" \
  --go-header-file="$BOILERPLATE" \
  --output-file="zz_generated.deepcopy.go" \
  "$PKG"