package main

import (
	"flag"
	"log"
	"net/http"
	"net/http/httputil"
)

func main() {
	port := flag.String("p", "8080", "listen port")
	flag.Parse()

	http.HandleFunc("/", handler)
	log.Println("Listening on localhost:", *port)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	dump, _ := httputil.DumpRequest(r, true)
	w.Write(dump)
}
