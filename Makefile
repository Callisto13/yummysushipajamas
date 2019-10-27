REPO = github.com/Callisto13/yummysushipajamas
BIN_DIR = bin
PB_DIR = pb

.PHONY: all proto server client

all: proto server client

server:
	@GOOS=linux go build -v -o $(BIN_DIR)/server $(PWD)/server/cmd

client:
	@go build -v -o $(BIN_DIR)/client $(PWD)/client/cmd

dep:
	@go mod vendor
	@go mod tidy

proto:
	@protoc -I $(PB_DIR)/ $(PB_DIR)/ysp.proto --go_out=plugins=grpc:pb

test: unit int

unit:
	@ginkgo -mod vendor -r client/
	@ginkgo -mod vendor -r server/

int:
	@echo "running integration tests in docker container"
	@docker run -it --rm -v $(PWD):$(PWD) -w $(PWD) callisto13/go-ginkgo ginkgo -mod vendor -p -r integration/

mock:
	@GOPRIVATE=$(REPO) GO111MODULE=on mockgen $(REPO)/$(PB_DIR) BasicClient,Basic_PrimeServer,Basic_PrimeClient > $(PB_DIR)/mocks/ysp_mock.go

clean:
	@rm bin/*

help:
	@echo "Usage:"
	@echo "  proto   ..................... regenerate grpc sources (will go to ./ysp/ysp.proto)"
	@echo "  client  ..................... build the client bin (will go to ./bin/client)"
	@echo "  server  ..................... build the server bin (will go to ./bin/server)"
	@echo "  clean   ..................... delete bins"
	@echo "  test    ..................... run all test suites"
	@echo "  unit    ..................... run server and client unit tests"
	@echo "  int     ..................... run integration tests"
	@echo "  dep     ..................... update dependencies"
	@echo "  mock    ..................... regenerate grpc testing mocks"
