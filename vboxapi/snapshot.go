package vboxapi

import "github.com/blacktop/go-vboxapi/vboxweb-v4"

// Snapshot is a snapshot of a virtual machine
type Snapshot struct {
	virtualbox      *VirtualBox
	managedObjectID string
	ID              string
	Name            string
	Description     string
	TimeStamp       string
	Online          bool
	machine         *Machine
	parent          *Snapshot
	children        []Snapshot
}

// GetID returns UUID of the snapshot
func (s *Snapshot) GetID() (string, error) {
	request := vboxweb4.ISnapshotgetId{This: s.managedObjectID}

	response, err := s.virtualbox.ISnapshotgetId(&request)
	if err != nil {
		return "", err // TODO: Wrap the error
	}

	// TODO: See if we need to do anything with the response
	return response.Returnval, nil
}

// GetName returns short name of the snapshot
func (s *Snapshot) GetName() (string, error) {
	request := vboxweb4.ISnapshotgetName{This: s.managedObjectID}

	response, err := s.virtualbox.ISnapshotgetName(&request)
	if err != nil {
		return "", err // TODO: Wrap the error
	}

	// TODO: See if we need to do anything with the response
	return response.Returnval, nil
}

// SetName sets short name of the snapshot
func (s *Snapshot) SetName(name string) error {
	request := vboxweb4.ISnapshotsetName{This: s.managedObjectID, Name: name}

	_, err := s.virtualbox.ISnapshotsetName(&request)
	if err != nil {
		return err
	}
	return nil
}

// GetDescription returns description of the snapshot
func (s *Snapshot) GetDescription() (string, error) {
	request := vboxweb4.ISnapshotgetDescription{This: s.managedObjectID}

	response, err := s.virtualbox.ISnapshotgetDescription(&request)
	if err != nil {
		return "", err // TODO: Wrap the error
	}

	// TODO: See if we need to do anything with the response
	return response.Returnval, nil
}

// SetDescription sets description of the snapshot
func (s *Snapshot) SetDescription(description string) error {
	request := vboxweb4.ISnapshotsetDescription{
		This:        s.managedObjectID,
		Description: description,
	}

	_, err := s.virtualbox.ISnapshotsetDescription(&request)
	if err != nil {
		return err
	}
	return nil
}

// TODO: Add the rest of the fields getters

// GetChildrenCount returns the number of direct children of this snapshot.
func (s *Snapshot) GetChildrenCount() (uint32, error) {
	request := vboxweb4.ISnapshotgetChildrenCount{This: s.managedObjectID}

	response, err := s.virtualbox.ISnapshotgetChildrenCount(&request)
	if err != nil {
		return 0, err // TODO: Wrap the error
	}

	// TODO: See if we need to do anything with the response
	return response.Returnval, nil
}
