package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAlidnsClient(t *testing.T) {
	conf := Config{
		EndPoint:        "alidns.cn-shenzhen.aliyuncs.com",
		AccessKeyId:     "LTAIS1a2d0baa2Xj",
		AccessKeySecret: "4IZsw0ea3Y1aZSdB8xDG0ZfvaBnYvH",
		Domain:          "example.com",
		Replace:         "vpn.example.com",
	}
	client, err := newAliDnsClient(conf)
	assert.Nil(t, err)
	assert.NotNil(t, client)
}

func TestNewClient(t *testing.T) {
	conf := Config{}
	client, err := NewClient(conf)
	assert.NotNil(t, err)
	assert.Nil(t, client)

	conf = Config{
		EndPoint:        "alidns.cn-shenzhen.aliyuncs.com",
		AccessKeyId:     "LTAIS1a2d0baa2Xj",
		AccessKeySecret: "4IZsw0ea3Y1aZSdB8xDG0ZfvaBnYvH",
		Domain:          "example.com",
		Replace:         "vpn.example.com",
	}

	client, err = NewClient(conf)
	assert.Nil(t, err)
	assert.NotNil(t, client)
}
