package vboxapi

import "github.com/blacktop/go-vboxapi/vboxweb-v4"

type VirtualBox struct {
	*vboxweb4.VboxPortType
	managedObjectId string
	basicAuth       *vboxweb4.BasicAuth
	controllerName  string
}

func New(username, password, url string, tls bool, controllerName string) *VirtualBox {
	basicAuth := &vboxweb4.BasicAuth{
		Login:    username,
		Password: password,
	}
	return &VirtualBox{
		VboxPortType:   vboxweb4.NewVboxPortType(url, tls, basicAuth),
		basicAuth:      basicAuth,
		controllerName: controllerName,
	}
}

// GetExistingSession opens a existing session.
// func (vb *VirtualBox) GetExistingSession(machineID string) error {
// 	vbSession, err := vb.GetSession()
// 	if err != nil {
// 		return err // TODO: Wrap the error
// 	}
// 	request := vboxweb4.IVirtualBoxopenExistingSession{
// 		This:      vb.managedObjectId,
// 		Session:   vbSession.managedObjectId,
// 		MachineId: machineID,
// 	}

// 	_, err = vb.IVirtualBoxopenExistingSession(&request)
// 	if err != nil {
// 		return err // TODO: Wrap the error
// 	}

// 	// TODO: See if we need to do anything with the response
// 	return nil
// }

// GetRemoteSession opens a remote session.
// func (vb *VirtualBox) GetRemoteSession(machineID string) (*Progress, error) {
// 	vbSession, err := vb.GetSession()
// 	if err != nil {
// 		return nil, err // TODO: Wrap the error
// 	}
// 	request := vboxweb4.IVirtualBoxopenRemoteSession{
// 		This:      vb.managedObjectId,
// 		Session:   vbSession.managedObjectId,
// 		MachineId: machineID,
// 		// Type_       string `xml:"type,omitempty"`
// 		// Environment string `xml:"environment,omitempty"`
// 	}

// 	response, err := vb.IVirtualBoxopenRemoteSession(&request)
// 	if err != nil {
// 		return nil, err // TODO: Wrap the error
// 	}

// 	// TODO: See if we need to do anything with the response
// 	return &Progress{virtualbox: vb, managedObjectId: response.Returnval}, nil
// }

// func (vb *VirtualBox) CreateHardDisk(format, location string) (*Medium, error) {
// 	var am vboxweb4.AccessMode
// 	am = "ReadWrite"
// 	var dt vboxweb4.DeviceType
// 	dt = "HardDisk"
// 	request := vboxweb4.IVirtualBoxcreateMedium{
// 		This: vb.managedObjectId, Format: format, Location: location,
// 		AccessMode:      &am,
// 		ADeviceTypeType: &dt,
// 	}
//
// 	response, err := vb.IVirtualBoxcreateMedium(&request)
// 	if err != nil {
// 		return nil, err // TODO: Wrap the error
// 	}
//
// 	return &Medium{virtualbox: vb, managedObjectId: response.Returnval}, nil
// }

func (vb *VirtualBox) GetMachines() ([]*Machine, error) {
	request := vboxweb4.IVirtualBoxgetMachines{This: vb.managedObjectId}

	response, err := vb.IVirtualBoxgetMachines(&request)
	if err != nil {
		return nil, err // TODO: Wrap the error
	}

	machines := make([]*Machine, len(response.Returnval))
	for n, oid := range response.Returnval {
		machines[n] = &Machine{virtualbox: vb, managedObjectId: oid}
	}

	return machines, nil
}

// func (vb *VirtualBox) GetSystemProperties() (*SystemProperties, error) {
// 	request := vboxweb4.IVirtualBoxgetSystemProperties{This: vb.managedObjectId}
//
// 	response, err := vb.IVirtualBoxgetSystemProperties(&request)
// 	if err != nil {
// 		return nil, err // TODO: Wrap the error
// 	}
//
// 	return &SystemProperties{vb, response.Returnval}, nil
// }

func (vb *VirtualBox) Logon() error {
	request := vboxweb4.IWebsessionManagerlogon{
		Username: vb.basicAuth.Login,
		Password: vb.basicAuth.Password,
	}

	response, err := vb.IWebsessionManagerlogon(&request)
	if err != nil {
		return err // TODO: Wrap the error
	}

	vb.managedObjectId = response.Returnval

	return nil
}

func (vb *VirtualBox) LogOff() error {
	request := vboxweb4.IWebsessionManagerlogoff{RefIVirtualBox: vb.managedObjectId}

	_, err := vb.IWebsessionManagerlogoff(&request)
	if err != nil {
		return err // TODO: Wrap the error
	}

	return nil
}

// func (vb *VirtualBox) GetHardDisk(objectID string) (*HardDisks, error) {
// 	request := vboxweb4.IVirtualBoxgetHardDisks{This: vb.managedObjectId}
//
// 	response, err := vb.IVirtualBoxgetHardDisks(&request)
// 	if err != nil {
// 		return nil, err // TODO: Wrap the error
// 	}
//
// 	var hardDisks []*HardDisk
// 	for _, oid := range response.Returnval {
// 		if objectID == "" || objectID == oid {
// 			hardDisks = append(hardDisks, &HardDisk{vb, oid})
// 		}
// 	}
//
// 	return &HardDisks{disks: hardDisks}, nil
// }

// func (vb *VirtualBox) CreateMedium(format string, location string, size int64) (*Medium, error) {
//
// 	medium, err := vb.CreateHardDisk(format, location)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer medium.Release()
//
// 	progress, err := medium.CreateBaseStorage(size, nil)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	if err := progress.WaitForCompletion(-1); err != nil {
// 		return nil, err
// 	}
//
// 	if p, err := progress.GetPercent(); err != nil {
// 		return nil, err
// 	} else if p != 100 {
// 		return nil, errors.New("failed to create medium")
// 	}
//
// 	return medium.Get()
// }

// func (vb *VirtualBox) GetMedium(mediumID, mediumName string) ([]*Medium, error) {
// 	hardDisks, err := vb.GetHardDisk("")
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	return hardDisks.GetMedium(mediumID, mediumName)
// }

// func (vb *VirtualBox) RemoveMedium(mediumID string) error {
// 	if mediumID == "" {
// 		return errors.New("mediumID is empty")
// 	}
//
// 	mediums, err := vb.GetMedium(mediumID, "")
// 	if err != nil {
// 		return err
// 	}
//
// 	if len(mediums) == 0 {
// 		return errors.New("no mediums returned")
// 	}
//
// 	progress, err := mediums[0].DeleteStorage()
// 	if err != nil {
// 		return err
// 	}
//
// 	if err := progress.WaitForCompletion(-1); err != nil {
// 		return err
// 	}
//
// 	if p, err := progress.GetPercent(); err != nil {
// 		return err
// 	} else if p != 100 {
// 		return errors.New("failed to remove medium")
// 	}
//
// 	return nil
// }

func (vb *VirtualBox) GetSession() (*Session, error) {
	request := vboxweb4.IWebsessionManagergetSessionObject{RefIVirtualBox: vb.managedObjectId}
	response, err := vb.IWebsessionManagergetSessionObject(&request)
	if err != nil {
		return nil, err // TODO: Wrap the error
	}

	// TODO: See if we need to do anything with the response

	return &Session{managedObjectId: response.Returnval, virtualbox: vb}, nil
}

func (vb *VirtualBox) FindMachine(nameOrID string) (*Machine, error) {
	request := vboxweb4.IVirtualBoxfindMachine{
		This:     vb.managedObjectId,
		NameOrId: nameOrID,
	}
	response, err := vb.IVirtualBoxfindMachine(&request)
	if err != nil {
		return nil, err // TODO: Wrap the error
	}

	return &Machine{managedObjectId: response.Returnval, virtualbox: vb}, nil
}

func (vb *VirtualBox) Release(managedObjectId string) error {
	request := vboxweb4.IManagedObjectRefrelease{This: managedObjectId}

	_, err := vb.IManagedObjectRefrelease(&request)
	if err != nil {
		return err // TODO: Wrap the error
	}

	// TODO: See if we need to do anything with the response
	return nil
}

func (vb *VirtualBox) GetMOID() string {
	return vb.managedObjectId
}

// func (vb *VirtualBox) NewMedium(moid string) *Medium {
// 	return &Medium{virtualbox: vb, managedObjectId: moid}
// }
