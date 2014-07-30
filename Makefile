default: test

deps:
	go get
	go get github.com/stretchr/testify/assert

build:
	go build

buildall:
	gox

test:
	go test

clean:
	rm -f ./envd
	rm -f ./envd_*