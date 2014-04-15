package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-martini/martini"
)

type Encoder interface {
	Encode(v ...interface{}) (string, error)
}

func Must(data string, err error) string {
	if err != nil {
		panic(err)
	}
	return data
}

type jsonEncoder struct{}

func (_ jsonEncoder) Encode(v ...interface{}) (string, error) {
	var data interface{} = v
	if v == nil {
		// So that empty results produces `[]` and not `null`
		data = []interface{}{}
	} else if len(v) == 1 {
		data = v[0]
	}
	b, err := json.Marshal(data)
	return string(b), err
}

func MapEncoder(c martini.Context, w http.ResponseWriter, r *http.Request) {
	c.MapTo(jsonEncoder{}, (*Encoder)(nil))
	w.Header().Set("Content-Type", "application/json")
}
