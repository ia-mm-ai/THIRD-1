package main

import (
	"context"
	"fmt"
	"os"

	"github.com/ava-labs/avalanchego/vms/rpcchainvm"
	"github.com/ia-mm-ai/THIRD-1/avalanchevm"
)

func main() {
	if err := rpcchainvm.Serve(context.Background(), &avalanchevm.VM{}); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
