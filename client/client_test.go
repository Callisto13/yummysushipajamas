package client_test

import (
	"errors"
	"io"

	"github.com/Callisto13/yummysushipajamas/client"
	ysp "github.com/Callisto13/yummysushipajamas/pb"
	yspmock "github.com/Callisto13/yummysushipajamas/pb/mocks"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Client", func() {
	var (
		mockYSP *yspmock.MockBasicClient
		c       client.Client
	)

	BeforeEach(func() {
		mockYSP = yspmock.NewMockBasicClient(ctrl)
		c = client.NewClient(mockYSP)
	})

	Describe("Sum", func() {
		It("retrieves the sum of two numbers from the server ", func() {
			mockYSP.EXPECT().Sum(gomock.Any(), gomock.Any()).Return(&ysp.Resp{Result: 42}, nil).Times(1)

			resp, err := c.Sum(10, 32)
			Expect(err).NotTo(HaveOccurred())
			Expect(resp).To(Equal(int32(42)))
		})

		Context("when the server call fails", func() {
			It("returns the error", func() {
				mockYSP.EXPECT().Sum(gomock.Any(), gomock.Any()).Return(nil, errors.New("boo")).Times(1)

				resp, err := c.Sum(10, 32)
				Expect(err).To(HaveOccurred())
				Expect(resp).To(BeZero())
			})
		})
	})

	Describe("Prime", func() {
		var stream *yspmock.MockBasic_PrimeClient

		BeforeEach(func() {
			stream = yspmock.NewMockBasic_PrimeClient(ctrl)
		})

		It("retrieves a stream of prime numbers from the server", func() {
			mockYSP.EXPECT().Prime(gomock.Any(), gomock.Any()).Return(stream, nil).Times(1)
			stream.EXPECT().Recv().Return(&ysp.Resp{Result: 5}, nil).Times(1)
			stream.EXPECT().Recv().Return(nil, io.EOF).Times(1)

			resp, err := c.Prime(5, 5)
			Expect(err).NotTo(HaveOccurred())
			Expect(resp).To(Equal([]int32{5}))
		})

		Context("when the server call fails", func() {
			It("returns the error", func() {
				mockYSP.EXPECT().Prime(gomock.Any(), gomock.Any()).Return(nil, errors.New("boo")).Times(1)

				resp, err := c.Prime(10, 32)
				Expect(err).To(HaveOccurred())
				Expect(resp).To(BeNil())
			})
		})

		Context("when the stream call fails", func() {
			It("returns the error", func() {
				mockYSP.EXPECT().Prime(gomock.Any(), gomock.Any()).Return(stream, nil).Times(1)
				stream.EXPECT().Recv().Return(nil, errors.New("boo")).Times(1)

				resp, err := c.Prime(10, 32)
				Expect(err).To(HaveOccurred())
				Expect(resp).To(BeNil())
			})
		})
	})
})
