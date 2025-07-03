# /*
# |    Protect your secrets, protect your sensitive data.
# :    Explore VMware Secrets Manager docs at https://vsecm.com/
# </
# <>/  keep your secrets... secret
# >/
# <>/' Copyright 2023-present VMware Secrets Manager contributors.
# >/'  SPDX-License-Identifier: BSD-2-Clause
# */

# Packages the "VSecM Sentinel" binary into a container image.
sentinel-bundle:
	./hack/bundle.sh "vsecm-sentinel" \
		$(VERSION) "dockerfiles/vsecm/sentinel.Dockerfile"

# Pushes the "VSecM Sentinel" container image the the local registry.
sentinel-push:
	./hack/push.sh "vsecm-sentinel" \
		$(VERSION) "$(VSECM_DOCKERHUB_REGISTRY_URL)/vsecm-sentinel"
