package main

import (
	"github.com/cmdrkeene/placebunny.com/conf"
	"github.com/cmdrkeene/placebunny.com/server"
	"runtime"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	server.Start(conf.Port)
}
