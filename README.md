1. Install grpc[https://grpc.io/docs/quickstart/go/]

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

   #### Golang Note
   1. no time struct in protobuf
   2. 

   #### References

   #### https://segmentfault.com/a/1190000016601810