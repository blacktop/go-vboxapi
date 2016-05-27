package main

import (
	"fmt"
	"log"
	"os"

	"github.com/blacktop/go-vboxapi/vboxapi"
)

func assert(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	url := "http://127.0.0.1:18083"
	if len(os.Args) >= 2 {
		url = os.Args[1]
	}

	client := vboxapi.New("", "", url, false, "")
	if err := client.Logon(); err != nil {
		log.Fatalf("Unable to log on to vboxweb: %v\n", err)
	}
	machines, err := client.GetMachines()
	assert(err)
	for _, machine := range machines {
		name, err := machine.GetName()
		assert(err)
		fmt.Println(name)
	}
}
