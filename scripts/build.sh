set -euo pipefail
script_dirpath="$(cd "$(dirname "${BASH_SOURCE[0]}")"; pwd)"
root_dirpath="$(dirname "${script_dirpath}")"

docker build -t "mieubrisse/datastore-army-lambda" --progress=plain -f "${root_dirpath}/Dockerfile" "${root_dirpath}"
