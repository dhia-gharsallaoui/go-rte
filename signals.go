package gorte

import (
	"net/http"
	"time"
)

type SignalValue struct {
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	Value       bool      `json:"value"`
	UpdatedDate time.Time `json:"updated_date"`
}

type Signals struct {
	Signal []struct {
		StartDate time.Time     `json:"start_date"`
		EndDate   time.Time     `json:"end_date"`
		Type      string        `json:"type"`
		Values    []SignalValue `json:"values"`
	} `json:"signals"`
}

func (s *market) GetSignals(opt *Period) (*Signals, *http.Response, error) {
	c := s.client
	req, err := c.NewRequest(http.MethodGet, "open_api/signal/v1/signals", opt)
	if err != nil {
		return nil, nil, err
	}
	sig := &Signals{}
	resp, err := c.Do(req, sig)
	if err != nil {
		return nil, resp, err
	}
	return sig, resp, err
}
