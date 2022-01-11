package meituan

import (
	"github.com/yzchan/meituan-union-go/request"
	"testing"
)

var client *Client

func init() {
	client = NewClient(Key, Secret)
	client.Debug = false
}

func TestGenerateUrlRequest(t *testing.T) {
	t.Log("test /api/generateUrl")

	req := request.NewGenerateLinkRequest()
	req.SetActId(33)
	req.SetSid("xxx")
	req.SetLinkType(1)
	req.SetShortLink(true)
	response, err := client.Do(req)

	if err != nil {
		t.Log(err.Error())
		t.Fatal("\x1b[31m测试失败！\x1b[0m")
	}
	t.Logf("response=%s\n", string(response))
	t.Log("\x1b[32mtest ok \x1b[0m")
}

func TestOrderListRequest(t *testing.T) {
	t.Log("test /api/orderList")

	req := request.NewOrderListRequest()
	req.SetActId(33)
	req.SetBusinessLine(2)
	req.SetStartTime(1634659200)
	req.SetEndTime(1634745600)
	req.SetQueryTimeType(1)
	req.SetPage(1)
	req.SetLimit(20)
	resp, err := client.Do(req)

	if err != nil {
		t.Log(err.Error())
		t.Fatal("\x1b[31m测试失败！\x1b[0m")
	}
	t.Logf("response=%s\n", string(resp))
	t.Log("\x1b[32mtest ok \x1b[0m")
}
