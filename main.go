package main

import (
	"net/http"
	"os"
	"strings"
)

func parse(source *http.Response) {
	//... todo tomorrow xd
}

func main() {
	var args = os.Args[1:]
	//https://www.google.sk/search?q=abc
	var query = make([]string, len(args))
	query = args
	queryStr := strings.Join(query, " ") // i.e wet coala (ignores program path)
	if len(queryStr) < 2 {
		return
	}
	resp, err := http.Get("//https://www.google.sk/search?q=" + queryStr)
	if err != nil {
		return
	}
	parse(resp)
	defer resp.Body.Close()
}
