package main

import (
	"adams549659584/go-proxy-百晓通AI/api"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	http.HandleFunc("/sysconf", api.SysConf)

	http.HandleFunc("/sydney/", api.Sydney)

	http.HandleFunc("/web/", api.WebStatic)

	http.HandleFunc("/", api.Index)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	addr := ":" + port

	log.Println("Starting 百晓通AI Proxy At " + addr)

	srv := &http.Server{
		Addr:         addr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
