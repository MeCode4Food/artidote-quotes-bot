clean:
	rm -rf /bin/artidote-quote;
build-debug:
	GO11MODULE=on GOOS=linux GOARCH=amd64 go build -gcflags "all=-N -l" -o ./debug/artidote-quote;
	zip -r function.zip .;
build:
	GO11MODULE=on GOARCH=amd64 GOOS=linux go build -o ./bin/artidote-quote;
	zip -r function.zip .;
