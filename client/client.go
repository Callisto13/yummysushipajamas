package client

import (
	"context"
	"fmt"
	"io"

	ysp "github.com/Callisto13/yummysushipajamas/pb"
)

type Client struct {
	conn ysp.BasicClient
}

func NewClient(yspConn ysp.BasicClient) Client {
	return Client{conn: yspConn}
}

func (c *Client) Call(name string, n1, n2 int) (interface{}, error) {
	var actions = map[string]func(int32, int32) (interface{}, error){
		"sum":   c.Sum,
		"prime": c.Prime,
	}

	return actions[name](int32(n1), int32(n2))
}

func (c *Client) Sum(n1, n2 int32) (interface{}, error) {
	response, err := c.conn.Sum(context.Background(), &ysp.Req{N1: n1, N2: n2})
	if err != nil {
		return nil, fmt.Errorf("Failed calling Sum: %s", err)
	}

	return response.Result, nil
}

func (c *Client) Prime(n1, n2 int32) (interface{}, error) {
	stream, err := c.conn.Prime(context.Background(), &ysp.Req{N1: n1, N2: n2})
	if err != nil {
		return nil, fmt.Errorf("Failed calling Prime: %s", err)
	}

	out := []int32{}

	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("Failed streaming Prime: %s", err)
		}

		out = append(out, resp.Result)
	}

	return out, nil
}
