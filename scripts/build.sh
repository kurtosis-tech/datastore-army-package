#!/usr/bin/env bash
set -euo pipefail # Bash "strict mode"
script_dirpath="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

# ==================================================================================================
#                                             Constants
# ==================================================================================================
source "${script_dirpath}/_constants.env"

# =============================================================================
#                                 Main Code
# =============================================================================
echo "'${IMAGE_ORG_AND_REPO}'' is now a Kurtosis module implemented purely in Kurtosis script language. Nothing to build."
