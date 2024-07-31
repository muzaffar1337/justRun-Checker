package main

import (
	"fmt"
	"justRun-Checker/pkg"
	"justRun-Checker/pkg/Discord"
	"os"

	"strings"
	"sync/atomic"
	"time"

	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpproxy"
)

func JustRun(vars *Objects) {
	go CheckUser(vars)
	go CheckUser2(vars)
	go CheckUser3(vars)
}

func JustRun2(vars *Objects) {
	go CheckUser(vars)
	go CheckUser2(vars)
}

func CheckUser(vars *Objects) {
	BodyC := ""
	for {
		if claiming {
			continue
		}
		var (
			Client   = vars.client.Get().(*fasthttp.Client)
			req      = username_suggestions(vars)
			res      = vars.response.Get().(*fasthttp.Response)
			Username = Users_Sorter()
		)
		req.SetBodyString(fmt.Sprintf("name=%s", Username))
		Client.Dial = fasthttpproxy.FasthttpHTTPDialer(Proxy_Sorter())
		if err := Client.Do(req, res); err == nil {
			BodyC = string(res.Body())
			// fmt.Println("User: ", Username)
			// fmt.Println(BodyC)
			atomic.AddInt32(&Counter.Attempt, 1)
			if strings.Contains(BodyC, fmt.Sprintf(`"username":"%s"`, Username)) {
				Variables.Avliable = Username
			} else if strings.Contains(BodyC, `Please wait a few minutes before you try again.`) || res.StatusCode() != 200 {
				atomic.AddInt32(&Counter.RateLimit, 1)
			}
		} else {
			atomic.AddInt32(&Counter.ProxyError, 1)
		}
		req.Reset()
		res.Reset()
		vars.request.Put(req)
		vars.response.Put(res)
		vars.client.Put(Client)
		time.Sleep(time.Millisecond)

	}
}

func CheckUser2(vars *Objects) {
	BodyC := ""
	for {
		if claiming {
			continue
		}

		if int(Counter.UsernamesManager+1) >= len(Variables.Users) {
			Counter.UsernamesManager = 0
		}
		var (
			Client    = vars.client.Get().(*fasthttp.Client)
			req       = username_suggestions(vars)
			res       = vars.response.Get().(*fasthttp.Response)
			username  = Users_Sorter()
			username2 = Users_Sorter()
			username3 = Variables.Users[Variables.Random.Intn(len(Variables.Users))]
		)

		req.SetBodyString(fmt.Sprintf("name=%s %s&email=%s@gmail.com", username, username2, username3))
		Client.Dial = fasthttpproxy.FasthttpHTTPDialer(Proxy_Sorter())
		if err := Client.Do(req, res); err == nil {
			BodyC = string(res.Body())
			// fmt.Printf("User1:%s ,User2:%s , User3:%s \n", username, username2, username3)
			// fmt.Println(BodyC)
			atomic.AddInt32(&Counter.Attempt, 1)
			if res.StatusCode() == 200 && !strings.Contains(BodyC, `Please wait a few minutes before you try again.`) {

				if strings.Contains(BodyC, fmt.Sprintf(`"username":"%s"`, username)) {
					Variables.Avliable = username
				}
				if strings.Contains(BodyC, fmt.Sprintf(`"username":"%s"`, username2)) {
					Variables.Avliable = username2
				}
				if strings.Contains(BodyC, fmt.Sprintf(`"username":"%s"`, username3)) {
					Variables.Avliable = username3
				}
			} else if strings.Contains(BodyC, `Please wait a few minutes before you try again.`) || res.StatusCode() != 200 {
				atomic.AddInt32(&Counter.RateLimit, 1)
			}
		} else {
			atomic.AddInt32(&Counter.ProxyError, 1)
		}
		req.Reset()
		res.Reset()
		vars.request.Put(req)
		vars.response.Put(res)
		vars.client.Put(Client)
		time.Sleep(time.Millisecond)

	}
}

func CheckUser3(vars *Objects) {

	BodyC := ""
	for {
		if claiming {
			continue
		}
		var (
			Client   = vars.client.Get().(*fasthttp.Client)
			req      = Checkhosts(vars)
			res      = vars.response.Get().(*fasthttp.Response)
			Username = Users_Sorter()
		)

		req.SetBodyString(fmt.Sprintf("username=%s", Username))

		Client.Dial = fasthttpproxy.FasthttpHTTPDialer(Proxy_Sorter())
		if err := Client.Do(req, res); err == nil {
			BodyC = string(res.Body())
			atomic.AddInt32(&Counter.Attempt, 1)

			if strings.Contains(BodyC, `"account_created": false, "errors": {"email": ["This field is required."], "device_id": ["This field is required."]}, "status": "ok"`) || strings.Contains(BodyC, `"account_created": false, "errors": {"phone_number": ["This field is required."], "device_id": ["This field is required."], "__all__": ["You need an email or confirmed phone number."]}, "existing_user": false, "status": "ok", "error_type": "required, required, no_contact_point_found"`) {
				Variables.Avliable = Username

			} else if strings.Contains(BodyC, `username_is_taken`) || strings.Contains(BodyC, `held`) {
				atomic.AddInt32(&Counter.Attempt, 1)
			} else {
				atomic.AddInt32(&Counter.RateLimit, 1)
			}
		} else {
			atomic.AddInt32(&Counter.ProxyError, 1)
		}
		req.Reset()
		res.Reset()
		vars.request.Put(req)
		vars.response.Put(res)
		vars.client.Put(Client)
		time.Sleep(time.Millisecond)
	}
}

func solver() {
	var err error
	Variables.Hosts, err = pkg.LoadFile("Hosts", "Files/hosts.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	h := "https://i.instagram.com/api/v1/accounts/create/"
	for {
		Variables.host = Hosts_Sorter()
		if Variables.host != h {
			h = Variables.host
			Variables.Hostlogger = "Host: " + Variables.host
		}
		time.Sleep(3500 * time.Millisecond)
	}
}

func Log() {
	TempLog := ""
	TempCounter := 0
	for {
		TempLog = fmt.Sprintf("%s\n%s\n", Variables.MonitorLogger, Variables.ClaimingLogger)
		Arr := strings.Split(TempLog, "\n")
		if TempCounter < len(Arr) {
			Arr = removeDuplicates(Arr)
			TempCounter = len(Arr)
			time.Sleep(1 * time.Millisecond)
		}
		TempLog = strings.Join(Arr, "\n")
		ActiveLogger = TempLog
		p2.Text = ActiveLogger
	}
}

func Checker_UpdateTextFile() {
	for {
		Variables.Users, _ = pkg.LoadFile("Users", "Files/users.txt")
		Variables.Sessions, _ = pkg.LoadFile("Accounts", "Files/Accounts.txt")
		for _, i := range Variables.Sessions {
			Acc := strings.Split(i, "|")
			go func(session string) {
				if isValidSession(Acc[0]) {
					Variables.Accounts = append(Accounts, AccountsInfo{SessionID: Acc[0], FB_Dtsg: Acc[1], Fbid: Acc[2], Account: session})
				} else {
					if len(Variables.Sessions) > 1 {
						pkg.RemoveFromFile(session, &Variables.Sessions, "Files/Accounts.txt")
					}
				}
			}(i)
			time.Sleep(12 * time.Hour)
		}
	}
}

func Checker() {
	for {
		if Variables.Avliable != "" {
			username := Variables.Avliable
			if Variables.Try != username {
				Variables.MonitorLogger = "\n@[" + username + "](fg:green) Is Avliable Now"
				lock.Lock()
				Start = time.Now()
				claiming = true
				go func() {
					description := fmt.Sprintf("[@" + username + "](https://www.instagram.com/" + username + ") Is Avliable Now")
					Discord.SendMessage(Settings.WebHook, Discord.Message{
						Embeds: &[]Discord.Embed{{
							Description: &description,
						}},
					})
				}()
				Claimed(username)
				lock.Unlock()
				Variables.Try = username // Try = x
			}
			Variables.Avliable = "" // Avliable = ""
		}
	}
}
