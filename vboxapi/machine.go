package vboxapi

import "github.com/blacktop/go-vboxapi/vboxweb-v4"

type Machine struct {
	virtualbox      *VirtualBox
	managedObjectId string
	ID              string
	Name            string
}

// func (m *Machine) GetNetworkAdapter(slot uint32) (*NetworkAdapter, error) {
// 	request := vboxweb4.IMachinegetNetworkAdapter{This: m.managedObjectId, Slot: slot}
//
// 	response, err := m.virtualbox.IMachinegetNetworkAdapter(&request)
// 	if err != nil {
// 		return nil, err // TODO: Wrap the error
// 	}
//
// 	return &NetworkAdapter{m.virtualbox, response.Returnval}, nil
// }

func (m *Machine) GetSettingsFilePath() (string, error) {
	request := vboxweb4.IMachinegetSettingsFilePath{This: m.managedObjectId}

	response, err := m.virtualbox.IMachinegetSettingsFilePath(&request)
	if err != nil {
		return "", err // TODO: Wrap the error
	}

	return response.Returnval, nil
}

func (m *Machine) SaveSettings() error {
	request := vboxweb4.IMachinesaveSettings{This: m.managedObjectId}

	_, err := m.virtualbox.IMachinesaveSettings(&request)
	if err != nil {
		defer m.DiscardSettings()
		return err // TODO: Wrap the error
	}

	return nil
}

func (m *Machine) DiscardSettings() error {
	request := vboxweb4.IMachinediscardSettings{This: m.managedObjectId}

	_, err := m.virtualbox.IMachinediscardSettings(&request)
	if err != nil {
		return err // TODO: Wrap the error
	}

	return nil
}

// func (m *Machine) Unlock(session *Session) error {
// 	if err := session.UnlockMachine(); err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (m *Machine) Lock(session *Session) error {
// 	request := vboxweb4.IMachineaddStorageController
// 	if err := session.LockMachine(m, 3); err != nil {
// 		return err
// 	}
// 	return nil
// }

func (m *Machine) GetID() (string, error) {
	request := vboxweb4.IMachinegetId{This: m.managedObjectId}

	response, err := m.virtualbox.IMachinegetId(&request)
	if err != nil {
		return "", err // TODO: Wrap the error
	}

	// TODO: See if we need to do anything with the response
	return response.Returnval, nil
}

func (m *Machine) GetName() (string, error) {
	request := vboxweb4.IMachinegetName{This: m.managedObjectId}

	response, err := m.virtualbox.IMachinegetName(&request)
	if err != nil {
		return "", err // TODO: Wrap the error
	}

	// TODO: See if we need to do anything with the response
	return response.Returnval, nil
}

func (m *Machine) Release() error {
	return m.virtualbox.Release(m.managedObjectId)
}

func (m *Machine) Refresh() error {
	if mr, err := m.virtualbox.FindMachine(m.ID); err != nil {
		return err
	} else {
		m.managedObjectId = mr.managedObjectId
	}
	return nil
}
