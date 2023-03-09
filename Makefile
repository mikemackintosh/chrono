build:
	release/build.sh

deps:
	go mod download

clean:
	go clean -modcache
	rm release/bin/chrono-*
