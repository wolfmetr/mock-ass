package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

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

	random_data.InitWithDefaults()
	server := http.Server{
		Addr: fmt.Sprintf(":%d", port),
		Handler: newAppHandler(
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
		log.Println(color.RedString("serve error: %v", err))
		os.Exit(1)
	}
}
