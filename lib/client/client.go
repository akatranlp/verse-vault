package client

import "net/http"

type Client interface {
	Do(r *http.Request) (*http.Response, error)
}
