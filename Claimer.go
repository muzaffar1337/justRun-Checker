package main

import (
	"fmt"
	"justRun-Checker/pkg"
	"justRun-Checker/pkg/Discord"
	"strings"
	"sync"
	"time"

	"github.com/valyala/fasthttp"
)

var Acwg sync.WaitGroup

func Claimed(user string) {
	els := time.Since(Start)
	for _, Acc := range Variables.Accounts {
		Variables.ClaimingLogger = fmt.Sprintf("Trying with fbid:%s", Acc.Fbid)
		req := FastAccCenter()
		res := fasthttp.AcquireResponse()
		req.Header.Set("Cookie", "sessionid="+Acc.SessionID)
		requestBody := "fb_dtsg=" + Acc.FB_Dtsg +
			"&variables=%7B%22client_mutation_id%22%3A%220a528653-8d8d-4f8d-b17d-e4ed9837e351%22%2C%22family_device_id%22%3A%22device_id_fetch_ig_did%22%2C%22identity_ids%22%3A%5B%22" +
			Acc.Fbid + "%22%5D%2C%22target_fx_identifier%22%3A%22FXACINFRAOBIDPERVIEWERAVPtUBICmBiZf5G0POEi_EUi2YhS-zqGG4gjQRxEIy2f4s3DwkVFYL5T577JIQiFcvahwL3fAuBZkLwpZ9qkzrWW6jaC%22%2C%22username%22%3A%22" +
			user + "%22%2C%22interface%22%3A%22IG_WEB%22%7D&doc_id=7081659258589245"
		req.SetBodyString(requestBody)
		client := &fasthttp.Client{
			MaxIdleConnDuration: 5 * time.Second,
			MaxConnsPerHost:     1000,
			ReadBufferSize:      8192,
		}
		if err := client.Do(req, res); err == nil {
			if strings.Contains(string(res.Body()), `"error":null,"ui_response"`) {
				Variables.ClaimingLogger = fmt.Sprintf("\n@[%s](fg:green) Claimed in %s", user, els.String())
				description := fmt.Sprintf("[@%s](https://www.instagram.com/%s) Claimed in %s, Att:%d", user, user, els.String(), Counter.Attempt)
				Discord.SendMessage(Settings.WebHook, Discord.Message{
					Embeds: &[]Discord.Embed{{
						Description: &description,
					}},
				})
				Writefile("@"+user+".txt", fmt.Sprint("Session : ", fmt.Sprint(Acc.SessionID), "\nFbid : ", fmt.Sprint(Acc.Fbid), "\nFB_Dtsg : ", fmt.Sprint(Acc.FB_Dtsg)))
				Accounts = RemoveSession(Accounts, Acc)
				pkg.RemoveFromFile(Acc.SessionID, &Variables.Sessions, "Files/Accounts.txt")
				break
			} else if strings.Contains(string(res.Body()), `"errorSummary":"Sorry`) {
				Accounts = RemoveSession(Accounts, Acc)
				pkg.RemoveFromFile(Acc.SessionID, &Variables.Sessions, "Files/Accounts.txt")
			}
		}
	}
	claiming = false // stop claiming start checking
}
