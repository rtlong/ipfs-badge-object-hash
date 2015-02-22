package main

import (
	"fmt"
	b58 "github.com/jbenet/go-base58"
	"github.com/jbenet/go-ipfs/merkledag"
	"os"
)

func main() {
	inputstr := os.Args[1]

	node := merkledag.Node{}
	node.Data = []byte(inputstr)

	if mh, err := node.Multihash(); err == nil {
		encoded := b58.Encode(mh)
		fmt.Printf("%s\n", encoded)
	} else {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
