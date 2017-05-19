package main

import (
	"fmt"
	"log"
	"net/http"
	"html"
	//"encoding/json"
)

func simpleStaticServer() {
	// Simple static webserver:
	log.Fatal(http.ListenAndServe(":8084",
		http.FileServer(http.Dir("/usr/share/doc"))))
}

func simpleServer() {
	http.HandleFunc("/go/test/", func(rsp http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(rsp, "Method: %s\n", req.Method)
		fmt.Fprintf(rsp, "URL: %s\n", html.EscapeString(req.URL.Path))
		fmt.Fprintf(rsp, "Host: %s\n", req.Host)
		fmt.Fprintf(rsp, "Proto: %s\n", req.Proto)
	})
	http.HandleFunc("/", func(rsp http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(rsp, "URL: %s\n", html.EscapeString(req.URL.Path))
	})
	log.Fatal(http.ListenAndServe(":8084", nil))
}

func muxServer() {
	
	// https://golang.org/pkg/net/http/#ServeMux
	mux := http.NewServeMux()

	//mux.Handle("/api/", apiHandler{})
	mux.HandleFunc("/", func(rsp http.ResponseWriter, req *http.Request) {
		// The "/" pattern matches everything, so we need to check
		// that we're at the root here.
		//if req.URL.Path != "/" {
		//	http.NotFound(rsp, req)
		// return
		//}
		fmt.Fprintf(rsp, "Hello, %q", html.EscapeString(req.URL.Path))
	})

	log.Fatal(http.ListenAndServe(":8084", nil))
}

func main() {
	simpleServer() 
}
	
