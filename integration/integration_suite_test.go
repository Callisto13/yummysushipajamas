package integration_test

import (
	"fmt"
	"os/exec"
	"testing"

	. "github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/config"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

var (
	cliBin    string
	serverBin string
)

func TestIntegration(t *testing.T) {
	RegisterFailHandler(Fail)

	var serverCmd *exec.Cmd

	BeforeSuite(func() {
		var err error
		cliBin, err = gexec.Build("github.com/Callisto13/yummysushipajamas/client/cmd", "-mod=vendor")
		Expect(err).NotTo(HaveOccurred())

		serverBin, err = gexec.Build("github.com/Callisto13/yummysushipajamas/server/cmd", "-mod=vendor")
		Expect(err).NotTo(HaveOccurred())

		serverCmd = exec.Command(serverBin, "-port", fmt.Sprintf("1430%d", config.GinkgoConfig.ParallelNode))
		Expect(serverCmd.Start()).To(Succeed())
	})

	AfterSuite(func() {
		//TODO: why is .Wait() being weird?
		Expect(serverCmd.Process.Kill()).To(Succeed())
		gexec.CleanupBuildArtifacts()
	})

	RunSpecs(t, "Integration Suite")
}
