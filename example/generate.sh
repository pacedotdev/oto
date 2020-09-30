#!/usr/bin/env bash

oto -template server.go.plush \
	-out server.gen.go \
	-pkg main \
	./def
gofmt -w server.gen.go server.gen.go
echo "generated server.gen.go"

oto -template client.js.plush \
	-out client.gen.js \
	-pkg main \
	./def
echo "generated client.gen.js"

oto -template client.swift.plush \
	-out ./swift/SwiftCLIExample/SwiftCLIExample/client.gen.swift \
	-pkg main \
	./def
echo "generated client.gen.swift"
