deps:
	go get

build:
	go build

buildall:
	gox

clean:
	rm -f ./envd
	rm -f ./envd_*