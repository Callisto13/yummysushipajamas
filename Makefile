REPO = github.com/Callisto13/yummysushipajamas
BIN_DIR = bin
PB_DIR = pb

.PHONY: all proto server client

all: proto server client docker
reload: all destroy deploy

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

slow-test:
	@echo "running all tests in docker container"
	@docker run -it --rm -v $(PWD):$(PWD) -w $(PWD) callisto13/go-ginkgo ginkgo -mod vendor -p -r

unit:
	@ginkgo -mod vendor -r client/
	@ginkgo -mod vendor -r server/

int:
	@echo "running integration tests in docker container"
	@docker run -it --rm -v $(PWD):$(PWD) -w $(PWD) callisto13/go-ginkgo ginkgo -mod vendor -p -r integration/

mock:
	@GOPRIVATE=$(REPO) GO111MODULE=on mockgen $(REPO)/$(PB_DIR) BasicClient,Basic_PrimeServer,Basic_PrimeClient > $(PB_DIR)/mocks/ysp_mock.go

docker:
	@docker build -t callisto13/ysp .
	@docker push callisto13/ysp

deploy:
	@ytt -f kube/deployment.yaml -f kube/values.yaml | kubectl apply -f-
	@ytt -f kube/service.yaml -f kube/values.yaml | kubectl apply -f-
	@echo "now run: \"export SERVICE_ADDR=$$(minikube service ysp-service --url | cut -d / -f 3)\""

destroy:
	@kubectl delete service/ysp-service
	@kubectl delete deployment/ysp

clean:
	@rm bin/*

help:
	@echo "Usage:"
	@echo "  proto     ..................... regenerate grpc sources (will go to ./ysp/ysp.proto)"
	@echo "  client    ..................... build the client bin (will go to ./bin/client)"
	@echo "  server    ..................... build the server bin (will go to ./bin/server)"
	@echo "  clean     ..................... delete bins"
	@echo "  test      ..................... run all test suites"
	@echo "  slow-test ..................... run all test suites in docker container"
	@echo "  unit      ..................... run server and client unit tests"
	@echo "  int       ..................... run integration tests"
	@echo "  dep       ..................... update dependencies"
	@echo "  mock      ..................... regenerate grpc testing mocks"
	@echo "  docker    ..................... rebuild and push callisto13/ysp docker image"
	@echo "  deploy    ..................... deploy server and service to minikube"
	@echo "  destroy   ..................... delete server and service"
	@echo "  reload    ..................... proto server client docker destroy deploy (aka rebuild and redeploy the lot)"
