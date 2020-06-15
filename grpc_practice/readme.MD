## Generate grpc code
./protoc.exe --go_out=plugins=grpc:../gravatar gravatar.proto  

## Run tests and coverage
go test -coverprofile cover.out && go tool cover -html=cover.out