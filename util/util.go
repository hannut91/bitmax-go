package util

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/google/uuid"
)

func UUID() string {
	return strings.ReplaceAll(uuid.New().String(), "-", "")
}

func ComputeHmac256(message string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func HTTP(
	method,
	path string,
	headers map[string]string,
	params map[string]string,
	body []byte,
) (resp *http.Response, err error) {
	req, err := http.NewRequest(method, path, bytes.NewBuffer(body))
	if err != nil {
		return
	}

	values := url.Values{}
	for k, v := range params {
		values.Set(k, v)
	}
	req.URL.RawQuery = values.Encode()

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp, err = http.DefaultClient.Do(req)
	return
}

func HttpGet(
	u string,
	queryParams map[string]string,
	options map[string]string,
) (data []byte, err error) {
	values := url.Values{}
	for k, v := range queryParams {
		values.Set(k, v)
	}

	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return
	}
	req.URL.RawQuery = values.Encode()

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}

	defer res.Body.Close()
	data, err = ioutil.ReadAll(res.Body)
	return
}
