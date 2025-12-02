package duckdns

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

type ConnectorInterface interface {
	SetTXTRecord(ctx context.Context, domain string, txt string) (*http.Response, error)
	CleanTXTRecord(ctx context.Context, domain string) (*http.Response, error)
}

const updateURL = "https://www.duckdns.org/update"

type Connector struct {
	httpClient *http.Client
	token      string
	updateURL  string
}

func NewConnector(token string) *Connector {
	return &Connector{
		httpClient: &http.Client{Timeout: 10 * time.Second},
		token:      token,
		updateURL:  updateURL,
	}
}

func (c *Connector) SetTXTRecord(ctx context.Context, domain, txt string) (*http.Response, error) {
	return c.updateTXTRecord(ctx, domain, txt)
}

func (c *Connector) CleanTXTRecord(ctx context.Context, domain string) (*http.Response, error) {
	return c.updateTXTRecord(ctx, domain, "")
}

func (c *Connector) updateTXTRecord(ctx context.Context, domain, txt string) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, updateURL, nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("domains", domain)
	q.Add("token", c.token)
	q.Add("verbose", "true")
	q.Add("txt", txt)
	if txt == "" {
		q.Add("clear", "true")
	}
	req.URL.RawQuery = q.Encode()

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to update TXT record for domain %s: %w", domain, err)
	}

	return resp, nil
}
