package lib

import (
	"strconv"

	"github.com/World-of-Cryptopups/atomicassets-go"
)

func GetAssetStats(wallet string) (*atomicassets.AccountCollectionProps, error) {
	a := atomicassets.New()

	return a.GetAccountCollection(wallet, COLLECTION_NAME)
}

// calculates the total assets a user has
func CountAssetStats(wallet string) (int, error) {
	q, err := GetAssetStats(wallet)
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
