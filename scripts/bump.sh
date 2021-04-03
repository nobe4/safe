#!/usr/bin/env bash
set -e

current_version=$(make version)

# awk -F. split by '.', increase the last number, add the '.' back
next_version=$(echo "$current_version" | awk -F. '{$NF = $NF + 1;} 1' | sed 's/ /./g')

echo "Bump version $current_version to version $next_version"

sed -i -e "s/$current_version/$next_version/" makefile

echo "Don't forget to push the new tag before you merge."
