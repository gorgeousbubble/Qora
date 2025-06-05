package main

import (
	"fmt"
	"net/http"
	"qora/app"
	"runtime"
)

func init() {
	// start multi-cpu
	core := runtime.NumCPU()
	runtime.GOMAXPROCS(core)
	// start debug pprof
	go func() {
		_ = http.ListenAndServe(":11080", nil)
	}()
}

func main() {
	fmt.Println("The Qora Project")
	qora := app.New()
	qora.Init()
	qora.Start()
}
