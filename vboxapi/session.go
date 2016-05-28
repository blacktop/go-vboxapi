package vboxapi

import "github.com/blacktop/go-vboxapi/vboxweb-v4"

type Session struct {
	virtualbox      *VirtualBox
	managedObjectId string
	console         *Console
	machine         *Machine
}

func (s *Session) UnlockMachine() error {
	request := vboxweb4.ISessionunlockMachine{This: s.managedObjectId}
	_, err := s.virtualbox.ISessionunlockMachine(&request)
	if err != nil {
		return err // TODO: Wrap the error
	}

	// TODO: See if we need to do anything with the response
	return nil
}

func (s *Session) LockMachine(m *Machine, l vboxweb4.LockType) error {
	request := vboxweb4.IMachinelockMachine{
		This:     m.managedObjectId,
		Session:  s.managedObjectId,
		LockType: &l,
	}
	_, err := s.virtualbox.IMachinelockMachine(&request)
	if err != nil {
		return err // TODO: Wrap the error
	}

	// TODO: See if we need to do anything with the response
	return nil
}

func (s *Session) LaunchVMProcess(m *Machine) (*Progress, error) {
	request := vboxweb4.IMachinelaunchVMProcess{
		This:    m.managedObjectId,
		Session: s.managedObjectId,
		Type_:   "headless",
		// Environment string `xml:"environment,omitempty"`
	}

	response, err := m.virtualbox.IMachinelaunchVMProcess(&request)
	if err != nil {
		return nil, err // TODO: Wrap the error
	}

	// TODO: See if we need to do anything with the response
	return &Progress{virtualbox: s.virtualbox, managedObjectId: response.Returnval}, nil
}

func (s *Session) GetMachine() (*Machine, error) {
	request := vboxweb4.ISessiongetMachine{This: s.managedObjectId}
	response, err := s.virtualbox.ISessiongetMachine(&request)
	if err != nil {
		return nil, err // TODO: Wrap the error
	}

	// TODO: See if we need to do anything with the response
	return &Machine{managedObjectId: response.Returnval, virtualbox: s.virtualbox}, nil
}

// GetConsole obtains the controls for the VM associated with this session.
// The call fails unless the VM associated with this session has started.
// It returns a new Console instance and any error encountered.
func (s *Session) GetConsole() (*Console, error) {
	request := vboxweb4.ISessiongetConsole{This: s.managedObjectId}
	response, err := s.virtualbox.ISessiongetConsole(&request)
	if err != nil {
		return nil, err // TODO: Wrap the error
	}

	// TODO: See if we need to do anything with the response
	return &Console{managedObjectID: response.Returnval, virtualbox: s.virtualbox}, nil
}

func (s *Session) Release() error {
	return s.virtualbox.Release(s.managedObjectId)
}
