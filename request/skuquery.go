package request

import (
	"net/url"
	"strconv"
)

type SkuQueryRequest struct {
	base
}

func NewSkuQueryRequest() *SkuQueryRequest {
	var uri url.URL
	q := uri.Query()
	q.Add("businessType", "")
	q.Add("sid", "")
	q.Add("pageSize", "20")
	q.Add("pageNo", "1")
	q.Add("longitude", "")
	q.Add("latitude", "")
	q.Add("deviceType", "")
	q.Add("deviceId", "")
	return &SkuQueryRequest{
		base{params: q},
	}
}

func (r SkuQueryRequest) GetApiPath() string {
	return "/sku/query"
}

func (r SkuQueryRequest) Params() url.Values {
	return r.params
}

func (r SkuQueryRequest) SetSid(v string) SkuQueryRequest {
	r.params.Set("sid", v)
	return r
}

func (r SkuQueryRequest) SetBusinessType(v int) SkuQueryRequest {
	r.params.Set("businessType", strconv.Itoa(v))
	return r
}

func (r SkuQueryRequest) SetPageSize(v int) SkuQueryRequest {
	r.params.Set("pageSize", strconv.Itoa(v))
	return r
}

func (r SkuQueryRequest) SetPageNo(v int) SkuQueryRequest {
	r.params.Set("pageNO", strconv.Itoa(v))
	return r
}

func (r SkuQueryRequest) SetGeo(long, lat string) SkuQueryRequest {
	r.params.Set("longitude", long)
	r.params.Set("latitude", lat)
	return r
}

func (r SkuQueryRequest) SetDevice(deviceType, deviceId string) SkuQueryRequest {
	r.params.Set("deviceType", deviceType)
	r.params.Set("deviceId", deviceId)
	return r
}
