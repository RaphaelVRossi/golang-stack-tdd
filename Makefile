test:
	@TMPDIR=.tmp/ CGO_ENABLED=0 go test -v ./...

setup:
	@go mod tidy
