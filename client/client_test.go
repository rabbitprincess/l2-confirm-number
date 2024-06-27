package client

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBlockNumber(t *testing.T) {
	for _, test := range []struct {
		name string
		url  string
	}{
		{"arbitrum", "https://arbitrum-one-rpc.publicnode.com"},
		{"base", "https://base-rpc.publicnode.com"},
		{"blast", "https://rpc.blast.io"},
		{"optimism", "https://optimism-rpc.publicnode.com"},
		{"manta", "https://pacific-rpc.manta.network/http"},
		{"metal", "https://rpc.metall2.com"},
		{"scroll", "https://rpc.ankr.com/scroll"},
	} {
		// var maxConfirmTime uint64 = 0
		// var maxConfirmNumber uint64 = 0
		c, err := NewClient(test.url)
		require.NoError(t, err)
		ctx := context.Background()

		var (
			blockNumberLatest    uint64
			blockNumberSafe      uint64
			blockNumberFinalized uint64
			blockTimeLatest      uint64
			blockTimeSafe        uint64
			blockTimeFinalized   uint64
		)

		latest, err := c.GetBestBlock(ctx)
		if err == nil {
			blockNumberLatest = latest.Number.Uint64()
			blockTimeLatest = latest.Time
		}
		safe, err := c.GetSafeBlock(ctx)
		if err == nil {
			blockNumberSafe = safe.Number.Uint64()
			blockTimeSafe = safe.Time
		}
		finalized, err := c.GetFinalizedBlock(ctx)
		if err == nil {
			blockNumberFinalized = finalized.Number.Uint64()
			blockTimeFinalized = finalized.Time
		}

		fmt.Printf("%s\n", test.name)
		fmt.Printf("\t%-20s block %-12vtimestamp %v\n", "latest", blockNumberLatest, blockTimeLatest)
		if blockNumberSafe != 0 {
			fmt.Printf("\t%-20s block %-12vtimestamp %v\n", "safe", blockNumberSafe, blockTimeSafe)
		}
		fmt.Printf("\t%-20s block %-12vtimestamp %v\n", "finalized", blockNumberFinalized, blockTimeFinalized)
		fmt.Printf("\t%-20s block %-12vtimestamp %v sec\n", "latest-finalized", blockNumberLatest-blockNumberFinalized, blockTimeLatest-blockTimeFinalized)
		if blockNumberSafe != 0 {
			fmt.Printf("\t%-20s block %-12vtimestamp %v sec\n", "latest-safe", blockNumberLatest-blockNumberSafe, blockTimeLatest-blockTimeSafe)
		}
	}
}
