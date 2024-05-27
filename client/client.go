package client

import (
	"context"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Client struct {
	client *ethclient.Client
}

func NewClient(url string) (*Client, error) {
	client, err := ethclient.Dial(url)
	if err != nil {
		return nil, err
	}

	return &Client{
		client: client,
	}, nil
}

func (c *Client) GetBestBlock(ctx context.Context) (*types.Header, error) {
	return c.getBlockHeader(ctx, "latest")
}

func (c *Client) GetFinalizedBlock(ctx context.Context) (*types.Header, error) {
	return c.getBlockHeader(ctx, "finalized")
}

func (c *Client) GetSafeBlock(ctx context.Context) (*types.Header, error) {
	return c.getBlockHeader(ctx, "safe")
}

func (c *Client) getBlockHeader(ctx context.Context, number string) (*types.Header, error) {
	head := &types.Header{}
	err := c.client.Client().CallContext(ctx, head, "eth_getBlockByNumber", number, false)
	if err != nil {
		return nil, err
	}

	return head, nil
}
