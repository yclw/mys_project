# xxx -- golang web server template

```
cd ./deployments && docker_compose up -d && cd ..

go work init

cd ./apps/api && go mod tidy && cd ../
cd ./apps/auth && go mod tidy && cd ../../
	
cd ./pkg/common && go mod tidy && cd ../../
cd ./pkg/model && go mod tidy && cd ../../
cd ./pkg/protobuf && go mod tidy && ./scripts/generate.sh && cd ../../
cd ./pkg/utils && go mod tidy && cd ../../

cd ./apps/auth && go run main.go
cd ./apps/api && go run main.go
```