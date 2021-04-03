#!/usr/bin/env bash
set -e

if [ "$(git symbolic-ref --short HEAD)" != "master" ]; then
	echo "Not on master branch, not tagging"
	exit 1
fi

release=$(make version)

if git rev-parse "$release" >/dev/null 2>&1
then
	echo "Tag $release already present, run 'make bump'"
else
	git tag "$release"
	git push origin "$release"
fi
