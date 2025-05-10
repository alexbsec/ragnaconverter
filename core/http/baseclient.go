package http

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type Method string

const (
	GET  Method = "GET"
	POST Method = "POST"
)

type Request struct {
	Client      http.Client
	Method      Method
	Host        string
	Headers     map[string]string
	QueryParams map[string]string
	Body        any
}

func NewRequest(
	client http.Client,
	method Method,
	host string,
	headers map[string]string,
	queryParams map[string]string,
	body any,
) Request {
	return Request{
		Client:      client,
		Method:      method,
		Host:        host,
		Headers:     headers,
		QueryParams: queryParams,
		Body:        body,
	}
}

func (r Request) Send() ([]byte, int, error) {
	data, err := json.Marshal(r.Body)
	if err != nil {
		return []byte{}, http.StatusInternalServerError, err
	}

	var requestBody io.Reader
	if r.Body != nil && r.Body != "" {
		requestBody = strings.NewReader(string(data))
	}

	if r.QueryParams != nil {
		r.Host = r.resolveQueryParams()
	}

	req, err := http.NewRequest(string(r.Method), r.Host, requestBody)
	if err != nil {
		return []byte{}, http.StatusInternalServerError, err
	}

	for key, value := range r.Headers {
		req.Header.Set(key, value)
	}

	res, err := r.Client.Do(req)
	if err != nil {
		if res != nil {
			return []byte{}, res.StatusCode, err
		}
		return []byte{}, http.StatusInternalServerError, err
	}

	defer func() {
		_ = res.Body.Close()
	}()

	raw, err := io.ReadAll(res.Body)
	if err != nil {
		return raw, http.StatusInternalServerError, err
	}

	return raw, res.StatusCode, nil
}

func (r Request) resolveQueryParams() string {
	params := url.Values{}

	for paramName, value := range r.QueryParams {
		if value != "" {
			params.Add(paramName, value)
		}
	}

	u, _ := url.ParseRequestURI(r.Host)
	u.RawQuery = params.Encode()
	return fmt.Sprintf("%v", u)
}
