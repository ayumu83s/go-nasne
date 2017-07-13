package nasne

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/url"
)

type Client struct {
	baseURL    *url.URL
	httpClient *http.Client

	common   service
	Status   *StatusService
	Recorded *RecordedService
}

type service struct {
	client *Client
}

func NewClient(ip string, httpClient *http.Client) (*Client, error) {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	validateIP := net.ParseIP(ip)
	if validateIP == nil {
		err := fmt.Errorf("failed to parse ip: %s", ip)
		return nil, err
	}

	URL := &url.URL{
		Scheme: "http",
		Host:   validateIP.String(),
	}

	client := &Client{
		baseURL:    URL,
		httpClient: httpClient,
	}

	client.common.client = client
	client.Status = (*StatusService)(&client.common)
	client.Recorded = (*RecordedService)(&client.common)
	return client, nil
}

func (client *Client) Request(ctx context.Context, req *http.Request) (*http.Response, error) {
	// TODO
	// if ctx != nil {
	// 	return ctxhttp.Do(ctx, client.httpClient, req)
	// }
	return client.httpClient.Do(req)
}

func (client *Client) Get(ctx context.Context, endpoint string) (*http.Response, error) {
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}
	return client.Request(ctx, req)
}

func decodeBody(res *http.Response, out interface{}) error {
	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)
	return decoder.Decode(out)
}
