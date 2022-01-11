package meituan

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/yzchan/meituan-union-go/request"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
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
		Debug:  false,
	}
}

func (c *Client) sign(params url.Values) string {
	params.Del("sign")

	qs := params.Encode() // url.Values的Encode方法会对key进行排序
	qs = strings.ReplaceAll(qs, "&", "")
	qs = strings.ReplaceAll(qs, "=", "")

	str := c.Secret + qs + c.Secret

	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func (c *Client) Do(request request.Requester) (content []byte, err error) {
	return c.request(request.GetApiPath(), request.Params())
}

func (c *Client) request(path string, params url.Values) (content []byte, err error) {
	var (
		body  []byte
		resp  *http.Response
		req   *http.Request
		proxy func(*http.Request) (*url.URL, error)
	)

	params.Add("appkey", c.Key)
	params.Add("ts", strconv.Itoa(int(time.Now().Unix())))
	params.Add("sign", c.sign(params))
	uri := GateWay + path + "?" + params.Encode()

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
