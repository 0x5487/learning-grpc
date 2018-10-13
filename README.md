1. Install [protobuf compiler](https://github.com/google/protobuf/blob/master/README.md#protocol-compiler-installation)

1. Install the protoc Go plugin

   ```
   $ go get -u github.com/golang/protobuf/protoc-gen-go
   ```

1. Rebuild the generated Go code

   ```
   $ go generate google.golang.org/grpc/examples/helloworld/...
   ```
   
   Or run `protoc` command (with the grpc plugin)
   
   ```
   $ protoc helloworld/helloworld.proto --go_out=.
   ```

   #### Golang Note
   1. no time struct in protobuf
   2. 