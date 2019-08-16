package types

type Product struct {
	Symbol     string `json:"symbol"`
	BaseAsset  string `json:"baseAsset"`
	QuoteAsset string `json:"quoteAsset"`
	PriceScale int    `json:"priceScale"`
	QtyScale   int    `json:"qtyScale"`
	Status     string `json:"status"`
}
