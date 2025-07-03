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
sidecar-bundle:
	./hack/bundle.sh "vsecm-sidecar" \
		$(VERSION) "dockerfiles/vsecm/sidecar.Dockerfile"

# Pushes the "VSecM Sidecar" container image to the public registry.
sidecar-push:
	./hack/push.sh "vsecm-sidecar" \
		$(VERSION) "$(VSECM_DOCKERHUB_REGISTRY_URL)/vsecm-sidecar"
