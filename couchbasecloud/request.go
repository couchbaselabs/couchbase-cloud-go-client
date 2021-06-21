package couchbasecloud

import (
	"net/url"
	"time"
)

const (
	contentTypeJSON = "application/json; charset=utf-8"
)

type request struct {
	endpoint        string
	timeout         time.Duration
	method          string
	contentType     string
	queryParameters url.Values
	body            []byte
}
