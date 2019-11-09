.PHONY: build

run: build
	./bin/go-ast-transform

build:
	mkdir -p bin/
	go build -v -o bin/

gen:
	go tool compile -S testdata/call_study.go > /tmp/call_study.S

