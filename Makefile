
BIN_NAME=gortb

setup:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

.PHONY: protobufs
protobufs:
	protoc -I=./ --go_out=./ ./protobufs/gtfs-realtime.proto

build:
	go build -o ./build/$(BIN_NAME)

.PHONY: clean
clean:
	rm -rf ./build
