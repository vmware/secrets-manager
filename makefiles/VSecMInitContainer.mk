# /*
# |    Protect your secrets, protect your sensitive data.
# :    Explore VMware Secrets Manager docs at https://vsecm.com/
# </
# <>/  keep your secrets... secret
# >/
# <>/' Copyright 2023-present VMware Secrets Manager contributors.
# >/'  SPDX-License-Identifier: BSD-2-Clause
# */

# Packages the "VSecM Init Container" binary into a container image.
init-container-bundle:
	./hack/bundle.sh "vsecm-init-container" \
		$(VERSION) "dockerfiles/vsecm/init-container.Dockerfile"

# Pushes the "VSecM Init Container" container image to the public registry.
init-container-push:
	./hack/push.sh "vsecm-init-container" \
		$(VERSION) "$(VSECM_DOCKERHUB_REGISTRY_URL)/vsecm-init-container"
