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

// TakeSnapshot takes a snapshot of the currently locked machine
// Saves the current execution state and all settings of the machine and
// creates differencing images for all normal (non-independent) media.
func (m *Machine) TakeSnapshot(name string, description string) (*Progress, error) {
	request := vboxweb4.IMachinetakeSnapshot{
		This:        m.managedObjectId,
		Name:        name,
		Description: description,
		Pause:       false,
	}

	response, err := m.virtualbox.IMachinetakeSnapshot(&request)
	if err != nil {
		return nil, err // TODO: Wrap the error
	}

	// TODO: See if we need to do anything with the response
	return &Progress{virtualbox: m.virtualbox, managedObjectId: response.Returnval}, nil
}

// FindSnapshot finds a snapshot matching a name or ID
// Returns a snapshot of this machine with the given UUID
func (m *Machine) FindSnapshot(nameOrID string) (*Snapshot, error) {
	request := vboxweb4.IMachinefindSnapshot{
		This:     m.managedObjectId,
		NameOrId: nameOrID,
	}

	response, err := m.virtualbox.IMachinefindSnapshot(&request)
	if err != nil {
		return nil, err // TODO: Wrap the error
	}

	// TODO: See if we need to do anything with the response
	return &Snapshot{virtualbox: m.virtualbox, managedObjectID: response.Returnval}, nil
}

// RestoreSnapshot starts resetting the machine's current state to the state
// contained in the given snapshot, asynchronously
// Returns a *Progress
func (m *Machine) RestoreSnapshot(snapshot *Snapshot) (*Progress, error) {
	request := vboxweb4.IMachinerestoreSnapshot{
		This:     m.managedObjectId,
		Snapshot: snapshot.managedObjectID,
	}

	response, err := m.virtualbox.IMachinerestoreSnapshot(&request)
	if err != nil {
		return nil, err // TODO: Wrap the error
	}

	// TODO: See if we need to do anything with the response
	return &Progress{virtualbox: m.virtualbox, managedObjectId: response.Returnval}, nil
}

// DeleteSnapshot starts deleting the specified snapshot asynchronously.
// Returns a *Progress
func (m *Machine) DeleteSnapshot(snapshot *Snapshot) (*Progress, error) {
	request := vboxweb4.IMachinedeleteSnapshot{
		This: m.managedObjectId,
		Id:   snapshot.managedObjectID,
	}

	response, err := m.virtualbox.IMachinedeleteSnapshot(&request)
	if err != nil {
		return nil, err // TODO: Wrap the error
	}

	// TODO: See if we need to do anything with the response
	return &Progress{virtualbox: m.virtualbox, managedObjectId: response.Returnval}, nil
}

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
