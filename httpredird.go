package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

// ConfigStruct stores config
type ConfigStruct struct {
	BindAddr     string `json:"bind_addr"`
	TargetURL    string `json:"target_url"`
	RedirectCode int    `json:"redirect_code"`
}

var config ConfigStruct

func init() {
	flag.StringVar(&config.TargetURL, "target", "http://127.0.0.1:1234/", "url to reditrect to")
	flag.StringVar(&config.BindAddr, "bind", ":80", "address to bind to")
	flag.IntVar(&config.RedirectCode, "code", 301, "redirect with this code")
}

func main() {
	flag.Parse()
	err := http.ListenAndServe(config.BindAddr, http.RedirectHandler(config.TargetURL, config.RedirectCode))
	log.Println(err)
	log.Fatal("Try setting CAP_NET_BIND_SERVICE: setcap 'cap_net_bind_service=+ep' " + os.Args[0])
}
