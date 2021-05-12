package main

import (
	"flag"
	"fmt"
	"net/http"
	"tt/fb"

	"github.com/golang/glog"
)

func main() {
	port := flag.Int("port", 8080, "http port")
	flag.Parse()
	glog.Info("Starting server at localhost:", *port)
	handler := &fb.FB{}
	glog.Fatal(http.ListenAndServe(fmt.Sprint(":", *port), handler))
}
