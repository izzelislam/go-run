package lib

import "net/url"

func UrlParser() {
	var urlString = "https://bandanbesar.com/berat?nama=john due&age=10"
	u, e := url.Parse(urlString)

	if e != nil {
		println(e.Error())
		return
	}

	println(u.Host)
	println(u.Path)
	println(u.Scheme)
	println(u.RawQuery)
	println(u.Query()["nama"][0])
	println(u.Query()["age"][0])
}