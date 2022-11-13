package main

import (
	"context"

	"github.com/blaqkube/internet-natives/api"
)

var (
	version = "dev"
)

func main() {
	app, err := api.NewApp(context.Background(), version)
	if err != nil {
		panic(err)
	}
	app.Start()
}
