.PHONY: init
init:
	go install github.com/google/wire/cmd/wire@latest
	go install github.com/golang/mock/mockgen@latest

.PHONY: mock
mock:
	mockgen -source=internal/service/user.go -destination mocks/service/user.go
	mockgen -source=internal/repository/user.go -destination mocks/repository/user.go

.PHONY: test
test:
	go test -coverpkg=./internal/handler,./internal/service,./internal/repository -coverprofile=./coverage.out ./test/server/...
	go tool cover -html=./coverage.out -o coverage.html

.PHONY: build
build:
	go build -ldflags="-s -w" -o ./bin/server ./cmd/server/...

.PHONY: docker
docker:
	docker build -f deploy/build/Dockerfile --build-arg APP_RELATIVE_PATH=./cmd/job/... -t 1.1.1.1:5000/demo-api:v1 .