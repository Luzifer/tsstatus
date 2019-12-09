package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/Luzifer/rconfig/v2"
	log "github.com/sirupsen/logrus"
)

var (
	cfg = struct {
		Listen         string `flag:"listen" default:":3000" description:"Port/IP to listen on"`
		LogLevel       string `flag:"log-level" default:"info" description:"Log level (debug, info, warn, error, fatal)"`
		QueryPass      string `flag:"query-pass" default:"" description:"Password for the server-query user" validate:"nonzero"`
		QueryUser      string `flag:"query-user" default:"" description:"Username for the server-query login" validate:"nonzero"`
		ServerAddress  string `flag:"server-address" default:"" description:"IP/Port of the Teamspeak server" validate:"nonzero"`
		ServerID       int    `flag:"server-id" default:"1" description:"ID of the virtual server to use for channel list"`
		VersionAndExit bool   `flag:"version" default:"false" description:"Prints current version and exits"`
	}{}

	version = "dev"
)

func init() {
	rconfig.AutoEnv(true)
	if err := rconfig.ParseAndValidate(&cfg); err != nil {
		log.Fatalf("Unable to parse commandline options: %s", err)
	}

	if cfg.VersionAndExit {
		fmt.Printf("tsstatus %s\n", version)
		os.Exit(0)
	}

	if l, err := log.ParseLevel(cfg.LogLevel); err != nil {
		log.WithError(err).Fatal("Unable to parse log level")
	} else {
		log.SetLevel(l)
	}
}

func main() {
	http.HandleFunc("/status", handleStatusRequest)
	http.ListenAndServe(cfg.Listen, nil)
}

func handleStatusRequest(w http.ResponseWriter, r *http.Request) {
	resp := struct {
		Info  *serverStats `json:"info,omitempty"`
		Error string       `json:"error,omitempty"`
	}{}

	info, err := getServerStats()
	resp.Info = info
	if err != nil {
		resp.Error = err.Error()
	}

	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	json.NewEncoder(w).Encode(resp)
}
