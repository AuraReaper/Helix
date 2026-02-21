package client

import (
	"context"
	"net"

	"github.com/AuraReaper/helix/protocol"
)

type Client struct {
	conn net.Conn
}

type Options struct{}

func New(endpoint string, opts Options) (*Client, error) {
	conn, err := net.Dial("tcp", endpoint)
	if err != nil {
		return nil, err
	}

	return &Client{
		conn: conn,
	}, nil
}

func (c *Client) Close() error {
	return c.conn.Close()
}

func (c *Client) Set(ctx context.Context, key, value []byte) (any, error) {
	cmd := &protocol.CommandSet{
		Key:   key,
		Value: value,
	}
	return nil, nil
}
