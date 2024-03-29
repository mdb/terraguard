#!/bin/sh

version="${1}"
repo="mdb/terraguard"

result="$(curl \
  --header "Accept: application/vnd.github.v3+json" \
  --write "%{http_code}" \
  --out "/dev/null" \
  "https://api.github.com/repos/${repo}/releases/tags/${version}")"

if [ "${result}" = "404" ]; then
  exit 0
fi

echo "${version} is an existing ${repo} release"
exit 1
