package main

import (
	"context"
	"fmt"

	"github.com/blaqkube/internet-natives/idxr"
)

func main() {
	ctx := context.Background()
	url := "http://localhost:8545"
	c, err := idxr.NewClient(ctx, url)
	if err != nil {
		fmt.Println(err)
	}
	n, err := c.BlockNumber(ctx)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("current block is", n)
}
