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

// NotifyUrl 通知回调地址
func (c *Client) NotifyUrl(ctx context.Context, params NotifyUrlParams, param gorequest.Params) error {

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
	XSign := sign(param, params.ApiKey, XTimestamp)

	// 创建请求
	client := gorequest.NewHttp()

	// 设置请求地址
	client.SetUri(params.NotifyUrl)

	// 设置格式
	client.SetContentTypeJson()

	// 设置参数
	client.SetParams(param)

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
