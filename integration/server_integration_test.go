package integration_test

import (
	"bytes"
	"fmt"
	"os/exec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Server CLI", func() {
	Context("when a too-high log level is set", func() {
		var out *bytes.Buffer

		BeforeEach(func() {
			cmd := exec.Command(serverBin, "-log-level=panic")
			out = bytes.NewBuffer([]byte{})
			cmd.Stdout = out
			cmd.Stderr = out
		})

		It("no server side logs are visible", func() {
			Eventually(out.String()).Should(Equal(""))
		})
	})

	Context("when the `port` flag is provided", func() {
		It("the server is started on the given port", func() {
			out := lsofCmd(serverPort)
			Expect(out).To(ContainSubstring(fmt.Sprintf("%d", serverCmd.Process.Pid)))
		})
	})
})

func lsofCmd(port string) string {
	cmd := exec.Command("lsof", "-Fp", "-i", fmt.Sprintf(":%s", port))
	out, err := cmd.CombinedOutput()
	Expect(err).NotTo(HaveOccurred())

	return string(out)
}
