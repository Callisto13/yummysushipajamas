package integration_test

import (
	"fmt"
	"os/exec"
	"time"

	. "github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/config"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("Server CLI", func() {
	// graceful shutdown
	Context("when the server is killed mid-call", func() {
		var cliCmd *exec.Cmd

		BeforeEach(func() {
			args := []string{"-action=prime", "0", "100000"}
			cliCmd = exec.Command(cliBin, args...)
			cliCmd.Env = []string{fmt.Sprintf("SERVICE_ADDR=localhost:1430%d", config.GinkgoConfig.ParallelNode)}
		})

		It("fulfils the current request before shutting down", func(done Done) {
			var session *gexec.Session
			cmdReturns := make(chan struct{})

			go func() {
				defer GinkgoRecover()

				var err error
				session, err = gexec.Start(cliCmd, GinkgoWriter, GinkgoWriter)
				Expect(err).NotTo(HaveOccurred())
				Eventually(session, "10s").Should(gexec.Exit(0))

				close(cmdReturns)
			}()

			time.Sleep(time.Second)
			serverSession.Interrupt()

			Eventually(cmdReturns, "10s").Should(BeClosed())
			Expect(string(session.Out.Contents())).To(ContainSubstring("99991"))

			close(done)
		}, 15.0)
	})

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
