package main

import (
	"flag"
	"go/application"
)

var (
	addr string
)

func init() {
	flag.StringVar(&addr, "addr", ":9000", "tcp host:port to connect")
	flag.Parse()
}

func main() {
	application.Serve(addr)
}
