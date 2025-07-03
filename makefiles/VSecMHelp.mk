# /*
# |    Protect your secrets, protect your sensitive data.
# :    Explore VMware Secrets Manager docs at https://vsecm.com/
# </
# <>/  keep your secrets... secret
# >/
# <>/' Copyright 2023-present VMware Secrets Manager contributors.
# >/'  SPDX-License-Identifier: BSD-2-Clause
# */

# TODO: need to rewrite.

#help:
#	@echo "--------------------------------------------------------------------"
#	@echo "          🛡️ VMware Secrets Manager: Keep your secrets... secret."
#	@echo "          🛡️ https://vsecm.com/"
#	@echo "--------------------------------------------------------------------"
#	@echo "        ℹ️ This Makefile assumes you use Minikube and Docker"
#	@echo "        ℹ️ for most operations."
#	@echo "--------------------------------------------------------------------"
#	@echo "If you are on the build server, stop the cronjob first: \`crontab -e\`"
#	@echo "Or \`sudo service cron stop\`"
#	@echo "--------------------------------------------------------------------"
#
#	@if [ "`uname`" = "Darwin" ]; then \
#		if type docker > /dev/null 2>&1; then \
#			echo "  Using Docker for Mac?"; \
#			echo "        ➡ 'make mac-tunnel' to proxy to the internal registry."; \
#		else \
#			echo "  Docker is not installed on this Mac."; \
#		fi; \
#	fi
