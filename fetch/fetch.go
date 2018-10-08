package fetch

import (
	"net/url"
	"net/http"
	"gospider/persist"
)

func Fetch(addr *url.URL) {
	response, err := http.Get(addr.String())
	if err != nil {
		panic(err)
	}
	persist.Persisit(response, addr)
}