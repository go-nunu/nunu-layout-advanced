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
	go test -coverpkg=./internal/handler,./internal/service,./internal/repository -coverprofile=./.nunu/coverage.out ./test/server/...
	go tool cover -html=./.nunu/coverage.out -o coverage.html

