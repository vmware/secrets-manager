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
safe-bundle:
	./hack/bundle.sh "vsecm-safe" \
		$(VERSION) "dockerfiles/vsecm/safe.Dockerfile"

# Pushes the "VSecM Safe" container to the public registry.
safe-push:
	./hack/push.sh "vsecm-safe" $(VERSION) "$(VSECM_DOCKERHUB_REGISTRY_URL)/vsecm-safe"
