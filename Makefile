project = ipdata

deps:
	-rm Gopkg.toml
	-rm Gopkg.lock
	-rm -r vendor
	dep init
test:
	go clean
	go test ./...
run:
	go run main.go
build:
	make deps
	go build -o bin/server
deploy:
	make deps
	make test
	env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/server
	heroku container:login
	heroku container:push web -a $(project)
	heroku container:release web -a $(project)
	rm ipdata