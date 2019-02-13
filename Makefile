setup: run_services
	@go run ./cmd/db/setup.go

run_services:
	@docker-compose up --build -d

run_server:
	@MONGO_URL=mongodb://0.0.0.0:27017/admin PORT=:4444 go run cmd/main.go

run_client:
	@/bin/bash -c "cd $$GOPATH/src/github.com/zhongweili/vue-go-spa/pkg/http/web/app && yarn serve"
