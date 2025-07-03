#!/usr/bin/env bash

# /*
# |    Protect your secrets, protect your sensitive data.
# :    Explore VMware Secrets Manager docs at https://vsecm.com/
# </
# <>/  keep your secrets... secret
# >/
# <>/' Copyright 2023-present VMware Secrets Manager contributors.
# >/'  SPDX-License-Identifier: BSD-2-Clause
# */

# This is a script to fix unsigned images.
# Normally, signing and pushing should be a single step
# and we should not need to pull the images and sign them again.
# So we'd rarely (if ever) need to use this script.

VERSION="0.40.0"

echo "TODO: docker pull will change since we will use gcr."
