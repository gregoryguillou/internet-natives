package api

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type Capture struct {
	BlockNumber uint64 `dynamodbav:"blockNumber" json:"blockNumber"`
	Network     string `dynamodbav:"network" json:"network"`
}

type CapturedBlock struct {
	BlockNumber  uint64
	BlockHash    common.Hash
	Transactions []CapturedTransaction
}

type CapturedTransaction struct {
	Hash     common.Hash
	Type     uint8
	From     common.Address
	To       *common.Address
	Value    *big.Int
	Data     []byte
	Receipts *int
}

func (c *Client) Current(ctx context.Context) (uint64, error) {
	return c.BlockNumber(ctx)
}

func (c *Client) BlockNumber(ctx context.Context) (uint64, error) {
	return c.Client.BlockNumber(ctx)
}

func (c *Client) CaptureBlock(ctx context.Context, blockNumber uint64) (*CapturedBlock, error) {
	b, err := c.BlockByNumber(ctx, big.NewInt(int64(blockNumber)))
	if err != nil {
		return nil, err
	}
	chainID, err := c.NetworkID(ctx)
	if err != nil {
		return nil, err
	}
	txns := b.Transactions()
	capturedTransactions := []CapturedTransaction{}
	for _, tx := range txns {
		msg, err := tx.AsMessage(types.NewEIP155Signer(chainID), b.BaseFee())
		if err != nil {
			return nil, err
		}
		capturedTransactions = append(
			capturedTransactions,
			CapturedTransaction{
				Hash:  tx.Hash(),
				Type:  tx.Type(),
				From:  msg.From(),
				To:    tx.To(),
				Value: tx.Value(),
				Data:  tx.Data(),
			},
		)
	}
	return &CapturedBlock{
		BlockNumber:  blockNumber,
		BlockHash:    b.Hash(),
		Transactions: capturedTransactions,
	}, nil
}

func (c *Client) CaptureTransactionEvents(ctx context.Context, hash common.Hash) ([]*types.Log, error) {
	receipts, err := c.TransactionReceipt(ctx, hash)
	if err != nil {
		return nil, err
	}
	return receipts.Logs, nil
}
