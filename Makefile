# Makefile
.PHONY: test doc
test:
	go test -v -cover -race ./...

doc:
	go install github.com/gopherguides/hype/cmd/hype@v0.3.0
	rm README.md
	hype export -f hype.md -format markdown > README.md
