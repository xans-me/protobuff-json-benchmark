  
#!/bin/bash

protoc grpcProtobuff/proto/*.proto --go_out=plugins=grpc:. --go_opt=paths=source_relative