package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
)

func parse(source *http.Response) (line string) {
	// read response to string
	sBytes, err := ioutil.ReadAll(source.Body)
	if err != nil {
		return "-1"
	}
	s := string(sBytes)
	// parse output to get description of query if it exists
	var re = regexp.MustCompile(`(?m)<div class="kno-rdesc r-iw3x82clRHMU".+jsl=".+"><span>(.+)<span>`)
	//indx := re.FindStringIndex(s)

	if len(re.FindStringIndex(s)) > 0 {
		fmt.Println(re.FindString(s), "found at index", re.FindStringIndex(s)[0])
	} else {
		fmt.Println(len(re.FindStringIndex(s)))
	}
	return ""
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
	resp, err := http.Get("http://www.google.sk/search?q=" + queryStr)
	fmt.Println("asked for: " + queryStr)
	if err != nil {
		return
	}
	fmt.Println("Beginning parse:")
	parse(resp)
	fmt.Println("End of parse")
	defer resp.Body.Close()
}
