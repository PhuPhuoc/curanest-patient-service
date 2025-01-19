package common

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type RequestOptions struct {
	Method  string      // "GET", "POST", "PUT", "DELETE"
	URL     string      // Endpoint
	Payload interface{} // data payload (optional)
	Token   string      // Token Bearer (optional)
}

type APIResponse struct {
	StatusCode int
	Body       []byte
}

func CallExternalAPI(ctx context.Context, opts RequestOptions) (*APIResponse, error) {
	if opts.Method == "" {
		return nil, errors.New("HTTP method is required")
	}
	if opts.URL == "" {
		return nil, errors.New("URL is required")
	}

	// mapping payload
	var body io.Reader
	if opts.Payload != nil {
		jsonData, err := json.Marshal(opts.Payload)
		if err != nil {
			return nil, err
		}
		body = bytes.NewBuffer(jsonData)
	}

	// create request
	req, err := http.NewRequestWithContext(ctx, opts.Method, opts.URL, body)
	if err != nil {
		return nil, err
	}

	// add headers
	req.Header.Set("Content-Type", "application/json")
	if opts.Token != "" {
		req.Header.Set("Authorization", "Bearer "+opts.Token)
	}

	// send request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// read response body
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return &APIResponse{
		StatusCode: resp.StatusCode,
		Body:       respBody,
	}, nil
}
