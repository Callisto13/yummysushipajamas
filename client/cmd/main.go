package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"google.golang.org/grpc"

	"github.com/Callisto13/yummysushipajamas/client"
	ysp "github.com/Callisto13/yummysushipajamas/pb"
)

func main() {
	serviceAddress := os.Getenv("SERVICE_ADDR")
	if serviceAddress == "" {
		fmt.Println("SERVICE_ADDR not set")
		os.Exit(1)
	}

	var action string
	flag.StringVar(&action, "action", "", "sum or prime")
	flag.Parse()

	if action == "" {
		fmt.Println("usage: -action=<sum|prime> n1 n2")
		os.Exit(1)
	}

	n1, n2, err := parseArgs(flag.Args())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	conn, err := grpc.Dial(serviceAddress, grpc.WithInsecure())
	if err != nil {
		fmt.Printf("Failed to connect to grpc server: %s\n", err)
		os.Exit(1)
	}
	defer conn.Close()

	c := client.NewClient(ysp.NewBasicClient(conn))
	response, err := c.Call(action, n1, n2)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(response)
}

func parseArgs(args []string) (int, int, error) {
	if len(args) < 2 {
		return 0, 0, fmt.Errorf("Please provide 2 integers")
	}

	n1, err := strconv.Atoi(args[0])
	if err != nil {
		return 0, 0, fmt.Errorf("Could not parse first arg: %s", err)
	}

	n2, err := strconv.Atoi(args[1])
	if err != nil {
		return 0, 0, fmt.Errorf("Could not parse second arg: %s", err)
	}

	return n1, n2, nil
}
