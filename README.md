## Learning GRPC
1. golang middleware to handle error (https://segmentfault.com/a/1190000016601823)
1. proto spec v3
1. grpc context metadata
1. Protocol Buffers naming style guide (https://developers.google.com/protocol-buffers/docs/style)
1. send metadata from client to server (https://github.com/grpc/grpc-go/blob/master/Documentation/grpc-metadata.md)
1. client-side Interceptor (https://xuanwo.io/2019/03/10/request-id-in-grpc/)


## installation 
1. Install grpc[https://grpc.io/docs/quickstart/go/] <br/>
   remember to copy `include` folder as well.

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
   $ protoc helloworld/proto/helloworld.proto --go_out=plugins=grpc:.
   $ protoc helloworld/proto/helloworld.proto --plugin=protoc-gen-dart=c:\Users\jason\AppData\Roaming\Pub\Cache\bin\protoc-gen-dart.bat --dart_out=grpc:dart\proto
   ```

## Golang Note
1. no time struct in protobuf
 
## References
1. https://segmentfault.com/a/1190000016601810