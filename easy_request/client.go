package easy_request

import (
	"bytes"
	"encoding/json"
	"github.com/QAQTOT/go_easy_function/quick_func"
	"io"
	"net/http"
	"net/url"
	"time"
)

type Client struct {
	Host   string
	Uri    string
	Method string
	Header map[string]string
	Data   map[string]string
	Client *http.Client
}

func NewClient() *Client {
	return &Client{
		Client: &http.Client{
			Timeout: 60 * time.Second,
		},
	}
}

func (this *Client) Request(formType uint) (string, error) {
	var params *bytes.Buffer
	switch formType {
	case 1: // url params
		params = nil
		this.Uri += "?" + quick_func.HttpBuildQuery(this.Data)
		break
	case 2: // form
		form := url.Values{}
		for key, val := range this.Data {
			form.Add(key, val)
		}
		params = bytes.NewBufferString(form.Encode())
		break
	case 3: // json
		marshal, err := json.Marshal(this.Data)
		if err != nil {
			return "", err
		}
		params.Write(marshal)
		break
	}

	req, err := http.NewRequest(this.Method, this.Host+this.Uri, params)
	if err != nil {
		return "", err
	}

	for k, v := range this.Header {
		req.Header.Set(k, v)
	}

	response, err := this.Client.Do(req)
	if err != nil {
		return "", err
	}

	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	return string(body), err
}

func (this *Client) SetHeader(headers map[string]string) *Client {
	this.Header = headers
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

func (this *Client) SetData(data map[string]string) *Client {
	this.Data = data
	return this
}
