package client

import (
	"encoding/json"
)

type Subscriber struct {
	ID       string `json:"id"`
	Callback string `json:"callback"`
	EventID  string `json:"event_id"`

	HookClient
}

func NewSubscriber(host, apiVer, token string) *Subscriber {
	return &Subscriber{
		HookClient: HookClient{
			host:  host,
			ver:   apiVer,
			token: token,
		},
	}
}

func (s *Subscriber) Subscribe(publisher, event, callback string) (*Subscriber, error) {
	s.Callback = callback
	retData, err := s.postData(s.api()+"/events/"+publisher+"/"+event+"/subscribers", s)
	if nil != err {
		return nil, err
	}
	err = json.Unmarshal(retData, s)
	return s, err
}
