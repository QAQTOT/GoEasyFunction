package easy_request

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"reflect"
)

func PostForm(host, uri string, data map[string]string) (string, error) {

	c := &http.Client{}

	form := url.Values{}
	for key, val := range data {
		form.Add(key, val)
	}

	req, err := http.NewRequest("POST", host+uri, bytes.NewBufferString(form.Encode()))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	response, err := c.Do(req)
	if err != nil {
		return "", err
	}

	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)

	return string(body), err
}

func PostJson(host, uri string, data any) (string, error) {

	t := reflect.TypeOf(data)

	buffer := bytes.NewBuffer(nil)

	switch t.Kind() {
	case reflect.Struct, reflect.Map:
		marshal, err := json.Marshal(data)
		if err != nil {
			return "", err
		}
		buffer.Write(marshal)
		break
	case reflect.String:
		buffer.WriteString(data.(string))
		break
	default:
		panic("unhandled default case")
	}

	c := &http.Client{}

	req, err := http.NewRequest("POST", host+uri, buffer)
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	response, err := c.Do(req)
	if err != nil {
		return "", err
	}

	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	return string(body), err
}
