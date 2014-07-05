deps:
	go get

build:
	go build

buildall:
	gox

clean:
	rm -f ./env
	rm -f ./env_*