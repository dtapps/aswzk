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

	// 创建请求
	client := gorequest.NewHttp()
	client.SetDebug()

	// 设置请求地址
	client.SetUri(c.GetApiUrl() + url)
	log.Printf("请求地址：%+v\n", c.GetApiUrl()+url)

	// 设置方式
	client.SetMethod(method)

	// 设置格式
	client.SetContentTypeJson()

	// 设置参数
	client.SetParams(param)

	// 添加请求头
	client.SetHeader("X-Timestamp", XTimestamp)
	client.SetHeader("X-UserId", c.GetUserID())
	client.SetHeader("X-Sign", XSign)

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
