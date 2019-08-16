package types

type Quote struct {
	Symbol   string `json:"symbol"`
	BidPrice string `json:"bidPrice"`
	BidSize  string `json:"bidSize"`
	AskPrice string `json:"askPrice"`
	AskSize  string `json:"askSize"`
}
