# /*
# |    Protect your secrets, protect your sensitive data.
# :    Explore VMware Secrets Manager docs at https://vsecm.com/
# </
# <>/  keep your secrets... secret
# >/
# <>/' Copyright 2023-present VMware Secrets Manager contributors.
# >/'  SPDX-License-Identifier: BSD-2-Clause
# */

# Packages the "VSecM Sidecar" binary into a container image.
sidecar-bundle-ist:
	./hack/bundle.sh "vsecm-sidecar" \
		$(VERSION) "dockerfiles/vsecm/sidecar.Dockerfile"

# Packages the "VSecM Sidecar" binary into a container image for FIPS.
sidecar-bundle-ist-fips:
	./hack/bundle.sh "vsecm-fips-sidecar" \
		$(VERSION) "dockerfiles/vsecm-fips/sidecar.Dockerfile"

# Pushes the "VSecM Sidecar" container image to the public registry.
sidecar-push-ist:
	./hack/push.sh "vsecm-sidecar" \
		$(VERSION) "$(VSECM_DOCKERHUB_REGISTRY_URL)/vsecm-sidecar"

# Pushes the "VSecM Sidecar" (FIPS) container image to the public registry.
sidecar-push-ist-fips:
	./hack/push.sh "vsecm-fips-sidecar" \
		$(VERSION) "$(VSECM_DOCKERHUB_REGISTRY_URL)/vsecm-fips-sidecar"

# Pushes the "VSecM Sidecar" container image to the local registry.
sidecar-push-ist-local:
	./hack/push.sh "vsecm-sidecar" \
		$(VERSION) "$(VSECM_LOCAL_REGISTRY_URL)/vsecm-sidecar"

# Pushes the "VSecM Sidecar" (FIPS) container image to the local registry.
sidecar-push-ist-fips-local:
	./hack/push.sh "vsecm-fips-sidecar" \
		$(VERSION) "$(VSECM_LOCAL_REGISTRY_URL)/vsecm-fips-sidecar"
