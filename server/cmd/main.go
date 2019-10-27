package main

import (
	"flag"
	"log"
	"net"

	ysp "github.com/Callisto13/yummysushipajamas/pb"
	"github.com/Callisto13/yummysushipajamas/server"
	"google.golang.org/grpc"
)

func main() {
	var port string
	flag.StringVar(&port, "port", "1430", "port to start server on. default: 1430")
	flag.Parse()

	l, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := server.Server{}

	//TODO: certs, needs SANs to work in kube i think
	grpcServer := grpc.NewServer()
	ysp.RegisterBasicServer(grpcServer, &s)
	if err := grpcServer.Serve(l); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
