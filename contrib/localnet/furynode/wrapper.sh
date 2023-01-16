#!/usr/bin/env sh

##
## Input parameters
##
ID=${ID:-0}
LOG=${LOG:-fury.log}

##
## Assert linux binary
##
if ! [ -f "/app/build/furyd" ]; then
	echo "The binary /app/build/furyd cannot be found. Please add the binary to the shared folder."
	exit 1
fi

#BINARY_CHECK="$(file "${BINARY}" | grep 'ELF 64-bit LSB executable, x86-64')"
#if [ -z "${BINARY_CHECK}" ]; then
#	echo "Binary needs to be OS linux, ARCH amd64"
#	exit 1
#fi

##
## Run binary with all parameters
##
export FYDHOME="/fury/node${ID}/fury"

if [ -d "$(dirname "${FYDHOME}"/"${LOG}")" ]; then
  "/app/build/furyd" --home "${FYDHOME}" "$@" | tee "${FYDHOME}/${LOG}"
else
  "/app/build/furyd" --home "${FYDHOME}" "$@"
fi