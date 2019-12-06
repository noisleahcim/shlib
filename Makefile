build:
	@go build -o shlib main.go

test:
	@/bin/sh tests/cli/import/all
	@/bin/sh tests/cli/import/module
	@/bin/sh tests/cli/import/function