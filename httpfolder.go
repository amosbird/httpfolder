package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
)

var ip string
var port string
var usr string
var passwd string

func main() {
	flag.StringVar(&ip, "i", "127.0.0.1", "ip to listen on")
	flag.StringVar(&port, "p", "8080", "port to listen on")
	flag.Parse()
	cwd := flag.Arg(0)
	if cwd == "" {
		cwd, _ = os.Getwd()
	}
	fmt.Printf("Serving: %s at http://%s:%s\n", cwd, ip, port)
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		FileServer(Dir(cwd)).ServeHTTP(res, req)
	})
	http.ListenAndServe((":" + port), nil)
}
