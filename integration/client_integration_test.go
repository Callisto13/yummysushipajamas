package integration_test

import (
	"bytes"
	"fmt"
	"os/exec"

	. "github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/config"
	. "github.com/onsi/gomega"
)

var _ = Describe("YSP Client", func() {
	var (
		cliCmd  *exec.Cmd
		cliOut  *bytes.Buffer
		cliArgs []string
		cliEnv  []string
	)

	JustBeforeEach(func() {
		cliCmd = exec.Command(cliBin, cliArgs...)
		cliCmd.Env = cliEnv
		cliOut = bytes.NewBuffer([]byte{})
		cliCmd.Stdout = cliOut
		cliCmd.Stderr = cliOut
	})

	BeforeEach(func() {
		cliEnv = []string{fmt.Sprintf("SERVICE_ADDR=localhost:1430%d", config.GinkgoConfig.ParallelNode)}
	})

	Context("when SERVICE_ADDR is not set for the client", func() {
		BeforeEach(func() {
			cliEnv = []string{}
		})

		It("client complains and exits", func() {
			err := cliCmd.Run()
			Expect(err).To(HaveOccurred())
			Expect(cliOut.String()).To(Equal("SERVICE_ADDR not set\n"))
		})
	})

	Context("no action is chosen", func() {
		BeforeEach(func() {
			cliArgs = []string{}
		})

		It("client prints the help and exits", func() {
			err := cliCmd.Run()
			Expect(err).To(HaveOccurred())
			Expect(cliOut.String()).To(Equal("usage: -action=<sum|prime> n1 n2\n"))
		})
	})

	Context("when fewer than two numbers are provided", func() {
		BeforeEach(func() {
			cliArgs = []string{"-action=sum"}
		})

		It("client complains and exits", func() {
			err := cliCmd.Run()
			Expect(err).To(HaveOccurred())
			Expect(cliOut.String()).To(Equal("Please provide 2 integers\n"))
		})
	})

	Context("calling sum", func() {
		BeforeEach(func() {
			cliArgs = []string{"-action=sum", "4", "5"}
		})

		It("prints the sum of two numbers", func() {
			err := cliCmd.Run()
			Expect(err).NotTo(HaveOccurred())
			Expect(cliOut.String()).To(Equal("9\n"))
		})
	})

	Context("calling prime", func() {
		BeforeEach(func() {
			cliArgs = []string{"-action=prime", "3", "15"}
		})

		It("prints all primes between two numbers", func() {
			err := cliCmd.Run()
			Expect(err).NotTo(HaveOccurred())
			Expect(cliOut.String()).To(Equal("[3 5 7 11 13]\n"))
		})
	})
})
