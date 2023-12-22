package aswzk

import (
	"fmt"
	"go.dtapp.net/gomd5"
	"go.dtapp.net/gorequest"
	"go.dtapp.net/gostring"
	"sort"
)

// 签名
func (c *Client) sign(param gorequest.Params) string {
	var keys []string
	for k := range param {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	signStr := fmt.Sprintf("%s", c.GetUserID())
	for _, key := range keys {
		signStr += fmt.Sprintf("%s%s", key, gostring.ToString(param.Get(key)))
	}
	signStr += fmt.Sprintf("%s", c.GetApiKey())
	return gomd5.ToLower(signStr)
}
