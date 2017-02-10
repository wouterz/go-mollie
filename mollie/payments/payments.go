package payments

import (
	"time"
	"strconv"
	"localhost/he/go-mollie-api/mollie/core"
)

type PaymentApi struct {
	core core.Core
}

type PaymentData struct {
	Amount float64  `json:"amount,string"`
	Description string `json:"description"`
	RedirectUrl string `json:"redirectUrl"`
	WebhookUrl string `json:"webhookUrl"`
	Method string `json:"method,omitempty"`
	Metadata interface{} `json:"metadata"`
	// One of de_DE en_US es_ES fr_FR nl_BE fr_BE nl_NL
	Locale string `json:"locale,omitempty"`

	Issuer string `json:"issuer,omitempty"`
}

type PaymentReply struct {
	Id string
	Mode string
	CreatedDatetime time.Time `json:"createdDatetime"`
	ExpiredDatetime time.Time `json:"expiredDatetime"`
	CancelledDatetime time.Time `json:"cancelledDatetime"`
	PaidDatetime time.Time `json:"paidDatetime"`
	Status string
	ExpiryPeriod string `json:"expiryPeriod"`
	Amount float64 `json:",string"`
	AmountRefunded float64 `json:"amountRefunded,string,omitempty"`
	AmountRemaining float64 `json:"amountRemaining,string,omitempty"`
	Description string
	Method string
	Metadata interface{}
	Details interface{} `json:",omitempty"`
	ProfileId string `json:"profileId"`
	Links map[string]string
}

type PaymentReplyWrapper struct {
	TotalCount int `json:"totalCount"`
	Offset     int
	Count      int
	Data       []PaymentReply
}

func NewPayments(c core.Core) *PaymentApi {
	return &PaymentApi{core: c}
}

func (a *PaymentApi) New(data PaymentData) (*PaymentReply, error) {
	p := PaymentReply{}

	err := a.core.Post("payments", &p, &data)
	if err != nil {
		return nil, err
	}

	return &p, nil
}

func (a *PaymentApi) Get(id string) (*PaymentReply, error) {
	p := PaymentReply{}

	err := a.core.Get("payments/" + id, &p)
	if err != nil {
		return nil, err
	}

	return &p, nil
}

func (a *PaymentApi) List(offset, limit uint64) ([]PaymentReply, error) {
	p := PaymentReplyWrapper{}

	uri := "payments?offset="
	uri += strconv.FormatUint(offset, 10)
	uri += "&count="
	uri += strconv.FormatUint(limit, 10)

	err := a.core.Get(uri, &p)
	if err != nil {
		return nil, err
	}

	return p.Data, nil
}
