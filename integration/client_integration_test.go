package integration_test

import (
	"fmt"
	"os/exec"

	. "github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/config"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("YSP Client", func() {
	var (
		cliCmd  *exec.Cmd
		cliArgs []string
		cliEnv  []string
	)

	JustBeforeEach(func() {
		cliCmd = exec.Command(cliBin, cliArgs...)
		cliCmd.Env = cliEnv
	})

	BeforeEach(func() {
		cliEnv = []string{fmt.Sprintf("SERVICE_ADDR=localhost:1430%d", config.GinkgoConfig.ParallelNode)}
	})

	Context("when SERVICE_ADDR is not set for the client", func() {
		BeforeEach(func() {
			cliEnv = []string{}
		})

		It("client complains and exits", func() {
			session, err := gexec.Start(cliCmd, GinkgoWriter, GinkgoWriter)
			Expect(err).NotTo(HaveOccurred())
			Eventually(session.Out).Should(gbytes.Say("SERVICE_ADDR not set\n"))
		})
	})

	Context("no action is chosen", func() {
		BeforeEach(func() {
			cliArgs = []string{}
		})

		It("client prints the help and exits", func() {
			session, err := gexec.Start(cliCmd, GinkgoWriter, GinkgoWriter)
			Expect(err).NotTo(HaveOccurred())
			Eventually(session.Out).Should(gbytes.Say("usage: -action=<sum|prime> n1 n2\n"))
		})
	})

	Context("when fewer than two numbers are provided", func() {
		BeforeEach(func() {
			cliArgs = []string{"-action=sum"}
		})

		It("client complains and exits", func() {
			session, err := gexec.Start(cliCmd, GinkgoWriter, GinkgoWriter)
			Expect(err).NotTo(HaveOccurred())
			Eventually(session.Out).Should(gbytes.Say("Please provide 2 integers\n"))
		})
	})

	Context("calling sum", func() {
		BeforeEach(func() {
			cliArgs = []string{"-action=sum", "4", "5"}
		})

		It("prints the sum of two numbers", func() {
			session, err := gexec.Start(cliCmd, GinkgoWriter, GinkgoWriter)
			Expect(err).NotTo(HaveOccurred())
			Eventually(session.Out).Should(gbytes.Say("9\n"))
		})
	})

	Context("calling prime", func() {
		BeforeEach(func() {
			cliArgs = []string{"-action=prime", "3", "15"}
		})

		It("prints all primes between two numbers", func() {
			session, err := gexec.Start(cliCmd, GinkgoWriter, GinkgoWriter)
			Expect(err).NotTo(HaveOccurred())
			Eventually(session.Out).Should(gbytes.Say("3\n5\n7\n11\n13\n"))
		})
	})
})
