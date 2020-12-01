#!/usr/bin/env bash

DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" >/dev/null 2>&1 && pwd)"

if command -v nix-shell /dev/null 2>&1; then
  printf "%s/nix-exec.sh" "${DIR}"
fi
