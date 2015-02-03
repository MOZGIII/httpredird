package main

import (
	"encoding/json"
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

// LoadConfig loads config from json file
func LoadConfig(filename string) (*ConfigStruct, error) {
	reader, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer reader.Close()

	decoder := json.NewDecoder(reader)

	config := new(ConfigStruct)
	err = decoder.Decode(config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

// MustLoadConfig calls LoadConfig and makes it a must-success
func MustLoadConfig(filename string) (config *ConfigStruct) {
	config, err := LoadConfig(filename)
	if err != nil {
		panic(err)
	}
	return
}

var configFile string

func init() {
	const (
		configFileDefault = "config.json"
		configFileUsage   = "config file to use"
	)
	flag.StringVar(&configFile, "config", configFileDefault, configFileUsage)
	flag.StringVar(&configFile, "c", configFileDefault, configFileUsage+" (shorthand)")
}

func main() {
	flag.Parse()
	config := MustLoadConfig(configFile)
	err := http.ListenAndServe(config.BindAddr, http.RedirectHandler(config.TargetURL, config.RedirectCode))
	log.Println(err)
	log.Fatal("Try setting CAP_NET_BIND_SERVICE: setcap 'cap_net_bind_service=+ep' " + os.Args[0])
}
