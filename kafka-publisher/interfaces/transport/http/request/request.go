package request

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
 "Company import path goes here"

"io/ioutil"
	"net/http"
)

type EventRequest struct {
	ID        string      `json:"id"`
	Type      string      `json:"type"`
	Body      interface{} `json:"body"`
	CreatedAt int64       `json:"created_at"`
	Expiry    int64       `json:"expiry"`
	Version   int         `json:"version"`
	TraceInfo struct {
		TraceID struct {
			High int64 `json:"high"`
			Low  int64 `json:"low"`
		} `json:"trace_id"`
		SpanID   int  `json:"span_id"`
		ParentID int  `json:"parent_id"`
		Sampled  bool `json:"sampled"`
	} `json:"trace_info"`
}

type PayLoad struct {
	Key int
	Value EventRequest
}

type Tri struct {
	Body struct {
		id int `json:"id"`
	} `json:"body"`
}

type Tri2 struct {
	Body struct {
		Tri struct {
			ID int `json:"id"`
		} `json:"trip"`
	} `json:"body"`
}

func DecodeEventRequest(_ context.Context, r *http.Request) (interface{}, error) {

//"DECODING of the request code belongs here I have striped them for privacy"

	}
	request.ID = uuid.New().String()
	payLoad := PayLoad{
		Key:key,
		Value:request,
	}
	return payLoad, nil
}
