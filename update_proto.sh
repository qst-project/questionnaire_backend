protoc -I api/proto --go_out=./pkg/proto --go-grpc_out=./pkg/proto --go_opt=paths=source_relative api/proto/questionnaire.proto
