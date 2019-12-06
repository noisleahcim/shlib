build:
	@go build -o shlib main.go

test:
	@/bin/sh tests/cli/import/all.sh
	@/bin/sh tests/cli/import/module.sh
	@/bin/sh tests/cli/import/function.sh