package main

import (
	"fmt"
	"reflect"
	"strings"
)

import (
	fs "github.com/eleven26/go-filesystem"
	"github.com/spf13/viper"
)

const File = "udns.yml"

type Config struct {
	EndPoint        string
	AccessKeyId     string
	AccessKeySecret string
	Domain          string
	Replace         string
	Keyword         string
}

func NewFromDefaultConfigPaths() (conf Config, err error) {
	p, err := DefaultConfigPath()
	if err != nil {
		return Config{}, err
	}

	return NewConfig(p)
}

func NewConfig(path string) (conf Config, err error) {
	err = readInConfig(path)
	if err != nil {
		return
	}

	conf = Config{
		EndPoint:        viper.GetString("endpoint"),
		AccessKeyId:     viper.GetString("access_key_id"),
		AccessKeySecret: viper.GetString("access_key_secret"),
		Domain:          viper.GetString("domain"),
		Replace:         viper.GetString("replace"),
	}

	err = conf.validate()
	if err != nil {
		return
	}

	conf.Keyword = conf.keyword()

	return conf, nil
}

// readInConfig Load the configuration file from the specified path.
func readInConfig(path string) error {
	exist, err := fs.Exists(path)
	if err != nil {
		return err
	}
	if !exist {
		return fmt.Errorf("Configuration file not exists：%s\n", path)
	}

	viper.SetConfigFile(path)

	return viper.ReadInConfig()
}

// Validate check if any configuration fields are empty
func (c *Config) validate() error {
	v := reflect.ValueOf(*c)
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		if t.Field(i).Name == "Keyword" {
			continue
		}

		if v.Field(i).Interface().(string) == "" {
			return fmt.Errorf("field %s can not be empty", t.Field(i).Name)
		}
	}

	return nil
}

func (c *Config) keyword() string {
	keyword := strings.ReplaceAll(c.Replace, c.Domain, "")

	return strings.TrimRight(keyword, ".")
}
