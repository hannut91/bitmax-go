package bitmax

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/hannut91/bitmax-go/types"
	"github.com/hannut91/bitmax-go/util"
)

const (
	Version = "v1"
	URL     = "https://bitmax.io"

	SuccessCode = 0
)

type Client struct {
	Key          string
	Secret       string
	HTTP         *http.Client
	AccountGroup int
}

func CreateClient(key, secret string) *Client {
	return &Client{Key: key, Secret: secret, HTTP: new(http.Client)}
}

func (c *Client) Products() (products []*types.Product, err error) {
	const apiPath = "products"
	apiURL := fmt.Sprintf("%s/api/%s/%s", URL, Version, apiPath)
	resp, err := util.HTTP("GET", apiURL, nil, nil, nil)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &products)
	return
}

func (c *Client) UserInfo() (userInfo *types.UserInfo, err error) {
	const apiPath = "user/info"
	apiURL := fmt.Sprintf("%s/api/%s/%s", URL, Version, apiPath)
	var res struct {
		types.Response
		*types.UserInfo
	}

	err = c.httpAuthGET(apiPath, apiURL, &res)
	if err != nil {
		return
	}

	if res.Code != SuccessCode {
		err = errors.New(res.Message)
		return
	}

	userInfo = res.UserInfo
	return
}

func (c *Client) Balances() (balances []*types.Balance, err error) {
	const apiPath = "balance"
	apiURL := fmt.Sprintf("%s/%d/api/%s/%s", URL, c.AccountGroup, Version,
		apiPath)
	var res struct {
		types.Response
		Balances []*types.Balance `json:"data"`
	}

	err = c.httpAuthGET(apiPath, apiURL, &res)
	if err != nil {
		return
	}

	if res.Code != SuccessCode {
		err = errors.New(res.Message)
		return
	}

	balances = res.Balances
	return
}

func (c *Client) Balance(
	symbol string,
) (balance *types.Balance, err error) {
	const apiPath = "balance"
	apiURL := fmt.Sprintf("%s/%d/api/%s/%s/%s", URL, c.AccountGroup, Version,
		apiPath, symbol)
	var res struct {
		types.Response
		Balance *types.Balance `json:"data"`
	}

	err = c.httpAuthGET(apiPath, apiURL, &res)
	if err != nil {
		return
	}

	if res.Code != SuccessCode {
		err = errors.New(res.Message)
		return
	}

	balance = res.Balance
	return
}

func (c *Client) Quote(symbol string) (quote *types.Quote, err error) {
	const apiPath = "quote"
	apiURL := fmt.Sprintf("%s/api/%s/%s", URL, Version, apiPath)
	var res struct {
		types.Response
		*types.Quote
	}

	queryParams := map[string]string{
		"symbol": symbol,
	}
	data, err := util.HttpGet(apiURL, queryParams, nil)
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &res)
	if err != nil {
		return
	}

	if res.Code != SuccessCode {
		err = errors.New(res.Message)
		return
	}

	quote = res.Quote
	return
}

func (c *Client) Order(
	side, symbol, orderPrice, amount string,
) (orderResponse *types.OrderResponse, err error) {
	const apiPath = "order"
	apiURL := fmt.Sprintf("%s/%d/api/%s/%s", URL, c.AccountGroup, Version,
		apiPath)
	var res struct {
		types.Response
		OrderResponse *types.OrderResponse `json:"data"`
	}

	coid := util.UUID()
	timestamp := time.Now().UnixNano() / int64(time.Millisecond)

	body := map[string]interface{}{
		"coid":       coid,
		"time":       timestamp,
		"symbol":     symbol,
		"orderPrice": orderPrice,
		"orderQty":   amount,
		"orderType":  "limit",
		"side":       side,
	}
	bytes, err := json.Marshal(body)
	if err != nil {
		return
	}

	headers := map[string]string{
		"x-auth-coid":      coid,
		"x-auth-timestamp": fmt.Sprint(timestamp),
		"Content-Type":     "application/json",
	}
	err = c.httpAuthPOST(apiPath, apiURL, bytes, headers, &res)
	if err != nil {
		return
	}

	if res.Code != SuccessCode {
		err = errors.New(res.Message)
		return
	}

	orderResponse = res.OrderResponse
	return
}

func (c *Client) Orders() (orders []*types.Order, err error) {
	const apiPath = "order/open"
	apiURL := fmt.Sprintf("%s/%d/api/%s/%s", URL, c.AccountGroup, Version,
		apiPath)
	var res struct {
		types.Response
		Orders []*types.Order `json:"data"`
	}

	err = c.httpAuthGET(apiPath, apiURL, &res)
	if err != nil {
		return
	}

	if res.Code != SuccessCode {
		err = errors.New(res.Message)
		return
	}

	orders = res.Orders
	return
}

func (c *Client) CancelOrder(
	symbol, orderCOID string,
) (orderResponse *types.OrderResponse, err error) {
	const apiPath = "order"
	apiURL := fmt.Sprintf("%s/%d/api/%s/%s", URL, c.AccountGroup, Version,
		apiPath)
	coid := util.UUID()
	timestamp := time.Now().UnixNano() / int64(time.Millisecond)
	msg := fmt.Sprint(timestamp) + "+" + apiPath + "+" + coid
	signMsg := util.ComputeHmac256(msg, c.Secret)

	headers := map[string]string{
		"x-auth-key":       c.Key,
		"x-auth-signature": signMsg,
		"x-auth-coid":      coid,
		"x-auth-timestamp": fmt.Sprint(timestamp),
		"Content-Type":     "application/json",
	}

	body := map[string]interface{}{
		"coid":     coid,
		"origCoid": orderCOID,
		"time":     timestamp,
		"symbol":   symbol,
	}
	bytes, err := json.Marshal(body)
	if err != nil {
		return
	}
	resp, err := util.HTTP("DELETE", apiURL, headers, nil, bytes)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	var res struct {
		types.Response
		OrderResponse *types.OrderResponse `json:"data"`
	}
	err = json.Unmarshal(data, &res)
	if err != nil {
		return
	}

	if res.Code != SuccessCode {
		err = errors.New(res.Message)
		return
	}

	orderResponse = res.OrderResponse
	return
}

func (c *Client) CancelOrderAll(
	symbol string,
	side string,
) (result *types.CancelAllResult, err error) {
	const apiPath = "order/all"
	apiURL := fmt.Sprintf("%s/%d/api/%s/%s", URL, c.AccountGroup, Version,
		apiPath)
	timestamp := time.Now().UnixNano() / int64(time.Millisecond)
	msg := fmt.Sprint(timestamp) + "+" + apiPath
	signMsg := util.ComputeHmac256(msg, c.Secret)

	headers := map[string]string{
		"x-auth-key":       c.Key,
		"x-auth-signature": signMsg,
		"x-auth-timestamp": fmt.Sprint(timestamp),
		"Content-Type":     "application/json",
	}

	params := map[string]string{
		"symbol": symbol,
	}

	if side != "" {
		params["side"] = side
	}

	resp, err := util.HTTP("DELETE", apiURL, headers, params, nil)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	var res struct {
		types.Response
		CancelAllResult *types.CancelAllResult `json:"data"`
	}
	err = json.Unmarshal(data, &res)
	if err != nil {
		return
	}

	if res.Code != SuccessCode {
		err = errors.New(res.Message)
		return
	}

	result = res.CancelAllResult
	return
}
func (c *Client) httpAuthPOST(
	apiPath,
	apiURL string,
	body []byte,
	headers map[string]string,
	res interface{},
) (err error) {
	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(body))
	if err != nil {
		return
	}

	msg := headers["x-auth-timestamp"] + "+" + apiPath + "+" + headers["x-auth-coid"]
	signMsg := util.ComputeHmac256(msg, c.Secret)

	req.Header.Set("x-auth-key", c.Key)
	req.Header.Set("x-auth-signature", signMsg)

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp, err := c.HTTP.Do(req)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &res)
	return
}

func (c *Client) httpAuthGET(apiPath, apiURL string, res interface{}) (err error) {
	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return
	}

	timeStamp := time.Now().UnixNano() / int64(time.Millisecond)
	msg := fmt.Sprint(timeStamp) + "+" + apiPath
	signMsg := util.ComputeHmac256(msg, c.Secret)

	req.Header.Set("x-auth-key", c.Key)
	req.Header.Set("x-auth-signature", signMsg)
	req.Header.Set("x-auth-timestamp", fmt.Sprint(timeStamp))

	resp, err := c.HTTP.Do(req)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &res)
	return
}
