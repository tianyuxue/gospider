package main

import (
	"log"
	"os"
	"bufio"
	"path/filepath"
	"net/url"
	"gospider/model"
	"gospider/fetch"
)

func loadUrl() {
	pwd, _ := os.Getwd()
	p := filepath.FromSlash(pwd+"/resources/urls.txt")
	file, err := os.Open(p)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		rawUrl := scanner.Text()
		addr, err := url.Parse(rawUrl)
		if err != nil {
			log.Fatal(err)
		}
		model.Tasks.Push(addr)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	loadUrl()
	for !model.Tasks.IsEmpty() {
		addr := model.Tasks.Pop()
		fetch.Fetch(addr)
	}
	return
}