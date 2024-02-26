package aswzk

import (
	"context"
	"errors"
	"fmt"
	"go.dtapp.net/gojson"
	"go.dtapp.net/gorequest"
	"go.dtapp.net/gotime"
	"net/url"
)

type NotifyUrlParams struct {
	NotifyUrl string `json:"notify_url"` // 回调地址
	UserID    string `json:"user_id"`    // 用户编号
	ApiKey    string `json:"api_key"`    // 秘钥
}

type NotifyUrlRequestBody struct {
	RechargeAccount string  `json:"recharge_account"` // 充值账号
	RechargeMoney   float64 `json:"recharge_money"`   // 充值金额 单位：元
	RechargeType    string  `json:"recharge_type"`    // 充值类型
	RechargeReason  string  `json:"recharge_reason"`  // 充值失败原因 只有充值失败才会返回的数据内容
	OrderID         string  `json:"order_id"`         // 订单编号
	OrderNo         string  `json:"order_no"`         // 商户订单编号
	OrderRemarks    string  `json:"order_remarks"`    // 订单备注
	OrderStatus     string  `json:"order_status"`     // 订单状态 SUCCESS=充值成功，FAILURE=充值失败，RECHARGE=充值中
	OrderCost       float64 `json:"order_cost"`       // 订单成本价 单位：元，只有充值成功才会返回的数据内容
}

// NotifyUrl 通知回调地址
func (c *Client) NotifyUrl(ctx context.Context, params NotifyUrlParams, requestBody NotifyUrlRequestBody) error {

	// 验证回调地址
	_, err := url.ParseRequestURI(params.NotifyUrl)
	if err != nil {
		return err
	}

	// 检查密钥
	if params.ApiKey != c.config.apiKey {
		return errors.New("api_key is not match")
	}

	// 获取时间戳
	XTimestamp := fmt.Sprintf("%v", gotime.Current().Timestamp())

	// 签名
	requestParams := gorequest.NewParams()
	requestParams.Set("recharge_account", requestBody.RechargeAccount)
	requestParams.Set("recharge_money", requestBody.RechargeMoney)
	requestParams.Set("recharge_type", requestBody.RechargeType)
	requestParams.Set("recharge_reason", requestBody.RechargeReason)
	requestParams.Set("order_id", requestBody.OrderID)
	requestParams.Set("order_no", requestBody.OrderNo)
	requestParams.Set("order_remarks", requestBody.OrderRemarks)
	requestParams.Set("order_status", requestBody.OrderStatus)
	requestParams.Set("order_cost", requestBody.OrderCost)

	// 签名
	XSign := sign(requestParams, params.ApiKey, XTimestamp)

	// 创建请求
	client := gorequest.NewHttp()

	// 设置请求地址
	client.SetUri(params.NotifyUrl)

	// 设置格式
	client.SetContentTypeJson()

	// 设置参数
	client.SetParams(requestParams)

	// 添加请求头
	client.SetHeader("X-Timestamp", XTimestamp)
	client.SetHeader("X-Sign", XSign)

	// 发起请求
	request, err := client.Post(ctx)
	if err != nil {
		return err
	}

	// 定义
	var response struct {
		Code int `json:"code"` // 状态码
	}
	err = gojson.Unmarshal(request.ResponseBody, &response)
	if err != nil {
		return err
	}

	// 日志
	if c.gormLog.status {
		go c.gormLog.client.Middleware(ctx, request)
	}
	if c.mongoLog.status {
		go c.mongoLog.client.Middleware(ctx, request)
	}

	if response.Code == CodeSuccess {
		return nil
	}

	return errors.New(fmt.Sprintf("code: %v", response.Code))
}
