package aswzk

import (
	"context"
	"fmt"
	"go.dtapp.net/gorequest"
	"go.dtapp.net/gotime"
	"log"
)

// 请求接口
func (c *Client) request(ctx context.Context, url string, param gorequest.Params, method string) (gorequest.Response, error) {

	// 获取时间戳
	XTimestamp := fmt.Sprintf("%v", gotime.Current().Timestamp())

	// 签名
	XSign := sign(param, c.GetApiKey(), XTimestamp)
	log.Printf("签名参数：%+v\n", param)

	// 设置请求地址
	c.httpClient.SetUri(c.GetApiUrl() + url)

	// 设置方式
	c.httpClient.SetMethod(method)

	// 设置格式
	c.httpClient.SetContentTypeJson()

	// 设置参数
	c.httpClient.SetParams(param)

	// 添加请求头
	c.httpClient.SetHeader("X-Timestamp", XTimestamp)
	c.httpClient.SetHeader("X-UserId", c.GetUserID())
	c.httpClient.SetHeader("X-Sign", XSign)

	// 发起请求
	request, err := c.httpClient.Request(ctx)
	if err != nil {
		return gorequest.Response{}, err
	}

	return request, err
}
