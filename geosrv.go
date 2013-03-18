package main

import (
	"flag"
	"jproj"
	"log"
	"net/http"
	"os"
	"webx"
)

var htdocs string
var logfile string

func init() {
	flag.StringVar(&htdocs, "htdocs", "./htdocs", "Location of static content")
	flag.StringVar(&logfile, "log", "", "Location for log file")
}

func main() {
	flag.Parse()
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile)
	log.SetOutput(os.Stderr)
	if logfile != "" {
		fp, err := os.OpenFile(logfile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
			os.Exit(-1)
		}
		log.SetOutput(fp)
		defer fp.Close()
	}
	log.Printf("%s starting (pid: %d)", os.Args[0], os.Getpid())
	http.HandleFunc("/api/proj/", jproj.JsonProj)
	http.Handle("/", http.FileServer(http.Dir(htdocs)))
	log.Fatal(http.ListenAndServe(":8080", webx.LogHandler{http.DefaultServeMux}))
}
