package request

import (
	"fmt"
	"net/url"
	"strconv"
)

type OrderListRequest struct {
	//BusinessLine  int `query:"businessLine"`
	//ActId         int `query:"actId"`
	//Page          int `query:"page"`
	//Limit         int `query:"limit"`
	//QueryTimeType int `query:"queryTimeType"`
	//StartTime     int `query:"startTime"`
	//EndTime       int `query:"endTime"`
	params map[string]string
	query  url.Values
}

func NewOrderListRequest() *OrderListRequest {
	var uri url.URL
	q := uri.Query()
	q.Add("businessLine", "")
	q.Add("actId", "")
	q.Add("page", "1")
	q.Add("limit", "20")
	q.Add("queryTimeType", "1")
	q.Add("startTime", "")
	q.Add("endTime", "")
	return &OrderListRequest{
		params: map[string]string{
			"businessLine":  "",
			"actId":         "",
			"page":          "",
			"limit":         "",
			"queryTimeType": "",
			"startTime":     "",
			"endTime":       "",
		},
		query: q,
	}
}

func (r OrderListRequest) GetApiPath() string {
	return "/api/orderList"
}

func (r OrderListRequest) SetBusinessLine(v int) OrderListRequest {
	r.params["businessLine"] = strconv.Itoa(v)
	r.query.Set("businessLine", strconv.Itoa(v))
	return r
}

func (r OrderListRequest) SetActId(v int) OrderListRequest {
	r.params["actId"] = strconv.Itoa(v)
	r.query.Set("actId", strconv.Itoa(v))
	return r
}

func (r OrderListRequest) SetStartTime(v int) OrderListRequest {
	r.params["startTime"] = strconv.Itoa(v)
	r.query.Set("startTime", strconv.Itoa(v))
	return r
}

func (r OrderListRequest) SetEndTime(v int) OrderListRequest {
	r.params["endTime"] = strconv.Itoa(v)
	r.query.Set("endTime", strconv.Itoa(v))
	return r
}

func (r OrderListRequest) Params() url.Values {
	fmt.Println(r.params)
	for k, v := range r.params {
		fmt.Println(k, v)
	}
	fmt.Println(r.query)
	for k, v := range r.params {
		fmt.Println(k, v)
	}
	fmt.Println(r.query.Encode())
	return r.query
}
