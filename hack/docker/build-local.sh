#!/usr/bin/env bash

# /*
# |    Protect your secrets, protect your sensitive data.
# :    Explore VMware Secrets Manager docs at https://vsecm.com/
# </
# <>/  keep your secrets... secret
# >/
# <>/' Copyright 2023-present VMware Secrets Manager contributors.
# >/'  SPDX-License-Identifier: BSD-2-Clause
# */

# Builds VSecM components and pushes them to the local registry
# for development.

set -e

export VSECM_VERSION="dev"
COMPONENTS=("initcontainer" "inspector" "safe" "sentinel" "sidecar")
PLATFORMS="linux/amd64,linux/arm64"

# Set up buildx with docker-container driver if not already set up
setup_buildx() {
  # Check if we have a builder that supports multi-platform
  if ! docker buildx inspect multiplatform-builder &>/dev/null; then
    echo "Creating a new buildx builder for multi-platform builds..."
    docker buildx create --name multiplatform-builder \
    --driver docker-container --use
  else
    # Use the existing builder
    docker buildx use multiplatform-builder
  fi

  # Verify the builder is working
  docker buildx inspect --bootstrap
}

# Function to build a component
build_component() {
  local component=$1
  local version=$2
  local platforms=$3
  local output_flag=$4  # --load for single platform, --push for registry

  echo "Building vsecm-${component}:${version}..."

  # If building for multiple platforms without pushing, we can only use
  # the --output=type=image flag
  # If building for a single platform, we can use --load to import
  # to Docker daemon
  docker buildx build \
    --platform "${platforms}" \
    --build-arg APPVERSION="${version}" \
    -t "vsecm-${component}:${version}" \
    "${output_flag}" \
    -f "./dockerfiles/${component}.Dockerfile" .

  echo "Finished building spike-${component}:${version}"
}

# For local only, we can only build for the current architecture with --load
PLATFORMS=$(docker info --format '{{.Architecture}}' | sed 's/^/linux\//')
OUTPUT_FLAG="--load"
echo "Building only for local platform: $PLATFORMS"

# Set up buildx
setup_buildx

# Build all components
for component in "${COMPONENTS[@]}"; do
  build_component "${component}" "${VSECM_VERSION}" "${PLATFORMS}" "${OUTPUT_FLAG}"
done

echo "All components built successfully!"
