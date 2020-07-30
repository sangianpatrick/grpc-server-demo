#protoc --proto_path=$GOPATH/src/github.com/google/protobuf/src --proto_path=./proto --go_out=plugins=grpc:./src/pb ./proto/*.proto

protoc --proto_path=./proto --go_out=plugins=grpc:./src/pb ./proto/*.proto