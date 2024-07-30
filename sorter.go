package main

// func Users_Sorter() string {
// 	Counter.UsernamesManager = (Counter.UsernamesManager + 1) % int32(len((Variables.Users)))
// 	return Variables.Users[Counter.UsernamesManager]
// }

//	func Proxy_Sorter() fasthttp.DialFunc {
//		Counter.ProxyManager = (Counter.ProxyManager + 1) % int32(len((Variables.Proxy)))
//		return Variables.Proxy[Counter.ProxyManager]
//	}
//
//	func Proxy_Sorter() string {
//		Counter.ProxyManager = (Counter.ProxyManager + 1) % int32(len((Variables.Proxies)))
//		return Variables.Proxies[Counter.ProxyManager]
//	}
func Users_Sorter() string {
	if len(Variables.Users) <= int(Counter.UsernamesManager) {

		Counter.UsernamesManager = 1
	} else {

		Counter.UsernamesManager++
	}
	return Variables.Users[Counter.UsernamesManager-1]
}

func Proxy_Sorter() string {
	if len(Variables.Proxies) <= int(Counter.ProxyManager) {
		Counter.ProxyManager = 1
	} else {
		Counter.ProxyManager++
	}
	return Variables.Proxies[Counter.ProxyManager-1]
}

func Hosts_Sorter() string {
	return Variables.Hosts[Variables.Random.Intn(len(Variables.Hosts))]
}
