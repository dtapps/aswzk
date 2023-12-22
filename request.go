package aswzk

import (
	"context"
	"fmt"
	"go.dtapp.net/gorequest"
	"go.dtapp.net/gotime"
)

// 请求接口
func (c *Client) request(ctx context.Context, url string, param gorequest.Params, method string) (gorequest.Response, error) {

	// 签名
	sign := c.sign(param)

	// 创建请求
	client := gorequest.NewHttp()

	// 设置请求地址
	timestamp := gotime.Current().Timestamp()
	client.SetUri(fmt.Sprintf("%s?user_id=%s&timestamp=%v&sign=%s", c.config.apiUrl+url, c.GetUserID(), timestamp, sign))

	// 设置方式
	client.SetMethod(method)

	// 设置格式
	client.SetContentTypeJson()

	// 设置用户代理
	client.SetUserAgent(gorequest.GetRandomUserAgentSystem())

	// 设置参数
	client.SetParams(param)

	// 发起请求
	request, err := client.Request(ctx)
	if err != nil {
		return gorequest.Response{}, err
	}

	// 日志
	if c.gormLog.status {
		go c.gormLog.client.Middleware(ctx, request)
	}

	return request, err
}
