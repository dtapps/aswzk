package aswzk

import (
	"errors"
	"go.dtapp.net/golog"
)

// ClientConfig 实例配置
type ClientConfig struct {
	ApiUrl string // 接口地址
	UserID string // 用户编号
	ApiKey string // 秘钥
}

// Client 实例
type Client struct {
	config struct {
		apiUrl string // 接口地址
		userID string // 用户编号
		apiKey string // 秘钥
	}
	gormLog struct {
		status bool           // 状态
		client *golog.ApiGorm // 日志服务
	}
	mongoLog struct {
		status bool            // 状态
		client *golog.ApiMongo // 日志服务
	}
}

// NewClient 创建实例化
func NewClient(config *ClientConfig) (*Client, error) {

	c := &Client{}
	if config.ApiUrl == "" {
		return nil, errors.New("ApiUrl is empty")
	}

	c.config.apiUrl = config.ApiUrl
	c.config.userID = config.UserID
	c.config.apiKey = config.ApiKey

	return c, nil
}
