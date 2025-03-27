default: check

tidy:
	go mod tidy

fmt: tidy
	trunk fmt

fmt-all: tidy
	trunk fmt --all

check: fmt
	trunk check

check-all: fmt-all
	trunk check --all

test:
	go test ./...

update: upgrade
upgrade: tidy
	go get -u
	trunk upgrade

## Git Hooks
pre-commit: clean check test
	git add go.mod go.sum
