package lib

import (
	"strconv"

	"github.com/World-of-Cryptopups/atomicassets-go"
)

var a = atomicassets.New()

// calculates the total assets a user has
func GetAssetStats(wallet string) (int, error) {
	q, err := a.GetAccountCollection(wallet, COLLECTION_NAME)
	if err != nil {
		return 0, err
	}

	total := 0
	for _, v := range q.Data.Schemas {
		value, _ := strconv.Atoi(v.Assets)

		total += value
	}

	return total, nil
}
