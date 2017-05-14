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

var dataPath string

func init() {
	dataPath = os.Getenv("MOCK_ASS_DATA_DIR")
}

func main() {
	flag.Parse()
	if *flagColor == false {
		color.NoColor = true
	}

	collection, err := random_data.InitCollection(dataPath)
	if err != nil {
		log.Fatalf(color.RedString("InitCollection error: %v", err))
	}
	log.Println(color.BlueString("Data collection successfully loaded"))
	server := http.Server{
		Addr: fmt.Sprintf(":%d", *flagPort),
		Handler: newAppHandler(
			collection,
			Route{
				path: "/session/",
				hand: generateResp,
			},
			Route{
				path: "/init",
				hand: get_session,
			},
		),
	}

	log.Println(color.BlueString("Start server port %d", *flagPort))
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf(color.RedString("serve error: %v", err))
	}
}
