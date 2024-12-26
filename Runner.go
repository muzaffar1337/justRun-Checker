package main

import (
	"fmt"
	"justRun-Checker/pkg/Discord"

	"strings"
	"sync/atomic"

	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpproxy"
)

func CheckUser(vars *Objects) {
	BodyC := ""

	var (
		Client   = vars.client.Get().(*fasthttp.Client)
		req      = CheckUserReq(vars)
		res      = vars.response.Get().(*fasthttp.Response)
		Username = Users_Sorter()
	)

	req.SetBodyString(fmt.Sprintf("username=%s", Username))

	Client.Dial = fasthttpproxy.FasthttpHTTPDialer(Proxy_Sorter())
	if err := Client.Do(req, res); err == nil {
		BodyC = string(res.Body())
		if strings.Contains(BodyC, `"error_type": "required, required"`) {
			Variables.Avliable = Username
			if Variables.Avliable != Variables.Try {
				Variables.MonitorLogger = "\n@[" + Username + "](fg:green) Is Avliable Now"
				lock.Lock()
				go func() {
					description := fmt.Sprintf("[@" + Username + "](https://www.instagram.com/" + Username + ") Is Avliable Now")
					Discord.SendMessage(Settings.WebHook, Discord.Message{
						Embeds: &[]Discord.Embed{{
							Description: &description,
						}},
					})
				}()
				Claimed(Username)
				lock.Unlock()
				Variables.Try = Username
			}
		} else if strings.Contains(BodyC, `Please wait a few minutes before you try again.`) || res.StatusCode() != 200 {
			atomic.AddInt32(&Counter.RateLimit, 1)
		} else if strings.Contains(BodyC, `username_is_taken`) {
			atomic.AddInt32(&Counter.Attempt, 1)
		}
	} else {
		atomic.AddInt32(&Counter.ProxyError, 1)
	}
	req.Reset()
	res.Reset()
	vars.request.Put(req)
	vars.response.Put(res)
	vars.client.Put(Client)
}
