package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

// To run: go run ex_1.8_1.9_fetch_url.go www.guimp.com
func main() {
	for _, url := range os.Args[1:] {

		if !strings.HasPrefix(url, "http") && !strings.HasPrefix(url, "https") {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}

		fmt.Printf("%s\n%s", resp.Status, b)
	}
}
