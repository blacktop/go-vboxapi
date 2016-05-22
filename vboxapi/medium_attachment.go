package vboxapi

import "github.com/blacktop/go-vboxapi/vboxweb"

type MediumAttachment struct {
	*vboxweb.IMediumAttachment
	virtualbox      *VirtualBox
	managedObjectId string
}

func (m *MediumAttachment) GetMedium() (*Medium, error) {
	return &Medium{virtualbox: m.virtualbox, managedObjectId: m.Medium}, nil
}
