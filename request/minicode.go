package request

import (
	"net/url"
	"strconv"
)

type MiniCodeRequest struct {
	base
}

func NewMiniCodeRequest() *MiniCodeRequest {
	var uri url.URL
	q := uri.Query()
	q.Add("sid", "")
	q.Add("actId", "")
	return &MiniCodeRequest{
		base{params: q},
	}
}

func (r *MiniCodeRequest) GetApiPath() string {
	return "/api/miniCode"
}

func (r *MiniCodeRequest) Params() url.Values {
	return r.params
}

func (r *MiniCodeRequest) SetSid(v string) *MiniCodeRequest {
	r.params.Set("sid", v)
	return r
}

func (r *MiniCodeRequest) SetActId(v int) *MiniCodeRequest {
	r.params.Set("actId", strconv.Itoa(v))
	return r
}
