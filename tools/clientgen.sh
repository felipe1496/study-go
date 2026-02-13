#!/usr/bin/env bash
set -euo pipefail

CLIENT_GEN="$1"
BOILERPLATE_REL="$2"
MODULE="$3"
OUTPUT_DIR_REL="$4"
OUTPUT_PKG="$5"

# Bazel define isso em `bazel run`. É o path real do seu workspace.
if [[ -z "${BUILD_WORKSPACE_DIRECTORY:-}" ]]; then
  echo "ERROR: BUILD_WORKSPACE_DIRECTORY not set. Run via 'bazel run'." >&2
  exit 1
fi

BOILERPLATE_ABS="${BUILD_WORKSPACE_DIRECTORY}/${BOILERPLATE_REL}"
OUTPUT_DIR_ABS="${BUILD_WORKSPACE_DIRECTORY}/${OUTPUT_DIR_REL}"

mkdir -p "$OUTPUT_DIR_ABS"

"$CLIENT_GEN" \
  --go-header-file="$BOILERPLATE_ABS" \
  --input-base="${MODULE}/pkg/apis" \
  --input="sample/v1alpha1" \
  --output-dir="$OUTPUT_DIR_ABS" \
  --output-pkg="${MODULE}/${OUTPUT_PKG}" \
  --clientset-name="versioned"