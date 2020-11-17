package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	port := flag.String("p", ":8080", "Port to listen")
	root := flag.String("r", ".", "Directory root")
	flag.Parse()
	http.Handle("/", http.FileServer(http.Dir(*root)))
	log.Fatal(http.ListenAndServe(*port, nil))
}
