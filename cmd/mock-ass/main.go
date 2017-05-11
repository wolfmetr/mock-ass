package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/wolfmetr/mock-ass/random_data"

	"github.com/fatih/color"
)

var (
	flagColor = flag.Bool("color", false, "enable color output")
	flagPort  = flag.Uint("port", 8000, "server start port")
)

func main() {
	flag.Parse()
	if *flagColor == false {
		color.NoColor = true
	}
	port := *flagPort

	collection, err := random_data.InitCollection()
	if err != nil {
		log.Fatalf(color.RedString("InitCollection error: %v", err))
	}
	log.Println(color.BlueString("Data collection successfully loaded"))
	server := http.Server{
		Addr: fmt.Sprintf(":%d", port),
		Handler: newAppHandler(
			collection,
			Route{
				path: "/session/",
				hand: hello,
			},
			Route{
				path: "/init",
				hand: get_session,
			},
		),
	}

	log.Println(color.BlueString("Start server port %d", port))
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf(color.RedString("serve error: %v", err))
	}
}
