#!/usr/bin/env bash

REPO=$(dirname "$0")
pushd "${REPO}" &> /dev/null

export GITHUB_TOKEN="$(gh auth token)"

if [[ -z "${GITHUB_TOKEN}" ]]; then
    echo "GITHUB_TOKEN is not set"
    popd &> /dev/null
    exit 1
fi

VERSION=$(grep '^const version' cmd/cmd.go | cut -d '"' -f2)

if git tag --list | grep -q "${VERSION}"; then
    echo "version found in tags"
    popd &> /dev/null
    exit 1
fi

git tag -a "v${VERSION}" -m "${VERSION}"
git push origin "v${VERSION}"

goreleaser release --clean

popd &> /dev/null
