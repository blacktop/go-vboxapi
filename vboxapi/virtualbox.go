package vboxapi

import "github.com/blacktop/go-vboxapi/vboxweb"

type VirtualBox struct {
	*vboxweb.VboxPortType
	managedObjectId string
	basicAuth       *vboxweb.BasicAuth
	controllerName  string
}

func New(username, password, url string, tls bool, controllerName string) *VirtualBox {
	basicAuth := &vboxweb.BasicAuth{
		Login:    username,
		Password: password,
	}
	return &VirtualBox{
		VboxPortType:   vboxweb.NewVboxPortType(url, tls, basicAuth),
		basicAuth:      basicAuth,
		controllerName: controllerName,
	}
}

func (vb *VirtualBox) Logon() error {
	request := vboxweb.IWebsessionManagerlogon{
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

func (vb *VirtualBox) Release(managedObjectId string) error {
	request := vboxweb.IManagedObjectRefrelease{This: managedObjectId}

	_, err := vb.IManagedObjectRefrelease(&request)
	if err != nil {
		return err // TODO: Wrap the error
	}

	// TODO: See if we need to do anything with the response
	return nil
}
