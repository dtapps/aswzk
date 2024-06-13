package aswzk

import (
	"context"
	"go.dtapp.net/gojson"
	"go.dtapp.net/gorequest"
	"go.opentelemetry.io/otel/codes"
	"net/http"
)

type PhoneBillOrderResponse struct {
	Code    int         `json:"code"`
	Info    string      `json:"info"`
	Data    interface{} `json:"data"`
	Time    int64       `json:"time"`
	TraceID string      `json:"trace_id"`
}

type PhoneBillOrderResult struct {
	Result PhoneBillOrderResponse // 结果
	Body   []byte                 // 内容
	Http   gorequest.Response     // 请求
}

func newPhoneBillOrderResult(result PhoneBillOrderResponse, body []byte, http gorequest.Response) *PhoneBillOrderResult {
	return &PhoneBillOrderResult{Result: result, Body: body, Http: http}
}

// PhoneBillOrder 话费订单下单
func (c *Client) PhoneBillOrder(ctx context.Context, notMustParams ...gorequest.Params) (*PhoneBillOrderResult, error) {

	// OpenTelemetry链路追踪
	ctx = c.TraceStartSpan(ctx, "phone_bill/order")
	defer c.TraceEndSpan()

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)

	// 请求
	request, err := c.request(ctx, "phone_bill/order", params, http.MethodPost)
	if err != nil {
		c.TraceSetStatus(codes.Error, err.Error())
		c.TraceRecordError(err)
		return newPhoneBillOrderResult(PhoneBillOrderResponse{}, request.ResponseBody, request), err
	}

	// 定义
	var response PhoneBillOrderResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	if err != nil {
		c.TraceSetStatus(codes.Error, err.Error())
		c.TraceRecordError(err)
	}
	return newPhoneBillOrderResult(response, request.ResponseBody, request), err
}
