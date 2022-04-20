package main

import (
	"net/http"
	"time"
)

type GetSignalsOptions struct {
	StartDate Time `url:"start_date"`
	EndDate   Time `url:"end_date"`
}

type Signals struct {
	Signal []struct {
		StartDate time.Time `json:"start_date"`
		EndDate   time.Time `json:"end_date"`
		Type      string    `json:"type"`
		Values    []struct {
			StartDate   time.Time `json:"start_date"`
			EndDate     time.Time `json:"end_date"`
			Value       bool      `json:"value"`
			UpdatedDate time.Time `json:"updated_date"`
		} `json:"values"`
	} `json:"signals"`
}

func (c *Client) GetSignals(opt *GetSignalsOptions) (*Signals, *http.Response, error) {

	req, err := c.NewRequest("GET", "open_api/signal/v1/signals", opt)

	if err != nil {
		return nil, nil, err
	}
	var sig *Signals
	resp, err := c.Do(req, &sig)
	if err != nil {
		return nil, resp, err
	}

	return sig, resp, err
}
