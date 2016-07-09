To build protobufffer messages:

# go fetch brings down the soruces
go fetch

protoc --go_out=plugins=grpc,import_prefix=,import_path=proto:$(pwd)/proto  -I=/usr/local/Cellar/protobuf/2.6.0/include -I /Users/michaelritsema/gocode/src/github.com/ziften/protocolbuffers/Protobuf /Users/michaelritsema/gocode/src/github.com/ziften/protocolbuffers/Protobuf/*.proto

