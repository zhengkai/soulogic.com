SHELL:=/bin/bash

NG := "./node_modules/@angular/cli/bin/ng"

DOMAIN := "dev.local.9farm.com"

local:
	$(NG) serve --port 22001 --host 127.0.0.1 --public-host $(DOMAIN)

init:
	NG_CLI_ANALYTICS=ci npm install
	npm audit fix

test:
	./dist/build.sh test.local.9farm.com test
