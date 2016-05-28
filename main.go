package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/blacktop/go-vboxapi/vboxapi"
)

func assert(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// PowerOn power on machine
// func PowerOn(client *vboxapi.VirtualBox, machineID string) error {
// 	progress, err := client.GetRemoteSession(machineID)
// 	assert(err)
// 	progress.WaitForCompletion(-1)
// 	return nil
// }

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
		machineID, err := machine.GetID()
		assert(err)
		fmt.Println(name, ": ", machineID)
		if strings.EqualFold(name, "default") {
			// assert(PowerOn(client, machineID))
			vbSession, err := client.GetSession()
			assert(err)

			// assert(vbSession.LockMachine(machine, "Write"))

			// console, err := vbSession.GetConsole()
			// assert(err)

			snapshot, err := machine.FindSnapshot("test")
			assert(err)

			progress, err := machine.RestoreSnapshot(snapshot)
			assert(err)
			assert(progress.WaitForCompletion(-1))

			progress, err = vbSession.LaunchVMProcess(machine)
			assert(err)

			assert(progress.WaitForCompletion(-1))

			sessionMachine, err := vbSession.GetMachine()
			assert(err)

			// progress, err = sessionMachine.TakeSnapshot("test3", "this is another test snapshot")
			// assert(err)
			// assert(progress.WaitForCompletion(-1))

			name, err := sessionMachine.GetName()
			assert(err)

			console, err := vbSession.GetConsole()
			assert(err)

			time.Sleep(5 * time.Second)

			progress, err = console.PowerDown()
			assert(err)
			assert(progress.WaitForCompletion(-1))

			fmt.Println("Session Machine: ", name)
			assert(vbSession.UnlockMachine())
		}
	}

	// machine, err := client.FindMachine("6e94d53e-5f78-4366-9b48-a5725ac6dbfb")
	// assert(err)
	// machineID, err := machine.GetID()
	// assert(err)

	// log off
	assert(client.LogOff())
}
