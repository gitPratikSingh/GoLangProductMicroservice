swagger: check_install
	GO111MODULE=off swagger generate spec -o ./swagger.yaml --scan-models
check_install:
	which swagger || GO111MODULE=off go get -u github.com/go-swagger/go-swagger/cmd/swagger

check_grpc_install:
	which protoc || brew install protobuf
	GO111MODULE=on go get github.com/golang/protobuf/protoc-gen-go

generate_proto: check_grpc_install
	PATH="${PATH}:${GOPATH}/go/bin"
	protoc -I grpc/ grpc/services.proto --go_out=plugins=grpc:grpc/currency
	
