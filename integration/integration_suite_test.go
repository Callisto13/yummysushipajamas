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
	cliBin        string
	serverBin     string
	serverPort    string
	serverSession *gexec.Session
)

func TestIntegration(t *testing.T) {
	RegisterFailHandler(Fail)

	BeforeSuite(func() {
		var err error
		cliBin, err = gexec.Build("github.com/Callisto13/yummysushipajamas/client/cmd", "-mod=vendor")
		Expect(err).NotTo(HaveOccurred())

		serverBin, err = gexec.Build("github.com/Callisto13/yummysushipajamas/server/cmd", "-mod=vendor")
		Expect(err).NotTo(HaveOccurred())

		serverPort = fmt.Sprintf("1430%d", config.GinkgoConfig.ParallelNode)

		serverCmd := exec.Command(serverBin, "-port", serverPort)
		serverSession, err = gexec.Start(serverCmd, GinkgoWriter, GinkgoWriter)
		Expect(err).NotTo(HaveOccurred())
	})

	AfterSuite(func() {
		serverSession.Terminate().Wait()
		gexec.Terminate()
		gexec.CleanupBuildArtifacts()
	})

	RunSpecs(t, "Integration Suite")
}
