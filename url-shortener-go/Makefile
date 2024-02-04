MAINFILE = ./cmd/url-shortener/main.go
OUTFILE = ./bin/url-shortener-go

help:
	@echo "Select a command to run from the makefile"
	@echo "\t run -> run the binary"
	@echo "\t build -> build the binary"
	@echo "\t build-prod -> build the binary with tidy"
	@echo "\t watch -> run the binary with CompileDaemon"
	@echo "\t clean -> clean the binary"
	@echo "\t ready -> get all the dependencies"

build:
	@go build -o ${OUTFILE} ${MAINFILE}

run: build
	@./${OUTFILE}

tidy:
	go mod tidy

build-prod: tidy build

test:
	@go test -v ./...

clean:
	@rm -rf bin
