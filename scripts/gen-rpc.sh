#!/bin/sh

oto -template ./templates/server.go.plush -out ./backend/rpc/rpcserver/rpc.server.go -ignore Ignorer -pkg rpcserver ./backend/definitions
oto -template ./templates/client.go.plush -out ./backend/rpc/rpcclient/rpc.client.go -ignore Ignorer -pkg rpcclient ./backend/definitions
oto -template ./templates/client.js.plush -out ./frontend/src/rpc/rpc.gen.js -ignore Ignorer ./backend/definitions

gofmt -w ./backend/rpc/