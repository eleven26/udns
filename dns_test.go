//go:build integration

package main

import (
	"errors"
	"testing"

	alidns20150109 "github.com/alibabacloud-go/alidns-20150109/v3/client"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type ClientStub struct {
	mock.Mock
}

func (c *ClientStub) DescribeDomainRecords(domain string, keyword string) ([]*Record, error) {
	args := c.Called(domain, keyword)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).([]*Record), args.Error(1)
}

func (c *ClientStub) UpdateRecord(id string, ip string, rr string) error {
	args := c.Called(id, ip, rr)

	return args.Error(0)
}

func (c *ClientStub) AddRecord(ip string, domain string, rr string) error {
	args := c.Called(ip, domain, rr)

	return args.Error(0)
}

func (c *ClientStub) Client() *alidns20150109.Client {
	args := c.Called()

	return args.Get(0).(*alidns20150109.Client)
}

func TestGetRecord(t *testing.T) {
	client := new(ClientStub)

	client.On("DescribeDomainRecords", "", "").Return(nil, errors.New("test"))

	dns := &DnsResolver{
		conf:   Config{},
		client: client,
	}

	records, err := dns.Record()
	assert.NotNil(t, err)
	assert.Empty(t, records)

	client.AssertExpectations(t)
}
