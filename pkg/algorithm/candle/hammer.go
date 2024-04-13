package candle

import (
	"log/slog"
	"strconv"

	"github.com/koustubh25/trader/pkg/algorithm"
	"github.com/koustubh25/trader/pkg/util"
)

type HammerCandleAgorithm struct {
	Candle
}

func (h *HammerCandleAgorithm) IsBuy() (bool, error) {
	open, _ := strconv.ParseFloat(h.Open, 64)
	close, _ := strconv.ParseFloat(h.Close, 64)
	high, _ := strconv.ParseFloat(h.High, 64)
	low, _ := strconv.ParseFloat(h.Low, 64)

	body := high - low
	upperShadow := high - util.Max(open, close)
	lowerShadow := util.Min(open, close) - low
	slog.Debug("HammerCandleAgorithm: IsBuy", "body", body, "upperShadow", upperShadow, "lowerShadow", lowerShadow)

	// Check if it's a hammer
	return body > upperShadow && lowerShadow > body*2, nil
}

func (h *HammerCandleAgorithm) IsSell() (bool, error) {
	open, _ := strconv.ParseFloat(h.Open, 64)
	close, _ := strconv.ParseFloat(h.Close, 64)
	high, _ := strconv.ParseFloat(h.High, 64)
	low, _ := strconv.ParseFloat(h.Low, 64)

	body := high - low
	upperShadow := high - util.Max(open, close)
	lowerShadow := util.Min(open, close) - low

	slog.Debug("HammerCandleAgorithm: IsSell", "body", body, "upperShadow", upperShadow, "lowerShadow", lowerShadow)

	// Check if it's an inverted hammer
	return body > lowerShadow && upperShadow > body*2, nil
}

// Ensure HammerCandle implements TradingAlgorithm
var _ algorithm.TradingAlgorithm = (*HammerCandleAgorithm)(nil)
