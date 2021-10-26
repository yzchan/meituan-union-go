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

const GateWay = "https://openapi.meituan.com"

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

	params["appkey"] = c.Key
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
	if content, err = c.request(GateWay+"/api/generateLink", params); err != nil {
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
	if content, err = c.request(GateWay+"/api/miniCode", params); err != nil {
		return
	}
	ret = &Resp{}
	if err = json.Unmarshal(content, ret); err != nil {
		return
	}
	return
}

type Order struct {
	BusinessLine     int         `json:"businessLine"`
	SubBusinessLine  int         `json:"subBusinessLine"`
	ActId            int         `json:"actId"`
	Quantity         int         `json:"quantity"`
	OrderId          string      `json:"orderId"`
	Paytime          string      `json:"paytime"`
	ModTime          string      `json:"modTime"`
	Payprice         string      `json:"payprice"`
	Profit           string      `json:"profit"`
	CpaProfit        string      `json:"cpaProfit"`
	Sid              string      `json:"sid"`
	Appkey           string      `json:"appkey"`
	Smstitle         string      `json:"smstitle"`
	Status           int         `json:"status"`
	TradeTypeList    []int       `json:"tradeTypeList"`
	RiskOrder        int `json:"riskOrder"`
	Refundprofit     string `json:"refundprofit"`
	CpaRefundProfit  string `json:"cpaRefundProfit"`
	RefundInfoList   interface{} `json:"refundInfoList"`
	RefundProfitList interface{} `json:"refundProfitList"`
	Extra            interface{} `json:"extra"`
}

type OrderListResp struct {
	Total    int     `json:"total"`
	DataList []Order `json:"dataList"`
}

func (c *Client) OrderList(params map[string]string) (ret *OrderListResp, err error) {
	var content []byte
	params["ts"] = strconv.Itoa(int(time.Now().Unix()))
	if content, err = c.request(GateWay+"/api/orderList", params); err != nil {
		return
	}
	ret = &OrderListResp{}
	if err = json.Unmarshal(content, ret); err != nil {
		return
	}
	return
}

type OrderResp struct {
	Status int    `json:"status"`
	Desc   string `json:"des"`
	Order  Order  `json:"data"`
}

func (c *Client) Order(params map[string]string) (ret *OrderResp, err error) {
	var content []byte
	if content, err = c.request(GateWay+"/api/order", params); err != nil {
		return
	}
	fmt.Println(string(content))
	ret = &OrderResp{}
	if err = json.Unmarshal(content, ret); err != nil {
		return
	}
	return
}

func main() {
	client := NewClient("xxx", "xxx")
	client.Debug = false
	//params := map[string]string{
	//	"actId":     "33",
	//	"sid":       "test",
	//	"linkType":  "1",
	//	"shortLink": "1",
	//}
	//fmt.Println(client.GenerateUrl(params))

	//params := map[string]string{
	//	"actId":     "33",
	//	"businessLine": "2",    // actId和businessLine至少有一个
	//	"startTime": "1634659200",
	//	"endTime": "1634745600",   // 不能超过1天
	//	"page": "1",
	//	"limit": "20",
	//	"queryTimeType": "1",
	//}
	//fmt.Println(client.OrderList(params))

	//params := map[string]string{
	//	"actId":        "33",
	//	"businessLine": "2", // actId和businessLine至少有一个
	//	"orderId":      "3233710410647623",
	//	"full":         "1",
	//}
	//fmt.Println(client.Order(params))
}
