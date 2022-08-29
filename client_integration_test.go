//go:build integration

package main

import (
	"log"
	"testing"

	alidns20150109 "github.com/alibabacloud-go/alidns-20150109/v3/client"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/stretchr/testify/assert"
)

var (
	c    Client
	conf Config
	rr   *string
	id   string
)

func init() {
	rr = tea.String("5910093801")

	var err error

	conf, err = NewFromDefaultConfigPaths()
	if err != nil {
		log.Fatal(err)
	}

	c, err = NewClient(conf)
	if err != nil {
		log.Fatal(err)
	}
}

func setup() {
	err := c.AddRecord("127.0.0.1", conf.Domain, *rr)
	if err != nil {
		log.Fatal(err)
	}
}

func teardown() {
	if id == "" {
		return
	}

	request := &alidns20150109.DeleteDomainRecordRequest{
		RecordId: &id,
	}
	runtime := &util.RuntimeOptions{}

	_, err := c.Client().DeleteDomainRecordWithOptions(request, runtime)
	if err != nil {
		log.Fatal(err)
	}
}

func TestDescribeDomainRecords(t *testing.T) {
	records, err := c.DescribeDomainRecords(conf.Domain, *rr)
	assert.Nil(t, err)
	assert.Len(t, records, 0)

	setup()
	defer teardown()

	records, err = c.DescribeDomainRecords(conf.Domain, *rr)
	assert.Nil(t, err)
	assert.Len(t, records, 1)

	id = *records[0].RecordId
}

func TestAddRecord(t *testing.T) {
	defer teardown()

	err := c.AddRecord("127.0.0.1", conf.Domain, *rr)

	records, err := c.DescribeDomainRecords(conf.Domain, *rr)
	assert.Nil(t, err)
	assert.Len(t, records, 1)

	record := records[0]
	assert.Equal(t, *rr, *record.RR)

	id = *record.RecordId
}

func TestUpdateRecord(t *testing.T) {
	setup()
	defer teardown()

	records, _ := c.DescribeDomainRecords(conf.Domain, *rr)
	record := records[0]

	var newIp string
	if *record.Value == "127.0.0.1" {
		newIp = "192.168.0.1"
	} else {
		newIp = "127.0.0.1"
	}

	id = *record.RecordId

	err := c.UpdateRecord(*record.RecordId, newIp, *rr)
	assert.Nil(t, err)

	records, err = c.DescribeDomainRecords(conf.Domain, *rr)
	assert.Nil(t, err)
	assert.Len(t, records, 1)

	r := records[0]
	assert.Equal(t, newIp, *r.Value)
}
