package server_test

import (
	"errors"

	ysp "github.com/Callisto13/yummysushipajamas/pb"
	yspmock "github.com/Callisto13/yummysushipajamas/pb/mocks"
	"github.com/Callisto13/yummysushipajamas/server"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"golang.org/x/net/context"
)

var _ = Describe("Server", func() {
	var s server.Server

	BeforeEach(func() {
		s = server.Server{}
	})

	Describe("Sum", func() {
		It("returns the sum of two given numbers", func() {
			response, err := s.Sum(context.TODO(), &ysp.Req{N1: 1, N2: 2})
			Expect(err).NotTo(HaveOccurred())
			Expect(response.Result).To(Equal(int32(3)))
		})
	})

	Describe("Prime", func() {
		var mockStream *yspmock.MockBasic_PrimeServer

		BeforeEach(func() {
			mockStream = yspmock.NewMockBasic_PrimeServer(ctrl)
		})

		It("streams the primes between two given numbers as a stream", func() {
			mockStream.EXPECT().Send(&ysp.Resp{Result: 5}).Times(1)
			mockStream.EXPECT().Send(&ysp.Resp{Result: 7}).Times(1)
			Expect(s.Prime(&ysp.Req{N1: 5, N2: 8}, mockStream)).To(Succeed())
		})

		It("streams nothing if no prime is found", func() {
			mockStream.EXPECT().Send(gomock.Any).Times(0)
			Expect(s.Prime(&ysp.Req{N1: 0, N2: 1}, mockStream)).To(Succeed())
		})

		Context("when sending data to the stream fails", func() {
			It("returns the error", func() {
				mockStream.EXPECT().Send(&ysp.Resp{Result: 5}).Return(errors.New("boo")).Times(1)
				Expect(s.Prime(&ysp.Req{N1: 5, N2: 7}, mockStream)).NotTo(Succeed())
			})
		})
	})
})
