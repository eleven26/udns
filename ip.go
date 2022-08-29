package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

type IpResponse struct {
	Ip string `json:"ip"`
}

// Ip get public ip
func Ip() (r *string, err error) {
	url := "https://api64.ipify.org?format=json"

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
	}(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	ipResponse := IpResponse{}
	err = json.Unmarshal(body, &ipResponse)
	if err != nil {
		return nil, err
	}

	if validIPAddress(ipResponse.Ip) != "IPv4" {
		return nil, fmt.Errorf("invalid ipv4 address")
	}

	r = &ipResponse.Ip

	return
}

// https://leetcode.cn/problems/validate-ip-address/solution/yan-zheng-ipdi-zhi-by-leetcode-solution-kge5/
func validIPAddress(queryIP string) string {
	if sp := strings.Split(queryIP, "."); len(sp) == 4 {
		for _, s := range sp {
			if len(s) > 1 && s[0] == '0' {
				return "Neither"
			}
			if v, err := strconv.Atoi(s); err != nil || v > 255 {
				return "Neither"
			}
		}
		return "IPv4"
	}
	if sp := strings.Split(queryIP, ":"); len(sp) == 8 {
		for _, s := range sp {
			if len(s) > 4 {
				return "Neither"
			}
			if _, err := strconv.ParseUint(s, 16, 64); err != nil {
				return "Neither"
			}
		}
		return "IPv6"
	}
	return "Neither"
}
