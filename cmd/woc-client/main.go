package main

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/scottjbarr/woc"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage : woc-client command args\n")
		os.Exit(1)
	}

	cmd := os.Args[1]

	ctx := context.Background()

	if err := execute(ctx, cmd, os.Args[2:]); err != nil {
		panic(err)
	}
}

func execute(ctx context.Context, cmd string, args []string) error {
	c := woc.New()

	// fmt.Printf("args = %+v\n", args)

	switch cmd {
	case "bulktxs":
		if len(args) == 0 {
			return errors.New("Usage: woc-client bulktxs hash,hash,hash,...")
		}

		hashes := strings.Split(args[0], ",")

		return bulkTXs(ctx, c, hashes)

	case "broadcast":
		if len(args) == 0 {
			return errors.New("Usage: woc-client broadcast txhex")
		}

		data := args[0]

		return broadcast(ctx, c, data)
	}

	return errors.New("Unknown command")
}

func bulkTXs(ctx context.Context, c *woc.Client, hashes []string) error {
	txs, err := c.BulkTXs(ctx, hashes)
	if err != nil {
		return err
	}

	for _, tx := range txs {
		fmt.Printf("%+v\n", tx)
	}

	return nil
}

func broadcast(ctx context.Context, c *woc.Client, data string) error {
	b, err := hex.DecodeString(data)
	if err != nil {
		return err
	}

	if err := c.Broadcast(ctx, b); err != nil {
		return err
	}

	return nil
}
