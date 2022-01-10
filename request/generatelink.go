package request

import (
	"net/url"
	"strconv"
)

const H5 = 1       // H5类型的链接
const DEEPLINK = 2 // DEEP类型的链接
const CENTER = 3   // 中间唤起页的链接
const WECHAT = 4   // 微信小程序Path
const GROUP = 5    // 团口令

type GenerateLinkRequest struct {
	base
}

func NewGenerateLinkRequest() *GenerateLinkRequest {
	var uri url.URL
	q := uri.Query()
	q.Add("sid", "")
	q.Add("actId", "")
	q.Add("linkType", "")
	q.Add("shortLink", "1") // 默认为取短链
	return &GenerateLinkRequest{
		base{params: q},
	}
}

func (r GenerateLinkRequest) GetApiPath() string {
	return "/api/generateLink"
}

func (r GenerateLinkRequest) Params() url.Values {
	return r.params
}

func (r GenerateLinkRequest) SetSid(v string) GenerateLinkRequest {
	r.params.Set("sid", v)
	return r
}

func (r GenerateLinkRequest) SetActId(v int) GenerateLinkRequest {
	r.params.Set("actId", strconv.Itoa(v))
	return r
}

func (r GenerateLinkRequest) SetLinkType(v int) GenerateLinkRequest {
	r.params.Set("linkType", strconv.Itoa(v))
	return r
}

func (r GenerateLinkRequest) SetShortLink(v bool) GenerateLinkRequest {
	if v {
		r.params.Set("shortLink", "1")
	} else {
		r.params.Set("shortLink", "0")
	}
	return r
}
