package persist

import (
	"net/http"
	"net/url"
	"io/ioutil"
	"os"
	"path/filepath"
	"fmt"
)

func Persisit(response *http.Response, url *url.URL) {
	html, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	pwd, _ := os.Getwd()
	p := filepath.FromSlash(pwd + "/result/" + url.Host + ".html")

	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	n, err := f.Write(html)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d bytes writen\n", n)
}