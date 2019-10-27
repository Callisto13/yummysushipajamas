REPO = github.com/Callisto13/yummysushipajamas
PB_DIR = pb

dep:
	@go mod vendor
	@go mod tidy

proto:
	@protoc -I $(PB_DIR)/ $(PB_DIR)/ysp.proto --go_out=plugins=grpc:pb

unit:
	@ginkgo -mod vendor -r client/
	@ginkgo -mod vendor -r server/

mock:
	@GOPRIVATE=$(REPO) GO111MODULE=on mockgen $(REPO)/$(PB_DIR) BasicClient,Basic_PrimeServer,Basic_PrimeClient > $(PB_DIR)/mocks/ysp_mock.go

help:
	@echo "Usage:"
	@echo "  proto   ..................... regenerate grpc sources (will go to ./ysp/ysp.proto)"
	@echo "  unit    ..................... run server and client unit tests"
	@echo "  dep     ..................... update dependencies"
	@echo "  mock    ..................... regenerate grpc testing mocks"
