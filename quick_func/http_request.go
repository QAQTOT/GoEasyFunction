package quick_func

import (
	"net/url"
	"strings"
)

func HttpBuildQuery(params map[string]string) string {
	var queryStr []string
	for k, v := range params {
		queryStr = append(queryStr, url.QueryEscape(k)+"="+url.QueryEscape(v))
	}
	return strings.Join(queryStr, "&")
}
