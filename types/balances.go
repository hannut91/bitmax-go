package types

type Balance struct {
	AssetCode       string `json:"assetCode"`
	AssetName       string `json:"assetName"`
	TotalAmount     string `json:"totalAmount"`
	AvailableAmount string `json:"availableAmount"`
	InOrderAmount   string `json:"inOrderAmount"`
	BTCValue        string `json:"btcValue"`
}
