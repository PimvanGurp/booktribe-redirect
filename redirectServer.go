package main

import (
	"io/ioutil"
	"net/http"
)

const webshopMainUrl = "https://www.bookdepository.com?a_aid=Booktribe"

var books map[string]string = make(map[string]string)

func redirectUrl(res http.ResponseWriter, req *http.Request) {
	url := findBookRedirect(req.URL.Path[1:])
	http.Redirect(res, req, url, 302)
}

func findBookRedirect(filename string) string {
	cacheRedirect, found := books[filename]
	if found {
		return cacheRedirect
	}
	fileRedirect, err := ioutil.ReadFile("./books/" + filename + ".txt")
	if err != nil {
		return webshopMainUrl
	}
	books[filename] = string(fileRedirect)
	return books[filename]
}

func main() {
	http.HandleFunc("/", redirectUrl)
	http.ListenAndServe(":9001", nil)
}
