package main

import (
	"log"
	"os"
	"bufio"
	"path/filepath"
	"net/http"
	"io/ioutil"
	"gospider/model"
	"strings"
	"fmt"
)
var tasks = model.Stack{}

func fetch(url string) {
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	persisit(response, url)
}

func persisit(response *http.Response, url string) {
	html, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	pwd, _ := os.Getwd()
	p := filepath.FromSlash(pwd + "/result/" + strings.Split(url, "//")[1] )

	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	n, err := f.Write(html)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d bytes writen", n)
}

func loadUrl() {
	pwd, _ := os.Getwd()
	p := filepath.FromSlash(pwd+"/config/urls.txt")
	file, err := os.Open(p)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		tasks.Push(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	loadUrl()
	for !tasks.IsEmpty() {
		url := tasks.Pop()
		fetch(url)
	}
	return
}