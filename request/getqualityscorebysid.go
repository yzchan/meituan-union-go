package request

import (
	"net/url"
	"strconv"
)

type GetQualityScoreBySidRequest struct {
	base
}

func NewGetQualityScoreBySidRequest() *GetQualityScoreBySidRequest {
	var uri url.URL
	q := uri.Query()
	q.Add("sid", "")
	q.Add("type", "1")
	q.Add("businessLine", "")
	q.Add("beginDate", "")
	q.Add("endDate", "")
	return &GetQualityScoreBySidRequest{
		base{params: q},
	}
}

func (r *GetQualityScoreBySidRequest) GetApiPath() string {
	return "/api/getqualityscorebysid"
}

func (r *GetQualityScoreBySidRequest) Params() url.Values {
	return r.params
}

func (r *GetQualityScoreBySidRequest) SetSid(v string) *GetQualityScoreBySidRequest {
	r.params.Set("sid", v)
	return r
}

func (r *GetQualityScoreBySidRequest) SetType(v int) *GetQualityScoreBySidRequest {
	r.params.Set("type", strconv.Itoa(v))
	return r
}

func (r *GetQualityScoreBySidRequest) SetBusinessLine(v int) *GetQualityScoreBySidRequest {
	r.params.Set("businessLine", strconv.Itoa(v))
	return r
}

func (r *GetQualityScoreBySidRequest) SetBeginDate(v string) *GetQualityScoreBySidRequest {
	r.params.Set("beginDate", v)
	return r
}
func (r *GetQualityScoreBySidRequest) SetEndDate(v string) *GetQualityScoreBySidRequest {
	r.params.Set("endDate", v)
	return r
}
