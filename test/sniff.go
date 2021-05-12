package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
)

func main() {

	opdsRootHandler := func(w http.ResponseWriter, req *http.Request) {
		qu, _ := url.QueryUnescape(req.URL.String())
		fmt.Printf("Header:\n%#v", req.Header)
		fmt.Printf("DUMB->URL: [%s]\n", qu)
	}

	http.HandleFunc("/opds", opdsRootHandler)
	log.Fatal(http.ListenAndServe(":8085", nil))
}
