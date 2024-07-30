package main

import (
	"crypto/tls"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpproxy"
)

func FastAccCenter() *fasthttp.Request {
	var Req *fasthttp.Request = fasthttp.AcquireRequest()
	Req.SetRequestURI("https://accountscenter.instagram.com/api/graphql/")
	Req.Header.SetMethod(fasthttp.MethodPost)
	Req.Header.SetContentType("application/x-www-form-urlencoded")
	Req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:125.0) Gecko/20100101 Firefox/125.0")
	Req.Header.Set("X-Csrftoken", "mfr")
	Req.Header.Set("Sec-Fetch-Site", "same-origin")
	return Req
}

func username_suggestions() *fasthttp.Request {
	Request := fasthttp.AcquireRequest()
	Request.Header.SetMethod(fasthttp.MethodPost)
	Request.SetRequestURI("https://i.instagram.com/api/v1/accounts/username_suggestions/")
	Request.Header.SetHost("i.instagram.com")
	Request.Header.SetUserAgent("Instagram 212.0.0.38.119 Android (25/7.1.2; 300dpi; 1600x900; Asus; ASUS_Z01QD; ASUS_Z01QD; intel; en_US; 329675731)")
	Request.Header.SetContentType("application/x-www-form-urlencoded; charset=UTF-8")
	Request.Header.Add("Accept-Language", "en;q=0.9")
	Request.Header.Add("X-CSRFTOKEN", "missing")
	return Request
}

func Checkhosts() *fasthttp.Request {
	Request := fasthttp.AcquireRequest()
	Request.Header.SetMethod(fasthttp.MethodPost)
	Request.SetRequestURI(Variables.host)
	Request.Header.SetHost(filterHost(Variables.host))
	Request.Header.SetUserAgent("Instagram 212.0.0.38.119 Android (25/7.1.2; 300dpi; 1600x900; Asus; ASUS_Z01QD; ASUS_Z01QD; intel; en_US; 329675731)")
	Request.Header.SetContentType("application/x-www-form-urlencoded; charset=UTF-8")
	Request.Header.Add("Accept-Language", "en;q=0.9")
	Request.Header.Add("X-CSRFTOKEN", "missing")
	return Request
}

func filterHost(host string) string {

	host = strings.TrimPrefix(host, "http://")
	host = strings.TrimPrefix(host, "https://")

	host = strings.Split(host, "/")[0]
	return host
}

func isValidSession(sessionID string) bool {
	req, err := http.NewRequest("GET", "https://i.instagram.com/api/v1/accounts/current_user/?edit=true", nil)
	if err != nil {
		return false
	}

	req.Header.Set("Host", "i.instagram.com")
	req.Header.Set("Cookie", "sessionid="+sessionID)
	req.Header.Set("User-Agent", "Instagram 275.0.0.27.98 Android")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {

		_ = resp.Body.Close()
		return false
	}

	if resp.StatusCode == http.StatusOK {
		_ = resp.Body.Close()
		return true
	} else if resp.StatusCode == http.StatusForbidden {
		_ = resp.Body.Close()
		return false
	}
	_ = resp.Body.Close()
	return false
}

func createClient() (*fasthttp.Client, error) {
	client := &fasthttp.Client{
		Dial: fasthttpproxy.FasthttpHTTPDialer(Proxy_Sorter()),
		TLSConfig: &tls.Config{
			InsecureSkipVerify: true, // Disable SSL verification (not recommended for production)
		},
		MaxConnsPerHost:     1000,
		MaxIdleConnDuration: 10 * time.Second,
		ReadBufferSize:      8192,
	}

	return client, nil
}

func RandomString() string {
	Random := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	Letters := []rune("abcdefghijklmnopqrstuwxyz123456789_")
	B := make([]rune, 23)
	for i := range B {
		B[i] = Letters[Random.Intn(len(Letters))]
	}
	return string(B)
}
