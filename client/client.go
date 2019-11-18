package client

import (
	"context"
	"fmt"
	"io"

	ysp "github.com/Callisto13/yummysushipajamas/pb"
)

type Client struct {
	conn ysp.BasicClient
	out  io.Writer
}

func NewClient(yspConn ysp.BasicClient, out io.Writer) Client {
	return Client{conn: yspConn, out: out}
}

func (c Client) Call(name string, n1, n2 int) error {
	var actions = map[string]func(int32, int32) error{
		"sum":   c.Sum,
		"prime": c.Prime,
	}

	return actions[name](int32(n1), int32(n2))
}

func (c Client) Sum(n1, n2 int32) error {
	response, err := c.conn.Sum(context.Background(), &ysp.Req{N1: n1, N2: n2})
	if err != nil {
		return fmt.Errorf("Failed calling Sum: %s", err)
	}

	fmt.Fprintln(c.out, response.Result)

	return nil
}

func (c Client) Prime(n1, n2 int32) error {
	stream, err := c.conn.Prime(context.Background(), &ysp.Req{N1: n1, N2: n2})
	if err != nil {
		return fmt.Errorf("Failed calling Prime: %s", err)
	}

	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("Failed streaming Prime: %s", err)
		}

		fmt.Fprintln(c.out, resp.Result)
	}

	return nil
}
