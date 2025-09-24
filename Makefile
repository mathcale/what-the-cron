.PHONY: all build clean

all: clean build

build:
	@CGO_ENABLED=0 go build -a -ldflags="-s -w" -trimpath -o bin/what-the-cron ./cli

clean:
	@rm -rf bin/
