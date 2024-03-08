package aswzk

import (
	"go.dtapp.net/golog"
)

// ConfigApp 配置
func (c *Client) ConfigApp(userID string, apiKey string) *Client {
	c.config.userID = userID
	c.config.apiKey = apiKey
	return c
}

// ConfigApiGormFun 接口日志配置
func (c *Client) ConfigApiGormFun(apiClientFun golog.ApiGormFun) {
	client := apiClientFun()
	if client != nil {
		c.gormLog.client = client
		c.gormLog.status = true
	}
}
