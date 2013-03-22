package main

import (
	"flag"
	"log"
	"log/syslog"
	"net/http"
	"os"
	"gallows.inf.ed.ac.uk/hug/web"
)

var htdocs string
var logfile string

func init() {
	flag.StringVar(&htdocs, "htdocs", "./htdocs", "Location of static content")
	flag.StringVar(&logfile, "log", "", "Location for log file, or the string \"syslog\"")
}

func main() {
	flag.Parse()
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile)
	log.SetOutput(os.Stderr)
	if logfile == "syslog" {
		logger, err := syslog.New(syslog.LOG_INFO, "geosrv")
		if err != nil {
			log.Fatal(err)
		}
		log.SetOutput(logger)
		log.SetFlags(log.Lshortfile)
	} else if logfile != "" {
		fp, err := os.OpenFile(logfile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
			os.Exit(-1)
		}
		log.SetOutput(fp)
		defer fp.Close()
	}
	log.Printf("%s starting (pid: %d)", os.Args[0], os.Getpid())
	http.HandleFunc("/api/proj/", web.JsonProj)
	http.Handle("/", http.FileServer(http.Dir(htdocs)))
	log.Fatal(http.ListenAndServe(":8080", web.LogHandler{http.DefaultServeMux}))
}
