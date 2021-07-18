set -euo pipefail
script_dirpath="$(cd "$(dirname "${BASH_SOURCE[0]}")"; pwd)"
root_dirpath="$(dirname "${script_dirpath}")"

docker build -t datastore-army --progress=plain -f "${root_dirpath}/Dockerfile" "${root_dirpath}"
