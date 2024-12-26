package main

import (
	"math/rand"
	"sync"

	"github.com/gizak/termui/v3/widgets"
)

var (
	lock sync.Mutex
)

var (
	p0 = widgets.NewParagraph()
	p2 = widgets.NewParagraph()
)
var (
	Variables       = &Global_Variable{}
	Counter         = &Counters{}
	Settings        = &ConfigerSettings{}
	Accounts        []AccountsInfo
	Synchronise     = &synchronization{}
	fasthttpObjects = &Objects{}
)

type ConfigerSettings struct {
	Routines  string
	Biography string
	WebHook   string
	Mode      string
}
type AccountsInfo struct {
	Account   string
	SessionID string
	Fbid      string
	FB_Dtsg   string
}
type Counters struct {
	Attempt          int32
	RateLimit        int32
	RequestPerSec    int32
	ProxyError       int32
	AccountManager   int32
	UsernamesManager int32
	ProxyManager     int32
}

type synchronization struct {
	Swappedmutex sync.Mutex
	WG           sync.WaitGroup
}

type Global_Variable struct {
	Random                                                   *rand.Rand
	Users, Proxies, Sessions, Hosts                          []string
	Proxy                                                    string
	Accounts                                                 []AccountsInfo
	Try, Avliable, MonitorLogger, ClaimingLogger, Hostlogger string
}
