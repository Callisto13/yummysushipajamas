package client_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var ctrl *gomock.Controller

func TestClient(t *testing.T) {
	RegisterFailHandler(Fail)

	BeforeSuite(func() {
		ctrl = gomock.NewController(t)
	})

	AfterSuite(func() {
		ctrl.Finish()
	})

	RunSpecs(t, "Client Suite")
}
