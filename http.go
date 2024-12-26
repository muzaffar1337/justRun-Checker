package main

import (
	"crypto/tls"
	"math/rand"
	"sync"
	"time"

	"github.com/valyala/fasthttp"
)

type Objects struct {
	client   *sync.Pool
	request  *sync.Pool
	response *sync.Pool
}

func (vars *Objects) create_pool() {
	vars.client = &sync.Pool{
		New: func() interface{} {
			return &fasthttp.Client{
				TLSConfig: &tls.Config{
					InsecureSkipVerify: true,
				},
				MaxConnsPerHost:     1000,
				MaxIdleConnDuration: 10 * time.Second,
				ReadBufferSize:      8192,
			}
		},
	}

	vars.request = &sync.Pool{
		New: func() interface{} {
			return fasthttp.AcquireRequest()
		},
	}

	vars.response = &sync.Pool{
		New: func() interface{} {
			return fasthttp.AcquireResponse()
		},
	}
}

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

func CheckUserReq(vars *Objects) *fasthttp.Request {
	Request := vars.request.Get().(*fasthttp.Request)
	Request.Header.SetMethod(fasthttp.MethodPost)
	Request.SetRequestURI("https://i.instagram.com/api/v1/accounts/create/")
	Request.Header.SetHost("i.instagram.com")
	Request.Header.SetUserAgent("Instagram 275.0.0.27.98 Android")
	Request.Header.SetContentType("application/x-www-form-urlencoded; charset=UTF-8")
	Request.Header.Add("Accept-Language", "en;q=0.9")
	Request.Header.Add("X-CSRFTOKEN", "missing")
	return Request
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
