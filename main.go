package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/skratchdot/open-golang/open"
)

func getIP(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "FROM WHERE MY GITHUB IS ACCED !")
	ip, port, err := net.SplitHostPort(req.RemoteAddr)
	if err != nil {
		fmt.Fprintf(w, "userip: %q is not IP:port", req.RemoteAddr)
	}
	userIP := net.ParseIP(ip)
	if userIP == nil {
		fmt.Fprintf(w, "userip: %q is not IP:port", req.RemoteAddr)
		return
	}

	// This will only be defined when site is accessed via non-anonymous proxy
	// and takes precedence over RemoteAddr
	// Header.Get is case-insensitive
	forward := req.Header.Get("X-Forwarded-For")

	fmt.Fprintf(w, "<p>IP: %s</p>", ip)
	fmt.Fprintf(w, "<p>Port: %s</p>", port)
	fmt.Fprintf(w, "<p>Forwarded for: %s</p>", forward)
}

func main() {
	myport := strconv.Itoa(10002)
	r := httprouter.New()
	r.GET("/", getIP)
	r.GET("/test", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		fmt.Fprint(w, "Welcome!\n")
	})
	l, err := net.Listen("tcp", "0.0.0.0:"+myport)
	if err != nil {
		log.Fatal(err)
	}
	err = open.Start("http://0.0.0.0:" + myport + "/")
	if err != nil {
		log.Println(err)
	}
	log.Fatal(http.Serve(l, r))
}
