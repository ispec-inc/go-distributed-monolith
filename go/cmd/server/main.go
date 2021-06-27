package main

import (
	"context"

	"github.com/ispec-inc/monorepo/go/runner/server"
)

func main() {
	server, sclnup, err := server.New()
	if err != nil {
		panic(err)
	}
	defer sclnup()

	ctx := context.Background()
	server.Run(ctx)
}
