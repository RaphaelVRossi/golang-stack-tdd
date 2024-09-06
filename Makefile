test:
	@TMPDIR=.tmp/ go test -v ./...

setup:
	@go mod tidy
