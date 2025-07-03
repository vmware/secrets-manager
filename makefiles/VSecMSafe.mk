# /*
# |    Protect your secrets, protect your sensitive data.
# :    Explore VMware Secrets Manager docs at https://vsecm.com/
# </
# <>/  keep your secrets... secret
# >/
# <>/' Copyright 2023-present VMware Secrets Manager contributors.
# >/'  SPDX-License-Identifier: BSD-2-Clause
# */

# Packages the "VSecM Safe" into a container image.
safe-bundle-ist:
	./hack/bundle.sh "vsecm-safe" \
		$(VERSION) "dockerfiles/vsecm/safe.Dockerfile"

# Packages the "VSecM Safe" into a container image for FIPS.
safe-bundle-ist-fips:
	./hack/bundle.sh "vsecm-fips-safe" \
		$(VERSION) "dockerfiles/vsecm-fips/safe.Dockerfile"

# Pushes the "VSecM Safe" container to the public registry.
safe-push-ist:
	./hack/push.sh "vsecm-safe" $(VERSION) "$(VSECM_DOCKERHUB_REGISTRY_URL)/vsecm-safe"

# Pushes the "VSecM Safe" (FIPS) container to the public registry.
safe-push-ist-fips:
	./hack/push.sh "vsecm-fips-safe" \
		$(VERSION) "$(VSECM_DOCKERHUB_REGISTRY_URL)/vsecm-fips-safe"

# Pushes the "VSecM Safe" container image to the local registry.
safe-push-ist-local:
	./hack/push.sh "vsecm-safe" $(VERSION) "$(VSECM_LOCAL_REGISTRY_URL)/vsecm-safe"

# Pushes the "VSecM Safe" container image to the local registry.
safe-push-ist-fips-local:
	./hack/push.sh "vsecm-fips-safe" \
		$(VERSION) "$(VSECM_LOCAL_REGISTRY_URL)/vsecm-fips-safe"
