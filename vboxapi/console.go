package vboxapi

import "github.com/blacktop/go-vboxapi/vboxweb-v4"

// Console is a VirtualBox console object
type Console struct {
	virtualbox      *VirtualBox
	managedObjectID string
	machine         *Machine
}

// PowerDown starts forcibly powering off the controlled VM.
// It returns a Progress and any error encountered.
func (c *Console) PowerDown() (*Progress, error) {
	request := vboxweb4.IConsolepowerDown{This: c.managedObjectID}

	response, err := c.virtualbox.IConsolepowerDown(&request)
	if err != nil {
		return nil, err // TODO: Wrap the error
	}

	return &Progress{virtualbox: c.virtualbox, managedObjectId: response.Returnval}, nil
}

// PowerUp starts powering on the controlled VM.
// It returns a Progress and any error encountered.
func (c *Console) PowerUp() (*Progress, error) {
	request := vboxweb4.IConsolepowerUp{This: c.managedObjectID}

	response, err := c.virtualbox.IConsolepowerUp(&request)
	if err != nil {
		return nil, err // TODO: Wrap the error
	}

	return &Progress{virtualbox: c.virtualbox, managedObjectId: response.Returnval}, nil
}

// TakeSnapshot takes a snapshot of the currently locked machine
// Saves the current execution state and all settings of the machine and
// creates differencing images for all normal (non-independent) media.
func (c *Console) TakeSnapshot(name string, description string) (string, error) {
	request := vboxweb4.IConsoletakeSnapshot{
		This:        c.managedObjectID,
		Name:        name,
		Description: description,
	}

	response, err := c.virtualbox.IConsoletakeSnapshot(&request)
	if err != nil {
		return "", err // TODO: Wrap the error
	}

	// TODO: See if we need to do anything with the response
	return response.Returnval, nil
}

// func (console *Console) PowerDown() (Progress, error) {
// 	var progress Progress
// 	result := C.GoVboxConsolePowerDown(console.cconsole, &progress.cprogress)
// 	if C.GoVboxFAILED(result) != 0 || progress.cprogress == nil {
// 		return progress, errors.New(
// 			fmt.Sprintf("Failed to power down VM via IConsole: %x", result))
// 	}
// 	return progress, nil
// }
