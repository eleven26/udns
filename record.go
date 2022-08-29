package main

import alidns20150109 "github.com/alibabacloud-go/alidns-20150109/v3/client"

type Record struct {
	DomainName *string
	Line       *string
	Locked     *bool
	RR         *string
	RecordId   *string
	Status     *string
	TTL        *int64
	Type       *string
	Value      *string
	Weight     *int32
}

func NewRecord(record *alidns20150109.DescribeDomainRecordsResponseBodyDomainRecordsRecord) *Record {
	return &Record{
		DomainName: record.DomainName,
		Line:       record.Line,
		Locked:     record.Locked,
		RR:         record.RR,
		RecordId:   record.RecordId,
		Status:     record.Status,
		TTL:        record.TTL,
		Type:       record.Type,
		Value:      record.Value,
		Weight:     record.Weight,
	}
}
