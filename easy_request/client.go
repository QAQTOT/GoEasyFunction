package easy_request

import (
	"net/http"
	"time"
)

type Client struct {
	Host   string
	Uri    string
	Method string
	header []string
	Data   any
	Client *http.Client
}

func NewClient() *Client {
	return &Client{
		Client: &http.Client{
			Timeout: 60 * time.Second,
		},
	}
}

func (this *Client) SetHeader(headers []string) *Client {
	this.header = headers
	return this
}

func (this *Client) SetMethod(method string) *Client {
	this.Method = method
	return this
}

func (this *Client) SetUri(uri string) *Client {
	this.Uri = uri
	return this
}

func (this *Client) SetHost(host string) *Client {
	this.Host = host
	return this
}

func (this *Client) SetTimeout(timeout time.Duration) *Client {
	this.Client.Timeout = timeout
	return this
}

func (this *Client) SetData(data any) *Client {
	this.Data = data
	return this
}
