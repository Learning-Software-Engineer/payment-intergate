package http

import "net/http"

type ClientWrapper struct {
	HTTPClient *http.Client
}

func NewClientWrapper(client *http.Client) *ClientWrapper {
	return &ClientWrapper{
		HTTPClient: client,
	}
}
