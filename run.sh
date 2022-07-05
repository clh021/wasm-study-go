#!/usr/bin/env bash
# leehom Chen clh021@gmail.com
set -ex
# 路径准备
# OldPath=$(pwd)
# SCRIPT_PATH=$(realpath "${BASH_SOURCE[0]}")
SCRIPT_PATH=$(realpath "$0")
ProjectPath="$(dirname "$SCRIPT_PATH")"

# prepare load lib
if [ ! -f "${ProjectPath}/wasm_exec.js" ] ; then
    cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" "${ProjectPath}/web/wasm_exec.js"
fi

# build wasm
go mod tidy
GOARCH=wasm GOOS=js go build -o web/alert.wasm .

# server html
defaultPort=8000
PORT=${1:-$defaultPort}
python3 -m http.server "$PORT" -b "0.0.0.0" --directory web