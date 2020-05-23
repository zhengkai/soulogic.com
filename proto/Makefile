SHELL := /bin/bash

all:
	# $(MAKE) server
	$(MAKE) client

server:
	# which protoc-gen-go >/dev/null || go get google.golang.org/protobuf/cmd/protoc-gen-go@v1.23.0
	echo server

client:
	[ ! -f ../client/node_modules/protobufjs/bin/pbjs ] && cd ../client && make init || :
	../client/node_modules/protobufjs/bin/pbjs -t static-module -w es6 --es6 --force-number -o ../client/src/pb/pb.js *.proto
	../client/node_modules/protobufjs/bin/pbts -o ../client/src/pb/pb.d.ts ../client/src/pb/pb.js
