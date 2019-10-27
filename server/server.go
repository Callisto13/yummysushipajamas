package server

import (
	ysp "github.com/Callisto13/yummysushipajamas/pb"
	"golang.org/x/net/context"
)

type Server struct{}

func (s *Server) Sum(ctx context.Context, in *ysp.Req) (*ysp.Resp, error) {
	return &ysp.Resp{Result: in.N1 + in.N2}, nil
}

func (s *Server) Prime(in *ysp.Req, stream ysp.Basic_PrimeServer) error {
	for i := in.N1; i <= in.N2; i++ {
		if isPrime(int(i)) {
			if err := stream.Send(&ysp.Resp{Result: i}); err != nil {
				return err
			}
		}
	}

	return nil
}

func isPrime(n int) bool {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return n > 1
}
