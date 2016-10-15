package client

import (
	"fmt"
	"golang.org/x/net/context"
	"net/http"
	"net/url"
)

// AddOperandsPath computes a request path to the add action of operands.
func AddOperandsPath(left int, right int) string {
	return fmt.Sprintf("/add/%v/%v", left, right)
}

// add returns the sum of the left and right parameters in the response body
func (c *Client) AddOperands(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewAddOperandsRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewAddOperandsRequest create the request corresponding to the add action endpoint of the operands resource.
func (c *Client) NewAddOperandsRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// MultiplyOperandsPath computes a request path to the multiply action of operands.
func MultiplyOperandsPath(left int, right int) string {
	return fmt.Sprintf("/multiply/%v/%v", left, right)
}

// add returns the product of the left and right parameters in the response body
func (c *Client) MultiplyOperands(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewMultiplyOperandsRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewMultiplyOperandsRequest create the request corresponding to the multiply action endpoint of the operands resource.
func (c *Client) NewMultiplyOperandsRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// SubtractOperandsPath computes a request path to the subtract action of operands.
func SubtractOperandsPath(left int, right int) string {
	return fmt.Sprintf("/subtract/%v/%v", left, right)
}

// subtract returns the difference of the left and right parameters in the response body
func (c *Client) SubtractOperands(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewSubtractOperandsRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewSubtractOperandsRequest create the request corresponding to the subtract action endpoint of the operands resource.
func (c *Client) NewSubtractOperandsRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
