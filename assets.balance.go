package aswzk

import (
	"context"
	"go.dtapp.net/gojson"
	"go.dtapp.net/gorequest"
	"net/http"
)

type AssetsBalanceResponse struct {
	Code int    `json:"code"`
	Info string `json:"info"`
	Data struct {
		Balance float64 `json:"balance"` // 余额
	} `json:"data,omitempty"`
	Time    int    `json:"time"`
	TraceId string `json:"trace_id"`
}

type AssetsBalanceResult struct {
	Result AssetsBalanceResponse // 结果
	Body   []byte                // 内容
	Http   gorequest.Response    // 请求
}

func newAssetsBalanceResult(result AssetsBalanceResponse, body []byte, http gorequest.Response) *AssetsBalanceResult {
	return &AssetsBalanceResult{Result: result, Body: body, Http: http}
}

// AssetsBalance 余额查询
func (c *Client) AssetsBalance(ctx context.Context, notMustParams ...gorequest.Params) (*AssetsBalanceResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(ctx, "/assets/balance", params, http.MethodGet)
	if err != nil {
		return newAssetsBalanceResult(AssetsBalanceResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response AssetsBalanceResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newAssetsBalanceResult(response, request.ResponseBody, request), err
}
