package main

import (
	"fmt"
	"io/ioutil"
	"justRun-Checker/pkg"
	"log"
	"math/rand"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/tidwall/gjson"
)

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
ReA:
	Variables.Sessions, err = pkg.LoadFile("Accounts", "Files/Accounts.txt")
	if err != nil {
		pkg.WHITE.Printf("[%s] can't Load Accounts.txt Press Enter to agin Load\n", color.HiRedString("x"))
		fmt.Scanln()
		goto ReA
	} else {
		for _, i := range Variables.Sessions {
			Acc := strings.Split(i, "|")
			if Acc[0] != "" {
				Variables.Accounts = append(Accounts, AccountsInfo{SessionID: Acc[0], FB_Dtsg: Acc[1], Fbid: Acc[3], Account: i})
			}
		}
	}

	runtime.GOMAXPROCS(runtime.NumCPU())
	routines, err := strconv.Atoi(Settings.Routines)
	if err != nil {
		log.Fatalf("Invalid number of routines: %v", err)
	}
	fasthttpObjects.create_pool()

	go Dashboard()
	go RequestPerSec()
	Workerpool(routines, fasthttpObjects)
	select {}
}

func Workerpool(numWorkers int, vars *Objects) {
	jobQueue := make(chan int, numWorkers)

	for i := 0; i < numWorkers; i++ {
		Synchronise.WG.Add(1)
		go worker(jobQueue, vars)
	}

	go func() {
		jobID := 0
		for {
			select {
			case jobQueue <- jobID:
				jobID++
			default:
				time.Sleep(time.Millisecond)
			}
		}
	}()

	Synchronise.WG.Wait()
}

func worker(jobQueue <-chan int, vars *Objects) {
	for jobID := range jobQueue {
		CheckUser(vars)
		_ = jobID
	}
	Synchronise.WG.Done()
}
