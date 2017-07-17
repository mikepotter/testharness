package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"os/signal"
	"syscall"
)

type handler struct{}

var (
	addr     = flag.String("addr", "http://127.0.0.1:8080", "listen address")
	listener net.Listener
	err      error
)

func main() {
	flag.Parse()
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, os.Interrupt, os.Kill, syscall.SIGTERM)

	h := &handler{}
	s := http.Server{Handler: h}

	url, err := url.Parse(*addr)
	if err != nil {
		log.Fatal(err)
	}

	proto := "tcp"
	laddr := url.Host

	if url.Scheme == "unix" {
		proto = "unix"
		laddr = fmt.Sprintf("%s%s", url.Host, url.RequestURI())
	}
	fmt.Printf("Listening on %s\n", laddr)
	listener, err := net.Listen(proto, laddr)
	if err != nil {
		log.Fatal(err)
	}

	go func(c chan os.Signal) {
		// Wait for a SIGINT or SIGKILL:
		sig := <-c
		log.Printf("Caught signal %s: shutting down.", sig)
		listener.Close()
		os.Exit(0)
	}(sigc)

	log.Fatal(s.Serve(listener))
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	dump, _ := httputil.DumpRequest(r, true)
	w.Write(dump)
}
