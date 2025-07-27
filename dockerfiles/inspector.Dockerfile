# /*
# |    Protect your secrets, protect your sensitive data.
# :    Explore VMware Secrets Manager docs at https://vsecm.com/
# </
# <>/  keep your secrets... secret
# >/
# <>/' Copyright 2023-present VMware Secrets Manager contributors.
# >/'  SPDX-License-Identifier: BSD-2-Clause
# */

FROM --platform=$BUILDPLATFORM golang:1.24.5 AS builder
ARG BUILDPLATFORM
ARG TARGETPLATFORM
ARG TARGETOS
ARG TARGETARCH
ARG APPVERSION

ENV GOOS=$TARGETOS \
    GOARCH=$TARGETARCH \
    APPVERSION=$APPVERSION

WORKDIR /workspace

# Install cross-compilation tools
RUN apt-get update && apt-get install -y \
    gcc-x86-64-linux-gnu \
    g++-x86-64-linux-gnu \
    gcc-aarch64-linux-gnu \
    g++-aarch64-linux-gnu \
    libc6-dev-arm64-cross \
    libc6-dev-amd64-cross \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

# Download dependencies first (better layer caching)
COPY go.mod go.sum ./
RUN go mod download

# Copy the app source code
COPY . .

# Build the app for the target architecture
RUN echo "Building VSecM Inspector on $BUILDPLATFORM targeting $TARGETPLATFORM"
RUN ./hack/docker/buildx.sh ${TARGETARCH} inspector

# Use BusyBox as the base image
# FROM busybox:1.36-musl AS spike
FROM busybox:1.36 AS inspector
# Redefine the ARG in this stage to make it available
ARG APPVERSION

# Create necessary directories and users
RUN adduser -D -H -u 1000 inspector

# Copy the binary from builder
COPY --from=builder /workspace/inspector /usr/local/bin/inspector

# Change ownership to initcontainer user
RUN chown inspector:inspector /usr/local/bin/inspector

# Ensure the binary is executable
RUN chmod +x /usr/local/bin/inspector

# Apply labels to the final image
LABEL maintainers="VSecM Maintainers <maintainers@vsecm.com" \
      version="${APPVERSION}" \
      website="https://vsecm.com/" \
      repo="https://github.com/vmware/secrets-manager" \
      documentation="https://vsecm.com/" \
      contact="https://vsecm.com/docs/contact/" \
      community="https://vsecm.com/community/hello/" \
      changelog="https://vsecm.com/timeline/changelog/"

ENTRYPOINT ["/usr/local/bin/inspector"]
