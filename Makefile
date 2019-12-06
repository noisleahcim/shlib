build:
	@go build -o shlib main.go

test:
	@/bin/sh tests/import.sh