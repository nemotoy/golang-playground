package client

import (
	"net/http"

	"context"

	"github.com/cenkalti/backoff"
)

type Client interface {
	ShouldRetryRequest(ctx context.Context, req *http.Request) (*http.Response, error)
}

type client struct {
	c *http.Client
}

// TODO: gives backoff configurations to the argument
func New(c *http.Client) Client {
	return &client{c}
}

func (c *client) ShouldRetryRequest(ctx context.Context, req *http.Request) (*http.Response, error) {
	var r *http.Response
	err := backoff.Retry(func() (err error) {
		r, err = c.c.Do(req)
		// TODO: implements error handling in detail
		if err == nil && r.StatusCode == http.StatusOK {
			return nil
		}
		r.Body.Close()
		return err
	}, backoff.WithContext(backoff.NewExponentialBackOff(), ctx))
	if err != nil {
		return nil, err
	}
	return r, nil
}
