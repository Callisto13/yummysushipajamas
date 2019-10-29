package server

import (
	"math"

	ysp "github.com/Callisto13/yummysushipajamas/pb"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
)

type Server struct {
	logger *logrus.Logger
}

func NewYSPServer(l *logrus.Logger) Server {
	return Server{logger: l}
}

func (s *Server) Sum(ctx context.Context, in *ysp.Req) (*ysp.Resp, error) {
	s.logger.WithFields(logrus.Fields{"N1": in.N1, "N2": in.N2}).Info("sum")

	return &ysp.Resp{Result: in.N1 + in.N2}, nil
}

func (s *Server) Prime(in *ysp.Req, stream ysp.Basic_PrimeServer) error {
	s.logger.WithFields(logrus.Fields{"N1": in.N1, "N2": in.N2}).Info("prime")

	for i := in.N1; i <= in.N2; i++ {
		if isPrime(int(i)) {
			if err := stream.Send(&ysp.Resp{Result: i}); err != nil {
				s.logger.WithFields(logrus.Fields{"data": i}).Error("Failed to stream data", err)
				return err
			}
		}
	}

	return nil
}

func isPrime(n int) bool {
	for i := 2; i <= int(math.Floor(math.Sqrt(float64(n)))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return n > 1
}
