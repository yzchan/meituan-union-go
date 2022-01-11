package request

import (
	"net/url"
	"strconv"
)

type OrderListRequest struct {
	base
}

func NewOrderListRequest() *OrderListRequest {
	var uri url.URL
	q := uri.Query()
	q.Add("businessLine", "")
	q.Add("actId", "")
	q.Add("startTime", "")
	q.Add("endTime", "")
	q.Add("page", "1")
	q.Add("limit", "20")
	q.Add("queryTimeType", "1")
	return &OrderListRequest{
		base{params: q},
	}
}

func (r *OrderListRequest) GetApiPath() string {
	return "/api/orderList"
}

func (r *OrderListRequest) Params() url.Values {
	return r.params
}

func (r *OrderListRequest) SetBusinessLine(v int) *OrderListRequest {
	r.params.Set("businessLine", strconv.Itoa(v))
	return r
}

func (r *OrderListRequest) SetActId(v int) *OrderListRequest {
	r.params.Set("actId", strconv.Itoa(v))
	return r
}

func (r *OrderListRequest) SetStartTime(v int) *OrderListRequest {
	r.params.Set("startTime", strconv.Itoa(v))
	return r
}

func (r *OrderListRequest) SetEndTime(v int) *OrderListRequest {
	r.params.Set("endTime", strconv.Itoa(v))
	return r
}

func (r *OrderListRequest) SetPage(v int) *OrderListRequest {
	r.params.Set("page", strconv.Itoa(v))
	return r
}
func (r *OrderListRequest) SetLimit(v int) *OrderListRequest {
	r.params.Set("limit", strconv.Itoa(v))
	return r
}
func (r *OrderListRequest) SetQueryTimeType(v int) *OrderListRequest {
	r.params.Set("queryTimeType", strconv.Itoa(v))
	return r
}
