.PHONY: usage start stop restart inspect log stopall tests cli requirements clean deploy


# Colors
NC=\x1b[0m
L_GREEN=\x1b[32;01m

## usage: print useful commands
usage:
	@echo "$(L_GREEN)Choose a command: $(PWD) $(NC)"
	@bash -c "sed -ne 's/^##//p' ./Makefile | column -t -s ':' |  sed -e 's/^/ /'"

## release: create a release
install:
	go install -a .

## release: create a release
release:
	goreleaser --rm-dist
	rm -rdf dist

## brew: create brew formulae
brew:
	brew create distpm_0.1.0_Darwin_x86_64.tar.gz
804 KB



## lint: Lint
lint:
	ls
