package main

import (
	alidns20150109 "github.com/alibabacloud-go/alidns-20150109/v3/client"
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/auth/credentials"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/alidns"
)

type Client interface {
	DescribeDomainRecords(domain string, keyword string) ([]*Record, error)
	UpdateRecord(id string, ip string, rr string) error
	AddRecord(ip string, domain string, rr string) error
	Client() *alidns20150109.Client
}

func NewClient(conf Config) (Client, error) {
	err := conf.validate()
	if err != nil {
		return nil, err
	}

	client, err := newAliDnsClient(conf)
	if err != nil {
		return nil, err
	}

	return &_client{
		conf:   conf,
		client: client,
	}, nil
}

func newAliDnsClient(conf Config) (*alidns20150109.Client, error) {
	config := &openapi.Config{
		AccessKeyId:     &conf.AccessKeyId,
		AccessKeySecret: &conf.AccessKeySecret,
		Endpoint:        &conf.EndPoint,
	}

	return alidns20150109.NewClient(config)
}

type _client struct {
	conf   Config
	client *alidns20150109.Client
}

func (c *_client) DescribeDomainRecords(domain string, keyword string) ([]*Record, error) {
	request := &alidns20150109.DescribeDomainRecordsRequest{
		DomainName: &domain,
		KeyWord:    &keyword,
	}
	runtime := &util.RuntimeOptions{}

	response, err := c.client.DescribeDomainRecordsWithOptions(request, runtime)
	if err != nil {
		return nil, err
	}

	var result []*Record

	var zero int64 = 0
	if *response.Body.TotalCount == zero {
		return result, nil
	}

	for _, record := range response.Body.DomainRecords.Record {
		result = append(result, NewRecord(record))
	}

	return result, nil
}

func (c *_client) UpdateRecord(id string, ip string, rr string) error {
	request := &alidns20150109.UpdateDomainRecordRequest{
		RecordId: tea.String(id),
		RR:       tea.String(rr),
		Type:     tea.String("A"),
		Value:    tea.String(ip),
	}
	runtime := &util.RuntimeOptions{}

	_, err := c.client.UpdateDomainRecordWithOptions(request, runtime)

	return err
}

func (c *_client) AddRecord(ip string, domain string, rr string) error {
	config := sdk.NewConfig()

	credential := credentials.NewAccessKeyCredential(c.conf.AccessKeyId, c.conf.AccessKeySecret)
	client, err := alidns.NewClientWithOptions("cn-shenzhen", config, credential)
	if err != nil {
		return err
	}

	request := alidns.CreateAddDomainRecordRequest()

	request.Scheme = "https"
	request.Type = "A"
	request.Value = ip
	request.RR = rr
	request.DomainName = domain

	_, err = client.AddDomainRecord(request)

	return err
}

func (c *_client) Client() *alidns20150109.Client {
	return c.client
}
