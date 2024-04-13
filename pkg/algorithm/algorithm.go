package algorithm

type TradingAlgorithm interface {
	IsBuy() (bool, error)
	IsSell() (bool, error)
}
