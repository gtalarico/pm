.PHONY: release(GO)

# Colors
NC=\x1b[0m
L_GREEN=\x1b[32;01m

## usage: print useful commands
usage:
	@echo "$(L_GREEN)Choose a command: $(PWD) $(NC)"
	@bash -c "sed -ne 's/^##//p' ./Makefile | column -t -s ':' |  sed -e 's/^/ /'"

# release: creates github release on gtalarico/pm and brew on gtalarico/homebrew-tap
lint:
	golangci-lint run ./...

# release: creates github release on gtalarico/pm and brew on gtalarico/homebrew-tap
release:
	goreleaser release --rm-dist

# coverage: show test coverage
coverage:
	go test ./... -cover -race -coverprofile=coverage.out
	go tool cover -func=coverage.out
