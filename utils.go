package main

import (
	"fmt"
	"strings"

	"os"
	"strconv"
	"time"

	tm "github.com/buger/goterm"
	ui "github.com/gizak/termui/v3"
)

func RequestPerSec() {
	for {
		Befor := Counter.Attempt
		time.Sleep(time.Second)
		Counter.RequestPerSec = (Counter.Attempt - Befor)
	}
}

var ActiveLogger string

func commaSeparate(s uint64) string {
	q := strconv.FormatUint(s, 10)
	n := len(q)
	if n <= 3 {
		return q
	}
	i, _ := strconv.ParseUint(q[:n-3], 10, 64)
	return commaSeparate(i) + "," + q[n-3:]
}
func Dashboard() {
    tm.Clear()
    ui.Init()

    logo := `
       _                 _     _____                  
      | |               | |   |  __ \                 
      | |  _   _   ___  | |_  | |__) |  _   _   _ __  
  _   | | | | | | / __| | __| |  _  /  | | | | | '_ \ 
 | |__| | | |_| | \__ \ | |_  | | \ \  | |_| | | | | |
  \____/   \__,_| |___/  \__| |_|  \_\  \__,_| |_| |_|
                        Developer: mfr@muzaffar1337                 
    `

    logoLines := len(strings.Split(logo, "\n"))

    for {
        p0.Border = false
        p0.Text = fmt.Sprintf("%s\n[$] Att [%s](fg:green) R/S [%s](fg:green) R/L [%s](fg:red) ProxyErr [%s](fg:red) List [%s](fg:green) Accounts [%s](fg:green) Endpoints [%s](fg:blue)",
            logo,
            commaSeparate(uint64(Counter.Attempt)),
            commaSeparate(uint64(Counter.RequestPerSec)),
            commaSeparate(uint64(Counter.RateLimit)),
            commaSeparate(uint64(Counter.ProxyError)),
            commaSeparate(uint64(len(Variables.Users))),
            commaSeparate(uint64(len(Variables.Accounts))),
            commaSeparate(uint64(len(Variables.Hosts))),
        )

        // Adjust p0 to fit the logo and the status information
        p0.SetRect(0, 1, 350, 3 + logoLines + 8) // 3 lines for padding and 8 lines for the status

        // Adjust p2 to fit below p0
        p2.SetRect(0, 3 + logoLines + 8, 300, 500)

        p2.Text = ActiveLogger
        p2.Border = false

        ui.Render(p0, p2)
        time.Sleep(10 * time.Millisecond)
    }
}


// func Dashboard() {
// 	tm.Clear()
// 	ui.Init()
// 	for {
// 		p0.Border = false
// 		p0.Text = fmt.Sprintf("[$] Att [%s](fg:green) R/S [%s](fg:green) R/L [%s](fg:red) ProxyErr [%s](fg:red) List [%s](fg:green) Accounts [%s](fg:green) Endpoints [%s](fg:blue)", commaSeparate(uint64(Counter.Attempt)), commaSeparate(uint64(Counter.RequestPerSec)), commaSeparate(uint64(Counter.RateLimit)), commaSeparate(uint64(Counter.ProxyError)), commaSeparate(uint64(len(Variables.Users))), commaSeparate(uint64(len(Variables.Accounts))), commaSeparate(uint64(len(Variables.Hosts))))

// 		p0.SetRect(0, 1, 350, 200)
// 		// p2.Title = "JustRun"
// 		p2.Text = ActiveLogger
// 		p2.Border = false

// 		p2.SetRect(0, 10, 300, 500)

// 		ui.Render(p0, p2)
// 		time.Sleep(10 * time.Millisecond)
// 	}
// }

func removeDuplicates(s []string) []string {
	bucket := make(map[string]bool)
	var result []string
	for _, str := range s {
		if _, ok := bucket[str]; !ok {
			bucket[str] = true
			result = append(result, str)
		}
	}
	return result
}

func Writefile(path, text string) os.File {
	File, _ := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer File.Close()
	File.WriteString(text)
	return *File

}

func RemoveSession(List []AccountsInfo, Item AccountsInfo) []AccountsInfo {

	var newarray []AccountsInfo
	for _, s := range List {
		if s != Item {
			newarray = append(newarray, s)
		}
	}
	return newarray
}
