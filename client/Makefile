SHELL:=/bin/bash

NG := "./node_modules/@angular/cli/bin/ng"

HOST := $(shell hostname | tr '[:upper:]' '[:lower:]')

DOMAIN := "dev.$(HOST).9farm.com"

local: protobuf
	$(NG) serve --port 22001 --host 127.0.0.1 --public-host $(DOMAIN)

init:
	NG_CLI_ANALYTICS=ci npm install
	npm audit fix

protobuf:
	@[ ! -f $(NG) ] && make init || :
	@[ ! -f src/pb/pb.js ] && cd ../proto && make client || :

test: protobuf
	./dist/build.sh test.local.9farm.com test
