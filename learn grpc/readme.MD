## generate grpc
./protoc.exe --go_out=plugins=grpc:chat chat.proto  

## Handy commands
go get -u github.com/golang/protobuf/{proto,protoc-gen-go}  
go get -u google.golang.org/grpc  
protoc --go_out=plugins=grpc:. *.proto  