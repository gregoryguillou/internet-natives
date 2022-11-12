package main

import (
	"context"

	"github.com/blaqkube/internet-natives/idxr"
)

var (
	version = "dev"
)

func main() {
	app, err := idxr.NewApp(context.Background(), version)
	if err != nil {
		panic(err)
	}
	app.Start()
}
