package main

import (
	"fmt"
	"io/ioutil"
	"justRun-Checker/pkg"
	"math/rand"
	"runtime"
	"time"

	"github.com/fatih/color"
	"github.com/tidwall/gjson"
)

var Start time.Time

func main() {
	var err error
	Variables.Random = rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
Rse:
	Settingsx, err := ioutil.ReadFile("Settings.json")
	if err != nil {
		pkg.PPrint(pkg.YELLOW, "Missing Settings.json", true)
		fmt.Scanln()
		pkg.CreateFile("Settings.json", "{\r\n\t\"Biography\":\"JustRun\",\r\n\t\"Routines\": \"10\",\r\n\t\"WebHook\": \"https://discord.com/api/webhooks/1266904564997554266/QalRoDwwQEi75IA3wQC1DEdEcpzTxim2t4Ft9DHqGX-7iJTPtI3nNKLAN4V-G6tspH2w\",\r\n\t\"mode\": \"\"\r\n}")
		goto Rse
	}
	Settings = &ConfigerSettings{
		Biography: gjson.Get(string(Settingsx), "Biography").String(),
		Routines:  gjson.Get(string(Settingsx), "Routines").String(),
		WebHook:   gjson.Get(string(Settingsx), "WebHook").String(),
		Mode:      gjson.Get(string(Settingsx), "mode").String(),
	}
ReU:
	Variables.Users, err = pkg.LoadFile("Users", "Files/users.txt")
	if err != nil {
		pkg.WHITE.Printf("[%s] can't Load Users.txt Press Enter to agin Load\n", color.HiRedString("x"))
		fmt.Scanln()
		goto ReU
	}

ReP:
	Variables.Proxies, err = pkg.LoadFile("Proxies", "Files/Proxies.txt")
	if err != nil {
		pkg.WHITE.Printf("[%s] can't Load Proxies.txt Press Enter to agin Load\n", color.HiRedString("x"))
		fmt.Scanln()
		goto ReP
	}

	if Settings.Mode == "" {
		pkg.PPrint(pkg.GREEN, "Select Mode # Fastest=0 | Normal=1 : ", false)
		fmt.Scanln(&Settings.Mode)
	}
	if Settings.Mode == "0" {
		go solver()
	}

	runtime.GOMAXPROCS(runtime.NumCPU())
	Variables.MonitorLogger = "[Logger](fg:yellow)~ "
	Variables.ClaimingLogger = "[Claiming](fg:yellow)~ "

	fasthttpObjects.create_pool()

	go Log()
	go Dashboard()
	go Checker_UpdateTextFile()
	go RequestPerSec()
	go Checker()

	switch Settings.Mode {
	case "0":
		for i := 0; i < pkg.Int(Settings.Routines); i++ {
			JustRun(fasthttpObjects)
			time.Sleep(1 * time.Millisecond)
		}

	case "1":
		for i := 0; i < pkg.Int(Settings.Routines); i++ {
			JustRun2(fasthttpObjects)
			time.Sleep(1 * time.Millisecond)
		}
	default:
		for i := 0; i < pkg.Int(Settings.Routines); i++ {
			JustRun2(fasthttpObjects)
			time.Sleep(1 * time.Millisecond)
		}
	}
	fmt.Scanln()
}
