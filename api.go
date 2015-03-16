package api

import (
	"errors"
	"fmt"
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

type Parameters struct {
	Sort      string
	Direction string
	PerPage   int
}

func (p *Parameters) sortBy() (string, error) {
	if p.Sort == "" || p.Sort == "created_at" {
		return "created_at", nil
	}
	return "", errors.New("You can only sort by created_at")
}

func (p *Parameters) direction() (string, error) {
	if p.Direction == "" || p.Direction == "desc" {
		return "desc", nil
	} else if p.Direction == "asc" {
		return "asc", nil
	}
	return "", errors.New("You can only sort 'asc' or 'desc'")
}

func (p *Parameters) perPage() (int, error) {
	if p.PerPage == 0 {
		return 30, nil
	}
	if p.PerPage < 0 {
		return 0, errors.New("You cannot show less than 1 result")
	} else {
		return p.PerPage, nil
	}
}

func (p *Parameters) paramString() (string, error) {
	sort, err := p.sortBy()
	if err != nil {
		return "", err
	}
	direction, err := p.direction()
	if err != nil {
		return "", err
	}
	perPage, err := p.perPage()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("?sort=%s&direction=%s&per_page=%v", sort, direction, perPage), nil

}
