package api

import (
	"net/http"
)

const bugsnagApi = "https://api.bugsnag.com"

type Connection struct {
	client *JSONClient
}

func New(creds *Credentials) *Connection {
	client := http.DefaultClient

	return &Connection{
		client: &JSONClient{
			Credentials: *creds,
			Client:      client,
			Endpoint:    bugsnagApi,
		},
	}
}
