package idxr

import (
	"context"

	"github.com/ethereum/go-ethereum/ethclient"
)

type Client struct {
	*ethclient.Client
}

func NewClient(ctx context.Context, rawurl string) (*Client, error) {
	c, err := ethclient.DialContext(ctx, rawurl)
	if err != nil {
		return nil, err
	}
	return &Client{c}, nil
}
