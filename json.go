package api

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// JSONClient is the underlying client for JSON APIs.
type JSONClient struct {
	Credentials Credentials
	Client      *http.Client
	Endpoint    string
}

// Do sends an HTTP request and returns an HTTP response, following policy
// (e.g. redirects, cookies, auth) as configured on the client.
func (c *JSONClient) Exec(method, uri string, req, resp interface{}) error {
	body, err := json.Marshal(req)
	if err != nil {
		return err
	}
	if !strings.Contains(uri, "http") {
		uri = c.Endpoint + uri
	}

	httpReq, err := http.NewRequest(method, uri, bytes.NewReader(body))
	if err != nil {
		return err
	}

	return c.sendMessage(httpReq, resp)
}

// Do sends an HTTP request and returns an HTTP response, following policy
// (e.g. redirects, cookies, auth) as configured on the client.
func (c *JSONClient) Query(method, uri string, resp interface{}) error {
	if !strings.Contains(uri, "http") {
		uri = c.Endpoint + uri
	}
	var body io.Reader
	httpReq, err := http.NewRequest(method, uri, body)
	if err != nil {
		return err
	}
	return c.sendMessage(httpReq, resp)
}

func (c *JSONClient) sendMessage(httpReq *http.Request, resp interface{}) error {
	httpReq.Header.Set("User-Agent", "bugsnag-go-sdk")
	httpReq.Header.Set("Content-Type", "application/json")
	if c.Credentials.ApiKey != "" {
		httpReq.Header.Set("Authorization", "token "+c.Credentials.ApiKey)
	} else {
		httpReq.SetBasicAuth(c.Credentials.Username, c.Credentials.Password)
	}

	httpResp, err := c.Client.Do(httpReq)
	if err != nil {
		return err
	}
	defer func() {
		_ = httpResp.Body.Close()
	}()

	if !isValidStatus(httpResp.StatusCode) {
		bodyBytes, err := ioutil.ReadAll(httpResp.Body)
		if err != nil {
			return err
		}
		if len(bodyBytes) == 0 {
			return APIError{
				StatusCode: httpResp.StatusCode,
				Message:    httpResp.Status,
			}
		}
		log.Println(httpResp.Status)
		var jsonErr jsonErrorResponse
		if err := json.Unmarshal(bodyBytes, &jsonErr); err != nil {
			return err
		}
		return jsonErr.Err(httpResp.StatusCode)
	}

	if resp != nil {
		return json.NewDecoder(httpResp.Body).Decode(resp)
	}
	return nil
}

func isValidStatus(statusCode int) bool {
	if statusCode == http.StatusOK || statusCode == http.StatusCreated || statusCode == http.StatusNoContent {
		return true
	}
	return false
}

type jsonErrorResponse struct {
	Type    string `json:"__type"`
	Message string `json:"message"`
}

func (e jsonErrorResponse) Err(StatusCode int) error {
	return APIError{
		StatusCode: StatusCode,
		Type:       e.Type,
		Message:    e.Message,
	}
}

type NoBody struct{}
