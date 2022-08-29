package main

type Dns interface {
	Record() (*Record, error)
	UpdateRecord(id string, ip string) error
	AddRecord(ip string) error
	Replace() error
}

type DnsResolver struct {
	conf   Config
	client Client
}

func NewFromDefaultPaths() (Dns, error) {
	conf, err := NewFromDefaultConfigPaths()
	if err != nil {
		return nil, err
	}

	client, err := NewDnsResolver(conf)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func NewDnsResolver(conf Config) (Dns, error) {
	client, err := NewClient(conf)
	if err != nil {
		return nil, err
	}

	return &DnsResolver{
		conf:   conf,
		client: client,
	}, nil
}

// Record 根据关键字匹配出需要替换的域名
func (c *DnsResolver) Record() (*Record, error) {
	records, err := c.client.DescribeDomainRecords(c.conf.Domain, c.conf.Keyword)
	if err != nil {
		return nil, err
	}

	var r *Record

	for _, record := range records {
		if *record.RR == c.conf.Keyword {
			r = record
			break
		}
	}

	if r == nil {
		return nil, nil
	}

	return r, err
}

func (c *DnsResolver) UpdateRecord(id string, ip string) error {
	return c.client.UpdateRecord(id, ip, c.conf.Keyword)
}

func (c *DnsResolver) AddRecord(ip string) error {
	return c.client.AddRecord(ip, c.conf.Domain, c.conf.Keyword)
}

func (c *DnsResolver) Replace() error {
	ip, err := Ip()
	if err != nil {
		return err
	}

	record, err := c.Record()
	if err != nil {
		return err
	}

	if record != nil {
		if *record.Value != *ip {
			return c.UpdateRecord(*record.RecordId, *ip)
		} else {
			return nil
		}
	}

	return c.AddRecord(*ip)
}
