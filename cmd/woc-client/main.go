package main

import (
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

	if err := execute(cmd, os.Args[2:]); err != nil {
		panic(err)
	}
}

func execute(cmd string, args []string) error {
	c := woc.New()

	// fmt.Printf("args = %+v\n", args)

	switch cmd {
	case "bulktxs":
		if len(args) == 0 {
			return errors.New("Usage: woc-client bulktxs hash,hash,hash,...")
		}

		hashes := strings.Split(args[0], ",")

		return bulkTXs(c, hashes)

	case "broadcast":
		if len(args) == 0 {
			return errors.New("Usage: woc-client broadcast txhex")
		}

		data := args[0]

		return broadcast(c, data)
	}

	return errors.New("Unknown command")
}

func bulkTXs(c *woc.Client, hashes []string) error {
	txs, err := c.BulkTXs(hashes)
	if err != nil {
		return err
	}

	for _, tx := range txs {
		fmt.Printf("%+v\n", tx)
	}

	return nil
}

func broadcast(c *woc.Client, data string) error {
	b, err := hex.DecodeString(data)
	if err != nil {
		return err
	}

	if err := c.Broadcast(b); err != nil {
		return err
	}

	return nil
}
