# /*
# |    Protect your secrets, protect your sensitive data.
# :    Explore VMware Secrets Manager docs at https://vsecm.com/
# </
# <>/  keep your secrets... secret
# >/
# <>/' Copyright 2023-present VMware Secrets Manager contributors.
# >/'  SPDX-License-Identifier: BSD-2-Clause
# */

serve-docs:
	./hack/web-serve.sh

sync-docs:
	./hack/web-sync.sh

bundle: \
	inspector-bundle \
	keygen-bundle \
	safe-bundle \
	sidecar-bundle \
	sentinel-bundle \
	init-container-bundle

push: \
	inspector-push \
	keygen-push \
	safe-push \
	sidecar-push \
	sentinel-push \
	init-container-push
