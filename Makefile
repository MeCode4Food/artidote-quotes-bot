clean:
	rm -rf /bin/artidote-quote;
build-debug:
	GOOS=linux GOARCH=amd64 go build -gcflags "all=-N -l" -o ./debug/artidote-quote;
	zip -r function.zip .;
build:
	GOARCH=amd64 GOOS=linux go build -o ./bin/artidote-quote;
	zip -r function.zip .;
