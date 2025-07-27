# /*
# |    Protect your secrets, protect your sensitive data.
# :    Explore VMware Secrets Manager docs at https://vsecm.com/
# </
# <>/  keep your secrets... secret
# >/
# <>/' Copyright 2023-present VMware Secrets Manager contributors.
# >/'  SPDX-License-Identifier: BSD-2-Clause
# */

# 0. Prune docker file system to save resources.
docker-cleanup:
	./hack/docker/cleanup.sh

# 1. Reset the test cluster.
k8s-delete:
	./hack/k8s/minikube-delete.sh

# 2. Start the test cluster.
k8s-start:
	./hack/k8s/minikube-start.sh

# Deletes and re-installs the Minikube cluster.
k8s-reset:
	k8s-delete
	k8s-start

# 3. Build container images.
docker-build:
	./hack/docker/build-local.sh

# 4. Forward registry.
docker-forward-registry:
	./hack/docker/minikube-forward-registry.sh

# 5. Push to the container registry.
docker-push:
	./hack/docker/push-local.sh