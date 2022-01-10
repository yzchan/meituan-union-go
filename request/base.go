package request

import "net/url"

type Requester interface {
	GetApiPath() string
	Params() url.Values
}

type base struct {
	params url.Values
}

func (b base) SetParams(data map[string]string) {
	for k, v := range data {
		b.params.Set(k, v)
	}
}
