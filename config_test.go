package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewConfig(t *testing.T) {
	path := "no_exists.yml"
	_, err := NewConfig(path)
	assert.NotNil(t, err)
	assert.ErrorIs(t, err, os.ErrNotExist)

	path = "config.example.yml"
	conf, err := NewConfig(path)
	assert.Nil(t, err)

	assert.IsType(t, Config{}, conf)
	assert.Equal(t, "alidns.cn-shenzhen.aliyuncs.com", conf.EndPoint)
	assert.Equal(t, "LTAIS1a2d0baa2Xj", conf.AccessKeyId)
	assert.Equal(t, "4IZsw0ea3Y1aZSdB8xDG0ZfvaBnYvH", conf.AccessKeySecret)
	assert.Equal(t, "example.com", conf.Domain)
	assert.Equal(t, "vpn.example.com", conf.Replace)
}

func TestValidate(t *testing.T) {
	var err error
	var conf Config

	conf = Config{
		EndPoint:        "1",
		AccessKeyId:     "1",
		AccessKeySecret: "1",
		Domain:          "1",
		Replace:         "1",
		Keyword:         "1",
	}

	err = conf.validate()
	assert.Nil(t, err)

	conf = Config{}
	err = conf.validate()
	assert.NotNil(t, err)
}

func TestKeyword(t *testing.T) {
	c := &Config{
		Domain:  "example.com",
		Replace: "vpn.example.com",
	}

	assert.Equal(t, "vpn", c.keyword())
}
