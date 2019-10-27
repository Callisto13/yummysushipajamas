package server_test

import (
	"testing"

	"github.com/golang/mock/gomock"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var ctrl *gomock.Controller

func TestServer(t *testing.T) {
	RegisterFailHandler(Fail)

	BeforeSuite(func() {
		ctrl = gomock.NewController(t)
	})

	AfterSuite(func() {
		ctrl.Finish()
	})

	RunSpecs(t, "Server Suite")
}
