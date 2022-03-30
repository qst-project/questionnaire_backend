# questionnaire_backend


### Set up integration tests

    docker-compose up -d

### Update proto file

    protoc -I api/proto --go_out=./pkg/api --go-grpc_out=./pkg/api --go_opt=paths=source_relative api/proto/questionnaire.proto