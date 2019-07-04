PROJECT=grpc-apps

gen: $(shell which protoc)
	protoc -I proto proto/hello.proto --go_out=plugins=grpc:./go/helloworld/
	protoc -I proto proto/hello.proto \
    --swift_out=./iosDemo/iosDemo/ \
    --swiftgrpc_out=Client=true,Server=false:./iosDemo/iosDemo/