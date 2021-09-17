package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"
)

const GateWay = "https://runion.meituan.com"

type Client struct {
	Key    string
	Secret string
	Debug  bool
}

func NewClient(key string, secret string) *Client {
	return &Client{
		Key:    key,
		Secret: secret,
		Debug:  true,
	}
}

func (c *Client) sign(params map[string]string) string {
	delete(params, "sign")

	keys := make([]string, 0)
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	str := c.Secret
	for _, key := range keys {
		str += key + params[key]
	}
	str += c.Secret

	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func (c *Client) request(uri string, params map[string]string) (content []byte, err error) {
	var (
		body  []byte
		resp  *http.Response
		req   *http.Request
		proxy func(*http.Request) (*url.URL, error)
	)

	params["key"] = c.Key
	params["sign"] = c.sign(params)
	qs := make([]string, 0)
	for k, v := range params {
		qs = append(qs, k+"="+v)
	}
	uri = uri + "?" + strings.Join(qs, "&")

	if c.Debug {
		log.Println("==========Meituan Debug Start [Request↓]==========")
		log.Printf("GET %s\n%s\n", uri, string(body))
	}

	client := &http.Client{
		Transport: &http.Transport{Proxy: proxy},
		Timeout:   time.Second * 5,
	}
	if req, err = http.NewRequest("GET", uri, nil); err != nil {
		return
	}

	if resp, err = client.Do(req); err != nil {
		return
	}
	defer resp.Body.Close()

	if content, err = ioutil.ReadAll(resp.Body); err != nil {
		return
	}
	if c.Debug {
		log.Println("==========Meituan Debug [Response↓]==========")
		log.Printf("Http Code:%d\n%s\n", resp.StatusCode, string(content))
		log.Println("==========Meituan Debug End==========")
	}
	// 统一处理非200响应
	if resp.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("[%d]:%s", resp.StatusCode, string(content)))
	}
	return
}

type Resp struct {
	Status int    `json:"status"`
	Desc   string `json:"des"`
	Data   string `json:"data"`
}

func (c *Client) GenerateUrl(params map[string]string) (ret *Resp, err error) {
	var content []byte
	if content, err = c.request(GateWay+"/generateLink", params); err != nil {
		return
	}
	ret = &Resp{}
	if err = json.Unmarshal(content, ret); err != nil {
		return
	}
	return
}

func (c *Client) MiniCode(params map[string]string) (ret *Resp, err error) {
	var content []byte
	if content, err = c.request(GateWay+"/miniCode", params); err != nil {
		return
	}
	ret = &Resp{}
	if err = json.Unmarshal(content, ret); err != nil {
		return
	}
	return
}

type Order struct {
	Smstitle string `json:"smstitle"`
	Orderid  string `json:"orderid"`
	Paytime  string `json:"paytime"`
	Appkey   string `json:"appkey"`
	Payprice string `json:"payprice"`
	Profit   string `json:"profit"`
	Sid      string `json:"sid"`
	Status   int    `json:"status"`
}

type OrderResp struct {
	Total    int     `json:"total"`
	DataList []Order `json:"dataList"`
}

func (c *Client) OrderList(params map[string]string) (ret *OrderResp, err error) {
	var content []byte
	params["ts"] = strconv.Itoa(int(time.Now().Unix()))
	if content, err = c.request(GateWay+"/api/orderList", params); err != nil {
		return
	}
	ret = &OrderResp{}
	if err = json.Unmarshal(content, ret); err != nil {
		return
	}
	return
}
