# /*
# |    Protect your secrets, protect your sensitive data.
# :    Explore VMware Secrets Manager docs at https://vsecm.com/
# </
# <>/  keep your secrets... secret
# >/
# <>/' Copyright 2023-present VMware Secrets Manager contributors.
# >/'  SPDX-License-Identifier: BSD-2-Clause
# */

relay-client-bundle-ist:
	./hack/bundle.sh "vsecm-relay-client" \
		$(VERSION) "dockerfiles/vsecm/relay-client.Dockerfile"

relay-client-bundle-ist-fips:
	./hack/bundle.sh "vsecm-fips-relay-client" \
		$(VERSION) "dockerfiles/vsecm-fips/relay-client.Dockerfile"

relay-client-push-ist:
	./hack/push.sh "vsecm-relay-client" \
		$(VERSION) "$(VSECM_DOCKERHUB_REGISTRY_URL)/vsecm-relay-client"

relay-client-push-ist-fips:
	./hack/push.sh "vsecm-fips-relay-client" \
		$(VERSION) "$(VSECM_DOCKERHUB_REGISTRY_URL)/vsecm-fips-relay-client"

relay-client-push-ist-local:
	./hack/push.sh "vsecm-relay-client" $(VERSION) \
		"$(VSECM_LOCAL_REGISTRY_URL)/vsecm-relay-client"

relay-client-push-ist-fips-local:
	./hack/push.sh "vsecm-fips-relay-client" $(VERSION) \
		"$(VSECM_LOCAL_REGISTRY_URL)/vsecm-fips-relay-client"
