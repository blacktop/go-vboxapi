package main

import (
	"log"
	"os"

	"github.com/blacktop/go-vboxapi/vboxapi"
)

func main() {
	url := "http://127.0.0.1:18083"
	if len(os.Args) >= 2 {
		url = os.Args[1]
	}

	client := vboxapi.New("", "", url)
	if err := client.Logon(); err != nil {
		log.Fatalf("Unable to log on to vboxwebsrv: %v\n", err)
	}
}
