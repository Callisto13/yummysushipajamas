package main

import (
	"flag"
	"fmt"
	"net"
	"os"

	ysp "github.com/Callisto13/yummysushipajamas/pb"
	"github.com/Callisto13/yummysushipajamas/server"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {
	var (
		port     string
		logLevel string
	)

	flag.StringVar(&port, "port", "1430", "port to start server on")
	flag.StringVar(&logLevel, "log-level", "info", "log level")
	flag.Parse()

	logger := setupLogger(logLevel)

	l, err := net.Listen("tcp", ":"+port)
	if err != nil {
		logger.Error(fmt.Sprintf("Failed to listen on localhost:%s", port), err)
		os.Exit(1)
	}

	s := server.NewYSPServer(logger)

	//TODO: certs, needs SANs to work in kube i think
	grpcServer := grpc.NewServer()
	ysp.RegisterBasicServer(grpcServer, &s)

	logger.Info("starting YSP Service on port " + port)
	if err := grpcServer.Serve(l); err != nil {
		logger.Error(fmt.Sprintf("Failed to listen on localhost:%s", port), err)
		os.Exit(1)
	}
}

func setupLogger(level string) *logrus.Logger {
	logger := logrus.New()
	logger.Out = os.Stdout

	var logLevel logrus.Level
	switch level {
	case "trace":
		logLevel = logrus.TraceLevel
	case "debug":
		logLevel = logrus.DebugLevel
	case "info":
		logLevel = logrus.InfoLevel
	case "warn":
		logLevel = logrus.WarnLevel
	case "error":
		logLevel = logrus.ErrorLevel
	case "fatal":
		logLevel = logrus.FatalLevel
	case "panic":
		logLevel = logrus.PanicLevel
	default:
		logLevel = logrus.InfoLevel
	}

	logger.SetLevel(logLevel)

	return logger
}
