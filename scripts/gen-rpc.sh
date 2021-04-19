#!/bin/sh

. ./scripts/.env

oto -template ./templates/server.go.plush -out ./backend/rpc/rpcserver/rpc.server.go -ignore Ignorer -pkg rpcserver ./backend/definitions
oto -template ./templates/client.go.plush -out ./backend/rpc/rpcclient/rpc.client.go -ignore Ignorer -pkg rpcclient ./backend/definitions
oto -template ./templates/client.js.plush -out ./frontend/src/rpc/rpc.client.js -ignore Ignorer -params "backend_scheme:${BACKEND_SCHEME},backend_host:${BACKEND_HOST},backend_port:${BACKEND_PORT},backend_path:${BACKEND_PATH}" ./backend/definitions

gofmt -w ./backend/rpc/