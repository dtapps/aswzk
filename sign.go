package aswzk

import (
	"fmt"
	"go.dtapp.net/gomd5"
	"go.dtapp.net/gorequest"
	"go.dtapp.net/gostring"
	"sort"
)

// 签名
func sign(param gorequest.Params, apiKey string, timestamp string) string {
	var keys []string
	for k := range param {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	signStr := fmt.Sprintf("api_key=%s&", apiKey)
	for _, key := range keys {
		signStr += fmt.Sprintf("%s=%s&", key, gostring.ToString(param.Get(key)))
	}
	signStr += fmt.Sprintf("timestamp=%s", timestamp)
	return gomd5.ToLower(signStr)
}
