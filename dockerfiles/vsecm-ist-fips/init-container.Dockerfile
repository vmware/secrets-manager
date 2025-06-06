# /*
# |    Protect your secrets, protect your sensitive data.
# :    Explore VMware Secrets Manager docs at https://vsecm.com/
# </
# <>/  keep your secrets... secret
# >/
# <>/' Copyright 2023-present VMware Secrets Manager contributors.
# >/'  SPDX-License-Identifier: BSD-2-Clause
# */

# builder image
FROM golang:1.23.2-alpine3.20 AS builder

RUN mkdir /build
COPY app /build/app
COPY core /build/core
COPY lib /build/lib
COPY vendor /build/vendor
COPY go.mod /build/go.mod
WORKDIR /build

# GOEXPERIMENT=boringcrypto is required for FIPS compliance.
RUN CGO_ENABLED=0 GOEXPERIMENT=boringcrypto GOOS=linux go build -mod vendor -a -o vsecm-init-container \
  ./app/init_container/cmd/main.go

# generate clean, final image for end users
FROM gcr.io/distroless/static-debian11

ENV APP_VERSION="0.28.1"

LABEL "maintainers"="VSecM Maintainers <maintainers@vsecm.com>"
LABEL "version"=$APP_VERSION
LABEL "website"="https://vsecm.com/"
LABEL "repo"="https://github.com/vmware/secrets-manager"
LABEL "documentation"="https://vsecm.com/"
LABEL "contact"="https://vsecm.com/docs/contact"
LABEL "community"="https://vsecm.com/community/hello/"
LABEL "changelog"="https://vsecm.com/timeline/changelog/"

COPY --from=builder /build/vsecm-init-container .

# executable
ENTRYPOINT [ "./vsecm-init-container" ]
CMD [ "" ]
