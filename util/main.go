package util

import (
	"net/http"
	"net/url"
)

var Transport *http.Transport
var HttpClient *http.Client

func SetProxy(proxy *url.URL) {
	Transport.Proxy = http.ProxyURL(proxy)
}

func init() {
	Transport = &http.Transport{}
	HttpClient = &http.Client{Transport: Transport}
}
