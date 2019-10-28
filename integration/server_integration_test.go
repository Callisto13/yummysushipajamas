package integration_test

import (
	"fmt"
	"os/exec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("Server CLI", func() {
	Context("when a too-high log level is set", func() {
		var session *gexec.Session

		BeforeEach(func() {
			cmd := exec.Command(serverBin, "-log-level=panic")
			var err error
			session, err = gexec.Start(cmd, GinkgoWriter, GinkgoWriter)
			Expect(err).NotTo(HaveOccurred())
		})

		It("no server side logs are visible", func() {
			Eventually(session.Out.Contents()).Should(BeEmpty())
		})
	})

	Context("when the `port` flag is provided", func() {
		It("the server is started on the given port", func() {
			out := lsofCmd(serverPort)
			Expect(out).To(ContainSubstring(fmt.Sprintf("%d", serverSession.Command.Process.Pid)))
		})
	})
})

func lsofCmd(port string) string {
	cmd := exec.Command("lsof", "-Fp", "-i", fmt.Sprintf(":%s", port))
	out, err := cmd.CombinedOutput()
	Expect(err).NotTo(HaveOccurred())

	return string(out)
}
