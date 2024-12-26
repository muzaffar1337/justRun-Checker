package main

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
