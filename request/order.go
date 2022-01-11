package request

import (
	"net/url"
	"strconv"
)

type OrderRequest struct {
	base
}

func NewOrderRequest() *OrderRequest {
	var uri url.URL
	q := uri.Query()
	q.Add("businessLine", "")
	q.Add("actId", "")
	q.Add("full", "1")
	q.Add("orderId", "")
	return &OrderRequest{
		base{params: q},
	}
}

func (r *OrderRequest) GetApiPath() string {
	return "/api/order"
}

func (r *OrderRequest) Params() url.Values {
	return r.params
}

func (r *OrderRequest) SetBusinessLine(v int) *OrderRequest {
	r.params.Set("businessLine", strconv.Itoa(v))
	return r
}

func (r *OrderRequest) SetActId(v int) *OrderRequest {
	r.params.Set("actId", strconv.Itoa(v))
	return r
}

func (r *OrderRequest) SetOrderId(v string) *OrderRequest {
	r.params.Set("orderId", v)
	return r
}

// SetFull 是否返回完整订单信息(即是否包含返佣、退款信息) 枚举值： 0-非全量查询  1-全量查询
func (r *OrderRequest) SetFull(v bool) *OrderRequest {
	if v {
		r.params.Set("full", "1")
	} else {
		r.params.Set("full", "0")
	}
	return r
}