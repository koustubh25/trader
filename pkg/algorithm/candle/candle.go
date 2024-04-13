package candle

type Candle struct {
	Start     string `json:"start"`
	High      string `json:"high"`
	Low       string `json:"low"`
	Open      string `json:"open"`
	Close     string `json:"close"`
	Volume    string `json:"volume"`
	ProductID string `json:"product_id"`
}

type Event struct {
	Type    string   `json:"type"`
	Candles []Candle `json:"candles"`
}

type CandlesMessage struct {
	Channel     string  `json:"channel"`
	ClientID    string  `json:"client_id"`
	Timestamp   string  `json:"timestamp"`
	SequenceNum int     `json:"sequence_num"`
	Events      []Event `json:"events"`
}
