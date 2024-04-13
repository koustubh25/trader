package trade

import (
	"log"

	"github.com/koustubh25/trader/pkg/algorithm"
)

func ExecuteTrade(algo algorithm.TradingAlgorithm) error {
	// Call the Buy and Sell methods
	if err := ExecuteBuy(algo); err != nil {
		return err
	}
	if err := ExecuteSell(algo); err != nil {
		return err
	}
	return nil
}

func ExecuteBuy(algo algorithm.TradingAlgorithm) error {
	isBuy, err := algo.IsBuy()
	if err != nil {
		return err
	}
	if isBuy {
		// do some trade
		log.Println("Buying")
	}
	return nil
}

func ExecuteSell(algo algorithm.TradingAlgorithm) error {
	isSell, err := algo.IsSell()
	if err != nil {
		return err
	}
	if isSell {
		// do some trade
		log.Println("Selling")
	}
	return nil
}
