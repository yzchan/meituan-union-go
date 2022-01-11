package meituan

type Resp struct {
	Status int    `json:"status"`
	Desc   string `json:"des"`
	Data   string `json:"data"`
}

type Order struct {
	BusinessLine     int         `json:"businessLine"`
	SubBusinessLine  int         `json:"subBusinessLine"`
	ActId            int         `json:"actId"`
	Quantity         int         `json:"quantity"`
	OrderId          string      `json:"orderId"`
	Paytime          string      `json:"paytime"`
	ModTime          string      `json:"modTime"`
	Payprice         string      `json:"payprice"`
	Profit           string      `json:"profit"`
	CpaProfit        string      `json:"cpaProfit"`
	Sid              string      `json:"sid"`
	Appkey           string      `json:"appkey"`
	Smstitle         string      `json:"smstitle"`
	Status           int         `json:"status"`
	TradeTypeList    []int       `json:"tradeTypeList"`
	RiskOrder        int         `json:"riskOrder"`
	Refundprofit     string      `json:"refundprofit"`
	CpaRefundProfit  string      `json:"cpaRefundProfit"`
	RefundInfoList   interface{} `json:"refundInfoList"`
	RefundProfitList interface{} `json:"refundProfitList"`
	Extra            interface{} `json:"extra"`
}

type OrderListResp struct {
	Total    int     `json:"total"`
	DataList []Order `json:"dataList"`
}

type OrderResp struct {
	Status int    `json:"status"`
	Desc   string `json:"des"`
	Order  Order  `json:"data"`
}
