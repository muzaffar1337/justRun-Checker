package main

import (
	"math/rand"
	"sync"

	"github.com/gizak/termui/v3/widgets"
	"github.com/valyala/fasthttp"
)

var (
	lock sync.Mutex

	claiming bool = false
)

type Claiming struct {
	Username string
	Wg       sync.WaitGroup
}

var (
	p0 = widgets.NewParagraph()
	p2 = widgets.NewParagraph()
)
var (
	Variables = &Global_Variable{}
	Counter   = &Counters{}
	Settings  = &ConfigerSettings{}
	Accounts  []AccountsInfo
	// fasthttpObjects = &Objects{}
	Clients_checker  []*fasthttp.Client
	Clients_checker2 []*fasthttp.Client
	Clients_checker3 []*fasthttp.Client
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
type Global_Variable struct {
	Random                                                   *rand.Rand
	Users, Proxies, Sessions, Hosts                          []string
	Proxy                                                    []fasthttp.DialFunc
	host                                                     string
	Accounts                                                 []AccountsInfo
	Try, Avliable, MonitorLogger, ClaimingLogger, Hostlogger string
}
