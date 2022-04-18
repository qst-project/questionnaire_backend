# questionnaire_backend

### Set up integration tests

    docker-compose up -d

### Update proto file

    protoc -I api/proto --go_out=./pkg/api --go-grpc_out=./pkg/api --go_opt=paths=source_relative api/proto/questionnaire.proto

### Architecture

    grpc <-> delegate <-> usecase (using core models) <-> gateway <-> postgres

delegate checks if the requested action is even possible and maps incoming data to "core" format use case uses core
models and executes business logic gateway combines multiple postgres queries to achieve requested result
