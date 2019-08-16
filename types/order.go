package types

type Order struct {
	AccountCategory string `json:"accountCategory"`
	AccountId       string `json:"accountId"`
	AvgPrice        string `json:"avgPrice"`
	BaseAsset       string `json:"baseAsset"`
	BtmxCommission  string `json:"btmxCommission"`
	COID            string `json:"coid"`
	ErrorCode       string `json:"errorCode"`
	ExecID          string `json:"execId"`
	ExecInst        string `json:"execInst"`
	Fee             string `json:"fee"`
	FeeAsset        string `json:"feeAsset"`
	FilledQty       string `json:"filledQty"`
	Notional        string `json:"notional"`
	OrderPrice      string `json:"orderPrice"`
	OrderQty        string `json:"orderQty"`
	OrderType       string `json:"orderType"`
	QuoteAsset      string `json:"quoteAsset"`
	Side            string `json:"side"`
	Status          string `json:"status"`
	StopPrice       string `json:"stopPrice"`
	Symbol          string `json:"symbol"`
	Time            int64  `json:"time"`
	UserID          string `json:"userId"`
}

type CancelAllResult struct {
	Total    int64 `json:"total"`
	Canceled int64 `json:"canceled"`
}
