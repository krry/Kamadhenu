.PHONY: build clean install test run

build:
	go build -o kamadhenu

clean:
	rm -f kamadhenu
	rm -rf dist/

install: build
	install -m 755 kamadhenu /usr/local/bin/
	install -m 644 kamadhenu.1 /usr/local/share/man/man1/

uninstall:
	rm -f /usr/local/bin/kamadhenu
	rm -f /usr/local/share/man/man1/kamadhenu.1

test:
	go test ./...

run: build
	./kamadhenu

release:
	./build.sh

dist: release
