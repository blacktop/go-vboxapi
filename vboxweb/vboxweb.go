package vboxweb

import (
	"bytes"
	"crypto/tls"
	"encoding/xml"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"time"
)

// against "unused imports"
var _ time.Time
var _ xml.Name

type SettingsVersion string

const (
	SettingsVersionNull SettingsVersion = "Null"

	SettingsVersionV10 SettingsVersion = "v10"

	SettingsVersionV11 SettingsVersion = "v11"

	SettingsVersionV12 SettingsVersion = "v12"

	SettingsVersionV13pre SettingsVersion = "v13pre"

	SettingsVersionV13 SettingsVersion = "v13"

	SettingsVersionV14 SettingsVersion = "v14"

	SettingsVersionV15 SettingsVersion = "v15"

	SettingsVersionV16 SettingsVersion = "v16"

	SettingsVersionV17 SettingsVersion = "v17"

	SettingsVersionV18 SettingsVersion = "v18"

	SettingsVersionV19 SettingsVersion = "v19"

	SettingsVersionV110 SettingsVersion = "v110"

	SettingsVersionFuture SettingsVersion = "Future"
)

type AccessMode string

const (
	AccessModeReadOnly AccessMode = "ReadOnly"

	AccessModeReadWrite AccessMode = "ReadWrite"
)

type MachineState string

const (
	MachineStateNull MachineState = "Null"

	MachineStatePoweredOff MachineState = "PoweredOff"

	MachineStateSaved MachineState = "Saved"

	MachineStateTeleported MachineState = "Teleported"

	MachineStateAborted MachineState = "Aborted"

	MachineStateRunning MachineState = "Running"

	MachineStatePaused MachineState = "Paused"

	MachineStateStuck MachineState = "Stuck"

	MachineStateTeleporting MachineState = "Teleporting"

	MachineStateLiveSnapshotting MachineState = "LiveSnapshotting"

	MachineStateStarting MachineState = "Starting"

	MachineStateStopping MachineState = "Stopping"

	MachineStateSaving MachineState = "Saving"

	MachineStateRestoring MachineState = "Restoring"

	MachineStateTeleportingPausedVM MachineState = "TeleportingPausedVM"

	MachineStateTeleportingIn MachineState = "TeleportingIn"

	MachineStateDeletingSnapshotOnline MachineState = "DeletingSnapshotOnline"

	MachineStateDeletingSnapshotPaused MachineState = "DeletingSnapshotPaused"

	MachineStateRestoringSnapshot MachineState = "RestoringSnapshot"

	MachineStateDeletingSnapshot MachineState = "DeletingSnapshot"

	MachineStateSettingUp MachineState = "SettingUp"

	MachineStateFirstOnline MachineState = "FirstOnline"

	MachineStateLastOnline MachineState = "LastOnline"

	MachineStateFirstTransient MachineState = "FirstTransient"

	MachineStateLastTransient MachineState = "LastTransient"
)

type SessionState string

const (
	SessionStateNull SessionState = "Null"

	SessionStateClosed SessionState = "Closed"

	SessionStateOpen SessionState = "Open"

	SessionStateSpawning SessionState = "Spawning"

	SessionStateClosing SessionState = "Closing"
)

type CPUPropertyType string

const (
	CPUPropertyTypeNull CPUPropertyType = "Null"

	CPUPropertyTypePAE CPUPropertyType = "PAE"

	CPUPropertyTypeSynthetic CPUPropertyType = "Synthetic"
)

type HWVirtExPropertyType string

const (
	HWVirtExPropertyTypeNull HWVirtExPropertyType = "Null"

	HWVirtExPropertyTypeEnabled HWVirtExPropertyType = "Enabled"

	HWVirtExPropertyTypeExclusive HWVirtExPropertyType = "Exclusive"

	HWVirtExPropertyTypeVPID HWVirtExPropertyType = "VPID"

	HWVirtExPropertyTypeNestedPaging HWVirtExPropertyType = "NestedPaging"

	HWVirtExPropertyTypeLargePages HWVirtExPropertyType = "LargePages"

	HWVirtExPropertyTypeForce HWVirtExPropertyType = "Force"
)

type SessionType string

const (
	SessionTypeNull SessionType = "Null"

	SessionTypeDirect SessionType = "Direct"

	SessionTypeRemote SessionType = "Remote"

	SessionTypeExisting SessionType = "Existing"
)

type DeviceType string

const (
	DeviceTypeNull DeviceType = "Null"

	DeviceTypeFloppy DeviceType = "Floppy"

	DeviceTypeDVD DeviceType = "DVD"

	DeviceTypeHardDisk DeviceType = "HardDisk"

	DeviceTypeNetwork DeviceType = "Network"

	DeviceTypeUSB DeviceType = "USB"

	DeviceTypeSharedFolder DeviceType = "SharedFolder"
)

type DeviceActivity string

const (
	DeviceActivityNull DeviceActivity = "Null"

	DeviceActivityIdle DeviceActivity = "Idle"

	DeviceActivityReading DeviceActivity = "Reading"

	DeviceActivityWriting DeviceActivity = "Writing"
)

type ClipboardMode string

const (
	ClipboardModeDisabled ClipboardMode = "Disabled"

	ClipboardModeHostToGuest ClipboardMode = "HostToGuest"

	ClipboardModeGuestToHost ClipboardMode = "GuestToHost"

	ClipboardModeBidirectional ClipboardMode = "Bidirectional"
)

type Scope string

const (
	ScopeGlobal Scope = "Global"

	ScopeMachine Scope = "Machine"

	ScopeSession Scope = "Session"
)

type BIOSBootMenuMode string

const (
	BIOSBootMenuModeDisabled BIOSBootMenuMode = "Disabled"

	BIOSBootMenuModeMenuOnly BIOSBootMenuMode = "MenuOnly"

	BIOSBootMenuModeMessageAndMenu BIOSBootMenuMode = "MessageAndMenu"
)

type ProcessorFeature string

const (
	ProcessorFeatureHWVirtEx ProcessorFeature = "HWVirtEx"

	ProcessorFeaturePAE ProcessorFeature = "PAE"

	ProcessorFeatureLongMode ProcessorFeature = "LongMode"

	ProcessorFeatureNestedPaging ProcessorFeature = "NestedPaging"
)

type FirmwareType string

const (
	FirmwareTypeBIOS FirmwareType = "BIOS"

	FirmwareTypeEFI FirmwareType = "EFI"

	FirmwareTypeEFI32 FirmwareType = "EFI32"

	FirmwareTypeEFI64 FirmwareType = "EFI64"

	FirmwareTypeEFIDUAL FirmwareType = "EFIDUAL"
)

type PointingHidType string

const (
	PointingHidTypeNone PointingHidType = "None"

	PointingHidTypePS2Mouse PointingHidType = "PS2Mouse"

	PointingHidTypeUSBMouse PointingHidType = "USBMouse"

	PointingHidTypeUSBTablet PointingHidType = "USBTablet"

	PointingHidTypeComboMouse PointingHidType = "ComboMouse"
)

type KeyboardHidType string

const (
	KeyboardHidTypeNone KeyboardHidType = "None"

	KeyboardHidTypePS2Keyboard KeyboardHidType = "PS2Keyboard"

	KeyboardHidTypeUSBKeyboard KeyboardHidType = "USBKeyboard"

	KeyboardHidTypeComboKeyboard KeyboardHidType = "ComboKeyboard"
)

type VFSType string

const (
	VFSTypeFile VFSType = "File"

	VFSTypeCloud VFSType = "Cloud"

	VFSTypeS3 VFSType = "S3"

	VFSTypeWebDav VFSType = "WebDav"
)

type VFSFileType string

const (
	VFSFileTypeUnknown VFSFileType = "Unknown"

	VFSFileTypeFifo VFSFileType = "Fifo"

	VFSFileTypeDevChar VFSFileType = "DevChar"

	VFSFileTypeDirectory VFSFileType = "Directory"

	VFSFileTypeDevBlock VFSFileType = "DevBlock"

	VFSFileTypeFile VFSFileType = "File"

	VFSFileTypeSymLink VFSFileType = "SymLink"

	VFSFileTypeSocket VFSFileType = "Socket"

	VFSFileTypeWhiteOut VFSFileType = "WhiteOut"
)

type VirtualSystemDescriptionType string

const (
	VirtualSystemDescriptionTypeIgnore VirtualSystemDescriptionType = "Ignore"

	VirtualSystemDescriptionTypeOS VirtualSystemDescriptionType = "OS"

	VirtualSystemDescriptionTypeName VirtualSystemDescriptionType = "Name"

	VirtualSystemDescriptionTypeProduct VirtualSystemDescriptionType = "Product"

	VirtualSystemDescriptionTypeVendor VirtualSystemDescriptionType = "Vendor"

	VirtualSystemDescriptionTypeVersion VirtualSystemDescriptionType = "Version"

	VirtualSystemDescriptionTypeProductUrl VirtualSystemDescriptionType = "ProductUrl"

	VirtualSystemDescriptionTypeVendorUrl VirtualSystemDescriptionType = "VendorUrl"

	VirtualSystemDescriptionTypeDescription VirtualSystemDescriptionType = "Description"

	VirtualSystemDescriptionTypeLicense VirtualSystemDescriptionType = "License"

	VirtualSystemDescriptionTypeMiscellaneous VirtualSystemDescriptionType = "Miscellaneous"

	VirtualSystemDescriptionTypeCPU VirtualSystemDescriptionType = "CPU"

	VirtualSystemDescriptionTypeMemory VirtualSystemDescriptionType = "Memory"

	VirtualSystemDescriptionTypeHardDiskControllerIDE VirtualSystemDescriptionType = "HardDiskControllerIDE"

	VirtualSystemDescriptionTypeHardDiskControllerSATA VirtualSystemDescriptionType = "HardDiskControllerSATA"

	VirtualSystemDescriptionTypeHardDiskControllerSCSI VirtualSystemDescriptionType = "HardDiskControllerSCSI"

	VirtualSystemDescriptionTypeHardDiskControllerSAS VirtualSystemDescriptionType = "HardDiskControllerSAS"

	VirtualSystemDescriptionTypeHardDiskImage VirtualSystemDescriptionType = "HardDiskImage"

	VirtualSystemDescriptionTypeFloppy VirtualSystemDescriptionType = "Floppy"

	VirtualSystemDescriptionTypeCDROM VirtualSystemDescriptionType = "CDROM"

	VirtualSystemDescriptionTypeNetworkAdapter VirtualSystemDescriptionType = "NetworkAdapter"

	VirtualSystemDescriptionTypeUSBController VirtualSystemDescriptionType = "USBController"

	VirtualSystemDescriptionTypeSoundCard VirtualSystemDescriptionType = "SoundCard"
)

type VirtualSystemDescriptionValueType string

const (
	VirtualSystemDescriptionValueTypeReference VirtualSystemDescriptionValueType = "Reference"

	VirtualSystemDescriptionValueTypeOriginal VirtualSystemDescriptionValueType = "Original"

	VirtualSystemDescriptionValueTypeAuto VirtualSystemDescriptionValueType = "Auto"

	VirtualSystemDescriptionValueTypeExtraConfig VirtualSystemDescriptionValueType = "ExtraConfig"
)

type HostNetworkInterfaceMediumType string

const (
	HostNetworkInterfaceMediumTypeUnknown HostNetworkInterfaceMediumType = "Unknown"

	HostNetworkInterfaceMediumTypeEthernet HostNetworkInterfaceMediumType = "Ethernet"

	HostNetworkInterfaceMediumTypePPP HostNetworkInterfaceMediumType = "PPP"

	HostNetworkInterfaceMediumTypeSLIP HostNetworkInterfaceMediumType = "SLIP"
)

type HostNetworkInterfaceStatus string

const (
	HostNetworkInterfaceStatusUnknown HostNetworkInterfaceStatus = "Unknown"

	HostNetworkInterfaceStatusUp HostNetworkInterfaceStatus = "Up"

	HostNetworkInterfaceStatusDown HostNetworkInterfaceStatus = "Down"
)

type HostNetworkInterfaceType string

const (
	HostNetworkInterfaceTypeBridged HostNetworkInterfaceType = "Bridged"

	HostNetworkInterfaceTypeHostOnly HostNetworkInterfaceType = "HostOnly"
)

type MediumState string

const (
	MediumStateNotCreated MediumState = "NotCreated"

	MediumStateCreated MediumState = "Created"

	MediumStateLockedRead MediumState = "LockedRead"

	MediumStateLockedWrite MediumState = "LockedWrite"

	MediumStateInaccessible MediumState = "Inaccessible"

	MediumStateCreating MediumState = "Creating"

	MediumStateDeleting MediumState = "Deleting"
)

type MediumType string

const (
	MediumTypeNormal MediumType = "Normal"

	MediumTypeImmutable MediumType = "Immutable"

	MediumTypeWritethrough MediumType = "Writethrough"

	MediumTypeShareable MediumType = "Shareable"
)

type MediumVariant string

const (
	MediumVariantStandard MediumVariant = "Standard"

	MediumVariantVmdkSplit2G MediumVariant = "VmdkSplit2G"

	MediumVariantVmdkStreamOptimized MediumVariant = "VmdkStreamOptimized"

	MediumVariantVmdkESX MediumVariant = "VmdkESX"

	MediumVariantFixed MediumVariant = "Fixed"

	MediumVariantDiff MediumVariant = "Diff"
)

type DataType string

const (
	DataTypeInt32 DataType = "Int32"

	DataTypeInt8 DataType = "Int8"

	DataTypeString DataType = "String"
)

type DataFlags string

const (
	DataFlagsNone DataFlags = "None"

	DataFlagsMandatory DataFlags = "Mandatory"

	DataFlagsExpert DataFlags = "Expert"

	DataFlagsArray DataFlags = "Array"

	DataFlagsFlagMask DataFlags = "FlagMask"
)

type MediumFormatCapabilities string

const (
	MediumFormatCapabilitiesUuid MediumFormatCapabilities = "Uuid"

	MediumFormatCapabilitiesCreateFixed MediumFormatCapabilities = "CreateFixed"

	MediumFormatCapabilitiesCreateDynamic MediumFormatCapabilities = "CreateDynamic"

	MediumFormatCapabilitiesCreateSplit2G MediumFormatCapabilities = "CreateSplit2G"

	MediumFormatCapabilitiesDifferencing MediumFormatCapabilities = "Differencing"

	MediumFormatCapabilitiesAsynchronous MediumFormatCapabilities = "Asynchronous"

	MediumFormatCapabilitiesFile MediumFormatCapabilities = "File"

	MediumFormatCapabilitiesProperties MediumFormatCapabilities = "Properties"

	MediumFormatCapabilitiesCapabilityMask MediumFormatCapabilities = "CapabilityMask"
)

type MouseButtonState string

const (
	MouseButtonStateLeftButton MouseButtonState = "LeftButton"

	MouseButtonStateRightButton MouseButtonState = "RightButton"

	MouseButtonStateMiddleButton MouseButtonState = "MiddleButton"

	MouseButtonStateWheelUp MouseButtonState = "WheelUp"

	MouseButtonStateWheelDown MouseButtonState = "WheelDown"

	MouseButtonStateXButton1 MouseButtonState = "XButton1"

	MouseButtonStateXButton2 MouseButtonState = "XButton2"

	MouseButtonStateMouseStateMask MouseButtonState = "MouseStateMask"
)

type FramebufferPixelFormat string

const (
	FramebufferPixelFormatOpaque FramebufferPixelFormat = "Opaque"

	FramebufferPixelFormatFOURCCRGB FramebufferPixelFormat = "FOURCCRGB"
)

type NetworkAttachmentType string

const (
	NetworkAttachmentTypeNull NetworkAttachmentType = "Null"

	NetworkAttachmentTypeNAT NetworkAttachmentType = "NAT"

	NetworkAttachmentTypeBridged NetworkAttachmentType = "Bridged"

	NetworkAttachmentTypeInternal NetworkAttachmentType = "Internal"

	NetworkAttachmentTypeHostOnly NetworkAttachmentType = "HostOnly"

	NetworkAttachmentTypeVDE NetworkAttachmentType = "VDE"
)

type NetworkAdapterType string

const (
	NetworkAdapterTypeNull NetworkAdapterType = "Null"

	NetworkAdapterTypeAm79C970A NetworkAdapterType = "Am79C970A"

	NetworkAdapterTypeAm79C973 NetworkAdapterType = "Am79C973"

	NetworkAdapterTypeI82540EM NetworkAdapterType = "I82540EM"

	NetworkAdapterTypeI82543GC NetworkAdapterType = "I82543GC"

	NetworkAdapterTypeI82545EM NetworkAdapterType = "I82545EM"

	NetworkAdapterTypeVirtio NetworkAdapterType = "Virtio"
)

type PortMode string

const (
	PortModeDisconnected PortMode = "Disconnected"

	PortModeHostPipe PortMode = "HostPipe"

	PortModeHostDevice PortMode = "HostDevice"

	PortModeRawFile PortMode = "RawFile"
)

type USBDeviceState string

const (
	USBDeviceStateNotSupported USBDeviceState = "NotSupported"

	USBDeviceStateUnavailable USBDeviceState = "Unavailable"

	USBDeviceStateBusy USBDeviceState = "Busy"

	USBDeviceStateAvailable USBDeviceState = "Available"

	USBDeviceStateHeld USBDeviceState = "Held"

	USBDeviceStateCaptured USBDeviceState = "Captured"
)

type USBDeviceFilterAction string

const (
	USBDeviceFilterActionNull USBDeviceFilterAction = "Null"

	USBDeviceFilterActionIgnore USBDeviceFilterAction = "Ignore"

	USBDeviceFilterActionHold USBDeviceFilterAction = "Hold"
)

type AudioDriverType string

const (
	AudioDriverTypeNull AudioDriverType = "Null"

	AudioDriverTypeWinMM AudioDriverType = "WinMM"

	AudioDriverTypeOSS AudioDriverType = "OSS"

	AudioDriverTypeALSA AudioDriverType = "ALSA"

	AudioDriverTypeDirectSound AudioDriverType = "DirectSound"

	AudioDriverTypeCoreAudio AudioDriverType = "CoreAudio"

	AudioDriverTypeMMPM AudioDriverType = "MMPM"

	AudioDriverTypePulse AudioDriverType = "Pulse"

	AudioDriverTypeSolAudio AudioDriverType = "SolAudio"
)

type AudioControllerType string

const (
	AudioControllerTypeAC97 AudioControllerType = "AC97"

	AudioControllerTypeSB16 AudioControllerType = "SB16"
)

type VRDPAuthType string

const (
	VRDPAuthTypeNull VRDPAuthType = "Null"

	VRDPAuthTypeExternal VRDPAuthType = "External"

	VRDPAuthTypeGuest VRDPAuthType = "Guest"
)

type StorageBus string

const (
	StorageBusNull StorageBus = "Null"

	StorageBusIDE StorageBus = "IDE"

	StorageBusSATA StorageBus = "SATA"

	StorageBusSCSI StorageBus = "SCSI"

	StorageBusFloppy StorageBus = "Floppy"

	StorageBusSAS StorageBus = "SAS"
)

type StorageControllerType string

const (
	StorageControllerTypeNull StorageControllerType = "Null"

	StorageControllerTypeLsiLogic StorageControllerType = "LsiLogic"

	StorageControllerTypeBusLogic StorageControllerType = "BusLogic"

	StorageControllerTypeIntelAhci StorageControllerType = "IntelAhci"

	StorageControllerTypePIIX3 StorageControllerType = "PIIX3"

	StorageControllerTypePIIX4 StorageControllerType = "PIIX4"

	StorageControllerTypeICH6 StorageControllerType = "ICH6"

	StorageControllerTypeI82078 StorageControllerType = "I82078"

	StorageControllerTypeLsiLogicSas StorageControllerType = "LsiLogicSas"
)

type NATAliasMode string

const (
	NATAliasModeAliasLog NATAliasMode = "AliasLog"

	NATAliasModeAliasProxyOnly NATAliasMode = "AliasProxyOnly"

	NATAliasModeAliasUseSamePorts NATAliasMode = "AliasUseSamePorts"
)

type NATProtocol string

const (
	NATProtocolUDP NATProtocol = "UDP"

	NATProtocolTCP NATProtocol = "TCP"
)

type IVirtualBoxErrorInfogetResultCode struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBoxErrorInfo_getResultCode"`

	This string `xml:"_this,omitempty"`
}

type IVirtualBoxErrorInfogetResultCodeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBoxErrorInfo_getResultCodeResponse"`

	Returnval int32 `xml:"returnval,omitempty"`
}

type IVirtualBoxErrorInfogetInterfaceID struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBoxErrorInfo_getInterfaceID"`

	This string `xml:"_this,omitempty"`
}

type IVirtualBoxErrorInfogetInterfaceIDResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBoxErrorInfo_getInterfaceIDResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IVirtualBoxErrorInfogetComponent struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBoxErrorInfo_getComponent"`

	This string `xml:"_this,omitempty"`
}

type IVirtualBoxErrorInfogetComponentResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBoxErrorInfo_getComponentResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IVirtualBoxErrorInfogetText struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBoxErrorInfo_getText"`

	This string `xml:"_this,omitempty"`
}

type IVirtualBoxErrorInfogetTextResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBoxErrorInfo_getTextResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IVirtualBoxErrorInfogetNext struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBoxErrorInfo_getNext"`

	This string `xml:"_this,omitempty"`
}

type IVirtualBoxErrorInfogetNextResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBoxErrorInfo_getNextResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IDHCPServergetEnabled struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IDHCPServer_getEnabled"`

	This string `xml:"_this,omitempty"`
}

type IDHCPServergetEnabledResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IDHCPServer_getEnabledResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IDHCPServersetEnabled struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IDHCPServer_setEnabled"`

	This    string `xml:"_this,omitempty"`
	Enabled bool   `xml:"enabled,omitempty"`
}

type IDHCPServersetEnabledResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IDHCPServer_setEnabledResponse"`
}

type IDHCPServergetIPAddress struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IDHCPServer_getIPAddress"`

	This string `xml:"_this,omitempty"`
}

type IDHCPServergetIPAddressResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IDHCPServer_getIPAddressResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IDHCPServergetNetworkMask struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IDHCPServer_getNetworkMask"`

	This string `xml:"_this,omitempty"`
}

type IDHCPServergetNetworkMaskResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IDHCPServer_getNetworkMaskResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IDHCPServergetNetworkName struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IDHCPServer_getNetworkName"`

	This string `xml:"_this,omitempty"`
}

type IDHCPServergetNetworkNameResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IDHCPServer_getNetworkNameResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IDHCPServergetLowerIP struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IDHCPServer_getLowerIP"`

	This string `xml:"_this,omitempty"`
}

type IDHCPServergetLowerIPResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IDHCPServer_getLowerIPResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IDHCPServergetUpperIP struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IDHCPServer_getUpperIP"`

	This string `xml:"_this,omitempty"`
}

type IDHCPServergetUpperIPResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IDHCPServer_getUpperIPResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IDHCPServersetConfiguration struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IDHCPServer_setConfiguration"`

	This          string `xml:"_this,omitempty"`
	IPAddress     string `xml:"IPAddress,omitempty"`
	NetworkMask   string `xml:"networkMask,omitempty"`
	FromIPAddress string `xml:"FromIPAddress,omitempty"`
	ToIPAddress   string `xml:"ToIPAddress,omitempty"`
}

type IDHCPServersetConfigurationResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IDHCPServer_setConfigurationResponse"`
}

type IDHCPServerstart struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IDHCPServer_start"`

	This        string `xml:"_this,omitempty"`
	NetworkName string `xml:"networkName,omitempty"`
	TrunkName   string `xml:"trunkName,omitempty"`
	TrunkType   string `xml:"trunkType,omitempty"`
}

type IDHCPServerstartResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IDHCPServer_startResponse"`
}

type IDHCPServerstop struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IDHCPServer_stop"`

	This string `xml:"_this,omitempty"`
}

type IDHCPServerstopResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IDHCPServer_stopResponse"`
}

type IVirtualBoxgetVersion struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_getVersion"`

	This string `xml:"_this,omitempty"`
}

type IVirtualBoxgetVersionResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_getVersionResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IVirtualBoxgetRevision struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_getRevision"`

	This string `xml:"_this,omitempty"`
}

type IVirtualBoxgetRevisionResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_getRevisionResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type IVirtualBoxgetPackageType struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_getPackageType"`

	This string `xml:"_this,omitempty"`
}

type IVirtualBoxgetPackageTypeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_getPackageTypeResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IVirtualBoxgetHomeFolder struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_getHomeFolder"`

	This string `xml:"_this,omitempty"`
}

type IVirtualBoxgetHomeFolderResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_getHomeFolderResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IVirtualBoxgetSettingsFilePath struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_getSettingsFilePath"`

	This string `xml:"_this,omitempty"`
}

type IVirtualBoxgetSettingsFilePathResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_getSettingsFilePathResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IVirtualBoxgetHost struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_getHost"`

	This string `xml:"_this,omitempty"`
}

type IVirtualBoxgetHostResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_getHostResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IVirtualBoxgetSystemProperties struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_getSystemProperties"`

	This string `xml:"_this,omitempty"`
}

type IVirtualBoxgetSystemPropertiesResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_getSystemPropertiesResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IVirtualBoxgetMachines struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_getMachines"`

	This string `xml:"_this,omitempty"`
}

type IVirtualBoxgetMachinesResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_getMachinesResponse"`

	Returnval []string `xml:"returnval,omitempty"`
}

type IVirtualBoxgetHardDisks struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_getHardDisks"`

	This string `xml:"_this,omitempty"`
}

type IVirtualBoxgetHardDisksResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_getHardDisksResponse"`

	Returnval []string `xml:"returnval,omitempty"`
}

type IVirtualBoxgetDVDImages struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_getDVDImages"`

	This string `xml:"_this,omitempty"`
}

type IVirtualBoxgetDVDImagesResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_getDVDImagesResponse"`

	Returnval []string `xml:"returnval,omitempty"`
}

type IVirtualBoxgetFloppyImages struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_getFloppyImages"`

	This string `xml:"_this,omitempty"`
}

type IVirtualBoxgetFloppyImagesResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_getFloppyImagesResponse"`

	Returnval []string `xml:"returnval,omitempty"`
}

type IVirtualBoxgetProgressOperations struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_getProgressOperations"`

	This string `xml:"_this,omitempty"`
}

type IVirtualBoxgetProgressOperationsResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_getProgressOperationsResponse"`

	Returnval []string `xml:"returnval,omitempty"`
}

type IVirtualBoxgetGuestOSTypes struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_getGuestOSTypes"`

	This string `xml:"_this,omitempty"`
}

type IVirtualBoxgetGuestOSTypesResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_getGuestOSTypesResponse"`

	Returnval []*IGuestOSType `xml:"returnval,omitempty"`
}

type IVirtualBoxgetSharedFolders struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_getSharedFolders"`

	This string `xml:"_this,omitempty"`
}

type IVirtualBoxgetSharedFoldersResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_getSharedFoldersResponse"`

	Returnval []*ISharedFolder `xml:"returnval,omitempty"`
}

type IVirtualBoxgetPerformanceCollector struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_getPerformanceCollector"`

	This string `xml:"_this,omitempty"`
}

type IVirtualBoxgetPerformanceCollectorResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_getPerformanceCollectorResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IVirtualBoxgetDHCPServers struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_getDHCPServers"`

	This string `xml:"_this,omitempty"`
}

type IVirtualBoxgetDHCPServersResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_getDHCPServersResponse"`

	Returnval []string `xml:"returnval,omitempty"`
}

type IVirtualBoxcreateMachine struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_createMachine"`

	This       string `xml:"_this,omitempty"`
	Name       string `xml:"name,omitempty"`
	OsTypeId   string `xml:"osTypeId,omitempty"`
	BaseFolder string `xml:"baseFolder,omitempty"`
	Id         string `xml:"id,omitempty"`
	Override   bool   `xml:"override,omitempty"`
}

type IVirtualBoxcreateMachineResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_createMachineResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IVirtualBoxcreateLegacyMachine struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_createLegacyMachine"`

	This         string `xml:"_this,omitempty"`
	Name         string `xml:"name,omitempty"`
	OsTypeId     string `xml:"osTypeId,omitempty"`
	SettingsFile string `xml:"settingsFile,omitempty"`
	Id           string `xml:"id,omitempty"`
}

type IVirtualBoxcreateLegacyMachineResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_createLegacyMachineResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IVirtualBoxopenMachine struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_openMachine"`

	This         string `xml:"_this,omitempty"`
	SettingsFile string `xml:"settingsFile,omitempty"`
}

type IVirtualBoxopenMachineResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_openMachineResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IVirtualBoxregisterMachine struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_registerMachine"`

	This    string `xml:"_this,omitempty"`
	Machine string `xml:"machine,omitempty"`
}

type IVirtualBoxregisterMachineResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_registerMachineResponse"`
}

type IVirtualBoxgetMachine struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_getMachine"`

	This string `xml:"_this,omitempty"`
	Id   string `xml:"id,omitempty"`
}

type IVirtualBoxgetMachineResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_getMachineResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IVirtualBoxfindMachine struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_findMachine"`

	This string `xml:"_this,omitempty"`
	Name string `xml:"name,omitempty"`
}

type IVirtualBoxfindMachineResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_findMachineResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IVirtualBoxunregisterMachine struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_unregisterMachine"`

	This string `xml:"_this,omitempty"`
	Id   string `xml:"id,omitempty"`
}

type IVirtualBoxunregisterMachineResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_unregisterMachineResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IVirtualBoxcreateAppliance struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_createAppliance"`

	This string `xml:"_this,omitempty"`
}

type IVirtualBoxcreateApplianceResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_createApplianceResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IVirtualBoxcreateHardDisk struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_createHardDisk"`

	This     string `xml:"_this,omitempty"`
	Format   string `xml:"format,omitempty"`
	Location string `xml:"location,omitempty"`
}

type IVirtualBoxcreateHardDiskResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_createHardDiskResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IVirtualBoxopenHardDisk struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_openHardDisk"`

	This        string      `xml:"_this,omitempty"`
	Location    string      `xml:"location,omitempty"`
	AccessMode  *AccessMode `xml:"accessMode,omitempty"`
	SetImageId  bool        `xml:"setImageId,omitempty"`
	ImageId     string      `xml:"imageId,omitempty"`
	SetParentId bool        `xml:"setParentId,omitempty"`
	ParentId    string      `xml:"parentId,omitempty"`
}

type IVirtualBoxopenHardDiskResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_openHardDiskResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IVirtualBoxgetHardDisk struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_getHardDisk"`

	This string `xml:"_this,omitempty"`
	Id   string `xml:"id,omitempty"`
}

type IVirtualBoxgetHardDiskResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_getHardDiskResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IVirtualBoxfindHardDisk struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_findHardDisk"`

	This     string `xml:"_this,omitempty"`
	Location string `xml:"location,omitempty"`
}

type IVirtualBoxfindHardDiskResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_findHardDiskResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IVirtualBoxopenDVDImage struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_openDVDImage"`

	This     string `xml:"_this,omitempty"`
	Location string `xml:"location,omitempty"`
	Id       string `xml:"id,omitempty"`
}

type IVirtualBoxopenDVDImageResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_openDVDImageResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IVirtualBoxgetDVDImage struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_getDVDImage"`

	This string `xml:"_this,omitempty"`
	Id   string `xml:"id,omitempty"`
}

type IVirtualBoxgetDVDImageResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_getDVDImageResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IVirtualBoxfindDVDImage struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_findDVDImage"`

	This     string `xml:"_this,omitempty"`
	Location string `xml:"location,omitempty"`
}

type IVirtualBoxfindDVDImageResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_findDVDImageResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IVirtualBoxopenFloppyImage struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_openFloppyImage"`

	This     string `xml:"_this,omitempty"`
	Location string `xml:"location,omitempty"`
	Id       string `xml:"id,omitempty"`
}

type IVirtualBoxopenFloppyImageResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_openFloppyImageResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IVirtualBoxgetFloppyImage struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_getFloppyImage"`

	This string `xml:"_this,omitempty"`
	Id   string `xml:"id,omitempty"`
}

type IVirtualBoxgetFloppyImageResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_getFloppyImageResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IVirtualBoxfindFloppyImage struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_findFloppyImage"`

	This     string `xml:"_this,omitempty"`
	Location string `xml:"location,omitempty"`
}

type IVirtualBoxfindFloppyImageResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_findFloppyImageResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IVirtualBoxgetGuestOSType struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_getGuestOSType"`

	This string `xml:"_this,omitempty"`
	Id   string `xml:"id,omitempty"`
}

type IVirtualBoxgetGuestOSTypeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_getGuestOSTypeResponse"`

	Returnval *IGuestOSType `xml:"returnval,omitempty"`
}

type IVirtualBoxcreateSharedFolder struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_createSharedFolder"`

	This     string `xml:"_this,omitempty"`
	Name     string `xml:"name,omitempty"`
	HostPath string `xml:"hostPath,omitempty"`
	Writable bool   `xml:"writable,omitempty"`
}

type IVirtualBoxcreateSharedFolderResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_createSharedFolderResponse"`
}

type IVirtualBoxremoveSharedFolder struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_removeSharedFolder"`

	This string `xml:"_this,omitempty"`
	Name string `xml:"name,omitempty"`
}

type IVirtualBoxremoveSharedFolderResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_removeSharedFolderResponse"`
}

type IVirtualBoxgetExtraDataKeys struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_getExtraDataKeys"`

	This string `xml:"_this,omitempty"`
}

type IVirtualBoxgetExtraDataKeysResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_getExtraDataKeysResponse"`

	Returnval []string `xml:"returnval,omitempty"`
}

type IVirtualBoxgetExtraData struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_getExtraData"`

	This string `xml:"_this,omitempty"`
	Key  string `xml:"key,omitempty"`
}

type IVirtualBoxgetExtraDataResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_getExtraDataResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IVirtualBoxsetExtraData struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_setExtraData"`

	This  string `xml:"_this,omitempty"`
	Key   string `xml:"key,omitempty"`
	Value string `xml:"value,omitempty"`
}

type IVirtualBoxsetExtraDataResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_setExtraDataResponse"`
}

type IVirtualBoxopenSession struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_openSession"`

	This      string `xml:"_this,omitempty"`
	Session   string `xml:"session,omitempty"`
	MachineId string `xml:"machineId,omitempty"`
}

type IVirtualBoxopenSessionResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_openSessionResponse"`
}

type IVirtualBoxopenRemoteSession struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_openRemoteSession"`

	This        string `xml:"_this,omitempty"`
	Session     string `xml:"session,omitempty"`
	MachineId   string `xml:"machineId,omitempty"`
	Type_       string `xml:"type,omitempty"`
	Environment string `xml:"environment,omitempty"`
}

type IVirtualBoxopenRemoteSessionResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_openRemoteSessionResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IVirtualBoxopenExistingSession struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_openExistingSession"`

	This      string `xml:"_this,omitempty"`
	Session   string `xml:"session,omitempty"`
	MachineId string `xml:"machineId,omitempty"`
}

type IVirtualBoxopenExistingSessionResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_openExistingSessionResponse"`
}

type IVirtualBoxwaitForPropertyChange struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_waitForPropertyChange"`

	This    string `xml:"_this,omitempty"`
	What    string `xml:"what,omitempty"`
	Timeout uint32 `xml:"timeout,omitempty"`
}

type IVirtualBoxwaitForPropertyChangeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_waitForPropertyChangeResponse"`

	Changed string `xml:"changed,omitempty"`
	Values  string `xml:"values,omitempty"`
}

type IVirtualBoxcreateDHCPServer struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_createDHCPServer"`

	This string `xml:"_this,omitempty"`
	Name string `xml:"name,omitempty"`
}

type IVirtualBoxcreateDHCPServerResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_createDHCPServerResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IVirtualBoxfindDHCPServerByNetworkName struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_findDHCPServerByNetworkName"`

	This string `xml:"_this,omitempty"`
	Name string `xml:"name,omitempty"`
}

type IVirtualBoxfindDHCPServerByNetworkNameResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_findDHCPServerByNetworkNameResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IVirtualBoxremoveDHCPServer struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_removeDHCPServer"`

	This   string `xml:"_this,omitempty"`
	Server string `xml:"server,omitempty"`
}

type IVirtualBoxremoveDHCPServerResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_removeDHCPServerResponse"`
}

type IVirtualBoxcheckFirmwarePresent struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_checkFirmwarePresent"`

	This         string        `xml:"_this,omitempty"`
	FirmwareType *FirmwareType `xml:"firmwareType,omitempty"`
	Version      string        `xml:"version,omitempty"`
}

type IVirtualBoxcheckFirmwarePresentResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_checkFirmwarePresentResponse"`

	Url       string `xml:"url,omitempty"`
	File      string `xml:"file,omitempty"`
	Returnval bool   `xml:"returnval,omitempty"`
}

type IVFSExplorergetPath struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVFSExplorer_getPath"`

	This string `xml:"_this,omitempty"`
}

type IVFSExplorergetPathResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVFSExplorer_getPathResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IVFSExplorergetType struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVFSExplorer_getType"`

	This string `xml:"_this,omitempty"`
}

type IVFSExplorergetTypeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVFSExplorer_getTypeResponse"`

	Returnval *VFSType `xml:"returnval,omitempty"`
}

type IVFSExplorerupdate struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVFSExplorer_update"`

	This string `xml:"_this,omitempty"`
}

type IVFSExplorerupdateResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVFSExplorer_updateResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IVFSExplorercd struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVFSExplorer_cd"`

	This string `xml:"_this,omitempty"`
	ADir string `xml:"aDir,omitempty"`
}

type IVFSExplorercdResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVFSExplorer_cdResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IVFSExplorercdUp struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVFSExplorer_cdUp"`

	This string `xml:"_this,omitempty"`
}

type IVFSExplorercdUpResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVFSExplorer_cdUpResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IVFSExplorerentryList struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVFSExplorer_entryList"`

	This string `xml:"_this,omitempty"`
}

type IVFSExplorerentryListResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVFSExplorer_entryListResponse"`

	ANames []string `xml:"aNames,omitempty"`
	ATypes []uint32 `xml:"aTypes,omitempty"`
}

type IVFSExplorerexists struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVFSExplorer_exists"`

	This   string   `xml:"_this,omitempty"`
	ANames []string `xml:"aNames,omitempty"`
}

type IVFSExplorerexistsResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVFSExplorer_existsResponse"`

	Returnval []string `xml:"returnval,omitempty"`
}

type IVFSExplorerremove struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVFSExplorer_remove"`

	This   string   `xml:"_this,omitempty"`
	ANames []string `xml:"aNames,omitempty"`
}

type IVFSExplorerremoveResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVFSExplorer_removeResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IAppliancegetPath struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IAppliance_getPath"`

	This string `xml:"_this,omitempty"`
}

type IAppliancegetPathResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IAppliance_getPathResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IAppliancegetDisks struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IAppliance_getDisks"`

	This string `xml:"_this,omitempty"`
}

type IAppliancegetDisksResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IAppliance_getDisksResponse"`

	Returnval []string `xml:"returnval,omitempty"`
}

type IAppliancegetVirtualSystemDescriptions struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IAppliance_getVirtualSystemDescriptions"`

	This string `xml:"_this,omitempty"`
}

type IAppliancegetVirtualSystemDescriptionsResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IAppliance_getVirtualSystemDescriptionsResponse"`

	Returnval []string `xml:"returnval,omitempty"`
}

type IApplianceread struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IAppliance_read"`

	This string `xml:"_this,omitempty"`
	File string `xml:"file,omitempty"`
}

type IAppliancereadResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IAppliance_readResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IApplianceinterpret struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IAppliance_interpret"`

	This string `xml:"_this,omitempty"`
}

type IApplianceinterpretResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IAppliance_interpretResponse"`
}

type IApplianceimportMachines struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IAppliance_importMachines"`

	This string `xml:"_this,omitempty"`
}

type IApplianceimportMachinesResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IAppliance_importMachinesResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IAppliancecreateVFSExplorer struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IAppliance_createVFSExplorer"`

	This string `xml:"_this,omitempty"`
	AUri string `xml:"aUri,omitempty"`
}

type IAppliancecreateVFSExplorerResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IAppliance_createVFSExplorerResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IAppliancewrite struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IAppliance_write"`

	This   string `xml:"_this,omitempty"`
	Format string `xml:"format,omitempty"`
	Path   string `xml:"path,omitempty"`
}

type IAppliancewriteResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IAppliance_writeResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IAppliancegetWarnings struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IAppliance_getWarnings"`

	This string `xml:"_this,omitempty"`
}

type IAppliancegetWarningsResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IAppliance_getWarningsResponse"`

	Returnval []string `xml:"returnval,omitempty"`
}

type IVirtualSystemDescriptiongetCount struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualSystemDescription_getCount"`

	This string `xml:"_this,omitempty"`
}

type IVirtualSystemDescriptiongetCountResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualSystemDescription_getCountResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type IVirtualSystemDescriptiongetDescription struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualSystemDescription_getDescription"`

	This string `xml:"_this,omitempty"`
}

type IVirtualSystemDescriptiongetDescriptionResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualSystemDescription_getDescriptionResponse"`

	ATypes             []*VirtualSystemDescriptionType `xml:"aTypes,omitempty"`
	ARefs              []string                        `xml:"aRefs,omitempty"`
	AOvfValues         []string                        `xml:"aOvfValues,omitempty"`
	AVBoxValues        []string                        `xml:"aVBoxValues,omitempty"`
	AExtraConfigValues []string                        `xml:"aExtraConfigValues,omitempty"`
}

type IVirtualSystemDescriptiongetDescriptionByType struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualSystemDescription_getDescriptionByType"`

	This  string                        `xml:"_this,omitempty"`
	AType *VirtualSystemDescriptionType `xml:"aType,omitempty"`
}

type IVirtualSystemDescriptiongetDescriptionByTypeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualSystemDescription_getDescriptionByTypeResponse"`

	ATypes             []*VirtualSystemDescriptionType `xml:"aTypes,omitempty"`
	ARefs              []string                        `xml:"aRefs,omitempty"`
	AOvfValues         []string                        `xml:"aOvfValues,omitempty"`
	AVBoxValues        []string                        `xml:"aVBoxValues,omitempty"`
	AExtraConfigValues []string                        `xml:"aExtraConfigValues,omitempty"`
}

type IVirtualSystemDescriptiongetValuesByType struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualSystemDescription_getValuesByType"`

	This   string                             `xml:"_this,omitempty"`
	AType  *VirtualSystemDescriptionType      `xml:"aType,omitempty"`
	AWhich *VirtualSystemDescriptionValueType `xml:"aWhich,omitempty"`
}

type IVirtualSystemDescriptiongetValuesByTypeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualSystemDescription_getValuesByTypeResponse"`

	Returnval []string `xml:"returnval,omitempty"`
}

type IVirtualSystemDescriptionsetFinalValues struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualSystemDescription_setFinalValues"`

	This               string   `xml:"_this,omitempty"`
	AEnabled           []bool   `xml:"aEnabled,omitempty"`
	AVBoxValues        []string `xml:"aVBoxValues,omitempty"`
	AExtraConfigValues []string `xml:"aExtraConfigValues,omitempty"`
}

type IVirtualSystemDescriptionsetFinalValuesResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualSystemDescription_setFinalValuesResponse"`
}

type IVirtualSystemDescriptionaddDescription struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualSystemDescription_addDescription"`

	This              string                        `xml:"_this,omitempty"`
	AType             *VirtualSystemDescriptionType `xml:"aType,omitempty"`
	AVBoxValue        string                        `xml:"aVBoxValue,omitempty"`
	AExtraConfigValue string                        `xml:"aExtraConfigValue,omitempty"`
}

type IVirtualSystemDescriptionaddDescriptionResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualSystemDescription_addDescriptionResponse"`
}

type IBIOSSettingsgetLogoFadeIn struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IBIOSSettings_getLogoFadeIn"`

	This string `xml:"_this,omitempty"`
}

type IBIOSSettingsgetLogoFadeInResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IBIOSSettings_getLogoFadeInResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IBIOSSettingssetLogoFadeIn struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IBIOSSettings_setLogoFadeIn"`

	This       string `xml:"_this,omitempty"`
	LogoFadeIn bool   `xml:"logoFadeIn,omitempty"`
}

type IBIOSSettingssetLogoFadeInResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IBIOSSettings_setLogoFadeInResponse"`
}

type IBIOSSettingsgetLogoFadeOut struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IBIOSSettings_getLogoFadeOut"`

	This string `xml:"_this,omitempty"`
}

type IBIOSSettingsgetLogoFadeOutResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IBIOSSettings_getLogoFadeOutResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IBIOSSettingssetLogoFadeOut struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IBIOSSettings_setLogoFadeOut"`

	This        string `xml:"_this,omitempty"`
	LogoFadeOut bool   `xml:"logoFadeOut,omitempty"`
}

type IBIOSSettingssetLogoFadeOutResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IBIOSSettings_setLogoFadeOutResponse"`
}

type IBIOSSettingsgetLogoDisplayTime struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IBIOSSettings_getLogoDisplayTime"`

	This string `xml:"_this,omitempty"`
}

type IBIOSSettingsgetLogoDisplayTimeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IBIOSSettings_getLogoDisplayTimeResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type IBIOSSettingssetLogoDisplayTime struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IBIOSSettings_setLogoDisplayTime"`

	This            string `xml:"_this,omitempty"`
	LogoDisplayTime uint32 `xml:"logoDisplayTime,omitempty"`
}

type IBIOSSettingssetLogoDisplayTimeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IBIOSSettings_setLogoDisplayTimeResponse"`
}

type IBIOSSettingsgetLogoImagePath struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IBIOSSettings_getLogoImagePath"`

	This string `xml:"_this,omitempty"`
}

type IBIOSSettingsgetLogoImagePathResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IBIOSSettings_getLogoImagePathResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IBIOSSettingssetLogoImagePath struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IBIOSSettings_setLogoImagePath"`

	This          string `xml:"_this,omitempty"`
	LogoImagePath string `xml:"logoImagePath,omitempty"`
}

type IBIOSSettingssetLogoImagePathResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IBIOSSettings_setLogoImagePathResponse"`
}

type IBIOSSettingsgetBootMenuMode struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IBIOSSettings_getBootMenuMode"`

	This string `xml:"_this,omitempty"`
}

type IBIOSSettingsgetBootMenuModeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IBIOSSettings_getBootMenuModeResponse"`

	Returnval *BIOSBootMenuMode `xml:"returnval,omitempty"`
}

type IBIOSSettingssetBootMenuMode struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IBIOSSettings_setBootMenuMode"`

	This         string            `xml:"_this,omitempty"`
	BootMenuMode *BIOSBootMenuMode `xml:"bootMenuMode,omitempty"`
}

type IBIOSSettingssetBootMenuModeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IBIOSSettings_setBootMenuModeResponse"`
}

type IBIOSSettingsgetACPIEnabled struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IBIOSSettings_getACPIEnabled"`

	This string `xml:"_this,omitempty"`
}

type IBIOSSettingsgetACPIEnabledResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IBIOSSettings_getACPIEnabledResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IBIOSSettingssetACPIEnabled struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IBIOSSettings_setACPIEnabled"`

	This        string `xml:"_this,omitempty"`
	ACPIEnabled bool   `xml:"ACPIEnabled,omitempty"`
}

type IBIOSSettingssetACPIEnabledResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IBIOSSettings_setACPIEnabledResponse"`
}

type IBIOSSettingsgetIOAPICEnabled struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IBIOSSettings_getIOAPICEnabled"`

	This string `xml:"_this,omitempty"`
}

type IBIOSSettingsgetIOAPICEnabledResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IBIOSSettings_getIOAPICEnabledResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IBIOSSettingssetIOAPICEnabled struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IBIOSSettings_setIOAPICEnabled"`

	This          string `xml:"_this,omitempty"`
	IOAPICEnabled bool   `xml:"IOAPICEnabled,omitempty"`
}

type IBIOSSettingssetIOAPICEnabledResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IBIOSSettings_setIOAPICEnabledResponse"`
}

type IBIOSSettingsgetTimeOffset struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IBIOSSettings_getTimeOffset"`

	This string `xml:"_this,omitempty"`
}

type IBIOSSettingsgetTimeOffsetResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IBIOSSettings_getTimeOffsetResponse"`

	Returnval int64 `xml:"returnval,omitempty"`
}

type IBIOSSettingssetTimeOffset struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IBIOSSettings_setTimeOffset"`

	This       string `xml:"_this,omitempty"`
	TimeOffset int64  `xml:"timeOffset,omitempty"`
}

type IBIOSSettingssetTimeOffsetResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IBIOSSettings_setTimeOffsetResponse"`
}

type IBIOSSettingsgetPXEDebugEnabled struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IBIOSSettings_getPXEDebugEnabled"`

	This string `xml:"_this,omitempty"`
}

type IBIOSSettingsgetPXEDebugEnabledResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IBIOSSettings_getPXEDebugEnabledResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IBIOSSettingssetPXEDebugEnabled struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IBIOSSettings_setPXEDebugEnabled"`

	This            string `xml:"_this,omitempty"`
	PXEDebugEnabled bool   `xml:"PXEDebugEnabled,omitempty"`
}

type IBIOSSettingssetPXEDebugEnabledResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IBIOSSettings_setPXEDebugEnabledResponse"`
}

type IMachinegetParent struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getParent"`

	This string `xml:"_this,omitempty"`
}

type IMachinegetParentResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getParentResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMachinegetAccessible struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getAccessible"`

	This string `xml:"_this,omitempty"`
}

type IMachinegetAccessibleResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getAccessibleResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IMachinegetAccessError struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getAccessError"`

	This string `xml:"_this,omitempty"`
}

type IMachinegetAccessErrorResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getAccessErrorResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMachinegetName struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getName"`

	This string `xml:"_this,omitempty"`
}

type IMachinegetNameResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getNameResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMachinesetName struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setName"`

	This string `xml:"_this,omitempty"`
	Name string `xml:"name,omitempty"`
}

type IMachinesetNameResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setNameResponse"`
}

type IMachinegetDescription struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getDescription"`

	This string `xml:"_this,omitempty"`
}

type IMachinegetDescriptionResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getDescriptionResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMachinesetDescription struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setDescription"`

	This        string `xml:"_this,omitempty"`
	Description string `xml:"description,omitempty"`
}

type IMachinesetDescriptionResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setDescriptionResponse"`
}

type IMachinegetId struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getId"`

	This string `xml:"_this,omitempty"`
}

type IMachinegetIdResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getIdResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMachinegetOSTypeId struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getOSTypeId"`

	This string `xml:"_this,omitempty"`
}

type IMachinegetOSTypeIdResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getOSTypeIdResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMachinesetOSTypeId struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setOSTypeId"`

	This     string `xml:"_this,omitempty"`
	OSTypeId string `xml:"OSTypeId,omitempty"`
}

type IMachinesetOSTypeIdResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setOSTypeIdResponse"`
}

type IMachinegetHardwareVersion struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getHardwareVersion"`

	This string `xml:"_this,omitempty"`
}

type IMachinegetHardwareVersionResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getHardwareVersionResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMachinesetHardwareVersion struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setHardwareVersion"`

	This            string `xml:"_this,omitempty"`
	HardwareVersion string `xml:"HardwareVersion,omitempty"`
}

type IMachinesetHardwareVersionResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setHardwareVersionResponse"`
}

type IMachinegetHardwareUUID struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getHardwareUUID"`

	This string `xml:"_this,omitempty"`
}

type IMachinegetHardwareUUIDResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getHardwareUUIDResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMachinesetHardwareUUID struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setHardwareUUID"`

	This         string `xml:"_this,omitempty"`
	HardwareUUID string `xml:"hardwareUUID,omitempty"`
}

type IMachinesetHardwareUUIDResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setHardwareUUIDResponse"`
}

type IMachinegetCPUCount struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getCPUCount"`

	This string `xml:"_this,omitempty"`
}

type IMachinegetCPUCountResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getCPUCountResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type IMachinesetCPUCount struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setCPUCount"`

	This     string `xml:"_this,omitempty"`
	CPUCount uint32 `xml:"CPUCount,omitempty"`
}

type IMachinesetCPUCountResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setCPUCountResponse"`
}

type IMachinegetCPUHotPlugEnabled struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getCPUHotPlugEnabled"`

	This string `xml:"_this,omitempty"`
}

type IMachinegetCPUHotPlugEnabledResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getCPUHotPlugEnabledResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IMachinesetCPUHotPlugEnabled struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setCPUHotPlugEnabled"`

	This              string `xml:"_this,omitempty"`
	CPUHotPlugEnabled bool   `xml:"CPUHotPlugEnabled,omitempty"`
}

type IMachinesetCPUHotPlugEnabledResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setCPUHotPlugEnabledResponse"`
}

type IMachinegetMemorySize struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getMemorySize"`

	This string `xml:"_this,omitempty"`
}

type IMachinegetMemorySizeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getMemorySizeResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type IMachinesetMemorySize struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setMemorySize"`

	This       string `xml:"_this,omitempty"`
	MemorySize uint32 `xml:"memorySize,omitempty"`
}

type IMachinesetMemorySizeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setMemorySizeResponse"`
}

type IMachinegetMemoryBalloonSize struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getMemoryBalloonSize"`

	This string `xml:"_this,omitempty"`
}

type IMachinegetMemoryBalloonSizeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getMemoryBalloonSizeResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type IMachinesetMemoryBalloonSize struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setMemoryBalloonSize"`

	This              string `xml:"_this,omitempty"`
	MemoryBalloonSize uint32 `xml:"memoryBalloonSize,omitempty"`
}

type IMachinesetMemoryBalloonSizeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setMemoryBalloonSizeResponse"`
}

type IMachinegetPageFusionEnabled struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getPageFusionEnabled"`

	This string `xml:"_this,omitempty"`
}

type IMachinegetPageFusionEnabledResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getPageFusionEnabledResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IMachinesetPageFusionEnabled struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setPageFusionEnabled"`

	This              string `xml:"_this,omitempty"`
	PageFusionEnabled bool   `xml:"PageFusionEnabled,omitempty"`
}

type IMachinesetPageFusionEnabledResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setPageFusionEnabledResponse"`
}

type IMachinegetVRAMSize struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getVRAMSize"`

	This string `xml:"_this,omitempty"`
}

type IMachinegetVRAMSizeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getVRAMSizeResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type IMachinesetVRAMSize struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setVRAMSize"`

	This     string `xml:"_this,omitempty"`
	VRAMSize uint32 `xml:"VRAMSize,omitempty"`
}

type IMachinesetVRAMSizeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setVRAMSizeResponse"`
}

type IMachinegetAccelerate3DEnabled struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getAccelerate3DEnabled"`

	This string `xml:"_this,omitempty"`
}

type IMachinegetAccelerate3DEnabledResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getAccelerate3DEnabledResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IMachinesetAccelerate3DEnabled struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setAccelerate3DEnabled"`

	This                string `xml:"_this,omitempty"`
	Accelerate3DEnabled bool   `xml:"accelerate3DEnabled,omitempty"`
}

type IMachinesetAccelerate3DEnabledResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setAccelerate3DEnabledResponse"`
}

type IMachinegetAccelerate2DVideoEnabled struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getAccelerate2DVideoEnabled"`

	This string `xml:"_this,omitempty"`
}

type IMachinegetAccelerate2DVideoEnabledResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getAccelerate2DVideoEnabledResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IMachinesetAccelerate2DVideoEnabled struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setAccelerate2DVideoEnabled"`

	This                     string `xml:"_this,omitempty"`
	Accelerate2DVideoEnabled bool   `xml:"accelerate2DVideoEnabled,omitempty"`
}

type IMachinesetAccelerate2DVideoEnabledResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setAccelerate2DVideoEnabledResponse"`
}

type IMachinegetMonitorCount struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getMonitorCount"`

	This string `xml:"_this,omitempty"`
}

type IMachinegetMonitorCountResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getMonitorCountResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type IMachinesetMonitorCount struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setMonitorCount"`

	This         string `xml:"_this,omitempty"`
	MonitorCount uint32 `xml:"monitorCount,omitempty"`
}

type IMachinesetMonitorCountResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setMonitorCountResponse"`
}

type IMachinegetBIOSSettings struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getBIOSSettings"`

	This string `xml:"_this,omitempty"`
}

type IMachinegetBIOSSettingsResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getBIOSSettingsResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMachinegetFirmwareType struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getFirmwareType"`

	This string `xml:"_this,omitempty"`
}

type IMachinegetFirmwareTypeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getFirmwareTypeResponse"`

	Returnval *FirmwareType `xml:"returnval,omitempty"`
}

type IMachinesetFirmwareType struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setFirmwareType"`

	This         string        `xml:"_this,omitempty"`
	FirmwareType *FirmwareType `xml:"firmwareType,omitempty"`
}

type IMachinesetFirmwareTypeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setFirmwareTypeResponse"`
}

type IMachinegetPointingHidType struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getPointingHidType"`

	This string `xml:"_this,omitempty"`
}

type IMachinegetPointingHidTypeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getPointingHidTypeResponse"`

	Returnval *PointingHidType `xml:"returnval,omitempty"`
}

type IMachinesetPointingHidType struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setPointingHidType"`

	This            string           `xml:"_this,omitempty"`
	PointingHidType *PointingHidType `xml:"pointingHidType,omitempty"`
}

type IMachinesetPointingHidTypeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setPointingHidTypeResponse"`
}

type IMachinegetKeyboardHidType struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getKeyboardHidType"`

	This string `xml:"_this,omitempty"`
}

type IMachinegetKeyboardHidTypeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getKeyboardHidTypeResponse"`

	Returnval *KeyboardHidType `xml:"returnval,omitempty"`
}

type IMachinesetKeyboardHidType struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setKeyboardHidType"`

	This            string           `xml:"_this,omitempty"`
	KeyboardHidType *KeyboardHidType `xml:"keyboardHidType,omitempty"`
}

type IMachinesetKeyboardHidTypeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setKeyboardHidTypeResponse"`
}

type IMachinegetHpetEnabled struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getHpetEnabled"`

	This string `xml:"_this,omitempty"`
}

type IMachinegetHpetEnabledResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getHpetEnabledResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IMachinesetHpetEnabled struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setHpetEnabled"`

	This        string `xml:"_this,omitempty"`
	HpetEnabled bool   `xml:"hpetEnabled,omitempty"`
}

type IMachinesetHpetEnabledResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setHpetEnabledResponse"`
}

type IMachinegetSnapshotFolder struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getSnapshotFolder"`

	This string `xml:"_this,omitempty"`
}

type IMachinegetSnapshotFolderResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getSnapshotFolderResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMachinesetSnapshotFolder struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setSnapshotFolder"`

	This           string `xml:"_this,omitempty"`
	SnapshotFolder string `xml:"snapshotFolder,omitempty"`
}

type IMachinesetSnapshotFolderResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setSnapshotFolderResponse"`
}

type IMachinegetVRDPServer struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getVRDPServer"`

	This string `xml:"_this,omitempty"`
}

type IMachinegetVRDPServerResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getVRDPServerResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMachinegetMediumAttachments struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getMediumAttachments"`

	This string `xml:"_this,omitempty"`
}

type IMachinegetMediumAttachmentsResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getMediumAttachmentsResponse"`

	Returnval []*IMediumAttachment `xml:"returnval,omitempty"`
}

type IMachinegetUSBController struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getUSBController"`

	This string `xml:"_this,omitempty"`
}

type IMachinegetUSBControllerResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getUSBControllerResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMachinegetAudioAdapter struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getAudioAdapter"`

	This string `xml:"_this,omitempty"`
}

type IMachinegetAudioAdapterResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getAudioAdapterResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMachinegetStorageControllers struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getStorageControllers"`

	This string `xml:"_this,omitempty"`
}

type IMachinegetStorageControllersResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getStorageControllersResponse"`

	Returnval []string `xml:"returnval,omitempty"`
}

type IMachinegetSettingsFilePath struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getSettingsFilePath"`

	This string `xml:"_this,omitempty"`
}

type IMachinegetSettingsFilePathResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getSettingsFilePathResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMachinegetSettingsModified struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getSettingsModified"`

	This string `xml:"_this,omitempty"`
}

type IMachinegetSettingsModifiedResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getSettingsModifiedResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IMachinegetSessionState struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getSessionState"`

	This string `xml:"_this,omitempty"`
}

type IMachinegetSessionStateResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getSessionStateResponse"`

	Returnval *SessionState `xml:"returnval,omitempty"`
}

type IMachinegetSessionType struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getSessionType"`

	This string `xml:"_this,omitempty"`
}

type IMachinegetSessionTypeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getSessionTypeResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMachinegetSessionPid struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getSessionPid"`

	This string `xml:"_this,omitempty"`
}

type IMachinegetSessionPidResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getSessionPidResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type IMachinegetState struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getState"`

	This string `xml:"_this,omitempty"`
}

type IMachinegetStateResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getStateResponse"`

	Returnval *MachineState `xml:"returnval,omitempty"`
}

type IMachinegetLastStateChange struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getLastStateChange"`

	This string `xml:"_this,omitempty"`
}

type IMachinegetLastStateChangeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getLastStateChangeResponse"`

	Returnval int64 `xml:"returnval,omitempty"`
}

type IMachinegetStateFilePath struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getStateFilePath"`

	This string `xml:"_this,omitempty"`
}

type IMachinegetStateFilePathResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getStateFilePathResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMachinegetLogFolder struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getLogFolder"`

	This string `xml:"_this,omitempty"`
}

type IMachinegetLogFolderResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getLogFolderResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMachinegetCurrentSnapshot struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getCurrentSnapshot"`

	This string `xml:"_this,omitempty"`
}

type IMachinegetCurrentSnapshotResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getCurrentSnapshotResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMachinegetSnapshotCount struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getSnapshotCount"`

	This string `xml:"_this,omitempty"`
}

type IMachinegetSnapshotCountResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getSnapshotCountResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type IMachinegetCurrentStateModified struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getCurrentStateModified"`

	This string `xml:"_this,omitempty"`
}

type IMachinegetCurrentStateModifiedResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getCurrentStateModifiedResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IMachinegetSharedFolders struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getSharedFolders"`

	This string `xml:"_this,omitempty"`
}

type IMachinegetSharedFoldersResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getSharedFoldersResponse"`

	Returnval []*ISharedFolder `xml:"returnval,omitempty"`
}

type IMachinegetClipboardMode struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getClipboardMode"`

	This string `xml:"_this,omitempty"`
}

type IMachinegetClipboardModeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getClipboardModeResponse"`

	Returnval *ClipboardMode `xml:"returnval,omitempty"`
}

type IMachinesetClipboardMode struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setClipboardMode"`

	This          string         `xml:"_this,omitempty"`
	ClipboardMode *ClipboardMode `xml:"clipboardMode,omitempty"`
}

type IMachinesetClipboardModeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setClipboardModeResponse"`
}

type IMachinegetGuestPropertyNotificationPatterns struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getGuestPropertyNotificationPatterns"`

	This string `xml:"_this,omitempty"`
}

type IMachinegetGuestPropertyNotificationPatternsResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getGuestPropertyNotificationPatternsResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMachinesetGuestPropertyNotificationPatterns struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setGuestPropertyNotificationPatterns"`

	This                              string `xml:"_this,omitempty"`
	GuestPropertyNotificationPatterns string `xml:"guestPropertyNotificationPatterns,omitempty"`
}

type IMachinesetGuestPropertyNotificationPatternsResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setGuestPropertyNotificationPatternsResponse"`
}

type IMachinegetTeleporterEnabled struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getTeleporterEnabled"`

	This string `xml:"_this,omitempty"`
}

type IMachinegetTeleporterEnabledResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getTeleporterEnabledResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IMachinesetTeleporterEnabled struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setTeleporterEnabled"`

	This              string `xml:"_this,omitempty"`
	TeleporterEnabled bool   `xml:"teleporterEnabled,omitempty"`
}

type IMachinesetTeleporterEnabledResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setTeleporterEnabledResponse"`
}

type IMachinegetTeleporterPort struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getTeleporterPort"`

	This string `xml:"_this,omitempty"`
}

type IMachinegetTeleporterPortResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getTeleporterPortResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type IMachinesetTeleporterPort struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setTeleporterPort"`

	This           string `xml:"_this,omitempty"`
	TeleporterPort uint32 `xml:"teleporterPort,omitempty"`
}

type IMachinesetTeleporterPortResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setTeleporterPortResponse"`
}

type IMachinegetTeleporterAddress struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getTeleporterAddress"`

	This string `xml:"_this,omitempty"`
}

type IMachinegetTeleporterAddressResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getTeleporterAddressResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMachinesetTeleporterAddress struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setTeleporterAddress"`

	This              string `xml:"_this,omitempty"`
	TeleporterAddress string `xml:"teleporterAddress,omitempty"`
}

type IMachinesetTeleporterAddressResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setTeleporterAddressResponse"`
}

type IMachinegetTeleporterPassword struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getTeleporterPassword"`

	This string `xml:"_this,omitempty"`
}

type IMachinegetTeleporterPasswordResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getTeleporterPasswordResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMachinesetTeleporterPassword struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setTeleporterPassword"`

	This               string `xml:"_this,omitempty"`
	TeleporterPassword string `xml:"teleporterPassword,omitempty"`
}

type IMachinesetTeleporterPasswordResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setTeleporterPasswordResponse"`
}

type IMachinegetRTCUseUTC struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getRTCUseUTC"`

	This string `xml:"_this,omitempty"`
}

type IMachinegetRTCUseUTCResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getRTCUseUTCResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IMachinesetRTCUseUTC struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setRTCUseUTC"`

	This      string `xml:"_this,omitempty"`
	RTCUseUTC bool   `xml:"RTCUseUTC,omitempty"`
}

type IMachinesetRTCUseUTCResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setRTCUseUTCResponse"`
}

type IMachinegetIoCacheEnabled struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getIoCacheEnabled"`

	This string `xml:"_this,omitempty"`
}

type IMachinegetIoCacheEnabledResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getIoCacheEnabledResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IMachinesetIoCacheEnabled struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setIoCacheEnabled"`

	This           string `xml:"_this,omitempty"`
	IoCacheEnabled bool   `xml:"ioCacheEnabled,omitempty"`
}

type IMachinesetIoCacheEnabledResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setIoCacheEnabledResponse"`
}

type IMachinegetIoCacheSize struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getIoCacheSize"`

	This string `xml:"_this,omitempty"`
}

type IMachinegetIoCacheSizeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getIoCacheSizeResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type IMachinesetIoCacheSize struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setIoCacheSize"`

	This        string `xml:"_this,omitempty"`
	IoCacheSize uint32 `xml:"ioCacheSize,omitempty"`
}

type IMachinesetIoCacheSizeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setIoCacheSizeResponse"`
}

type IMachinegetIoBandwidthMax struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getIoBandwidthMax"`

	This string `xml:"_this,omitempty"`
}

type IMachinegetIoBandwidthMaxResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getIoBandwidthMaxResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type IMachinesetIoBandwidthMax struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setIoBandwidthMax"`

	This           string `xml:"_this,omitempty"`
	IoBandwidthMax uint32 `xml:"ioBandwidthMax,omitempty"`
}

type IMachinesetIoBandwidthMaxResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setIoBandwidthMaxResponse"`
}

type IMachinesetBootOrder struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setBootOrder"`

	This     string      `xml:"_this,omitempty"`
	Position uint32      `xml:"position,omitempty"`
	Device   *DeviceType `xml:"device,omitempty"`
}

type IMachinesetBootOrderResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setBootOrderResponse"`
}

type IMachinegetBootOrder struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getBootOrder"`

	This     string `xml:"_this,omitempty"`
	Position uint32 `xml:"position,omitempty"`
}

type IMachinegetBootOrderResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getBootOrderResponse"`

	Returnval *DeviceType `xml:"returnval,omitempty"`
}

type IMachineattachDevice struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_attachDevice"`

	This           string      `xml:"_this,omitempty"`
	Name           string      `xml:"name,omitempty"`
	ControllerPort int32       `xml:"controllerPort,omitempty"`
	Device         int32       `xml:"device,omitempty"`
	Type_          *DeviceType `xml:"type,omitempty"`
	Id             string      `xml:"id,omitempty"`
}

type IMachineattachDeviceResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_attachDeviceResponse"`
}

type IMachinedetachDevice struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_detachDevice"`

	This           string `xml:"_this,omitempty"`
	Name           string `xml:"name,omitempty"`
	ControllerPort int32  `xml:"controllerPort,omitempty"`
	Device         int32  `xml:"device,omitempty"`
}

type IMachinedetachDeviceResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_detachDeviceResponse"`
}

type IMachinepassthroughDevice struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_passthroughDevice"`

	This           string `xml:"_this,omitempty"`
	Name           string `xml:"name,omitempty"`
	ControllerPort int32  `xml:"controllerPort,omitempty"`
	Device         int32  `xml:"device,omitempty"`
	Passthrough    bool   `xml:"passthrough,omitempty"`
}

type IMachinepassthroughDeviceResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_passthroughDeviceResponse"`
}

type IMachinemountMedium struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_mountMedium"`

	This           string `xml:"_this,omitempty"`
	Name           string `xml:"name,omitempty"`
	ControllerPort int32  `xml:"controllerPort,omitempty"`
	Device         int32  `xml:"device,omitempty"`
	Medium         string `xml:"medium,omitempty"`
	Force          bool   `xml:"force,omitempty"`
}

type IMachinemountMediumResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_mountMediumResponse"`
}

type IMachinegetMedium struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getMedium"`

	This           string `xml:"_this,omitempty"`
	Name           string `xml:"name,omitempty"`
	ControllerPort int32  `xml:"controllerPort,omitempty"`
	Device         int32  `xml:"device,omitempty"`
}

type IMachinegetMediumResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getMediumResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMachinegetMediumAttachmentsOfController struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getMediumAttachmentsOfController"`

	This string `xml:"_this,omitempty"`
	Name string `xml:"name,omitempty"`
}

type IMachinegetMediumAttachmentsOfControllerResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getMediumAttachmentsOfControllerResponse"`

	Returnval []*IMediumAttachment `xml:"returnval,omitempty"`
}

type IMachinegetMediumAttachment struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getMediumAttachment"`

	This           string `xml:"_this,omitempty"`
	Name           string `xml:"name,omitempty"`
	ControllerPort int32  `xml:"controllerPort,omitempty"`
	Device         int32  `xml:"device,omitempty"`
}

type IMachinegetMediumAttachmentResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getMediumAttachmentResponse"`

	Returnval *IMediumAttachment `xml:"returnval,omitempty"`
}

type IMachinegetNetworkAdapter struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getNetworkAdapter"`

	This string `xml:"_this,omitempty"`
	Slot uint32 `xml:"slot,omitempty"`
}

type IMachinegetNetworkAdapterResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getNetworkAdapterResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMachineaddStorageController struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_addStorageController"`

	This           string      `xml:"_this,omitempty"`
	Name           string      `xml:"name,omitempty"`
	ConnectionType *StorageBus `xml:"connectionType,omitempty"`
}

type IMachineaddStorageControllerResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_addStorageControllerResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMachinegetStorageControllerByName struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getStorageControllerByName"`

	This string `xml:"_this,omitempty"`
	Name string `xml:"name,omitempty"`
}

type IMachinegetStorageControllerByNameResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getStorageControllerByNameResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMachinegetStorageControllerByInstance struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getStorageControllerByInstance"`

	This     string `xml:"_this,omitempty"`
	Instance uint32 `xml:"instance,omitempty"`
}

type IMachinegetStorageControllerByInstanceResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getStorageControllerByInstanceResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMachineremoveStorageController struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_removeStorageController"`

	This string `xml:"_this,omitempty"`
	Name string `xml:"name,omitempty"`
}

type IMachineremoveStorageControllerResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_removeStorageControllerResponse"`
}

type IMachinegetSerialPort struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getSerialPort"`

	This string `xml:"_this,omitempty"`
	Slot uint32 `xml:"slot,omitempty"`
}

type IMachinegetSerialPortResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getSerialPortResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMachinegetParallelPort struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getParallelPort"`

	This string `xml:"_this,omitempty"`
	Slot uint32 `xml:"slot,omitempty"`
}

type IMachinegetParallelPortResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getParallelPortResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMachinegetExtraDataKeys struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getExtraDataKeys"`

	This string `xml:"_this,omitempty"`
}

type IMachinegetExtraDataKeysResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getExtraDataKeysResponse"`

	Returnval []string `xml:"returnval,omitempty"`
}

type IMachinegetExtraData struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getExtraData"`

	This string `xml:"_this,omitempty"`
	Key  string `xml:"key,omitempty"`
}

type IMachinegetExtraDataResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getExtraDataResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMachinesetExtraData struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setExtraData"`

	This  string `xml:"_this,omitempty"`
	Key   string `xml:"key,omitempty"`
	Value string `xml:"value,omitempty"`
}

type IMachinesetExtraDataResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setExtraDataResponse"`
}

type IMachinegetCPUProperty struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getCPUProperty"`

	This     string           `xml:"_this,omitempty"`
	Property *CPUPropertyType `xml:"property,omitempty"`
}

type IMachinegetCPUPropertyResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getCPUPropertyResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IMachinesetCPUProperty struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setCPUProperty"`

	This     string           `xml:"_this,omitempty"`
	Property *CPUPropertyType `xml:"property,omitempty"`
	Value    bool             `xml:"value,omitempty"`
}

type IMachinesetCPUPropertyResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setCPUPropertyResponse"`
}

type IMachinegetCPUIDLeaf struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getCPUIDLeaf"`

	This string `xml:"_this,omitempty"`
	Id   uint32 `xml:"id,omitempty"`
}

type IMachinegetCPUIDLeafResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getCPUIDLeafResponse"`

	ValEax uint32 `xml:"valEax,omitempty"`
	ValEbx uint32 `xml:"valEbx,omitempty"`
	ValEcx uint32 `xml:"valEcx,omitempty"`
	ValEdx uint32 `xml:"valEdx,omitempty"`
}

type IMachinesetCPUIDLeaf struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setCPUIDLeaf"`

	This   string `xml:"_this,omitempty"`
	Id     uint32 `xml:"id,omitempty"`
	ValEax uint32 `xml:"valEax,omitempty"`
	ValEbx uint32 `xml:"valEbx,omitempty"`
	ValEcx uint32 `xml:"valEcx,omitempty"`
	ValEdx uint32 `xml:"valEdx,omitempty"`
}

type IMachinesetCPUIDLeafResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setCPUIDLeafResponse"`
}

type IMachineremoveCPUIDLeaf struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_removeCPUIDLeaf"`

	This string `xml:"_this,omitempty"`
	Id   uint32 `xml:"id,omitempty"`
}

type IMachineremoveCPUIDLeafResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_removeCPUIDLeafResponse"`
}

type IMachineremoveAllCPUIDLeaves struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_removeAllCPUIDLeaves"`

	This string `xml:"_this,omitempty"`
}

type IMachineremoveAllCPUIDLeavesResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_removeAllCPUIDLeavesResponse"`
}

type IMachinegetHWVirtExProperty struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getHWVirtExProperty"`

	This     string                `xml:"_this,omitempty"`
	Property *HWVirtExPropertyType `xml:"property,omitempty"`
}

type IMachinegetHWVirtExPropertyResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getHWVirtExPropertyResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IMachinesetHWVirtExProperty struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setHWVirtExProperty"`

	This     string                `xml:"_this,omitempty"`
	Property *HWVirtExPropertyType `xml:"property,omitempty"`
	Value    bool                  `xml:"value,omitempty"`
}

type IMachinesetHWVirtExPropertyResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setHWVirtExPropertyResponse"`
}

type IMachinesaveSettings struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_saveSettings"`

	This string `xml:"_this,omitempty"`
}

type IMachinesaveSettingsResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_saveSettingsResponse"`
}

type IMachinediscardSettings struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_discardSettings"`

	This string `xml:"_this,omitempty"`
}

type IMachinediscardSettingsResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_discardSettingsResponse"`
}

type IMachinedeleteSettings struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_deleteSettings"`

	This string `xml:"_this,omitempty"`
}

type IMachinedeleteSettingsResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_deleteSettingsResponse"`
}

type IMachineexport struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_export"`

	This       string `xml:"_this,omitempty"`
	AAppliance string `xml:"aAppliance,omitempty"`
}

type IMachineexportResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_exportResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMachinegetSnapshot struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getSnapshot"`

	This string `xml:"_this,omitempty"`
	Id   string `xml:"id,omitempty"`
}

type IMachinegetSnapshotResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getSnapshotResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMachinefindSnapshot struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_findSnapshot"`

	This string `xml:"_this,omitempty"`
	Name string `xml:"name,omitempty"`
}

type IMachinefindSnapshotResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_findSnapshotResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMachinesetCurrentSnapshot struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setCurrentSnapshot"`

	This string `xml:"_this,omitempty"`
	Id   string `xml:"id,omitempty"`
}

type IMachinesetCurrentSnapshotResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setCurrentSnapshotResponse"`
}

type IMachinecreateSharedFolder struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_createSharedFolder"`

	This     string `xml:"_this,omitempty"`
	Name     string `xml:"name,omitempty"`
	HostPath string `xml:"hostPath,omitempty"`
	Writable bool   `xml:"writable,omitempty"`
}

type IMachinecreateSharedFolderResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_createSharedFolderResponse"`
}

type IMachineremoveSharedFolder struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_removeSharedFolder"`

	This string `xml:"_this,omitempty"`
	Name string `xml:"name,omitempty"`
}

type IMachineremoveSharedFolderResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_removeSharedFolderResponse"`
}

type IMachinecanShowConsoleWindow struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_canShowConsoleWindow"`

	This string `xml:"_this,omitempty"`
}

type IMachinecanShowConsoleWindowResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_canShowConsoleWindowResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IMachineshowConsoleWindow struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_showConsoleWindow"`

	This string `xml:"_this,omitempty"`
}

type IMachineshowConsoleWindowResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_showConsoleWindowResponse"`

	Returnval uint64 `xml:"returnval,omitempty"`
}

type IMachinegetGuestProperty struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getGuestProperty"`

	This string `xml:"_this,omitempty"`
	Name string `xml:"name,omitempty"`
}

type IMachinegetGuestPropertyResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getGuestPropertyResponse"`

	Value     string `xml:"value,omitempty"`
	Timestamp uint64 `xml:"timestamp,omitempty"`
	Flags     string `xml:"flags,omitempty"`
}

type IMachinegetGuestPropertyValue struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getGuestPropertyValue"`

	This     string `xml:"_this,omitempty"`
	Property string `xml:"property,omitempty"`
}

type IMachinegetGuestPropertyValueResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getGuestPropertyValueResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMachinegetGuestPropertyTimestamp struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getGuestPropertyTimestamp"`

	This     string `xml:"_this,omitempty"`
	Property string `xml:"property,omitempty"`
}

type IMachinegetGuestPropertyTimestampResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getGuestPropertyTimestampResponse"`

	Returnval uint64 `xml:"returnval,omitempty"`
}

type IMachinesetGuestProperty struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setGuestProperty"`

	This     string `xml:"_this,omitempty"`
	Property string `xml:"property,omitempty"`
	Value    string `xml:"value,omitempty"`
	Flags    string `xml:"flags,omitempty"`
}

type IMachinesetGuestPropertyResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setGuestPropertyResponse"`
}

type IMachinesetGuestPropertyValue struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setGuestPropertyValue"`

	This     string `xml:"_this,omitempty"`
	Property string `xml:"property,omitempty"`
	Value    string `xml:"value,omitempty"`
}

type IMachinesetGuestPropertyValueResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setGuestPropertyValueResponse"`
}

type IMachineenumerateGuestProperties struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_enumerateGuestProperties"`

	This     string `xml:"_this,omitempty"`
	Patterns string `xml:"patterns,omitempty"`
}

type IMachineenumerateGuestPropertiesResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_enumerateGuestPropertiesResponse"`

	Name      []string `xml:"name,omitempty"`
	Value     []string `xml:"value,omitempty"`
	Timestamp []uint64 `xml:"timestamp,omitempty"`
	Flags     []string `xml:"flags,omitempty"`
}

type IMachinequerySavedThumbnailSize struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_querySavedThumbnailSize"`

	This     string `xml:"_this,omitempty"`
	ScreenId uint32 `xml:"screenId,omitempty"`
}

type IMachinequerySavedThumbnailSizeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_querySavedThumbnailSizeResponse"`

	Size   uint32 `xml:"size,omitempty"`
	Width  uint32 `xml:"width,omitempty"`
	Height uint32 `xml:"height,omitempty"`
}

type IMachinereadSavedThumbnailToArray struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_readSavedThumbnailToArray"`

	This     string `xml:"_this,omitempty"`
	ScreenId uint32 `xml:"screenId,omitempty"`
	BGR      bool   `xml:"BGR,omitempty"`
}

type IMachinereadSavedThumbnailToArrayResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_readSavedThumbnailToArrayResponse"`

	Width     uint32 `xml:"width,omitempty"`
	Height    uint32 `xml:"height,omitempty"`
	Returnval []byte `xml:"returnval,omitempty"`
}

type IMachinequerySavedScreenshotPNGSize struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_querySavedScreenshotPNGSize"`

	This     string `xml:"_this,omitempty"`
	ScreenId uint32 `xml:"screenId,omitempty"`
}

type IMachinequerySavedScreenshotPNGSizeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_querySavedScreenshotPNGSizeResponse"`

	Size   uint32 `xml:"size,omitempty"`
	Width  uint32 `xml:"width,omitempty"`
	Height uint32 `xml:"height,omitempty"`
}

type IMachinereadSavedScreenshotPNGToArray struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_readSavedScreenshotPNGToArray"`

	This     string `xml:"_this,omitempty"`
	ScreenId uint32 `xml:"screenId,omitempty"`
}

type IMachinereadSavedScreenshotPNGToArrayResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_readSavedScreenshotPNGToArrayResponse"`

	Width     uint32 `xml:"width,omitempty"`
	Height    uint32 `xml:"height,omitempty"`
	Returnval []byte `xml:"returnval,omitempty"`
}

type IMachinehotPlugCPU struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_hotPlugCPU"`

	This string `xml:"_this,omitempty"`
	Cpu  uint32 `xml:"cpu,omitempty"`
}

type IMachinehotPlugCPUResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_hotPlugCPUResponse"`
}

type IMachinehotUnplugCPU struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_hotUnplugCPU"`

	This string `xml:"_this,omitempty"`
	Cpu  uint32 `xml:"cpu,omitempty"`
}

type IMachinehotUnplugCPUResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_hotUnplugCPUResponse"`
}

type IMachinegetCPUStatus struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getCPUStatus"`

	This string `xml:"_this,omitempty"`
	Cpu  uint32 `xml:"cpu,omitempty"`
}

type IMachinegetCPUStatusResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getCPUStatusResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IMachinequeryLogFilename struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_queryLogFilename"`

	This string `xml:"_this,omitempty"`
	Idx  uint32 `xml:"idx,omitempty"`
}

type IMachinequeryLogFilenameResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_queryLogFilenameResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMachinereadLog struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_readLog"`

	This   string `xml:"_this,omitempty"`
	Idx    uint32 `xml:"idx,omitempty"`
	Offset uint64 `xml:"offset,omitempty"`
	Size   uint64 `xml:"size,omitempty"`
}

type IMachinereadLogResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_readLogResponse"`

	Returnval []byte `xml:"returnval,omitempty"`
}

type IConsolegetMachine struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_getMachine"`

	This string `xml:"_this,omitempty"`
}

type IConsolegetMachineResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_getMachineResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IConsolegetState struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_getState"`

	This string `xml:"_this,omitempty"`
}

type IConsolegetStateResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_getStateResponse"`

	Returnval *MachineState `xml:"returnval,omitempty"`
}

type IConsolegetGuest struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_getGuest"`

	This string `xml:"_this,omitempty"`
}

type IConsolegetGuestResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_getGuestResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IConsolegetKeyboard struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_getKeyboard"`

	This string `xml:"_this,omitempty"`
}

type IConsolegetKeyboardResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_getKeyboardResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IConsolegetMouse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_getMouse"`

	This string `xml:"_this,omitempty"`
}

type IConsolegetMouseResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_getMouseResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IConsolegetDisplay struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_getDisplay"`

	This string `xml:"_this,omitempty"`
}

type IConsolegetDisplayResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_getDisplayResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IConsolegetUSBDevices struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_getUSBDevices"`

	This string `xml:"_this,omitempty"`
}

type IConsolegetUSBDevicesResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_getUSBDevicesResponse"`

	Returnval []string `xml:"returnval,omitempty"`
}

type IConsolegetRemoteUSBDevices struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_getRemoteUSBDevices"`

	This string `xml:"_this,omitempty"`
}

type IConsolegetRemoteUSBDevicesResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_getRemoteUSBDevicesResponse"`

	Returnval []string `xml:"returnval,omitempty"`
}

type IConsolegetSharedFolders struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_getSharedFolders"`

	This string `xml:"_this,omitempty"`
}

type IConsolegetSharedFoldersResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_getSharedFoldersResponse"`

	Returnval []*ISharedFolder `xml:"returnval,omitempty"`
}

type IConsolegetRemoteDisplayInfo struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_getRemoteDisplayInfo"`

	This string `xml:"_this,omitempty"`
}

type IConsolegetRemoteDisplayInfoResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_getRemoteDisplayInfoResponse"`

	Returnval *IRemoteDisplayInfo `xml:"returnval,omitempty"`
}

type IConsolepowerUp struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_powerUp"`

	This string `xml:"_this,omitempty"`
}

type IConsolepowerUpResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_powerUpResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IConsolepowerUpPaused struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_powerUpPaused"`

	This string `xml:"_this,omitempty"`
}

type IConsolepowerUpPausedResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_powerUpPausedResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IConsolepowerDown struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_powerDown"`

	This string `xml:"_this,omitempty"`
}

type IConsolepowerDownResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_powerDownResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IConsolereset struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_reset"`

	This string `xml:"_this,omitempty"`
}

type IConsoleresetResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_resetResponse"`
}

type IConsolepause struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_pause"`

	This string `xml:"_this,omitempty"`
}

type IConsolepauseResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_pauseResponse"`
}

type IConsoleresume struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_resume"`

	This string `xml:"_this,omitempty"`
}

type IConsoleresumeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_resumeResponse"`
}

type IConsolepowerButton struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_powerButton"`

	This string `xml:"_this,omitempty"`
}

type IConsolepowerButtonResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_powerButtonResponse"`
}

type IConsolesleepButton struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_sleepButton"`

	This string `xml:"_this,omitempty"`
}

type IConsolesleepButtonResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_sleepButtonResponse"`
}

type IConsolegetPowerButtonHandled struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_getPowerButtonHandled"`

	This string `xml:"_this,omitempty"`
}

type IConsolegetPowerButtonHandledResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_getPowerButtonHandledResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IConsolegetGuestEnteredACPIMode struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_getGuestEnteredACPIMode"`

	This string `xml:"_this,omitempty"`
}

type IConsolegetGuestEnteredACPIModeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_getGuestEnteredACPIModeResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IConsolesaveState struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_saveState"`

	This string `xml:"_this,omitempty"`
}

type IConsolesaveStateResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_saveStateResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IConsoleadoptSavedState struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_adoptSavedState"`

	This           string `xml:"_this,omitempty"`
	SavedStateFile string `xml:"savedStateFile,omitempty"`
}

type IConsoleadoptSavedStateResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_adoptSavedStateResponse"`
}

type IConsoleforgetSavedState struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_forgetSavedState"`

	This   string `xml:"_this,omitempty"`
	Remove bool   `xml:"remove,omitempty"`
}

type IConsoleforgetSavedStateResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_forgetSavedStateResponse"`
}

type IConsolegetDeviceActivity struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_getDeviceActivity"`

	This  string      `xml:"_this,omitempty"`
	Type_ *DeviceType `xml:"type,omitempty"`
}

type IConsolegetDeviceActivityResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_getDeviceActivityResponse"`

	Returnval *DeviceActivity `xml:"returnval,omitempty"`
}

type IConsoleattachUSBDevice struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_attachUSBDevice"`

	This string `xml:"_this,omitempty"`
	Id   string `xml:"id,omitempty"`
}

type IConsoleattachUSBDeviceResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_attachUSBDeviceResponse"`
}

type IConsoledetachUSBDevice struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_detachUSBDevice"`

	This string `xml:"_this,omitempty"`
	Id   string `xml:"id,omitempty"`
}

type IConsoledetachUSBDeviceResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_detachUSBDeviceResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IConsolefindUSBDeviceByAddress struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_findUSBDeviceByAddress"`

	This string `xml:"_this,omitempty"`
	Name string `xml:"name,omitempty"`
}

type IConsolefindUSBDeviceByAddressResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_findUSBDeviceByAddressResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IConsolefindUSBDeviceById struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_findUSBDeviceById"`

	This string `xml:"_this,omitempty"`
	Id   string `xml:"id,omitempty"`
}

type IConsolefindUSBDeviceByIdResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_findUSBDeviceByIdResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IConsolecreateSharedFolder struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_createSharedFolder"`

	This     string `xml:"_this,omitempty"`
	Name     string `xml:"name,omitempty"`
	HostPath string `xml:"hostPath,omitempty"`
	Writable bool   `xml:"writable,omitempty"`
}

type IConsolecreateSharedFolderResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_createSharedFolderResponse"`
}

type IConsoleremoveSharedFolder struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_removeSharedFolder"`

	This string `xml:"_this,omitempty"`
	Name string `xml:"name,omitempty"`
}

type IConsoleremoveSharedFolderResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_removeSharedFolderResponse"`
}

type IConsoletakeSnapshot struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_takeSnapshot"`

	This        string `xml:"_this,omitempty"`
	Name        string `xml:"name,omitempty"`
	Description string `xml:"description,omitempty"`
}

type IConsoletakeSnapshotResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_takeSnapshotResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IConsoledeleteSnapshot struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_deleteSnapshot"`

	This string `xml:"_this,omitempty"`
	Id   string `xml:"id,omitempty"`
}

type IConsoledeleteSnapshotResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_deleteSnapshotResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IConsolerestoreSnapshot struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_restoreSnapshot"`

	This     string `xml:"_this,omitempty"`
	Snapshot string `xml:"snapshot,omitempty"`
}

type IConsolerestoreSnapshotResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_restoreSnapshotResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IConsoleteleport struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_teleport"`

	This        string `xml:"_this,omitempty"`
	Hostname    string `xml:"hostname,omitempty"`
	Tcpport     uint32 `xml:"tcpport,omitempty"`
	Password    string `xml:"password,omitempty"`
	MaxDowntime uint32 `xml:"maxDowntime,omitempty"`
}

type IConsoleteleportResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_teleportResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IHostNetworkInterfacegetName struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHostNetworkInterface_getName"`

	This string `xml:"_this,omitempty"`
}

type IHostNetworkInterfacegetNameResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHostNetworkInterface_getNameResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IHostNetworkInterfacegetId struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHostNetworkInterface_getId"`

	This string `xml:"_this,omitempty"`
}

type IHostNetworkInterfacegetIdResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHostNetworkInterface_getIdResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IHostNetworkInterfacegetNetworkName struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHostNetworkInterface_getNetworkName"`

	This string `xml:"_this,omitempty"`
}

type IHostNetworkInterfacegetNetworkNameResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHostNetworkInterface_getNetworkNameResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IHostNetworkInterfacegetDhcpEnabled struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHostNetworkInterface_getDhcpEnabled"`

	This string `xml:"_this,omitempty"`
}

type IHostNetworkInterfacegetDhcpEnabledResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHostNetworkInterface_getDhcpEnabledResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IHostNetworkInterfacegetIPAddress struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHostNetworkInterface_getIPAddress"`

	This string `xml:"_this,omitempty"`
}

type IHostNetworkInterfacegetIPAddressResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHostNetworkInterface_getIPAddressResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IHostNetworkInterfacegetNetworkMask struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHostNetworkInterface_getNetworkMask"`

	This string `xml:"_this,omitempty"`
}

type IHostNetworkInterfacegetNetworkMaskResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHostNetworkInterface_getNetworkMaskResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IHostNetworkInterfacegetIPV6Supported struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHostNetworkInterface_getIPV6Supported"`

	This string `xml:"_this,omitempty"`
}

type IHostNetworkInterfacegetIPV6SupportedResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHostNetworkInterface_getIPV6SupportedResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IHostNetworkInterfacegetIPV6Address struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHostNetworkInterface_getIPV6Address"`

	This string `xml:"_this,omitempty"`
}

type IHostNetworkInterfacegetIPV6AddressResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHostNetworkInterface_getIPV6AddressResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IHostNetworkInterfacegetIPV6NetworkMaskPrefixLength struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHostNetworkInterface_getIPV6NetworkMaskPrefixLength"`

	This string `xml:"_this,omitempty"`
}

type IHostNetworkInterfacegetIPV6NetworkMaskPrefixLengthResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHostNetworkInterface_getIPV6NetworkMaskPrefixLengthResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type IHostNetworkInterfacegetHardwareAddress struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHostNetworkInterface_getHardwareAddress"`

	This string `xml:"_this,omitempty"`
}

type IHostNetworkInterfacegetHardwareAddressResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHostNetworkInterface_getHardwareAddressResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IHostNetworkInterfacegetMediumType struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHostNetworkInterface_getMediumType"`

	This string `xml:"_this,omitempty"`
}

type IHostNetworkInterfacegetMediumTypeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHostNetworkInterface_getMediumTypeResponse"`

	Returnval *HostNetworkInterfaceMediumType `xml:"returnval,omitempty"`
}

type IHostNetworkInterfacegetStatus struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHostNetworkInterface_getStatus"`

	This string `xml:"_this,omitempty"`
}

type IHostNetworkInterfacegetStatusResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHostNetworkInterface_getStatusResponse"`

	Returnval *HostNetworkInterfaceStatus `xml:"returnval,omitempty"`
}

type IHostNetworkInterfacegetInterfaceType struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHostNetworkInterface_getInterfaceType"`

	This string `xml:"_this,omitempty"`
}

type IHostNetworkInterfacegetInterfaceTypeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHostNetworkInterface_getInterfaceTypeResponse"`

	Returnval *HostNetworkInterfaceType `xml:"returnval,omitempty"`
}

type IHostNetworkInterfaceenableStaticIpConfig struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHostNetworkInterface_enableStaticIpConfig"`

	This        string `xml:"_this,omitempty"`
	IPAddress   string `xml:"IPAddress,omitempty"`
	NetworkMask string `xml:"networkMask,omitempty"`
}

type IHostNetworkInterfaceenableStaticIpConfigResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHostNetworkInterface_enableStaticIpConfigResponse"`
}

type IHostNetworkInterfaceenableStaticIpConfigV6 struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHostNetworkInterface_enableStaticIpConfigV6"`

	This                        string `xml:"_this,omitempty"`
	IPV6Address                 string `xml:"IPV6Address,omitempty"`
	IPV6NetworkMaskPrefixLength uint32 `xml:"IPV6NetworkMaskPrefixLength,omitempty"`
}

type IHostNetworkInterfaceenableStaticIpConfigV6Response struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHostNetworkInterface_enableStaticIpConfigV6Response"`
}

type IHostNetworkInterfaceenableDynamicIpConfig struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHostNetworkInterface_enableDynamicIpConfig"`

	This string `xml:"_this,omitempty"`
}

type IHostNetworkInterfaceenableDynamicIpConfigResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHostNetworkInterface_enableDynamicIpConfigResponse"`
}

type IHostNetworkInterfacedhcpRediscover struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHostNetworkInterface_dhcpRediscover"`

	This string `xml:"_this,omitempty"`
}

type IHostNetworkInterfacedhcpRediscoverResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHostNetworkInterface_dhcpRediscoverResponse"`
}

type IHostgetDVDDrives struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHost_getDVDDrives"`

	This string `xml:"_this,omitempty"`
}

type IHostgetDVDDrivesResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHost_getDVDDrivesResponse"`

	Returnval []string `xml:"returnval,omitempty"`
}

type IHostgetFloppyDrives struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHost_getFloppyDrives"`

	This string `xml:"_this,omitempty"`
}

type IHostgetFloppyDrivesResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHost_getFloppyDrivesResponse"`

	Returnval []string `xml:"returnval,omitempty"`
}

type IHostgetUSBDevices struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHost_getUSBDevices"`

	This string `xml:"_this,omitempty"`
}

type IHostgetUSBDevicesResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHost_getUSBDevicesResponse"`

	Returnval []string `xml:"returnval,omitempty"`
}

type IHostgetUSBDeviceFilters struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHost_getUSBDeviceFilters"`

	This string `xml:"_this,omitempty"`
}

type IHostgetUSBDeviceFiltersResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHost_getUSBDeviceFiltersResponse"`

	Returnval []string `xml:"returnval,omitempty"`
}

type IHostgetNetworkInterfaces struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHost_getNetworkInterfaces"`

	This string `xml:"_this,omitempty"`
}

type IHostgetNetworkInterfacesResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHost_getNetworkInterfacesResponse"`

	Returnval []string `xml:"returnval,omitempty"`
}

type IHostgetProcessorCount struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHost_getProcessorCount"`

	This string `xml:"_this,omitempty"`
}

type IHostgetProcessorCountResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHost_getProcessorCountResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type IHostgetProcessorOnlineCount struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHost_getProcessorOnlineCount"`

	This string `xml:"_this,omitempty"`
}

type IHostgetProcessorOnlineCountResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHost_getProcessorOnlineCountResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type IHostgetProcessorCoreCount struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHost_getProcessorCoreCount"`

	This string `xml:"_this,omitempty"`
}

type IHostgetProcessorCoreCountResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHost_getProcessorCoreCountResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type IHostgetMemorySize struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHost_getMemorySize"`

	This string `xml:"_this,omitempty"`
}

type IHostgetMemorySizeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHost_getMemorySizeResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type IHostgetMemoryAvailable struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHost_getMemoryAvailable"`

	This string `xml:"_this,omitempty"`
}

type IHostgetMemoryAvailableResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHost_getMemoryAvailableResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type IHostgetOperatingSystem struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHost_getOperatingSystem"`

	This string `xml:"_this,omitempty"`
}

type IHostgetOperatingSystemResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHost_getOperatingSystemResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IHostgetOSVersion struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHost_getOSVersion"`

	This string `xml:"_this,omitempty"`
}

type IHostgetOSVersionResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHost_getOSVersionResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IHostgetUTCTime struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHost_getUTCTime"`

	This string `xml:"_this,omitempty"`
}

type IHostgetUTCTimeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHost_getUTCTimeResponse"`

	Returnval int64 `xml:"returnval,omitempty"`
}

type IHostgetAcceleration3DAvailable struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHost_getAcceleration3DAvailable"`

	This string `xml:"_this,omitempty"`
}

type IHostgetAcceleration3DAvailableResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHost_getAcceleration3DAvailableResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IHostgetProcessorSpeed struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHost_getProcessorSpeed"`

	This  string `xml:"_this,omitempty"`
	CpuId uint32 `xml:"cpuId,omitempty"`
}

type IHostgetProcessorSpeedResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHost_getProcessorSpeedResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type IHostgetProcessorFeature struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHost_getProcessorFeature"`

	This    string            `xml:"_this,omitempty"`
	Feature *ProcessorFeature `xml:"feature,omitempty"`
}

type IHostgetProcessorFeatureResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHost_getProcessorFeatureResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IHostgetProcessorDescription struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHost_getProcessorDescription"`

	This  string `xml:"_this,omitempty"`
	CpuId uint32 `xml:"cpuId,omitempty"`
}

type IHostgetProcessorDescriptionResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHost_getProcessorDescriptionResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IHostgetProcessorCPUIDLeaf struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHost_getProcessorCPUIDLeaf"`

	This    string `xml:"_this,omitempty"`
	CpuId   uint32 `xml:"cpuId,omitempty"`
	Leaf    uint32 `xml:"leaf,omitempty"`
	SubLeaf uint32 `xml:"subLeaf,omitempty"`
}

type IHostgetProcessorCPUIDLeafResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHost_getProcessorCPUIDLeafResponse"`

	ValEax uint32 `xml:"valEax,omitempty"`
	ValEbx uint32 `xml:"valEbx,omitempty"`
	ValEcx uint32 `xml:"valEcx,omitempty"`
	ValEdx uint32 `xml:"valEdx,omitempty"`
}

type IHostcreateHostOnlyNetworkInterface struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHost_createHostOnlyNetworkInterface"`

	This string `xml:"_this,omitempty"`
}

type IHostcreateHostOnlyNetworkInterfaceResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHost_createHostOnlyNetworkInterfaceResponse"`

	HostInterface string `xml:"hostInterface,omitempty"`
	Returnval     string `xml:"returnval,omitempty"`
}

type IHostremoveHostOnlyNetworkInterface struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHost_removeHostOnlyNetworkInterface"`

	This string `xml:"_this,omitempty"`
	Id   string `xml:"id,omitempty"`
}

type IHostremoveHostOnlyNetworkInterfaceResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHost_removeHostOnlyNetworkInterfaceResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IHostcreateUSBDeviceFilter struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHost_createUSBDeviceFilter"`

	This string `xml:"_this,omitempty"`
	Name string `xml:"name,omitempty"`
}

type IHostcreateUSBDeviceFilterResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHost_createUSBDeviceFilterResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IHostinsertUSBDeviceFilter struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHost_insertUSBDeviceFilter"`

	This     string `xml:"_this,omitempty"`
	Position uint32 `xml:"position,omitempty"`
	Filter   string `xml:"filter,omitempty"`
}

type IHostinsertUSBDeviceFilterResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHost_insertUSBDeviceFilterResponse"`
}

type IHostremoveUSBDeviceFilter struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHost_removeUSBDeviceFilter"`

	This     string `xml:"_this,omitempty"`
	Position uint32 `xml:"position,omitempty"`
}

type IHostremoveUSBDeviceFilterResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHost_removeUSBDeviceFilterResponse"`
}

type IHostfindHostDVDDrive struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHost_findHostDVDDrive"`

	This string `xml:"_this,omitempty"`
	Name string `xml:"name,omitempty"`
}

type IHostfindHostDVDDriveResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHost_findHostDVDDriveResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IHostfindHostFloppyDrive struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHost_findHostFloppyDrive"`

	This string `xml:"_this,omitempty"`
	Name string `xml:"name,omitempty"`
}

type IHostfindHostFloppyDriveResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHost_findHostFloppyDriveResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IHostfindHostNetworkInterfaceByName struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHost_findHostNetworkInterfaceByName"`

	This string `xml:"_this,omitempty"`
	Name string `xml:"name,omitempty"`
}

type IHostfindHostNetworkInterfaceByNameResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHost_findHostNetworkInterfaceByNameResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IHostfindHostNetworkInterfaceById struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHost_findHostNetworkInterfaceById"`

	This string `xml:"_this,omitempty"`
	Id   string `xml:"id,omitempty"`
}

type IHostfindHostNetworkInterfaceByIdResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHost_findHostNetworkInterfaceByIdResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IHostfindHostNetworkInterfacesOfType struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHost_findHostNetworkInterfacesOfType"`

	This  string                    `xml:"_this,omitempty"`
	Type_ *HostNetworkInterfaceType `xml:"type,omitempty"`
}

type IHostfindHostNetworkInterfacesOfTypeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHost_findHostNetworkInterfacesOfTypeResponse"`

	Returnval []string `xml:"returnval,omitempty"`
}

type IHostfindUSBDeviceById struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHost_findUSBDeviceById"`

	This string `xml:"_this,omitempty"`
	Id   string `xml:"id,omitempty"`
}

type IHostfindUSBDeviceByIdResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHost_findUSBDeviceByIdResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IHostfindUSBDeviceByAddress struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHost_findUSBDeviceByAddress"`

	This string `xml:"_this,omitempty"`
	Name string `xml:"name,omitempty"`
}

type IHostfindUSBDeviceByAddressResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHost_findUSBDeviceByAddressResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type ISystemPropertiesgetMinGuestRAM struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_getMinGuestRAM"`

	This string `xml:"_this,omitempty"`
}

type ISystemPropertiesgetMinGuestRAMResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_getMinGuestRAMResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type ISystemPropertiesgetMaxGuestRAM struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_getMaxGuestRAM"`

	This string `xml:"_this,omitempty"`
}

type ISystemPropertiesgetMaxGuestRAMResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_getMaxGuestRAMResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type ISystemPropertiesgetMinGuestVRAM struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_getMinGuestVRAM"`

	This string `xml:"_this,omitempty"`
}

type ISystemPropertiesgetMinGuestVRAMResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_getMinGuestVRAMResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type ISystemPropertiesgetMaxGuestVRAM struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_getMaxGuestVRAM"`

	This string `xml:"_this,omitempty"`
}

type ISystemPropertiesgetMaxGuestVRAMResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_getMaxGuestVRAMResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type ISystemPropertiesgetMinGuestCPUCount struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_getMinGuestCPUCount"`

	This string `xml:"_this,omitempty"`
}

type ISystemPropertiesgetMinGuestCPUCountResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_getMinGuestCPUCountResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type ISystemPropertiesgetMaxGuestCPUCount struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_getMaxGuestCPUCount"`

	This string `xml:"_this,omitempty"`
}

type ISystemPropertiesgetMaxGuestCPUCountResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_getMaxGuestCPUCountResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type ISystemPropertiesgetMaxGuestMonitors struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_getMaxGuestMonitors"`

	This string `xml:"_this,omitempty"`
}

type ISystemPropertiesgetMaxGuestMonitorsResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_getMaxGuestMonitorsResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type ISystemPropertiesgetMaxVDISize struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_getMaxVDISize"`

	This string `xml:"_this,omitempty"`
}

type ISystemPropertiesgetMaxVDISizeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_getMaxVDISizeResponse"`

	Returnval uint64 `xml:"returnval,omitempty"`
}

type ISystemPropertiesgetNetworkAdapterCount struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_getNetworkAdapterCount"`

	This string `xml:"_this,omitempty"`
}

type ISystemPropertiesgetNetworkAdapterCountResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_getNetworkAdapterCountResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type ISystemPropertiesgetSerialPortCount struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_getSerialPortCount"`

	This string `xml:"_this,omitempty"`
}

type ISystemPropertiesgetSerialPortCountResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_getSerialPortCountResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type ISystemPropertiesgetParallelPortCount struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_getParallelPortCount"`

	This string `xml:"_this,omitempty"`
}

type ISystemPropertiesgetParallelPortCountResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_getParallelPortCountResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type ISystemPropertiesgetMaxBootPosition struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_getMaxBootPosition"`

	This string `xml:"_this,omitempty"`
}

type ISystemPropertiesgetMaxBootPositionResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_getMaxBootPositionResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type ISystemPropertiesgetDefaultMachineFolder struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_getDefaultMachineFolder"`

	This string `xml:"_this,omitempty"`
}

type ISystemPropertiesgetDefaultMachineFolderResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_getDefaultMachineFolderResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type ISystemPropertiessetDefaultMachineFolder struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_setDefaultMachineFolder"`

	This                 string `xml:"_this,omitempty"`
	DefaultMachineFolder string `xml:"defaultMachineFolder,omitempty"`
}

type ISystemPropertiessetDefaultMachineFolderResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_setDefaultMachineFolderResponse"`
}

type ISystemPropertiesgetDefaultHardDiskFolder struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_getDefaultHardDiskFolder"`

	This string `xml:"_this,omitempty"`
}

type ISystemPropertiesgetDefaultHardDiskFolderResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_getDefaultHardDiskFolderResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type ISystemPropertiessetDefaultHardDiskFolder struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_setDefaultHardDiskFolder"`

	This                  string `xml:"_this,omitempty"`
	DefaultHardDiskFolder string `xml:"defaultHardDiskFolder,omitempty"`
}

type ISystemPropertiessetDefaultHardDiskFolderResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_setDefaultHardDiskFolderResponse"`
}

type ISystemPropertiesgetMediumFormats struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_getMediumFormats"`

	This string `xml:"_this,omitempty"`
}

type ISystemPropertiesgetMediumFormatsResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_getMediumFormatsResponse"`

	Returnval []string `xml:"returnval,omitempty"`
}

type ISystemPropertiesgetDefaultHardDiskFormat struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_getDefaultHardDiskFormat"`

	This string `xml:"_this,omitempty"`
}

type ISystemPropertiesgetDefaultHardDiskFormatResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_getDefaultHardDiskFormatResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type ISystemPropertiessetDefaultHardDiskFormat struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_setDefaultHardDiskFormat"`

	This                  string `xml:"_this,omitempty"`
	DefaultHardDiskFormat string `xml:"defaultHardDiskFormat,omitempty"`
}

type ISystemPropertiessetDefaultHardDiskFormatResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_setDefaultHardDiskFormatResponse"`
}

type ISystemPropertiesgetFreeDiskSpaceWarning struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_getFreeDiskSpaceWarning"`

	This string `xml:"_this,omitempty"`
}

type ISystemPropertiesgetFreeDiskSpaceWarningResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_getFreeDiskSpaceWarningResponse"`

	Returnval uint64 `xml:"returnval,omitempty"`
}

type ISystemPropertiessetFreeDiskSpaceWarning struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_setFreeDiskSpaceWarning"`

	This                 string `xml:"_this,omitempty"`
	FreeDiskSpaceWarning uint64 `xml:"freeDiskSpaceWarning,omitempty"`
}

type ISystemPropertiessetFreeDiskSpaceWarningResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_setFreeDiskSpaceWarningResponse"`
}

type ISystemPropertiesgetFreeDiskSpacePercentWarning struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_getFreeDiskSpacePercentWarning"`

	This string `xml:"_this,omitempty"`
}

type ISystemPropertiesgetFreeDiskSpacePercentWarningResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_getFreeDiskSpacePercentWarningResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type ISystemPropertiessetFreeDiskSpacePercentWarning struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_setFreeDiskSpacePercentWarning"`

	This                        string `xml:"_this,omitempty"`
	FreeDiskSpacePercentWarning uint32 `xml:"freeDiskSpacePercentWarning,omitempty"`
}

type ISystemPropertiessetFreeDiskSpacePercentWarningResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_setFreeDiskSpacePercentWarningResponse"`
}

type ISystemPropertiesgetFreeDiskSpaceError struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_getFreeDiskSpaceError"`

	This string `xml:"_this,omitempty"`
}

type ISystemPropertiesgetFreeDiskSpaceErrorResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_getFreeDiskSpaceErrorResponse"`

	Returnval uint64 `xml:"returnval,omitempty"`
}

type ISystemPropertiessetFreeDiskSpaceError struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_setFreeDiskSpaceError"`

	This               string `xml:"_this,omitempty"`
	FreeDiskSpaceError uint64 `xml:"freeDiskSpaceError,omitempty"`
}

type ISystemPropertiessetFreeDiskSpaceErrorResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_setFreeDiskSpaceErrorResponse"`
}

type ISystemPropertiesgetFreeDiskSpacePercentError struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_getFreeDiskSpacePercentError"`

	This string `xml:"_this,omitempty"`
}

type ISystemPropertiesgetFreeDiskSpacePercentErrorResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_getFreeDiskSpacePercentErrorResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type ISystemPropertiessetFreeDiskSpacePercentError struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_setFreeDiskSpacePercentError"`

	This                      string `xml:"_this,omitempty"`
	FreeDiskSpacePercentError uint32 `xml:"freeDiskSpacePercentError,omitempty"`
}

type ISystemPropertiessetFreeDiskSpacePercentErrorResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_setFreeDiskSpacePercentErrorResponse"`
}

type ISystemPropertiesgetRemoteDisplayAuthLibrary struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_getRemoteDisplayAuthLibrary"`

	This string `xml:"_this,omitempty"`
}

type ISystemPropertiesgetRemoteDisplayAuthLibraryResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_getRemoteDisplayAuthLibraryResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type ISystemPropertiessetRemoteDisplayAuthLibrary struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_setRemoteDisplayAuthLibrary"`

	This                     string `xml:"_this,omitempty"`
	RemoteDisplayAuthLibrary string `xml:"remoteDisplayAuthLibrary,omitempty"`
}

type ISystemPropertiessetRemoteDisplayAuthLibraryResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_setRemoteDisplayAuthLibraryResponse"`
}

type ISystemPropertiesgetWebServiceAuthLibrary struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_getWebServiceAuthLibrary"`

	This string `xml:"_this,omitempty"`
}

type ISystemPropertiesgetWebServiceAuthLibraryResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_getWebServiceAuthLibraryResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type ISystemPropertiessetWebServiceAuthLibrary struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_setWebServiceAuthLibrary"`

	This                  string `xml:"_this,omitempty"`
	WebServiceAuthLibrary string `xml:"webServiceAuthLibrary,omitempty"`
}

type ISystemPropertiessetWebServiceAuthLibraryResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_setWebServiceAuthLibraryResponse"`
}

type ISystemPropertiesgetLogHistoryCount struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_getLogHistoryCount"`

	This string `xml:"_this,omitempty"`
}

type ISystemPropertiesgetLogHistoryCountResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_getLogHistoryCountResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type ISystemPropertiessetLogHistoryCount struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_setLogHistoryCount"`

	This            string `xml:"_this,omitempty"`
	LogHistoryCount uint32 `xml:"LogHistoryCount,omitempty"`
}

type ISystemPropertiessetLogHistoryCountResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_setLogHistoryCountResponse"`
}

type ISystemPropertiesgetDefaultAudioDriver struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_getDefaultAudioDriver"`

	This string `xml:"_this,omitempty"`
}

type ISystemPropertiesgetDefaultAudioDriverResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_getDefaultAudioDriverResponse"`

	Returnval *AudioDriverType `xml:"returnval,omitempty"`
}

type ISystemPropertiesgetMaxDevicesPerPortForStorageBus struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_getMaxDevicesPerPortForStorageBus"`

	This string      `xml:"_this,omitempty"`
	Bus  *StorageBus `xml:"bus,omitempty"`
}

type ISystemPropertiesgetMaxDevicesPerPortForStorageBusResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_getMaxDevicesPerPortForStorageBusResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type ISystemPropertiesgetMinPortCountForStorageBus struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_getMinPortCountForStorageBus"`

	This string      `xml:"_this,omitempty"`
	Bus  *StorageBus `xml:"bus,omitempty"`
}

type ISystemPropertiesgetMinPortCountForStorageBusResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_getMinPortCountForStorageBusResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type ISystemPropertiesgetMaxPortCountForStorageBus struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_getMaxPortCountForStorageBus"`

	This string      `xml:"_this,omitempty"`
	Bus  *StorageBus `xml:"bus,omitempty"`
}

type ISystemPropertiesgetMaxPortCountForStorageBusResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_getMaxPortCountForStorageBusResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type ISystemPropertiesgetMaxInstancesOfStorageBus struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_getMaxInstancesOfStorageBus"`

	This string      `xml:"_this,omitempty"`
	Bus  *StorageBus `xml:"bus,omitempty"`
}

type ISystemPropertiesgetMaxInstancesOfStorageBusResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_getMaxInstancesOfStorageBusResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type ISystemPropertiesgetDeviceTypesForStorageBus struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_getDeviceTypesForStorageBus"`

	This string      `xml:"_this,omitempty"`
	Bus  *StorageBus `xml:"bus,omitempty"`
}

type ISystemPropertiesgetDeviceTypesForStorageBusResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_getDeviceTypesForStorageBusResponse"`

	Returnval []*DeviceType `xml:"returnval,omitempty"`
}

type IGuestgetOSTypeId struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuest_getOSTypeId"`

	This string `xml:"_this,omitempty"`
}

type IGuestgetOSTypeIdResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuest_getOSTypeIdResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IGuestgetAdditionsActive struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuest_getAdditionsActive"`

	This string `xml:"_this,omitempty"`
}

type IGuestgetAdditionsActiveResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuest_getAdditionsActiveResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IGuestgetAdditionsVersion struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuest_getAdditionsVersion"`

	This string `xml:"_this,omitempty"`
}

type IGuestgetAdditionsVersionResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuest_getAdditionsVersionResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IGuestgetSupportsSeamless struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuest_getSupportsSeamless"`

	This string `xml:"_this,omitempty"`
}

type IGuestgetSupportsSeamlessResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuest_getSupportsSeamlessResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IGuestgetSupportsGraphics struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuest_getSupportsGraphics"`

	This string `xml:"_this,omitempty"`
}

type IGuestgetSupportsGraphicsResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuest_getSupportsGraphicsResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IGuestgetMemoryBalloonSize struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuest_getMemoryBalloonSize"`

	This string `xml:"_this,omitempty"`
}

type IGuestgetMemoryBalloonSizeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuest_getMemoryBalloonSizeResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type IGuestsetMemoryBalloonSize struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuest_setMemoryBalloonSize"`

	This              string `xml:"_this,omitempty"`
	MemoryBalloonSize uint32 `xml:"memoryBalloonSize,omitempty"`
}

type IGuestsetMemoryBalloonSizeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuest_setMemoryBalloonSizeResponse"`
}

type IGuestgetPageFusionEnabled struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuest_getPageFusionEnabled"`

	This string `xml:"_this,omitempty"`
}

type IGuestgetPageFusionEnabledResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuest_getPageFusionEnabledResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IGuestsetPageFusionEnabled struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuest_setPageFusionEnabled"`

	This              string `xml:"_this,omitempty"`
	PageFusionEnabled bool   `xml:"pageFusionEnabled,omitempty"`
}

type IGuestsetPageFusionEnabledResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuest_setPageFusionEnabledResponse"`
}

type IGuestgetStatisticsUpdateInterval struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuest_getStatisticsUpdateInterval"`

	This string `xml:"_this,omitempty"`
}

type IGuestgetStatisticsUpdateIntervalResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuest_getStatisticsUpdateIntervalResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type IGuestsetStatisticsUpdateInterval struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuest_setStatisticsUpdateInterval"`

	This                     string `xml:"_this,omitempty"`
	StatisticsUpdateInterval uint32 `xml:"statisticsUpdateInterval,omitempty"`
}

type IGuestsetStatisticsUpdateIntervalResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuest_setStatisticsUpdateIntervalResponse"`
}

type IGuestinternalGetStatistics struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuest_internalGetStatistics"`

	This string `xml:"_this,omitempty"`
}

type IGuestinternalGetStatisticsResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuest_internalGetStatisticsResponse"`

	CpuUser         uint32 `xml:"cpuUser,omitempty"`
	CpuKernel       uint32 `xml:"cpuKernel,omitempty"`
	CpuIdle         uint32 `xml:"cpuIdle,omitempty"`
	MemTotal        uint32 `xml:"memTotal,omitempty"`
	MemFree         uint32 `xml:"memFree,omitempty"`
	MemBalloon      uint32 `xml:"memBalloon,omitempty"`
	MemShared       uint32 `xml:"memShared,omitempty"`
	MemCache        uint32 `xml:"memCache,omitempty"`
	PagedTotal      uint32 `xml:"pagedTotal,omitempty"`
	MemAllocTotal   uint32 `xml:"memAllocTotal,omitempty"`
	MemFreeTotal    uint32 `xml:"memFreeTotal,omitempty"`
	MemBalloonTotal uint32 `xml:"memBalloonTotal,omitempty"`
	MemSharedTotal  uint32 `xml:"memSharedTotal,omitempty"`
}

type IGuestsetCredentials struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuest_setCredentials"`

	This                  string `xml:"_this,omitempty"`
	UserName              string `xml:"userName,omitempty"`
	Password              string `xml:"password,omitempty"`
	Domain                string `xml:"domain,omitempty"`
	AllowInteractiveLogon bool   `xml:"allowInteractiveLogon,omitempty"`
}

type IGuestsetCredentialsResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuest_setCredentialsResponse"`
}

type IGuestexecuteProcess struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuest_executeProcess"`

	This        string   `xml:"_this,omitempty"`
	ExecName    string   `xml:"execName,omitempty"`
	Flags       uint32   `xml:"flags,omitempty"`
	Arguments   []string `xml:"arguments,omitempty"`
	Environment []string `xml:"environment,omitempty"`
	UserName    string   `xml:"userName,omitempty"`
	Password    string   `xml:"password,omitempty"`
	TimeoutMS   uint32   `xml:"timeoutMS,omitempty"`
}

type IGuestexecuteProcessResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuest_executeProcessResponse"`

	Pid       uint32 `xml:"pid,omitempty"`
	Returnval string `xml:"returnval,omitempty"`
}

type IGuestgetProcessOutput struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuest_getProcessOutput"`

	This      string `xml:"_this,omitempty"`
	Pid       uint32 `xml:"pid,omitempty"`
	Flags     uint32 `xml:"flags,omitempty"`
	TimeoutMS uint32 `xml:"timeoutMS,omitempty"`
	Size      uint64 `xml:"size,omitempty"`
}

type IGuestgetProcessOutputResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuest_getProcessOutputResponse"`

	Returnval []byte `xml:"returnval,omitempty"`
}

type IGuestgetProcessStatus struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuest_getProcessStatus"`

	This string `xml:"_this,omitempty"`
	Pid  uint32 `xml:"pid,omitempty"`
}

type IGuestgetProcessStatusResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuest_getProcessStatusResponse"`

	Exitcode  uint32 `xml:"exitcode,omitempty"`
	Flags     uint32 `xml:"flags,omitempty"`
	Returnval uint32 `xml:"returnval,omitempty"`
}

type IProgressgetId struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IProgress_getId"`

	This string `xml:"_this,omitempty"`
}

type IProgressgetIdResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IProgress_getIdResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IProgressgetDescription struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IProgress_getDescription"`

	This string `xml:"_this,omitempty"`
}

type IProgressgetDescriptionResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IProgress_getDescriptionResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IProgressgetInitiator struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IProgress_getInitiator"`

	This string `xml:"_this,omitempty"`
}

type IProgressgetInitiatorResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IProgress_getInitiatorResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IProgressgetCancelable struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IProgress_getCancelable"`

	This string `xml:"_this,omitempty"`
}

type IProgressgetCancelableResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IProgress_getCancelableResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IProgressgetPercent struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IProgress_getPercent"`

	This string `xml:"_this,omitempty"`
}

type IProgressgetPercentResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IProgress_getPercentResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type IProgressgetTimeRemaining struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IProgress_getTimeRemaining"`

	This string `xml:"_this,omitempty"`
}

type IProgressgetTimeRemainingResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IProgress_getTimeRemainingResponse"`

	Returnval int32 `xml:"returnval,omitempty"`
}

type IProgressgetCompleted struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IProgress_getCompleted"`

	This string `xml:"_this,omitempty"`
}

type IProgressgetCompletedResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IProgress_getCompletedResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IProgressgetCanceled struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IProgress_getCanceled"`

	This string `xml:"_this,omitempty"`
}

type IProgressgetCanceledResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IProgress_getCanceledResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IProgressgetResultCode struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IProgress_getResultCode"`

	This string `xml:"_this,omitempty"`
}

type IProgressgetResultCodeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IProgress_getResultCodeResponse"`

	Returnval int32 `xml:"returnval,omitempty"`
}

type IProgressgetErrorInfo struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IProgress_getErrorInfo"`

	This string `xml:"_this,omitempty"`
}

type IProgressgetErrorInfoResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IProgress_getErrorInfoResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IProgressgetOperationCount struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IProgress_getOperationCount"`

	This string `xml:"_this,omitempty"`
}

type IProgressgetOperationCountResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IProgress_getOperationCountResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type IProgressgetOperation struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IProgress_getOperation"`

	This string `xml:"_this,omitempty"`
}

type IProgressgetOperationResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IProgress_getOperationResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type IProgressgetOperationDescription struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IProgress_getOperationDescription"`

	This string `xml:"_this,omitempty"`
}

type IProgressgetOperationDescriptionResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IProgress_getOperationDescriptionResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IProgressgetOperationPercent struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IProgress_getOperationPercent"`

	This string `xml:"_this,omitempty"`
}

type IProgressgetOperationPercentResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IProgress_getOperationPercentResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type IProgressgetTimeout struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IProgress_getTimeout"`

	This string `xml:"_this,omitempty"`
}

type IProgressgetTimeoutResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IProgress_getTimeoutResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type IProgresssetTimeout struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IProgress_setTimeout"`

	This    string `xml:"_this,omitempty"`
	Timeout uint32 `xml:"timeout,omitempty"`
}

type IProgresssetTimeoutResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IProgress_setTimeoutResponse"`
}

type IProgresssetCurrentOperationProgress struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IProgress_setCurrentOperationProgress"`

	This    string `xml:"_this,omitempty"`
	Percent uint32 `xml:"percent,omitempty"`
}

type IProgresssetCurrentOperationProgressResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IProgress_setCurrentOperationProgressResponse"`
}

type IProgresssetNextOperation struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IProgress_setNextOperation"`

	This                     string `xml:"_this,omitempty"`
	NextOperationDescription string `xml:"nextOperationDescription,omitempty"`
	NextOperationsWeight     uint32 `xml:"nextOperationsWeight,omitempty"`
}

type IProgresssetNextOperationResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IProgress_setNextOperationResponse"`
}

type IProgresswaitForCompletion struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IProgress_waitForCompletion"`

	This    string `xml:"_this,omitempty"`
	Timeout int32  `xml:"timeout,omitempty"`
}

type IProgresswaitForCompletionResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IProgress_waitForCompletionResponse"`
}

type IProgresswaitForOperationCompletion struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IProgress_waitForOperationCompletion"`

	This      string `xml:"_this,omitempty"`
	Operation uint32 `xml:"operation,omitempty"`
	Timeout   int32  `xml:"timeout,omitempty"`
}

type IProgresswaitForOperationCompletionResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IProgress_waitForOperationCompletionResponse"`
}

type IProgresscancel struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IProgress_cancel"`

	This string `xml:"_this,omitempty"`
}

type IProgresscancelResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IProgress_cancelResponse"`
}

type ISnapshotgetId struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISnapshot_getId"`

	This string `xml:"_this,omitempty"`
}

type ISnapshotgetIdResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISnapshot_getIdResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type ISnapshotgetName struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISnapshot_getName"`

	This string `xml:"_this,omitempty"`
}

type ISnapshotgetNameResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISnapshot_getNameResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type ISnapshotsetName struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISnapshot_setName"`

	This string `xml:"_this,omitempty"`
	Name string `xml:"name,omitempty"`
}

type ISnapshotsetNameResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISnapshot_setNameResponse"`
}

type ISnapshotgetDescription struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISnapshot_getDescription"`

	This string `xml:"_this,omitempty"`
}

type ISnapshotgetDescriptionResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISnapshot_getDescriptionResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type ISnapshotsetDescription struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISnapshot_setDescription"`

	This        string `xml:"_this,omitempty"`
	Description string `xml:"description,omitempty"`
}

type ISnapshotsetDescriptionResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISnapshot_setDescriptionResponse"`
}

type ISnapshotgetTimeStamp struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISnapshot_getTimeStamp"`

	This string `xml:"_this,omitempty"`
}

type ISnapshotgetTimeStampResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISnapshot_getTimeStampResponse"`

	Returnval int64 `xml:"returnval,omitempty"`
}

type ISnapshotgetOnline struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISnapshot_getOnline"`

	This string `xml:"_this,omitempty"`
}

type ISnapshotgetOnlineResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISnapshot_getOnlineResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type ISnapshotgetMachine struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISnapshot_getMachine"`

	This string `xml:"_this,omitempty"`
}

type ISnapshotgetMachineResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISnapshot_getMachineResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type ISnapshotgetParent struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISnapshot_getParent"`

	This string `xml:"_this,omitempty"`
}

type ISnapshotgetParentResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISnapshot_getParentResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type ISnapshotgetChildren struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISnapshot_getChildren"`

	This string `xml:"_this,omitempty"`
}

type ISnapshotgetChildrenResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISnapshot_getChildrenResponse"`

	Returnval []string `xml:"returnval,omitempty"`
}

type IMediumgetId struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_getId"`

	This string `xml:"_this,omitempty"`
}

type IMediumgetIdResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_getIdResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMediumgetDescription struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_getDescription"`

	This string `xml:"_this,omitempty"`
}

type IMediumgetDescriptionResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_getDescriptionResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMediumsetDescription struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_setDescription"`

	This        string `xml:"_this,omitempty"`
	Description string `xml:"description,omitempty"`
}

type IMediumsetDescriptionResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_setDescriptionResponse"`
}

type IMediumgetState struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_getState"`

	This string `xml:"_this,omitempty"`
}

type IMediumgetStateResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_getStateResponse"`

	Returnval *MediumState `xml:"returnval,omitempty"`
}

type IMediumgetLocation struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_getLocation"`

	This string `xml:"_this,omitempty"`
}

type IMediumgetLocationResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_getLocationResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMediumsetLocation struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_setLocation"`

	This     string `xml:"_this,omitempty"`
	Location string `xml:"location,omitempty"`
}

type IMediumsetLocationResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_setLocationResponse"`
}

type IMediumgetName struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_getName"`

	This string `xml:"_this,omitempty"`
}

type IMediumgetNameResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_getNameResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMediumgetDeviceType struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_getDeviceType"`

	This string `xml:"_this,omitempty"`
}

type IMediumgetDeviceTypeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_getDeviceTypeResponse"`

	Returnval *DeviceType `xml:"returnval,omitempty"`
}

type IMediumgetHostDrive struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_getHostDrive"`

	This string `xml:"_this,omitempty"`
}

type IMediumgetHostDriveResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_getHostDriveResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IMediumgetSize struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_getSize"`

	This string `xml:"_this,omitempty"`
}

type IMediumgetSizeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_getSizeResponse"`

	Returnval uint64 `xml:"returnval,omitempty"`
}

type IMediumgetFormat struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_getFormat"`

	This string `xml:"_this,omitempty"`
}

type IMediumgetFormatResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_getFormatResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMediumgetMediumFormat struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_getMediumFormat"`

	This string `xml:"_this,omitempty"`
}

type IMediumgetMediumFormatResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_getMediumFormatResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMediumgetType struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_getType"`

	This string `xml:"_this,omitempty"`
}

type IMediumgetTypeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_getTypeResponse"`

	Returnval *MediumType `xml:"returnval,omitempty"`
}

type IMediumsetType struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_setType"`

	This  string      `xml:"_this,omitempty"`
	Type_ *MediumType `xml:"type,omitempty"`
}

type IMediumsetTypeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_setTypeResponse"`
}

type IMediumgetParent struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_getParent"`

	This string `xml:"_this,omitempty"`
}

type IMediumgetParentResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_getParentResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMediumgetChildren struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_getChildren"`

	This string `xml:"_this,omitempty"`
}

type IMediumgetChildrenResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_getChildrenResponse"`

	Returnval []string `xml:"returnval,omitempty"`
}

type IMediumgetBase struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_getBase"`

	This string `xml:"_this,omitempty"`
}

type IMediumgetBaseResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_getBaseResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMediumgetReadOnly struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_getReadOnly"`

	This string `xml:"_this,omitempty"`
}

type IMediumgetReadOnlyResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_getReadOnlyResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IMediumgetLogicalSize struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_getLogicalSize"`

	This string `xml:"_this,omitempty"`
}

type IMediumgetLogicalSizeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_getLogicalSizeResponse"`

	Returnval uint64 `xml:"returnval,omitempty"`
}

type IMediumgetAutoReset struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_getAutoReset"`

	This string `xml:"_this,omitempty"`
}

type IMediumgetAutoResetResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_getAutoResetResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IMediumsetAutoReset struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_setAutoReset"`

	This      string `xml:"_this,omitempty"`
	AutoReset bool   `xml:"autoReset,omitempty"`
}

type IMediumsetAutoResetResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_setAutoResetResponse"`
}

type IMediumgetLastAccessError struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_getLastAccessError"`

	This string `xml:"_this,omitempty"`
}

type IMediumgetLastAccessErrorResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_getLastAccessErrorResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMediumgetMachineIds struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_getMachineIds"`

	This string `xml:"_this,omitempty"`
}

type IMediumgetMachineIdsResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_getMachineIdsResponse"`

	Returnval []string `xml:"returnval,omitempty"`
}

type IMediumrefreshState struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_refreshState"`

	This string `xml:"_this,omitempty"`
}

type IMediumrefreshStateResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_refreshStateResponse"`

	Returnval *MediumState `xml:"returnval,omitempty"`
}

type IMediumgetSnapshotIds struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_getSnapshotIds"`

	This      string `xml:"_this,omitempty"`
	MachineId string `xml:"machineId,omitempty"`
}

type IMediumgetSnapshotIdsResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_getSnapshotIdsResponse"`

	Returnval []string `xml:"returnval,omitempty"`
}

type IMediumlockRead struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_lockRead"`

	This string `xml:"_this,omitempty"`
}

type IMediumlockReadResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_lockReadResponse"`

	Returnval *MediumState `xml:"returnval,omitempty"`
}

type IMediumunlockRead struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_unlockRead"`

	This string `xml:"_this,omitempty"`
}

type IMediumunlockReadResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_unlockReadResponse"`

	Returnval *MediumState `xml:"returnval,omitempty"`
}

type IMediumlockWrite struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_lockWrite"`

	This string `xml:"_this,omitempty"`
}

type IMediumlockWriteResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_lockWriteResponse"`

	Returnval *MediumState `xml:"returnval,omitempty"`
}

type IMediumunlockWrite struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_unlockWrite"`

	This string `xml:"_this,omitempty"`
}

type IMediumunlockWriteResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_unlockWriteResponse"`

	Returnval *MediumState `xml:"returnval,omitempty"`
}

type IMediumclose struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_close"`

	This string `xml:"_this,omitempty"`
}

type IMediumcloseResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_closeResponse"`
}

type IMediumgetProperty struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_getProperty"`

	This string `xml:"_this,omitempty"`
	Name string `xml:"name,omitempty"`
}

type IMediumgetPropertyResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_getPropertyResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMediumsetProperty struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_setProperty"`

	This  string `xml:"_this,omitempty"`
	Name  string `xml:"name,omitempty"`
	Value string `xml:"value,omitempty"`
}

type IMediumsetPropertyResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_setPropertyResponse"`
}

type IMediumgetProperties struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_getProperties"`

	This  string `xml:"_this,omitempty"`
	Names string `xml:"names,omitempty"`
}

type IMediumgetPropertiesResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_getPropertiesResponse"`

	ReturnNames []string `xml:"returnNames,omitempty"`
	Returnval   []string `xml:"returnval,omitempty"`
}

type IMediumsetProperties struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_setProperties"`

	This   string   `xml:"_this,omitempty"`
	Names  []string `xml:"names,omitempty"`
	Values []string `xml:"values,omitempty"`
}

type IMediumsetPropertiesResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_setPropertiesResponse"`
}

type IMediumcreateBaseStorage struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_createBaseStorage"`

	This        string         `xml:"_this,omitempty"`
	LogicalSize uint64         `xml:"logicalSize,omitempty"`
	Variant     *MediumVariant `xml:"variant,omitempty"`
}

type IMediumcreateBaseStorageResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_createBaseStorageResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMediumdeleteStorage struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_deleteStorage"`

	This string `xml:"_this,omitempty"`
}

type IMediumdeleteStorageResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_deleteStorageResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMediumcreateDiffStorage struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_createDiffStorage"`

	This    string         `xml:"_this,omitempty"`
	Target  string         `xml:"target,omitempty"`
	Variant *MediumVariant `xml:"variant,omitempty"`
}

type IMediumcreateDiffStorageResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_createDiffStorageResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMediummergeTo struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_mergeTo"`

	This   string `xml:"_this,omitempty"`
	Target string `xml:"target,omitempty"`
}

type IMediummergeToResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_mergeToResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMediumcloneTo struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_cloneTo"`

	This    string         `xml:"_this,omitempty"`
	Target  string         `xml:"target,omitempty"`
	Variant *MediumVariant `xml:"variant,omitempty"`
	Parent  string         `xml:"parent,omitempty"`
}

type IMediumcloneToResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_cloneToResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMediumcompact struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_compact"`

	This string `xml:"_this,omitempty"`
}

type IMediumcompactResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_compactResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMediumresize struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_resize"`

	This        string `xml:"_this,omitempty"`
	LogicalSize uint64 `xml:"logicalSize,omitempty"`
}

type IMediumresizeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_resizeResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMediumreset struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_reset"`

	This string `xml:"_this,omitempty"`
}

type IMediumresetResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_resetResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMediumFormatgetId struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMediumFormat_getId"`

	This string `xml:"_this,omitempty"`
}

type IMediumFormatgetIdResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMediumFormat_getIdResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMediumFormatgetName struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMediumFormat_getName"`

	This string `xml:"_this,omitempty"`
}

type IMediumFormatgetNameResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMediumFormat_getNameResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMediumFormatgetFileExtensions struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMediumFormat_getFileExtensions"`

	This string `xml:"_this,omitempty"`
}

type IMediumFormatgetFileExtensionsResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMediumFormat_getFileExtensionsResponse"`

	Returnval []string `xml:"returnval,omitempty"`
}

type IMediumFormatgetCapabilities struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMediumFormat_getCapabilities"`

	This string `xml:"_this,omitempty"`
}

type IMediumFormatgetCapabilitiesResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMediumFormat_getCapabilitiesResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type IMediumFormatdescribeProperties struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMediumFormat_describeProperties"`

	This string `xml:"_this,omitempty"`
}

type IMediumFormatdescribePropertiesResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMediumFormat_describePropertiesResponse"`

	Names       []string    `xml:"names,omitempty"`
	Description []string    `xml:"description,omitempty"`
	Types       []*DataType `xml:"types,omitempty"`
	Flags       []uint32    `xml:"flags,omitempty"`
	Defaults    []string    `xml:"defaults,omitempty"`
}

type IKeyboardputScancode struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IKeyboard_putScancode"`

	This     string `xml:"_this,omitempty"`
	Scancode int32  `xml:"scancode,omitempty"`
}

type IKeyboardputScancodeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IKeyboard_putScancodeResponse"`
}

type IKeyboardputScancodes struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IKeyboard_putScancodes"`

	This      string  `xml:"_this,omitempty"`
	Scancodes []int32 `xml:"scancodes,omitempty"`
}

type IKeyboardputScancodesResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IKeyboard_putScancodesResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type IKeyboardputCAD struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IKeyboard_putCAD"`

	This string `xml:"_this,omitempty"`
}

type IKeyboardputCADResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IKeyboard_putCADResponse"`
}

type IMousegetAbsoluteSupported struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMouse_getAbsoluteSupported"`

	This string `xml:"_this,omitempty"`
}

type IMousegetAbsoluteSupportedResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMouse_getAbsoluteSupportedResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IMousegetRelativeSupported struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMouse_getRelativeSupported"`

	This string `xml:"_this,omitempty"`
}

type IMousegetRelativeSupportedResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMouse_getRelativeSupportedResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IMousegetNeedsHostCursor struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMouse_getNeedsHostCursor"`

	This string `xml:"_this,omitempty"`
}

type IMousegetNeedsHostCursorResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMouse_getNeedsHostCursorResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IMouseputMouseEvent struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMouse_putMouseEvent"`

	This        string `xml:"_this,omitempty"`
	Dx          int32  `xml:"dx,omitempty"`
	Dy          int32  `xml:"dy,omitempty"`
	Dz          int32  `xml:"dz,omitempty"`
	Dw          int32  `xml:"dw,omitempty"`
	ButtonState int32  `xml:"buttonState,omitempty"`
}

type IMouseputMouseEventResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMouse_putMouseEventResponse"`
}

type IMouseputMouseEventAbsolute struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMouse_putMouseEventAbsolute"`

	This        string `xml:"_this,omitempty"`
	X           int32  `xml:"x,omitempty"`
	Y           int32  `xml:"y,omitempty"`
	Dz          int32  `xml:"dz,omitempty"`
	Dw          int32  `xml:"dw,omitempty"`
	ButtonState int32  `xml:"buttonState,omitempty"`
}

type IMouseputMouseEventAbsoluteResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMouse_putMouseEventAbsoluteResponse"`
}

type IDisplaygetScreenResolution struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IDisplay_getScreenResolution"`

	This     string `xml:"_this,omitempty"`
	ScreenId uint32 `xml:"screenId,omitempty"`
}

type IDisplaygetScreenResolutionResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IDisplay_getScreenResolutionResponse"`

	Width        uint32 `xml:"width,omitempty"`
	Height       uint32 `xml:"height,omitempty"`
	BitsPerPixel uint32 `xml:"bitsPerPixel,omitempty"`
}

type IDisplaysetVideoModeHint struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IDisplay_setVideoModeHint"`

	This         string `xml:"_this,omitempty"`
	Width        uint32 `xml:"width,omitempty"`
	Height       uint32 `xml:"height,omitempty"`
	BitsPerPixel uint32 `xml:"bitsPerPixel,omitempty"`
	Display      uint32 `xml:"display,omitempty"`
}

type IDisplaysetVideoModeHintResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IDisplay_setVideoModeHintResponse"`
}

type IDisplaysetSeamlessMode struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IDisplay_setSeamlessMode"`

	This    string `xml:"_this,omitempty"`
	Enabled bool   `xml:"enabled,omitempty"`
}

type IDisplaysetSeamlessModeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IDisplay_setSeamlessModeResponse"`
}

type IDisplaytakeScreenShotToArray struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IDisplay_takeScreenShotToArray"`

	This     string `xml:"_this,omitempty"`
	ScreenId uint32 `xml:"screenId,omitempty"`
	Width    uint32 `xml:"width,omitempty"`
	Height   uint32 `xml:"height,omitempty"`
}

type IDisplaytakeScreenShotToArrayResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IDisplay_takeScreenShotToArrayResponse"`

	Returnval []byte `xml:"returnval,omitempty"`
}

type IDisplayinvalidateAndUpdate struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IDisplay_invalidateAndUpdate"`

	This string `xml:"_this,omitempty"`
}

type IDisplayinvalidateAndUpdateResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IDisplay_invalidateAndUpdateResponse"`
}

type IDisplayresizeCompleted struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IDisplay_resizeCompleted"`

	This     string `xml:"_this,omitempty"`
	ScreenId uint32 `xml:"screenId,omitempty"`
}

type IDisplayresizeCompletedResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IDisplay_resizeCompletedResponse"`
}

type INetworkAdaptergetAdapterType struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapter_getAdapterType"`

	This string `xml:"_this,omitempty"`
}

type INetworkAdaptergetAdapterTypeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapter_getAdapterTypeResponse"`

	Returnval *NetworkAdapterType `xml:"returnval,omitempty"`
}

type INetworkAdaptersetAdapterType struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapter_setAdapterType"`

	This        string              `xml:"_this,omitempty"`
	AdapterType *NetworkAdapterType `xml:"adapterType,omitempty"`
}

type INetworkAdaptersetAdapterTypeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapter_setAdapterTypeResponse"`
}

type INetworkAdaptergetSlot struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapter_getSlot"`

	This string `xml:"_this,omitempty"`
}

type INetworkAdaptergetSlotResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapter_getSlotResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type INetworkAdaptergetEnabled struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapter_getEnabled"`

	This string `xml:"_this,omitempty"`
}

type INetworkAdaptergetEnabledResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapter_getEnabledResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type INetworkAdaptersetEnabled struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapter_setEnabled"`

	This    string `xml:"_this,omitempty"`
	Enabled bool   `xml:"enabled,omitempty"`
}

type INetworkAdaptersetEnabledResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapter_setEnabledResponse"`
}

type INetworkAdaptergetMACAddress struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapter_getMACAddress"`

	This string `xml:"_this,omitempty"`
}

type INetworkAdaptergetMACAddressResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapter_getMACAddressResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type INetworkAdaptersetMACAddress struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapter_setMACAddress"`

	This       string `xml:"_this,omitempty"`
	MACAddress string `xml:"MACAddress,omitempty"`
}

type INetworkAdaptersetMACAddressResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapter_setMACAddressResponse"`
}

type INetworkAdaptergetAttachmentType struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapter_getAttachmentType"`

	This string `xml:"_this,omitempty"`
}

type INetworkAdaptergetAttachmentTypeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapter_getAttachmentTypeResponse"`

	Returnval *NetworkAttachmentType `xml:"returnval,omitempty"`
}

type INetworkAdaptergetHostInterface struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapter_getHostInterface"`

	This string `xml:"_this,omitempty"`
}

type INetworkAdaptergetHostInterfaceResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapter_getHostInterfaceResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type INetworkAdaptersetHostInterface struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapter_setHostInterface"`

	This          string `xml:"_this,omitempty"`
	HostInterface string `xml:"hostInterface,omitempty"`
}

type INetworkAdaptersetHostInterfaceResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapter_setHostInterfaceResponse"`
}

type INetworkAdaptergetInternalNetwork struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapter_getInternalNetwork"`

	This string `xml:"_this,omitempty"`
}

type INetworkAdaptergetInternalNetworkResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapter_getInternalNetworkResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type INetworkAdaptersetInternalNetwork struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapter_setInternalNetwork"`

	This            string `xml:"_this,omitempty"`
	InternalNetwork string `xml:"internalNetwork,omitempty"`
}

type INetworkAdaptersetInternalNetworkResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapter_setInternalNetworkResponse"`
}

type INetworkAdaptergetNATNetwork struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapter_getNATNetwork"`

	This string `xml:"_this,omitempty"`
}

type INetworkAdaptergetNATNetworkResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapter_getNATNetworkResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type INetworkAdaptersetNATNetwork struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapter_setNATNetwork"`

	This       string `xml:"_this,omitempty"`
	NATNetwork string `xml:"NATNetwork,omitempty"`
}

type INetworkAdaptersetNATNetworkResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapter_setNATNetworkResponse"`
}

type INetworkAdaptergetVDENetwork struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapter_getVDENetwork"`

	This string `xml:"_this,omitempty"`
}

type INetworkAdaptergetVDENetworkResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapter_getVDENetworkResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type INetworkAdaptersetVDENetwork struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapter_setVDENetwork"`

	This       string `xml:"_this,omitempty"`
	VDENetwork string `xml:"VDENetwork,omitempty"`
}

type INetworkAdaptersetVDENetworkResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapter_setVDENetworkResponse"`
}

type INetworkAdaptergetCableConnected struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapter_getCableConnected"`

	This string `xml:"_this,omitempty"`
}

type INetworkAdaptergetCableConnectedResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapter_getCableConnectedResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type INetworkAdaptersetCableConnected struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapter_setCableConnected"`

	This           string `xml:"_this,omitempty"`
	CableConnected bool   `xml:"cableConnected,omitempty"`
}

type INetworkAdaptersetCableConnectedResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapter_setCableConnectedResponse"`
}

type INetworkAdaptergetLineSpeed struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapter_getLineSpeed"`

	This string `xml:"_this,omitempty"`
}

type INetworkAdaptergetLineSpeedResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapter_getLineSpeedResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type INetworkAdaptersetLineSpeed struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapter_setLineSpeed"`

	This      string `xml:"_this,omitempty"`
	LineSpeed uint32 `xml:"lineSpeed,omitempty"`
}

type INetworkAdaptersetLineSpeedResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapter_setLineSpeedResponse"`
}

type INetworkAdaptergetTraceEnabled struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapter_getTraceEnabled"`

	This string `xml:"_this,omitempty"`
}

type INetworkAdaptergetTraceEnabledResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapter_getTraceEnabledResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type INetworkAdaptersetTraceEnabled struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapter_setTraceEnabled"`

	This         string `xml:"_this,omitempty"`
	TraceEnabled bool   `xml:"traceEnabled,omitempty"`
}

type INetworkAdaptersetTraceEnabledResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapter_setTraceEnabledResponse"`
}

type INetworkAdaptergetTraceFile struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapter_getTraceFile"`

	This string `xml:"_this,omitempty"`
}

type INetworkAdaptergetTraceFileResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapter_getTraceFileResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type INetworkAdaptersetTraceFile struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapter_setTraceFile"`

	This      string `xml:"_this,omitempty"`
	TraceFile string `xml:"traceFile,omitempty"`
}

type INetworkAdaptersetTraceFileResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapter_setTraceFileResponse"`
}

type INetworkAdaptergetNatDriver struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapter_getNatDriver"`

	This string `xml:"_this,omitempty"`
}

type INetworkAdaptergetNatDriverResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapter_getNatDriverResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type INetworkAdaptergetBootPriority struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapter_getBootPriority"`

	This string `xml:"_this,omitempty"`
}

type INetworkAdaptergetBootPriorityResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapter_getBootPriorityResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type INetworkAdaptersetBootPriority struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapter_setBootPriority"`

	This         string `xml:"_this,omitempty"`
	BootPriority uint32 `xml:"bootPriority,omitempty"`
}

type INetworkAdaptersetBootPriorityResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapter_setBootPriorityResponse"`
}

type INetworkAdapterattachToNAT struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapter_attachToNAT"`

	This string `xml:"_this,omitempty"`
}

type INetworkAdapterattachToNATResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapter_attachToNATResponse"`
}

type INetworkAdapterattachToBridgedInterface struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapter_attachToBridgedInterface"`

	This string `xml:"_this,omitempty"`
}

type INetworkAdapterattachToBridgedInterfaceResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapter_attachToBridgedInterfaceResponse"`
}

type INetworkAdapterattachToInternalNetwork struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapter_attachToInternalNetwork"`

	This string `xml:"_this,omitempty"`
}

type INetworkAdapterattachToInternalNetworkResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapter_attachToInternalNetworkResponse"`
}

type INetworkAdapterattachToHostOnlyInterface struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapter_attachToHostOnlyInterface"`

	This string `xml:"_this,omitempty"`
}

type INetworkAdapterattachToHostOnlyInterfaceResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapter_attachToHostOnlyInterfaceResponse"`
}

type INetworkAdapterattachToVDE struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapter_attachToVDE"`

	This string `xml:"_this,omitempty"`
}

type INetworkAdapterattachToVDEResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapter_attachToVDEResponse"`
}

type INetworkAdapterdetach struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapter_detach"`

	This string `xml:"_this,omitempty"`
}

type INetworkAdapterdetachResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapter_detachResponse"`
}

type ISerialPortgetSlot struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISerialPort_getSlot"`

	This string `xml:"_this,omitempty"`
}

type ISerialPortgetSlotResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISerialPort_getSlotResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type ISerialPortgetEnabled struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISerialPort_getEnabled"`

	This string `xml:"_this,omitempty"`
}

type ISerialPortgetEnabledResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISerialPort_getEnabledResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type ISerialPortsetEnabled struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISerialPort_setEnabled"`

	This    string `xml:"_this,omitempty"`
	Enabled bool   `xml:"enabled,omitempty"`
}

type ISerialPortsetEnabledResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISerialPort_setEnabledResponse"`
}

type ISerialPortgetIOBase struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISerialPort_getIOBase"`

	This string `xml:"_this,omitempty"`
}

type ISerialPortgetIOBaseResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISerialPort_getIOBaseResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type ISerialPortsetIOBase struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISerialPort_setIOBase"`

	This   string `xml:"_this,omitempty"`
	IOBase uint32 `xml:"IOBase,omitempty"`
}

type ISerialPortsetIOBaseResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISerialPort_setIOBaseResponse"`
}

type ISerialPortgetIRQ struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISerialPort_getIRQ"`

	This string `xml:"_this,omitempty"`
}

type ISerialPortgetIRQResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISerialPort_getIRQResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type ISerialPortsetIRQ struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISerialPort_setIRQ"`

	This string `xml:"_this,omitempty"`
	IRQ  uint32 `xml:"IRQ,omitempty"`
}

type ISerialPortsetIRQResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISerialPort_setIRQResponse"`
}

type ISerialPortgetHostMode struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISerialPort_getHostMode"`

	This string `xml:"_this,omitempty"`
}

type ISerialPortgetHostModeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISerialPort_getHostModeResponse"`

	Returnval *PortMode `xml:"returnval,omitempty"`
}

type ISerialPortsetHostMode struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISerialPort_setHostMode"`

	This     string    `xml:"_this,omitempty"`
	HostMode *PortMode `xml:"hostMode,omitempty"`
}

type ISerialPortsetHostModeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISerialPort_setHostModeResponse"`
}

type ISerialPortgetServer struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISerialPort_getServer"`

	This string `xml:"_this,omitempty"`
}

type ISerialPortgetServerResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISerialPort_getServerResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type ISerialPortsetServer struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISerialPort_setServer"`

	This   string `xml:"_this,omitempty"`
	Server bool   `xml:"server,omitempty"`
}

type ISerialPortsetServerResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISerialPort_setServerResponse"`
}

type ISerialPortgetPath struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISerialPort_getPath"`

	This string `xml:"_this,omitempty"`
}

type ISerialPortgetPathResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISerialPort_getPathResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type ISerialPortsetPath struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISerialPort_setPath"`

	This string `xml:"_this,omitempty"`
	Path string `xml:"path,omitempty"`
}

type ISerialPortsetPathResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISerialPort_setPathResponse"`
}

type IParallelPortgetSlot struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IParallelPort_getSlot"`

	This string `xml:"_this,omitempty"`
}

type IParallelPortgetSlotResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IParallelPort_getSlotResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type IParallelPortgetEnabled struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IParallelPort_getEnabled"`

	This string `xml:"_this,omitempty"`
}

type IParallelPortgetEnabledResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IParallelPort_getEnabledResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IParallelPortsetEnabled struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IParallelPort_setEnabled"`

	This    string `xml:"_this,omitempty"`
	Enabled bool   `xml:"enabled,omitempty"`
}

type IParallelPortsetEnabledResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IParallelPort_setEnabledResponse"`
}

type IParallelPortgetIOBase struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IParallelPort_getIOBase"`

	This string `xml:"_this,omitempty"`
}

type IParallelPortgetIOBaseResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IParallelPort_getIOBaseResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type IParallelPortsetIOBase struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IParallelPort_setIOBase"`

	This   string `xml:"_this,omitempty"`
	IOBase uint32 `xml:"IOBase,omitempty"`
}

type IParallelPortsetIOBaseResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IParallelPort_setIOBaseResponse"`
}

type IParallelPortgetIRQ struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IParallelPort_getIRQ"`

	This string `xml:"_this,omitempty"`
}

type IParallelPortgetIRQResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IParallelPort_getIRQResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type IParallelPortsetIRQ struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IParallelPort_setIRQ"`

	This string `xml:"_this,omitempty"`
	IRQ  uint32 `xml:"IRQ,omitempty"`
}

type IParallelPortsetIRQResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IParallelPort_setIRQResponse"`
}

type IParallelPortgetPath struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IParallelPort_getPath"`

	This string `xml:"_this,omitempty"`
}

type IParallelPortgetPathResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IParallelPort_getPathResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IParallelPortsetPath struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IParallelPort_setPath"`

	This string `xml:"_this,omitempty"`
	Path string `xml:"path,omitempty"`
}

type IParallelPortsetPathResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IParallelPort_setPathResponse"`
}

type IUSBControllergetEnabled struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBController_getEnabled"`

	This string `xml:"_this,omitempty"`
}

type IUSBControllergetEnabledResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBController_getEnabledResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IUSBControllersetEnabled struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBController_setEnabled"`

	This    string `xml:"_this,omitempty"`
	Enabled bool   `xml:"enabled,omitempty"`
}

type IUSBControllersetEnabledResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBController_setEnabledResponse"`
}

type IUSBControllergetEnabledEhci struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBController_getEnabledEhci"`

	This string `xml:"_this,omitempty"`
}

type IUSBControllergetEnabledEhciResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBController_getEnabledEhciResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IUSBControllersetEnabledEhci struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBController_setEnabledEhci"`

	This        string `xml:"_this,omitempty"`
	EnabledEhci bool   `xml:"enabledEhci,omitempty"`
}

type IUSBControllersetEnabledEhciResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBController_setEnabledEhciResponse"`
}

type IUSBControllergetProxyAvailable struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBController_getProxyAvailable"`

	This string `xml:"_this,omitempty"`
}

type IUSBControllergetProxyAvailableResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBController_getProxyAvailableResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IUSBControllergetUSBStandard struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBController_getUSBStandard"`

	This string `xml:"_this,omitempty"`
}

type IUSBControllergetUSBStandardResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBController_getUSBStandardResponse"`

	Returnval uint16 `xml:"returnval,omitempty"`
}

type IUSBControllergetDeviceFilters struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBController_getDeviceFilters"`

	This string `xml:"_this,omitempty"`
}

type IUSBControllergetDeviceFiltersResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBController_getDeviceFiltersResponse"`

	Returnval []string `xml:"returnval,omitempty"`
}

type IUSBControllercreateDeviceFilter struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBController_createDeviceFilter"`

	This string `xml:"_this,omitempty"`
	Name string `xml:"name,omitempty"`
}

type IUSBControllercreateDeviceFilterResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBController_createDeviceFilterResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IUSBControllerinsertDeviceFilter struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBController_insertDeviceFilter"`

	This     string `xml:"_this,omitempty"`
	Position uint32 `xml:"position,omitempty"`
	Filter   string `xml:"filter,omitempty"`
}

type IUSBControllerinsertDeviceFilterResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBController_insertDeviceFilterResponse"`
}

type IUSBControllerremoveDeviceFilter struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBController_removeDeviceFilter"`

	This     string `xml:"_this,omitempty"`
	Position uint32 `xml:"position,omitempty"`
}

type IUSBControllerremoveDeviceFilterResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBController_removeDeviceFilterResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IUSBDevicegetId struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDevice_getId"`

	This string `xml:"_this,omitempty"`
}

type IUSBDevicegetIdResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDevice_getIdResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IUSBDevicegetVendorId struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDevice_getVendorId"`

	This string `xml:"_this,omitempty"`
}

type IUSBDevicegetVendorIdResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDevice_getVendorIdResponse"`

	Returnval uint16 `xml:"returnval,omitempty"`
}

type IUSBDevicegetProductId struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDevice_getProductId"`

	This string `xml:"_this,omitempty"`
}

type IUSBDevicegetProductIdResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDevice_getProductIdResponse"`

	Returnval uint16 `xml:"returnval,omitempty"`
}

type IUSBDevicegetRevision struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDevice_getRevision"`

	This string `xml:"_this,omitempty"`
}

type IUSBDevicegetRevisionResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDevice_getRevisionResponse"`

	Returnval uint16 `xml:"returnval,omitempty"`
}

type IUSBDevicegetManufacturer struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDevice_getManufacturer"`

	This string `xml:"_this,omitempty"`
}

type IUSBDevicegetManufacturerResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDevice_getManufacturerResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IUSBDevicegetProduct struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDevice_getProduct"`

	This string `xml:"_this,omitempty"`
}

type IUSBDevicegetProductResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDevice_getProductResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IUSBDevicegetSerialNumber struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDevice_getSerialNumber"`

	This string `xml:"_this,omitempty"`
}

type IUSBDevicegetSerialNumberResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDevice_getSerialNumberResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IUSBDevicegetAddress struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDevice_getAddress"`

	This string `xml:"_this,omitempty"`
}

type IUSBDevicegetAddressResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDevice_getAddressResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IUSBDevicegetPort struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDevice_getPort"`

	This string `xml:"_this,omitempty"`
}

type IUSBDevicegetPortResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDevice_getPortResponse"`

	Returnval uint16 `xml:"returnval,omitempty"`
}

type IUSBDevicegetVersion struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDevice_getVersion"`

	This string `xml:"_this,omitempty"`
}

type IUSBDevicegetVersionResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDevice_getVersionResponse"`

	Returnval uint16 `xml:"returnval,omitempty"`
}

type IUSBDevicegetPortVersion struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDevice_getPortVersion"`

	This string `xml:"_this,omitempty"`
}

type IUSBDevicegetPortVersionResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDevice_getPortVersionResponse"`

	Returnval uint16 `xml:"returnval,omitempty"`
}

type IUSBDevicegetRemote struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDevice_getRemote"`

	This string `xml:"_this,omitempty"`
}

type IUSBDevicegetRemoteResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDevice_getRemoteResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IUSBDeviceFiltergetName struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDeviceFilter_getName"`

	This string `xml:"_this,omitempty"`
}

type IUSBDeviceFiltergetNameResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDeviceFilter_getNameResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IUSBDeviceFiltersetName struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDeviceFilter_setName"`

	This string `xml:"_this,omitempty"`
	Name string `xml:"name,omitempty"`
}

type IUSBDeviceFiltersetNameResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDeviceFilter_setNameResponse"`
}

type IUSBDeviceFiltergetActive struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDeviceFilter_getActive"`

	This string `xml:"_this,omitempty"`
}

type IUSBDeviceFiltergetActiveResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDeviceFilter_getActiveResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IUSBDeviceFiltersetActive struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDeviceFilter_setActive"`

	This   string `xml:"_this,omitempty"`
	Active bool   `xml:"active,omitempty"`
}

type IUSBDeviceFiltersetActiveResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDeviceFilter_setActiveResponse"`
}

type IUSBDeviceFiltergetVendorId struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDeviceFilter_getVendorId"`

	This string `xml:"_this,omitempty"`
}

type IUSBDeviceFiltergetVendorIdResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDeviceFilter_getVendorIdResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IUSBDeviceFiltersetVendorId struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDeviceFilter_setVendorId"`

	This     string `xml:"_this,omitempty"`
	VendorId string `xml:"vendorId,omitempty"`
}

type IUSBDeviceFiltersetVendorIdResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDeviceFilter_setVendorIdResponse"`
}

type IUSBDeviceFiltergetProductId struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDeviceFilter_getProductId"`

	This string `xml:"_this,omitempty"`
}

type IUSBDeviceFiltergetProductIdResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDeviceFilter_getProductIdResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IUSBDeviceFiltersetProductId struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDeviceFilter_setProductId"`

	This      string `xml:"_this,omitempty"`
	ProductId string `xml:"productId,omitempty"`
}

type IUSBDeviceFiltersetProductIdResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDeviceFilter_setProductIdResponse"`
}

type IUSBDeviceFiltergetRevision struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDeviceFilter_getRevision"`

	This string `xml:"_this,omitempty"`
}

type IUSBDeviceFiltergetRevisionResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDeviceFilter_getRevisionResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IUSBDeviceFiltersetRevision struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDeviceFilter_setRevision"`

	This     string `xml:"_this,omitempty"`
	Revision string `xml:"revision,omitempty"`
}

type IUSBDeviceFiltersetRevisionResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDeviceFilter_setRevisionResponse"`
}

type IUSBDeviceFiltergetManufacturer struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDeviceFilter_getManufacturer"`

	This string `xml:"_this,omitempty"`
}

type IUSBDeviceFiltergetManufacturerResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDeviceFilter_getManufacturerResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IUSBDeviceFiltersetManufacturer struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDeviceFilter_setManufacturer"`

	This         string `xml:"_this,omitempty"`
	Manufacturer string `xml:"manufacturer,omitempty"`
}

type IUSBDeviceFiltersetManufacturerResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDeviceFilter_setManufacturerResponse"`
}

type IUSBDeviceFiltergetProduct struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDeviceFilter_getProduct"`

	This string `xml:"_this,omitempty"`
}

type IUSBDeviceFiltergetProductResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDeviceFilter_getProductResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IUSBDeviceFiltersetProduct struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDeviceFilter_setProduct"`

	This    string `xml:"_this,omitempty"`
	Product string `xml:"product,omitempty"`
}

type IUSBDeviceFiltersetProductResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDeviceFilter_setProductResponse"`
}

type IUSBDeviceFiltergetSerialNumber struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDeviceFilter_getSerialNumber"`

	This string `xml:"_this,omitempty"`
}

type IUSBDeviceFiltergetSerialNumberResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDeviceFilter_getSerialNumberResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IUSBDeviceFiltersetSerialNumber struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDeviceFilter_setSerialNumber"`

	This         string `xml:"_this,omitempty"`
	SerialNumber string `xml:"serialNumber,omitempty"`
}

type IUSBDeviceFiltersetSerialNumberResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDeviceFilter_setSerialNumberResponse"`
}

type IUSBDeviceFiltergetPort struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDeviceFilter_getPort"`

	This string `xml:"_this,omitempty"`
}

type IUSBDeviceFiltergetPortResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDeviceFilter_getPortResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IUSBDeviceFiltersetPort struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDeviceFilter_setPort"`

	This string `xml:"_this,omitempty"`
	Port string `xml:"port,omitempty"`
}

type IUSBDeviceFiltersetPortResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDeviceFilter_setPortResponse"`
}

type IUSBDeviceFiltergetRemote struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDeviceFilter_getRemote"`

	This string `xml:"_this,omitempty"`
}

type IUSBDeviceFiltergetRemoteResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDeviceFilter_getRemoteResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IUSBDeviceFiltersetRemote struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDeviceFilter_setRemote"`

	This   string `xml:"_this,omitempty"`
	Remote string `xml:"remote,omitempty"`
}

type IUSBDeviceFiltersetRemoteResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDeviceFilter_setRemoteResponse"`
}

type IUSBDeviceFiltergetMaskedInterfaces struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDeviceFilter_getMaskedInterfaces"`

	This string `xml:"_this,omitempty"`
}

type IUSBDeviceFiltergetMaskedInterfacesResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDeviceFilter_getMaskedInterfacesResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type IUSBDeviceFiltersetMaskedInterfaces struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDeviceFilter_setMaskedInterfaces"`

	This             string `xml:"_this,omitempty"`
	MaskedInterfaces uint32 `xml:"maskedInterfaces,omitempty"`
}

type IUSBDeviceFiltersetMaskedInterfacesResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDeviceFilter_setMaskedInterfacesResponse"`
}

type IHostUSBDevicegetState struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHostUSBDevice_getState"`

	This string `xml:"_this,omitempty"`
}

type IHostUSBDevicegetStateResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHostUSBDevice_getStateResponse"`

	Returnval *USBDeviceState `xml:"returnval,omitempty"`
}

type IHostUSBDeviceFiltergetAction struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHostUSBDeviceFilter_getAction"`

	This string `xml:"_this,omitempty"`
}

type IHostUSBDeviceFiltergetActionResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHostUSBDeviceFilter_getActionResponse"`

	Returnval *USBDeviceFilterAction `xml:"returnval,omitempty"`
}

type IHostUSBDeviceFiltersetAction struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHostUSBDeviceFilter_setAction"`

	This   string                 `xml:"_this,omitempty"`
	Action *USBDeviceFilterAction `xml:"action,omitempty"`
}

type IHostUSBDeviceFiltersetActionResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHostUSBDeviceFilter_setActionResponse"`
}

type IAudioAdaptergetEnabled struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IAudioAdapter_getEnabled"`

	This string `xml:"_this,omitempty"`
}

type IAudioAdaptergetEnabledResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IAudioAdapter_getEnabledResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IAudioAdaptersetEnabled struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IAudioAdapter_setEnabled"`

	This    string `xml:"_this,omitempty"`
	Enabled bool   `xml:"enabled,omitempty"`
}

type IAudioAdaptersetEnabledResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IAudioAdapter_setEnabledResponse"`
}

type IAudioAdaptergetAudioController struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IAudioAdapter_getAudioController"`

	This string `xml:"_this,omitempty"`
}

type IAudioAdaptergetAudioControllerResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IAudioAdapter_getAudioControllerResponse"`

	Returnval *AudioControllerType `xml:"returnval,omitempty"`
}

type IAudioAdaptersetAudioController struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IAudioAdapter_setAudioController"`

	This            string               `xml:"_this,omitempty"`
	AudioController *AudioControllerType `xml:"audioController,omitempty"`
}

type IAudioAdaptersetAudioControllerResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IAudioAdapter_setAudioControllerResponse"`
}

type IAudioAdaptergetAudioDriver struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IAudioAdapter_getAudioDriver"`

	This string `xml:"_this,omitempty"`
}

type IAudioAdaptergetAudioDriverResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IAudioAdapter_getAudioDriverResponse"`

	Returnval *AudioDriverType `xml:"returnval,omitempty"`
}

type IAudioAdaptersetAudioDriver struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IAudioAdapter_setAudioDriver"`

	This        string           `xml:"_this,omitempty"`
	AudioDriver *AudioDriverType `xml:"audioDriver,omitempty"`
}

type IAudioAdaptersetAudioDriverResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IAudioAdapter_setAudioDriverResponse"`
}

type IVRDPServergetEnabled struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVRDPServer_getEnabled"`

	This string `xml:"_this,omitempty"`
}

type IVRDPServergetEnabledResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVRDPServer_getEnabledResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IVRDPServersetEnabled struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVRDPServer_setEnabled"`

	This    string `xml:"_this,omitempty"`
	Enabled bool   `xml:"enabled,omitempty"`
}

type IVRDPServersetEnabledResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVRDPServer_setEnabledResponse"`
}

type IVRDPServergetPorts struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVRDPServer_getPorts"`

	This string `xml:"_this,omitempty"`
}

type IVRDPServergetPortsResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVRDPServer_getPortsResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IVRDPServersetPorts struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVRDPServer_setPorts"`

	This  string `xml:"_this,omitempty"`
	Ports string `xml:"ports,omitempty"`
}

type IVRDPServersetPortsResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVRDPServer_setPortsResponse"`
}

type IVRDPServergetNetAddress struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVRDPServer_getNetAddress"`

	This string `xml:"_this,omitempty"`
}

type IVRDPServergetNetAddressResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVRDPServer_getNetAddressResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IVRDPServersetNetAddress struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVRDPServer_setNetAddress"`

	This       string `xml:"_this,omitempty"`
	NetAddress string `xml:"netAddress,omitempty"`
}

type IVRDPServersetNetAddressResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVRDPServer_setNetAddressResponse"`
}

type IVRDPServergetAuthType struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVRDPServer_getAuthType"`

	This string `xml:"_this,omitempty"`
}

type IVRDPServergetAuthTypeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVRDPServer_getAuthTypeResponse"`

	Returnval *VRDPAuthType `xml:"returnval,omitempty"`
}

type IVRDPServersetAuthType struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVRDPServer_setAuthType"`

	This     string        `xml:"_this,omitempty"`
	AuthType *VRDPAuthType `xml:"authType,omitempty"`
}

type IVRDPServersetAuthTypeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVRDPServer_setAuthTypeResponse"`
}

type IVRDPServergetAuthTimeout struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVRDPServer_getAuthTimeout"`

	This string `xml:"_this,omitempty"`
}

type IVRDPServergetAuthTimeoutResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVRDPServer_getAuthTimeoutResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type IVRDPServersetAuthTimeout struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVRDPServer_setAuthTimeout"`

	This        string `xml:"_this,omitempty"`
	AuthTimeout uint32 `xml:"authTimeout,omitempty"`
}

type IVRDPServersetAuthTimeoutResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVRDPServer_setAuthTimeoutResponse"`
}

type IVRDPServergetAllowMultiConnection struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVRDPServer_getAllowMultiConnection"`

	This string `xml:"_this,omitempty"`
}

type IVRDPServergetAllowMultiConnectionResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVRDPServer_getAllowMultiConnectionResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IVRDPServersetAllowMultiConnection struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVRDPServer_setAllowMultiConnection"`

	This                 string `xml:"_this,omitempty"`
	AllowMultiConnection bool   `xml:"allowMultiConnection,omitempty"`
}

type IVRDPServersetAllowMultiConnectionResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVRDPServer_setAllowMultiConnectionResponse"`
}

type IVRDPServergetReuseSingleConnection struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVRDPServer_getReuseSingleConnection"`

	This string `xml:"_this,omitempty"`
}

type IVRDPServergetReuseSingleConnectionResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVRDPServer_getReuseSingleConnectionResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IVRDPServersetReuseSingleConnection struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVRDPServer_setReuseSingleConnection"`

	This                  string `xml:"_this,omitempty"`
	ReuseSingleConnection bool   `xml:"reuseSingleConnection,omitempty"`
}

type IVRDPServersetReuseSingleConnectionResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVRDPServer_setReuseSingleConnectionResponse"`
}

type IVRDPServergetVideoChannel struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVRDPServer_getVideoChannel"`

	This string `xml:"_this,omitempty"`
}

type IVRDPServergetVideoChannelResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVRDPServer_getVideoChannelResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IVRDPServersetVideoChannel struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVRDPServer_setVideoChannel"`

	This         string `xml:"_this,omitempty"`
	VideoChannel bool   `xml:"videoChannel,omitempty"`
}

type IVRDPServersetVideoChannelResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVRDPServer_setVideoChannelResponse"`
}

type IVRDPServergetVideoChannelQuality struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVRDPServer_getVideoChannelQuality"`

	This string `xml:"_this,omitempty"`
}

type IVRDPServergetVideoChannelQualityResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVRDPServer_getVideoChannelQualityResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type IVRDPServersetVideoChannelQuality struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVRDPServer_setVideoChannelQuality"`

	This                string `xml:"_this,omitempty"`
	VideoChannelQuality uint32 `xml:"videoChannelQuality,omitempty"`
}

type IVRDPServersetVideoChannelQualityResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVRDPServer_setVideoChannelQualityResponse"`
}

type ISessiongetState struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISession_getState"`

	This string `xml:"_this,omitempty"`
}

type ISessiongetStateResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISession_getStateResponse"`

	Returnval *SessionState `xml:"returnval,omitempty"`
}

type ISessiongetType struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISession_getType"`

	This string `xml:"_this,omitempty"`
}

type ISessiongetTypeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISession_getTypeResponse"`

	Returnval *SessionType `xml:"returnval,omitempty"`
}

type ISessiongetMachine struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISession_getMachine"`

	This string `xml:"_this,omitempty"`
}

type ISessiongetMachineResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISession_getMachineResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type ISessiongetConsole struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISession_getConsole"`

	This string `xml:"_this,omitempty"`
}

type ISessiongetConsoleResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISession_getConsoleResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type ISessionclose struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISession_close"`

	This string `xml:"_this,omitempty"`
}

type ISessioncloseResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISession_closeResponse"`
}

type IStorageControllergetName struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IStorageController_getName"`

	This string `xml:"_this,omitempty"`
}

type IStorageControllergetNameResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IStorageController_getNameResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IStorageControllergetMaxDevicesPerPortCount struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IStorageController_getMaxDevicesPerPortCount"`

	This string `xml:"_this,omitempty"`
}

type IStorageControllergetMaxDevicesPerPortCountResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IStorageController_getMaxDevicesPerPortCountResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type IStorageControllergetMinPortCount struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IStorageController_getMinPortCount"`

	This string `xml:"_this,omitempty"`
}

type IStorageControllergetMinPortCountResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IStorageController_getMinPortCountResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type IStorageControllergetMaxPortCount struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IStorageController_getMaxPortCount"`

	This string `xml:"_this,omitempty"`
}

type IStorageControllergetMaxPortCountResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IStorageController_getMaxPortCountResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type IStorageControllergetInstance struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IStorageController_getInstance"`

	This string `xml:"_this,omitempty"`
}

type IStorageControllergetInstanceResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IStorageController_getInstanceResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type IStorageControllersetInstance struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IStorageController_setInstance"`

	This     string `xml:"_this,omitempty"`
	Instance uint32 `xml:"instance,omitempty"`
}

type IStorageControllersetInstanceResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IStorageController_setInstanceResponse"`
}

type IStorageControllergetPortCount struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IStorageController_getPortCount"`

	This string `xml:"_this,omitempty"`
}

type IStorageControllergetPortCountResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IStorageController_getPortCountResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type IStorageControllersetPortCount struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IStorageController_setPortCount"`

	This      string `xml:"_this,omitempty"`
	PortCount uint32 `xml:"portCount,omitempty"`
}

type IStorageControllersetPortCountResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IStorageController_setPortCountResponse"`
}

type IStorageControllergetBus struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IStorageController_getBus"`

	This string `xml:"_this,omitempty"`
}

type IStorageControllergetBusResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IStorageController_getBusResponse"`

	Returnval *StorageBus `xml:"returnval,omitempty"`
}

type IStorageControllergetControllerType struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IStorageController_getControllerType"`

	This string `xml:"_this,omitempty"`
}

type IStorageControllergetControllerTypeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IStorageController_getControllerTypeResponse"`

	Returnval *StorageControllerType `xml:"returnval,omitempty"`
}

type IStorageControllersetControllerType struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IStorageController_setControllerType"`

	This           string                 `xml:"_this,omitempty"`
	ControllerType *StorageControllerType `xml:"controllerType,omitempty"`
}

type IStorageControllersetControllerTypeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IStorageController_setControllerTypeResponse"`
}

type IStorageControllergetUseHostIOCache struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IStorageController_getUseHostIOCache"`

	This string `xml:"_this,omitempty"`
}

type IStorageControllergetUseHostIOCacheResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IStorageController_getUseHostIOCacheResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IStorageControllersetUseHostIOCache struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IStorageController_setUseHostIOCache"`

	This           string `xml:"_this,omitempty"`
	UseHostIOCache bool   `xml:"useHostIOCache,omitempty"`
}

type IStorageControllersetUseHostIOCacheResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IStorageController_setUseHostIOCacheResponse"`
}

type IStorageControllergetIDEEmulationPort struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IStorageController_getIDEEmulationPort"`

	This           string `xml:"_this,omitempty"`
	DevicePosition int32  `xml:"devicePosition,omitempty"`
}

type IStorageControllergetIDEEmulationPortResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IStorageController_getIDEEmulationPortResponse"`

	Returnval int32 `xml:"returnval,omitempty"`
}

type IStorageControllersetIDEEmulationPort struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IStorageController_setIDEEmulationPort"`

	This           string `xml:"_this,omitempty"`
	DevicePosition int32  `xml:"devicePosition,omitempty"`
	PortNumber     int32  `xml:"portNumber,omitempty"`
}

type IStorageControllersetIDEEmulationPortResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IStorageController_setIDEEmulationPortResponse"`
}

type IManagedObjectRefgetInterfaceName struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IManagedObjectRef_getInterfaceName"`

	This string `xml:"_this,omitempty"`
}

type IManagedObjectRefgetInterfaceNameResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IManagedObjectRef_getInterfaceNameResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IManagedObjectRefrelease struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IManagedObjectRef_release"`

	This string `xml:"_this,omitempty"`
}

type IManagedObjectRefreleaseResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IManagedObjectRef_releaseResponse"`
}

type IWebsessionManagerlogon struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IWebsessionManager_logon"`

	Username string `xml:"username,omitempty"`
	Password string `xml:"password,omitempty"`
}

type IWebsessionManagerlogonResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IWebsessionManager_logonResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IWebsessionManagergetSessionObject struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IWebsessionManager_getSessionObject"`

	RefIVirtualBox string `xml:"refIVirtualBox,omitempty"`
}

type IWebsessionManagergetSessionObjectResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IWebsessionManager_getSessionObjectResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IWebsessionManagerlogoff struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IWebsessionManager_logoff"`

	RefIVirtualBox string `xml:"refIVirtualBox,omitempty"`
}

type IWebsessionManagerlogoffResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IWebsessionManager_logoffResponse"`
}

type IPerformanceMetricgetMetricName struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IPerformanceMetric_getMetricName"`

	This string `xml:"_this,omitempty"`
}

type IPerformanceMetricgetMetricNameResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IPerformanceMetric_getMetricNameResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IPerformanceMetricgetObject struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IPerformanceMetric_getObject"`

	This string `xml:"_this,omitempty"`
}

type IPerformanceMetricgetObjectResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IPerformanceMetric_getObjectResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IPerformanceMetricgetDescription struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IPerformanceMetric_getDescription"`

	This string `xml:"_this,omitempty"`
}

type IPerformanceMetricgetDescriptionResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IPerformanceMetric_getDescriptionResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IPerformanceMetricgetPeriod struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IPerformanceMetric_getPeriod"`

	This string `xml:"_this,omitempty"`
}

type IPerformanceMetricgetPeriodResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IPerformanceMetric_getPeriodResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type IPerformanceMetricgetCount struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IPerformanceMetric_getCount"`

	This string `xml:"_this,omitempty"`
}

type IPerformanceMetricgetCountResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IPerformanceMetric_getCountResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type IPerformanceMetricgetUnit struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IPerformanceMetric_getUnit"`

	This string `xml:"_this,omitempty"`
}

type IPerformanceMetricgetUnitResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IPerformanceMetric_getUnitResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IPerformanceMetricgetMinimumValue struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IPerformanceMetric_getMinimumValue"`

	This string `xml:"_this,omitempty"`
}

type IPerformanceMetricgetMinimumValueResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IPerformanceMetric_getMinimumValueResponse"`

	Returnval int32 `xml:"returnval,omitempty"`
}

type IPerformanceMetricgetMaximumValue struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IPerformanceMetric_getMaximumValue"`

	This string `xml:"_this,omitempty"`
}

type IPerformanceMetricgetMaximumValueResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IPerformanceMetric_getMaximumValueResponse"`

	Returnval int32 `xml:"returnval,omitempty"`
}

type IPerformanceCollectorgetMetricNames struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IPerformanceCollector_getMetricNames"`

	This string `xml:"_this,omitempty"`
}

type IPerformanceCollectorgetMetricNamesResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IPerformanceCollector_getMetricNamesResponse"`

	Returnval []string `xml:"returnval,omitempty"`
}

type IPerformanceCollectorgetMetrics struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IPerformanceCollector_getMetrics"`

	This        string   `xml:"_this,omitempty"`
	MetricNames []string `xml:"metricNames,omitempty"`
	Objects     []string `xml:"objects,omitempty"`
}

type IPerformanceCollectorgetMetricsResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IPerformanceCollector_getMetricsResponse"`

	Returnval []string `xml:"returnval,omitempty"`
}

type IPerformanceCollectorsetupMetrics struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IPerformanceCollector_setupMetrics"`

	This        string   `xml:"_this,omitempty"`
	MetricNames []string `xml:"metricNames,omitempty"`
	Objects     []string `xml:"objects,omitempty"`
	Period      uint32   `xml:"period,omitempty"`
	Count       uint32   `xml:"count,omitempty"`
}

type IPerformanceCollectorsetupMetricsResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IPerformanceCollector_setupMetricsResponse"`

	Returnval []string `xml:"returnval,omitempty"`
}

type IPerformanceCollectorenableMetrics struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IPerformanceCollector_enableMetrics"`

	This        string   `xml:"_this,omitempty"`
	MetricNames []string `xml:"metricNames,omitempty"`
	Objects     []string `xml:"objects,omitempty"`
}

type IPerformanceCollectorenableMetricsResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IPerformanceCollector_enableMetricsResponse"`

	Returnval []string `xml:"returnval,omitempty"`
}

type IPerformanceCollectordisableMetrics struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IPerformanceCollector_disableMetrics"`

	This        string   `xml:"_this,omitempty"`
	MetricNames []string `xml:"metricNames,omitempty"`
	Objects     []string `xml:"objects,omitempty"`
}

type IPerformanceCollectordisableMetricsResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IPerformanceCollector_disableMetricsResponse"`

	Returnval []string `xml:"returnval,omitempty"`
}

type IPerformanceCollectorqueryMetricsData struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IPerformanceCollector_queryMetricsData"`

	This        string   `xml:"_this,omitempty"`
	MetricNames []string `xml:"metricNames,omitempty"`
	Objects     []string `xml:"objects,omitempty"`
}

type IPerformanceCollectorqueryMetricsDataResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IPerformanceCollector_queryMetricsDataResponse"`

	ReturnMetricNames     []string `xml:"returnMetricNames,omitempty"`
	ReturnObjects         []string `xml:"returnObjects,omitempty"`
	ReturnUnits           []string `xml:"returnUnits,omitempty"`
	ReturnScales          []uint32 `xml:"returnScales,omitempty"`
	ReturnSequenceNumbers []uint32 `xml:"returnSequenceNumbers,omitempty"`
	ReturnDataIndices     []uint32 `xml:"returnDataIndices,omitempty"`
	ReturnDataLengths     []uint32 `xml:"returnDataLengths,omitempty"`
	Returnval             []int32  `xml:"returnval,omitempty"`
}

type INATEnginegetNetwork struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATEngine_getNetwork"`

	This string `xml:"_this,omitempty"`
}

type INATEnginegetNetworkResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATEngine_getNetworkResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type INATEnginesetNetwork struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATEngine_setNetwork"`

	This    string `xml:"_this,omitempty"`
	Network string `xml:"network,omitempty"`
}

type INATEnginesetNetworkResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATEngine_setNetworkResponse"`
}

type INATEnginegetHostIP struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATEngine_getHostIP"`

	This string `xml:"_this,omitempty"`
}

type INATEnginegetHostIPResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATEngine_getHostIPResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type INATEnginesetHostIP struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATEngine_setHostIP"`

	This   string `xml:"_this,omitempty"`
	HostIP string `xml:"hostIP,omitempty"`
}

type INATEnginesetHostIPResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATEngine_setHostIPResponse"`
}

type INATEnginegetTftpPrefix struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATEngine_getTftpPrefix"`

	This string `xml:"_this,omitempty"`
}

type INATEnginegetTftpPrefixResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATEngine_getTftpPrefixResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type INATEnginesetTftpPrefix struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATEngine_setTftpPrefix"`

	This       string `xml:"_this,omitempty"`
	TftpPrefix string `xml:"tftpPrefix,omitempty"`
}

type INATEnginesetTftpPrefixResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATEngine_setTftpPrefixResponse"`
}

type INATEnginegetTftpBootFile struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATEngine_getTftpBootFile"`

	This string `xml:"_this,omitempty"`
}

type INATEnginegetTftpBootFileResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATEngine_getTftpBootFileResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type INATEnginesetTftpBootFile struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATEngine_setTftpBootFile"`

	This         string `xml:"_this,omitempty"`
	TftpBootFile string `xml:"tftpBootFile,omitempty"`
}

type INATEnginesetTftpBootFileResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATEngine_setTftpBootFileResponse"`
}

type INATEnginegetTftpNextServer struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATEngine_getTftpNextServer"`

	This string `xml:"_this,omitempty"`
}

type INATEnginegetTftpNextServerResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATEngine_getTftpNextServerResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type INATEnginesetTftpNextServer struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATEngine_setTftpNextServer"`

	This           string `xml:"_this,omitempty"`
	TftpNextServer string `xml:"tftpNextServer,omitempty"`
}

type INATEnginesetTftpNextServerResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATEngine_setTftpNextServerResponse"`
}

type INATEnginegetAliasMode struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATEngine_getAliasMode"`

	This string `xml:"_this,omitempty"`
}

type INATEnginegetAliasModeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATEngine_getAliasModeResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type INATEnginesetAliasMode struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATEngine_setAliasMode"`

	This      string `xml:"_this,omitempty"`
	AliasMode uint32 `xml:"aliasMode,omitempty"`
}

type INATEnginesetAliasModeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATEngine_setAliasModeResponse"`
}

type INATEnginegetDnsPassDomain struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATEngine_getDnsPassDomain"`

	This string `xml:"_this,omitempty"`
}

type INATEnginegetDnsPassDomainResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATEngine_getDnsPassDomainResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type INATEnginesetDnsPassDomain struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATEngine_setDnsPassDomain"`

	This          string `xml:"_this,omitempty"`
	DnsPassDomain bool   `xml:"dnsPassDomain,omitempty"`
}

type INATEnginesetDnsPassDomainResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATEngine_setDnsPassDomainResponse"`
}

type INATEnginegetDnsProxy struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATEngine_getDnsProxy"`

	This string `xml:"_this,omitempty"`
}

type INATEnginegetDnsProxyResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATEngine_getDnsProxyResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type INATEnginesetDnsProxy struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATEngine_setDnsProxy"`

	This     string `xml:"_this,omitempty"`
	DnsProxy bool   `xml:"dnsProxy,omitempty"`
}

type INATEnginesetDnsProxyResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATEngine_setDnsProxyResponse"`
}

type INATEnginegetDnsUseHostResolver struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATEngine_getDnsUseHostResolver"`

	This string `xml:"_this,omitempty"`
}

type INATEnginegetDnsUseHostResolverResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATEngine_getDnsUseHostResolverResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type INATEnginesetDnsUseHostResolver struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATEngine_setDnsUseHostResolver"`

	This               string `xml:"_this,omitempty"`
	DnsUseHostResolver bool   `xml:"dnsUseHostResolver,omitempty"`
}

type INATEnginesetDnsUseHostResolverResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATEngine_setDnsUseHostResolverResponse"`
}

type INATEnginegetRedirects struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATEngine_getRedirects"`

	This string `xml:"_this,omitempty"`
}

type INATEnginegetRedirectsResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATEngine_getRedirectsResponse"`

	Returnval []string `xml:"returnval,omitempty"`
}

type INATEnginesetNetworkSettings struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATEngine_setNetworkSettings"`

	This      string `xml:"_this,omitempty"`
	Mtu       uint32 `xml:"mtu,omitempty"`
	SockSnd   uint32 `xml:"sockSnd,omitempty"`
	SockRcv   uint32 `xml:"sockRcv,omitempty"`
	TcpWndSnd uint32 `xml:"TcpWndSnd,omitempty"`
	TcpWndRcv uint32 `xml:"TcpWndRcv,omitempty"`
}

type INATEnginesetNetworkSettingsResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATEngine_setNetworkSettingsResponse"`
}

type INATEnginegetNetworkSettings struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATEngine_getNetworkSettings"`

	This string `xml:"_this,omitempty"`
}

type INATEnginegetNetworkSettingsResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATEngine_getNetworkSettingsResponse"`

	Mtu       uint32 `xml:"mtu,omitempty"`
	SockSnd   uint32 `xml:"sockSnd,omitempty"`
	SockRcv   uint32 `xml:"sockRcv,omitempty"`
	TcpWndSnd uint32 `xml:"TcpWndSnd,omitempty"`
	TcpWndRcv uint32 `xml:"TcpWndRcv,omitempty"`
}

type INATEngineaddRedirect struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATEngine_addRedirect"`

	This      string       `xml:"_this,omitempty"`
	Name      string       `xml:"name,omitempty"`
	Proto     *NATProtocol `xml:"proto,omitempty"`
	HostIp    string       `xml:"hostIp,omitempty"`
	HostPort  uint16       `xml:"hostPort,omitempty"`
	GuestIp   string       `xml:"guestIp,omitempty"`
	GuestPort uint16       `xml:"guestPort,omitempty"`
}

type INATEngineaddRedirectResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATEngine_addRedirectResponse"`
}

type INATEngineremoveRedirect struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATEngine_removeRedirect"`

	This string `xml:"_this,omitempty"`
	Name string `xml:"name,omitempty"`
}

type INATEngineremoveRedirectResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATEngine_removeRedirectResponse"`
}

type InvalidObjectFault struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ InvalidObjectFault"`

	BadObjectID string `xml:"badObjectID,omitempty"`
}

type RuntimeFault struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ RuntimeFault"`

	ResultCode  int32  `xml:"resultCode,omitempty"`
	InterfaceID string `xml:"interfaceID,omitempty"`
	Component   string `xml:"component,omitempty"`
	Text        string `xml:"text,omitempty"`
}

type IRemoteDisplayInfo struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IRemoteDisplayInfo"`

	Active             bool   `xml:"active,omitempty"`
	Port               int32  `xml:"port,omitempty"`
	NumberOfClients    uint32 `xml:"numberOfClients,omitempty"`
	BeginTime          int64  `xml:"beginTime,omitempty"`
	EndTime            int64  `xml:"endTime,omitempty"`
	BytesSent          uint64 `xml:"bytesSent,omitempty"`
	BytesSentTotal     uint64 `xml:"bytesSentTotal,omitempty"`
	BytesReceived      uint64 `xml:"bytesReceived,omitempty"`
	BytesReceivedTotal uint64 `xml:"bytesReceivedTotal,omitempty"`
	User               string `xml:"user,omitempty"`
	Domain             string `xml:"domain,omitempty"`
	ClientName         string `xml:"clientName,omitempty"`
	ClientIP           string `xml:"clientIP,omitempty"`
	ClientVersion      uint32 `xml:"clientVersion,omitempty"`
	EncryptionStyle    uint32 `xml:"encryptionStyle,omitempty"`
}

type IGuestOSType struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestOSType"`

	FamilyId                        string                 `xml:"familyId,omitempty"`
	FamilyDescription               string                 `xml:"familyDescription,omitempty"`
	Id                              string                 `xml:"id,omitempty"`
	Description                     string                 `xml:"description,omitempty"`
	Is64Bit                         bool                   `xml:"is64Bit,omitempty"`
	RecommendedIOAPIC               bool                   `xml:"recommendedIOAPIC,omitempty"`
	RecommendedVirtEx               bool                   `xml:"recommendedVirtEx,omitempty"`
	RecommendedRAM                  uint32                 `xml:"recommendedRAM,omitempty"`
	RecommendedVRAM                 uint32                 `xml:"recommendedVRAM,omitempty"`
	RecommendedHDD                  uint32                 `xml:"recommendedHDD,omitempty"`
	AdapterType                     *NetworkAdapterType    `xml:"adapterType,omitempty"`
	RecommendedPae                  bool                   `xml:"recommendedPae,omitempty"`
	RecommendedDvdStorageController *StorageControllerType `xml:"recommendedDvdStorageController,omitempty"`
	RecommendedDvdStorageBus        *StorageBus            `xml:"recommendedDvdStorageBus,omitempty"`
	RecommendedHdStorageController  *StorageControllerType `xml:"recommendedHdStorageController,omitempty"`
	RecommendedHdStorageBus         *StorageBus            `xml:"recommendedHdStorageBus,omitempty"`
	RecommendedFirmware             *FirmwareType          `xml:"recommendedFirmware,omitempty"`
	RecommendedUsbHid               bool                   `xml:"recommendedUsbHid,omitempty"`
	RecommendedHpet                 bool                   `xml:"recommendedHpet,omitempty"`
	RecommendedUsbTablet            bool                   `xml:"recommendedUsbTablet,omitempty"`
	RecommendedRtcUseUtc            bool                   `xml:"recommendedRtcUseUtc,omitempty"`
}

type IMediumAttachment struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMediumAttachment"`

	Medium      string      `xml:"medium,omitempty"`
	Controller  string      `xml:"controller,omitempty"`
	Port        int32       `xml:"port,omitempty"`
	Device      int32       `xml:"device,omitempty"`
	Type_       *DeviceType `xml:"type,omitempty"`
	Passthrough bool        `xml:"passthrough,omitempty"`
}

type ISharedFolder struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISharedFolder"`

	Name            string `xml:"name,omitempty"`
	HostPath        string `xml:"hostPath,omitempty"`
	Accessible      bool   `xml:"accessible,omitempty"`
	Writable        bool   `xml:"writable,omitempty"`
	LastAccessError string `xml:"lastAccessError,omitempty"`
}

type VboxPortType struct {
	client *SOAPClient
}

func NewVboxPortType(url string, tls bool, auth *BasicAuth) *VboxPortType {
	if url == "" {
		url = ""
	}
	client := NewSOAPClient(url, tls, auth)

	return &VboxPortType{
		client: client,
	}
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxErrorInfogetResultCode(request *IVirtualBoxErrorInfogetResultCode) (*IVirtualBoxErrorInfogetResultCodeResponse, error) {
	response := new(IVirtualBoxErrorInfogetResultCodeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxErrorInfogetInterfaceID(request *IVirtualBoxErrorInfogetInterfaceID) (*IVirtualBoxErrorInfogetInterfaceIDResponse, error) {
	response := new(IVirtualBoxErrorInfogetInterfaceIDResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxErrorInfogetComponent(request *IVirtualBoxErrorInfogetComponent) (*IVirtualBoxErrorInfogetComponentResponse, error) {
	response := new(IVirtualBoxErrorInfogetComponentResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxErrorInfogetText(request *IVirtualBoxErrorInfogetText) (*IVirtualBoxErrorInfogetTextResponse, error) {
	response := new(IVirtualBoxErrorInfogetTextResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxErrorInfogetNext(request *IVirtualBoxErrorInfogetNext) (*IVirtualBoxErrorInfogetNextResponse, error) {
	response := new(IVirtualBoxErrorInfogetNextResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IDHCPServergetEnabled(request *IDHCPServergetEnabled) (*IDHCPServergetEnabledResponse, error) {
	response := new(IDHCPServergetEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IDHCPServersetEnabled(request *IDHCPServersetEnabled) (*IDHCPServersetEnabledResponse, error) {
	response := new(IDHCPServersetEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IDHCPServergetIPAddress(request *IDHCPServergetIPAddress) (*IDHCPServergetIPAddressResponse, error) {
	response := new(IDHCPServergetIPAddressResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IDHCPServergetNetworkMask(request *IDHCPServergetNetworkMask) (*IDHCPServergetNetworkMaskResponse, error) {
	response := new(IDHCPServergetNetworkMaskResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IDHCPServergetNetworkName(request *IDHCPServergetNetworkName) (*IDHCPServergetNetworkNameResponse, error) {
	response := new(IDHCPServergetNetworkNameResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IDHCPServergetLowerIP(request *IDHCPServergetLowerIP) (*IDHCPServergetLowerIPResponse, error) {
	response := new(IDHCPServergetLowerIPResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IDHCPServergetUpperIP(request *IDHCPServergetUpperIP) (*IDHCPServergetUpperIPResponse, error) {
	response := new(IDHCPServergetUpperIPResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IDHCPServersetConfiguration(request *IDHCPServersetConfiguration) (*IDHCPServersetConfigurationResponse, error) {
	response := new(IDHCPServersetConfigurationResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IDHCPServerstart(request *IDHCPServerstart) (*IDHCPServerstartResponse, error) {
	response := new(IDHCPServerstartResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IDHCPServerstop(request *IDHCPServerstop) (*IDHCPServerstopResponse, error) {
	response := new(IDHCPServerstopResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxgetVersion(request *IVirtualBoxgetVersion) (*IVirtualBoxgetVersionResponse, error) {
	response := new(IVirtualBoxgetVersionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxgetRevision(request *IVirtualBoxgetRevision) (*IVirtualBoxgetRevisionResponse, error) {
	response := new(IVirtualBoxgetRevisionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxgetPackageType(request *IVirtualBoxgetPackageType) (*IVirtualBoxgetPackageTypeResponse, error) {
	response := new(IVirtualBoxgetPackageTypeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxgetHomeFolder(request *IVirtualBoxgetHomeFolder) (*IVirtualBoxgetHomeFolderResponse, error) {
	response := new(IVirtualBoxgetHomeFolderResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxgetSettingsFilePath(request *IVirtualBoxgetSettingsFilePath) (*IVirtualBoxgetSettingsFilePathResponse, error) {
	response := new(IVirtualBoxgetSettingsFilePathResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxgetHost(request *IVirtualBoxgetHost) (*IVirtualBoxgetHostResponse, error) {
	response := new(IVirtualBoxgetHostResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxgetSystemProperties(request *IVirtualBoxgetSystemProperties) (*IVirtualBoxgetSystemPropertiesResponse, error) {
	response := new(IVirtualBoxgetSystemPropertiesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxgetMachines(request *IVirtualBoxgetMachines) (*IVirtualBoxgetMachinesResponse, error) {
	response := new(IVirtualBoxgetMachinesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxgetHardDisks(request *IVirtualBoxgetHardDisks) (*IVirtualBoxgetHardDisksResponse, error) {
	response := new(IVirtualBoxgetHardDisksResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxgetDVDImages(request *IVirtualBoxgetDVDImages) (*IVirtualBoxgetDVDImagesResponse, error) {
	response := new(IVirtualBoxgetDVDImagesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxgetFloppyImages(request *IVirtualBoxgetFloppyImages) (*IVirtualBoxgetFloppyImagesResponse, error) {
	response := new(IVirtualBoxgetFloppyImagesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxgetProgressOperations(request *IVirtualBoxgetProgressOperations) (*IVirtualBoxgetProgressOperationsResponse, error) {
	response := new(IVirtualBoxgetProgressOperationsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxgetGuestOSTypes(request *IVirtualBoxgetGuestOSTypes) (*IVirtualBoxgetGuestOSTypesResponse, error) {
	response := new(IVirtualBoxgetGuestOSTypesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxgetSharedFolders(request *IVirtualBoxgetSharedFolders) (*IVirtualBoxgetSharedFoldersResponse, error) {
	response := new(IVirtualBoxgetSharedFoldersResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxgetPerformanceCollector(request *IVirtualBoxgetPerformanceCollector) (*IVirtualBoxgetPerformanceCollectorResponse, error) {
	response := new(IVirtualBoxgetPerformanceCollectorResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxgetDHCPServers(request *IVirtualBoxgetDHCPServers) (*IVirtualBoxgetDHCPServersResponse, error) {
	response := new(IVirtualBoxgetDHCPServersResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxcreateMachine(request *IVirtualBoxcreateMachine) (*IVirtualBoxcreateMachineResponse, error) {
	response := new(IVirtualBoxcreateMachineResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxcreateLegacyMachine(request *IVirtualBoxcreateLegacyMachine) (*IVirtualBoxcreateLegacyMachineResponse, error) {
	response := new(IVirtualBoxcreateLegacyMachineResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxopenMachine(request *IVirtualBoxopenMachine) (*IVirtualBoxopenMachineResponse, error) {
	response := new(IVirtualBoxopenMachineResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxregisterMachine(request *IVirtualBoxregisterMachine) (*IVirtualBoxregisterMachineResponse, error) {
	response := new(IVirtualBoxregisterMachineResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxgetMachine(request *IVirtualBoxgetMachine) (*IVirtualBoxgetMachineResponse, error) {
	response := new(IVirtualBoxgetMachineResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxfindMachine(request *IVirtualBoxfindMachine) (*IVirtualBoxfindMachineResponse, error) {
	response := new(IVirtualBoxfindMachineResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxunregisterMachine(request *IVirtualBoxunregisterMachine) (*IVirtualBoxunregisterMachineResponse, error) {
	response := new(IVirtualBoxunregisterMachineResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxcreateAppliance(request *IVirtualBoxcreateAppliance) (*IVirtualBoxcreateApplianceResponse, error) {
	response := new(IVirtualBoxcreateApplianceResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxcreateHardDisk(request *IVirtualBoxcreateHardDisk) (*IVirtualBoxcreateHardDiskResponse, error) {
	response := new(IVirtualBoxcreateHardDiskResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxopenHardDisk(request *IVirtualBoxopenHardDisk) (*IVirtualBoxopenHardDiskResponse, error) {
	response := new(IVirtualBoxopenHardDiskResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxgetHardDisk(request *IVirtualBoxgetHardDisk) (*IVirtualBoxgetHardDiskResponse, error) {
	response := new(IVirtualBoxgetHardDiskResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxfindHardDisk(request *IVirtualBoxfindHardDisk) (*IVirtualBoxfindHardDiskResponse, error) {
	response := new(IVirtualBoxfindHardDiskResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxopenDVDImage(request *IVirtualBoxopenDVDImage) (*IVirtualBoxopenDVDImageResponse, error) {
	response := new(IVirtualBoxopenDVDImageResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxgetDVDImage(request *IVirtualBoxgetDVDImage) (*IVirtualBoxgetDVDImageResponse, error) {
	response := new(IVirtualBoxgetDVDImageResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxfindDVDImage(request *IVirtualBoxfindDVDImage) (*IVirtualBoxfindDVDImageResponse, error) {
	response := new(IVirtualBoxfindDVDImageResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxopenFloppyImage(request *IVirtualBoxopenFloppyImage) (*IVirtualBoxopenFloppyImageResponse, error) {
	response := new(IVirtualBoxopenFloppyImageResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxgetFloppyImage(request *IVirtualBoxgetFloppyImage) (*IVirtualBoxgetFloppyImageResponse, error) {
	response := new(IVirtualBoxgetFloppyImageResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxfindFloppyImage(request *IVirtualBoxfindFloppyImage) (*IVirtualBoxfindFloppyImageResponse, error) {
	response := new(IVirtualBoxfindFloppyImageResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxgetGuestOSType(request *IVirtualBoxgetGuestOSType) (*IVirtualBoxgetGuestOSTypeResponse, error) {
	response := new(IVirtualBoxgetGuestOSTypeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxcreateSharedFolder(request *IVirtualBoxcreateSharedFolder) (*IVirtualBoxcreateSharedFolderResponse, error) {
	response := new(IVirtualBoxcreateSharedFolderResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxremoveSharedFolder(request *IVirtualBoxremoveSharedFolder) (*IVirtualBoxremoveSharedFolderResponse, error) {
	response := new(IVirtualBoxremoveSharedFolderResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxgetExtraDataKeys(request *IVirtualBoxgetExtraDataKeys) (*IVirtualBoxgetExtraDataKeysResponse, error) {
	response := new(IVirtualBoxgetExtraDataKeysResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxgetExtraData(request *IVirtualBoxgetExtraData) (*IVirtualBoxgetExtraDataResponse, error) {
	response := new(IVirtualBoxgetExtraDataResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxsetExtraData(request *IVirtualBoxsetExtraData) (*IVirtualBoxsetExtraDataResponse, error) {
	response := new(IVirtualBoxsetExtraDataResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxopenSession(request *IVirtualBoxopenSession) (*IVirtualBoxopenSessionResponse, error) {
	response := new(IVirtualBoxopenSessionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxopenRemoteSession(request *IVirtualBoxopenRemoteSession) (*IVirtualBoxopenRemoteSessionResponse, error) {
	response := new(IVirtualBoxopenRemoteSessionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxopenExistingSession(request *IVirtualBoxopenExistingSession) (*IVirtualBoxopenExistingSessionResponse, error) {
	response := new(IVirtualBoxopenExistingSessionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxwaitForPropertyChange(request *IVirtualBoxwaitForPropertyChange) (*IVirtualBoxwaitForPropertyChangeResponse, error) {
	response := new(IVirtualBoxwaitForPropertyChangeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxcreateDHCPServer(request *IVirtualBoxcreateDHCPServer) (*IVirtualBoxcreateDHCPServerResponse, error) {
	response := new(IVirtualBoxcreateDHCPServerResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxfindDHCPServerByNetworkName(request *IVirtualBoxfindDHCPServerByNetworkName) (*IVirtualBoxfindDHCPServerByNetworkNameResponse, error) {
	response := new(IVirtualBoxfindDHCPServerByNetworkNameResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxremoveDHCPServer(request *IVirtualBoxremoveDHCPServer) (*IVirtualBoxremoveDHCPServerResponse, error) {
	response := new(IVirtualBoxremoveDHCPServerResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxcheckFirmwarePresent(request *IVirtualBoxcheckFirmwarePresent) (*IVirtualBoxcheckFirmwarePresentResponse, error) {
	response := new(IVirtualBoxcheckFirmwarePresentResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVFSExplorergetPath(request *IVFSExplorergetPath) (*IVFSExplorergetPathResponse, error) {
	response := new(IVFSExplorergetPathResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVFSExplorergetType(request *IVFSExplorergetType) (*IVFSExplorergetTypeResponse, error) {
	response := new(IVFSExplorergetTypeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVFSExplorerupdate(request *IVFSExplorerupdate) (*IVFSExplorerupdateResponse, error) {
	response := new(IVFSExplorerupdateResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVFSExplorercd(request *IVFSExplorercd) (*IVFSExplorercdResponse, error) {
	response := new(IVFSExplorercdResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVFSExplorercdUp(request *IVFSExplorercdUp) (*IVFSExplorercdUpResponse, error) {
	response := new(IVFSExplorercdUpResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVFSExplorerentryList(request *IVFSExplorerentryList) (*IVFSExplorerentryListResponse, error) {
	response := new(IVFSExplorerentryListResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVFSExplorerexists(request *IVFSExplorerexists) (*IVFSExplorerexistsResponse, error) {
	response := new(IVFSExplorerexistsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVFSExplorerremove(request *IVFSExplorerremove) (*IVFSExplorerremoveResponse, error) {
	response := new(IVFSExplorerremoveResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IAppliancegetPath(request *IAppliancegetPath) (*IAppliancegetPathResponse, error) {
	response := new(IAppliancegetPathResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IAppliancegetDisks(request *IAppliancegetDisks) (*IAppliancegetDisksResponse, error) {
	response := new(IAppliancegetDisksResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IAppliancegetVirtualSystemDescriptions(request *IAppliancegetVirtualSystemDescriptions) (*IAppliancegetVirtualSystemDescriptionsResponse, error) {
	response := new(IAppliancegetVirtualSystemDescriptionsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IApplianceread(request *IApplianceread) (*IAppliancereadResponse, error) {
	response := new(IAppliancereadResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IApplianceinterpret(request *IApplianceinterpret) (*IApplianceinterpretResponse, error) {
	response := new(IApplianceinterpretResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IApplianceimportMachines(request *IApplianceimportMachines) (*IApplianceimportMachinesResponse, error) {
	response := new(IApplianceimportMachinesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IAppliancecreateVFSExplorer(request *IAppliancecreateVFSExplorer) (*IAppliancecreateVFSExplorerResponse, error) {
	response := new(IAppliancecreateVFSExplorerResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IAppliancewrite(request *IAppliancewrite) (*IAppliancewriteResponse, error) {
	response := new(IAppliancewriteResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IAppliancegetWarnings(request *IAppliancegetWarnings) (*IAppliancegetWarningsResponse, error) {
	response := new(IAppliancegetWarningsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualSystemDescriptiongetCount(request *IVirtualSystemDescriptiongetCount) (*IVirtualSystemDescriptiongetCountResponse, error) {
	response := new(IVirtualSystemDescriptiongetCountResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualSystemDescriptiongetDescription(request *IVirtualSystemDescriptiongetDescription) (*IVirtualSystemDescriptiongetDescriptionResponse, error) {
	response := new(IVirtualSystemDescriptiongetDescriptionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualSystemDescriptiongetDescriptionByType(request *IVirtualSystemDescriptiongetDescriptionByType) (*IVirtualSystemDescriptiongetDescriptionByTypeResponse, error) {
	response := new(IVirtualSystemDescriptiongetDescriptionByTypeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualSystemDescriptiongetValuesByType(request *IVirtualSystemDescriptiongetValuesByType) (*IVirtualSystemDescriptiongetValuesByTypeResponse, error) {
	response := new(IVirtualSystemDescriptiongetValuesByTypeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualSystemDescriptionsetFinalValues(request *IVirtualSystemDescriptionsetFinalValues) (*IVirtualSystemDescriptionsetFinalValuesResponse, error) {
	response := new(IVirtualSystemDescriptionsetFinalValuesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualSystemDescriptionaddDescription(request *IVirtualSystemDescriptionaddDescription) (*IVirtualSystemDescriptionaddDescriptionResponse, error) {
	response := new(IVirtualSystemDescriptionaddDescriptionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IBIOSSettingsgetLogoFadeIn(request *IBIOSSettingsgetLogoFadeIn) (*IBIOSSettingsgetLogoFadeInResponse, error) {
	response := new(IBIOSSettingsgetLogoFadeInResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IBIOSSettingssetLogoFadeIn(request *IBIOSSettingssetLogoFadeIn) (*IBIOSSettingssetLogoFadeInResponse, error) {
	response := new(IBIOSSettingssetLogoFadeInResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IBIOSSettingsgetLogoFadeOut(request *IBIOSSettingsgetLogoFadeOut) (*IBIOSSettingsgetLogoFadeOutResponse, error) {
	response := new(IBIOSSettingsgetLogoFadeOutResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IBIOSSettingssetLogoFadeOut(request *IBIOSSettingssetLogoFadeOut) (*IBIOSSettingssetLogoFadeOutResponse, error) {
	response := new(IBIOSSettingssetLogoFadeOutResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IBIOSSettingsgetLogoDisplayTime(request *IBIOSSettingsgetLogoDisplayTime) (*IBIOSSettingsgetLogoDisplayTimeResponse, error) {
	response := new(IBIOSSettingsgetLogoDisplayTimeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IBIOSSettingssetLogoDisplayTime(request *IBIOSSettingssetLogoDisplayTime) (*IBIOSSettingssetLogoDisplayTimeResponse, error) {
	response := new(IBIOSSettingssetLogoDisplayTimeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IBIOSSettingsgetLogoImagePath(request *IBIOSSettingsgetLogoImagePath) (*IBIOSSettingsgetLogoImagePathResponse, error) {
	response := new(IBIOSSettingsgetLogoImagePathResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IBIOSSettingssetLogoImagePath(request *IBIOSSettingssetLogoImagePath) (*IBIOSSettingssetLogoImagePathResponse, error) {
	response := new(IBIOSSettingssetLogoImagePathResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IBIOSSettingsgetBootMenuMode(request *IBIOSSettingsgetBootMenuMode) (*IBIOSSettingsgetBootMenuModeResponse, error) {
	response := new(IBIOSSettingsgetBootMenuModeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IBIOSSettingssetBootMenuMode(request *IBIOSSettingssetBootMenuMode) (*IBIOSSettingssetBootMenuModeResponse, error) {
	response := new(IBIOSSettingssetBootMenuModeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IBIOSSettingsgetACPIEnabled(request *IBIOSSettingsgetACPIEnabled) (*IBIOSSettingsgetACPIEnabledResponse, error) {
	response := new(IBIOSSettingsgetACPIEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IBIOSSettingssetACPIEnabled(request *IBIOSSettingssetACPIEnabled) (*IBIOSSettingssetACPIEnabledResponse, error) {
	response := new(IBIOSSettingssetACPIEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IBIOSSettingsgetIOAPICEnabled(request *IBIOSSettingsgetIOAPICEnabled) (*IBIOSSettingsgetIOAPICEnabledResponse, error) {
	response := new(IBIOSSettingsgetIOAPICEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IBIOSSettingssetIOAPICEnabled(request *IBIOSSettingssetIOAPICEnabled) (*IBIOSSettingssetIOAPICEnabledResponse, error) {
	response := new(IBIOSSettingssetIOAPICEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IBIOSSettingsgetTimeOffset(request *IBIOSSettingsgetTimeOffset) (*IBIOSSettingsgetTimeOffsetResponse, error) {
	response := new(IBIOSSettingsgetTimeOffsetResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IBIOSSettingssetTimeOffset(request *IBIOSSettingssetTimeOffset) (*IBIOSSettingssetTimeOffsetResponse, error) {
	response := new(IBIOSSettingssetTimeOffsetResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IBIOSSettingsgetPXEDebugEnabled(request *IBIOSSettingsgetPXEDebugEnabled) (*IBIOSSettingsgetPXEDebugEnabledResponse, error) {
	response := new(IBIOSSettingsgetPXEDebugEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IBIOSSettingssetPXEDebugEnabled(request *IBIOSSettingssetPXEDebugEnabled) (*IBIOSSettingssetPXEDebugEnabledResponse, error) {
	response := new(IBIOSSettingssetPXEDebugEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetParent(request *IMachinegetParent) (*IMachinegetParentResponse, error) {
	response := new(IMachinegetParentResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetAccessible(request *IMachinegetAccessible) (*IMachinegetAccessibleResponse, error) {
	response := new(IMachinegetAccessibleResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetAccessError(request *IMachinegetAccessError) (*IMachinegetAccessErrorResponse, error) {
	response := new(IMachinegetAccessErrorResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetName(request *IMachinegetName) (*IMachinegetNameResponse, error) {
	response := new(IMachinegetNameResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetName(request *IMachinesetName) (*IMachinesetNameResponse, error) {
	response := new(IMachinesetNameResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetDescription(request *IMachinegetDescription) (*IMachinegetDescriptionResponse, error) {
	response := new(IMachinegetDescriptionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetDescription(request *IMachinesetDescription) (*IMachinesetDescriptionResponse, error) {
	response := new(IMachinesetDescriptionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetId(request *IMachinegetId) (*IMachinegetIdResponse, error) {
	response := new(IMachinegetIdResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetOSTypeId(request *IMachinegetOSTypeId) (*IMachinegetOSTypeIdResponse, error) {
	response := new(IMachinegetOSTypeIdResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetOSTypeId(request *IMachinesetOSTypeId) (*IMachinesetOSTypeIdResponse, error) {
	response := new(IMachinesetOSTypeIdResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetHardwareVersion(request *IMachinegetHardwareVersion) (*IMachinegetHardwareVersionResponse, error) {
	response := new(IMachinegetHardwareVersionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetHardwareVersion(request *IMachinesetHardwareVersion) (*IMachinesetHardwareVersionResponse, error) {
	response := new(IMachinesetHardwareVersionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetHardwareUUID(request *IMachinegetHardwareUUID) (*IMachinegetHardwareUUIDResponse, error) {
	response := new(IMachinegetHardwareUUIDResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetHardwareUUID(request *IMachinesetHardwareUUID) (*IMachinesetHardwareUUIDResponse, error) {
	response := new(IMachinesetHardwareUUIDResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetCPUCount(request *IMachinegetCPUCount) (*IMachinegetCPUCountResponse, error) {
	response := new(IMachinegetCPUCountResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetCPUCount(request *IMachinesetCPUCount) (*IMachinesetCPUCountResponse, error) {
	response := new(IMachinesetCPUCountResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetCPUHotPlugEnabled(request *IMachinegetCPUHotPlugEnabled) (*IMachinegetCPUHotPlugEnabledResponse, error) {
	response := new(IMachinegetCPUHotPlugEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetCPUHotPlugEnabled(request *IMachinesetCPUHotPlugEnabled) (*IMachinesetCPUHotPlugEnabledResponse, error) {
	response := new(IMachinesetCPUHotPlugEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetMemorySize(request *IMachinegetMemorySize) (*IMachinegetMemorySizeResponse, error) {
	response := new(IMachinegetMemorySizeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetMemorySize(request *IMachinesetMemorySize) (*IMachinesetMemorySizeResponse, error) {
	response := new(IMachinesetMemorySizeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetMemoryBalloonSize(request *IMachinegetMemoryBalloonSize) (*IMachinegetMemoryBalloonSizeResponse, error) {
	response := new(IMachinegetMemoryBalloonSizeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetMemoryBalloonSize(request *IMachinesetMemoryBalloonSize) (*IMachinesetMemoryBalloonSizeResponse, error) {
	response := new(IMachinesetMemoryBalloonSizeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetPageFusionEnabled(request *IMachinegetPageFusionEnabled) (*IMachinegetPageFusionEnabledResponse, error) {
	response := new(IMachinegetPageFusionEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetPageFusionEnabled(request *IMachinesetPageFusionEnabled) (*IMachinesetPageFusionEnabledResponse, error) {
	response := new(IMachinesetPageFusionEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetVRAMSize(request *IMachinegetVRAMSize) (*IMachinegetVRAMSizeResponse, error) {
	response := new(IMachinegetVRAMSizeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetVRAMSize(request *IMachinesetVRAMSize) (*IMachinesetVRAMSizeResponse, error) {
	response := new(IMachinesetVRAMSizeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetAccelerate3DEnabled(request *IMachinegetAccelerate3DEnabled) (*IMachinegetAccelerate3DEnabledResponse, error) {
	response := new(IMachinegetAccelerate3DEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetAccelerate3DEnabled(request *IMachinesetAccelerate3DEnabled) (*IMachinesetAccelerate3DEnabledResponse, error) {
	response := new(IMachinesetAccelerate3DEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetAccelerate2DVideoEnabled(request *IMachinegetAccelerate2DVideoEnabled) (*IMachinegetAccelerate2DVideoEnabledResponse, error) {
	response := new(IMachinegetAccelerate2DVideoEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetAccelerate2DVideoEnabled(request *IMachinesetAccelerate2DVideoEnabled) (*IMachinesetAccelerate2DVideoEnabledResponse, error) {
	response := new(IMachinesetAccelerate2DVideoEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetMonitorCount(request *IMachinegetMonitorCount) (*IMachinegetMonitorCountResponse, error) {
	response := new(IMachinegetMonitorCountResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetMonitorCount(request *IMachinesetMonitorCount) (*IMachinesetMonitorCountResponse, error) {
	response := new(IMachinesetMonitorCountResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetBIOSSettings(request *IMachinegetBIOSSettings) (*IMachinegetBIOSSettingsResponse, error) {
	response := new(IMachinegetBIOSSettingsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetFirmwareType(request *IMachinegetFirmwareType) (*IMachinegetFirmwareTypeResponse, error) {
	response := new(IMachinegetFirmwareTypeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetFirmwareType(request *IMachinesetFirmwareType) (*IMachinesetFirmwareTypeResponse, error) {
	response := new(IMachinesetFirmwareTypeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetPointingHidType(request *IMachinegetPointingHidType) (*IMachinegetPointingHidTypeResponse, error) {
	response := new(IMachinegetPointingHidTypeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetPointingHidType(request *IMachinesetPointingHidType) (*IMachinesetPointingHidTypeResponse, error) {
	response := new(IMachinesetPointingHidTypeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetKeyboardHidType(request *IMachinegetKeyboardHidType) (*IMachinegetKeyboardHidTypeResponse, error) {
	response := new(IMachinegetKeyboardHidTypeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetKeyboardHidType(request *IMachinesetKeyboardHidType) (*IMachinesetKeyboardHidTypeResponse, error) {
	response := new(IMachinesetKeyboardHidTypeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetHpetEnabled(request *IMachinegetHpetEnabled) (*IMachinegetHpetEnabledResponse, error) {
	response := new(IMachinegetHpetEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetHpetEnabled(request *IMachinesetHpetEnabled) (*IMachinesetHpetEnabledResponse, error) {
	response := new(IMachinesetHpetEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetSnapshotFolder(request *IMachinegetSnapshotFolder) (*IMachinegetSnapshotFolderResponse, error) {
	response := new(IMachinegetSnapshotFolderResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetSnapshotFolder(request *IMachinesetSnapshotFolder) (*IMachinesetSnapshotFolderResponse, error) {
	response := new(IMachinesetSnapshotFolderResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetVRDPServer(request *IMachinegetVRDPServer) (*IMachinegetVRDPServerResponse, error) {
	response := new(IMachinegetVRDPServerResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetMediumAttachments(request *IMachinegetMediumAttachments) (*IMachinegetMediumAttachmentsResponse, error) {
	response := new(IMachinegetMediumAttachmentsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetUSBController(request *IMachinegetUSBController) (*IMachinegetUSBControllerResponse, error) {
	response := new(IMachinegetUSBControllerResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetAudioAdapter(request *IMachinegetAudioAdapter) (*IMachinegetAudioAdapterResponse, error) {
	response := new(IMachinegetAudioAdapterResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetStorageControllers(request *IMachinegetStorageControllers) (*IMachinegetStorageControllersResponse, error) {
	response := new(IMachinegetStorageControllersResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetSettingsFilePath(request *IMachinegetSettingsFilePath) (*IMachinegetSettingsFilePathResponse, error) {
	response := new(IMachinegetSettingsFilePathResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetSettingsModified(request *IMachinegetSettingsModified) (*IMachinegetSettingsModifiedResponse, error) {
	response := new(IMachinegetSettingsModifiedResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetSessionState(request *IMachinegetSessionState) (*IMachinegetSessionStateResponse, error) {
	response := new(IMachinegetSessionStateResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetSessionType(request *IMachinegetSessionType) (*IMachinegetSessionTypeResponse, error) {
	response := new(IMachinegetSessionTypeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetSessionPid(request *IMachinegetSessionPid) (*IMachinegetSessionPidResponse, error) {
	response := new(IMachinegetSessionPidResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetState(request *IMachinegetState) (*IMachinegetStateResponse, error) {
	response := new(IMachinegetStateResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetLastStateChange(request *IMachinegetLastStateChange) (*IMachinegetLastStateChangeResponse, error) {
	response := new(IMachinegetLastStateChangeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetStateFilePath(request *IMachinegetStateFilePath) (*IMachinegetStateFilePathResponse, error) {
	response := new(IMachinegetStateFilePathResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetLogFolder(request *IMachinegetLogFolder) (*IMachinegetLogFolderResponse, error) {
	response := new(IMachinegetLogFolderResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetCurrentSnapshot(request *IMachinegetCurrentSnapshot) (*IMachinegetCurrentSnapshotResponse, error) {
	response := new(IMachinegetCurrentSnapshotResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetSnapshotCount(request *IMachinegetSnapshotCount) (*IMachinegetSnapshotCountResponse, error) {
	response := new(IMachinegetSnapshotCountResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetCurrentStateModified(request *IMachinegetCurrentStateModified) (*IMachinegetCurrentStateModifiedResponse, error) {
	response := new(IMachinegetCurrentStateModifiedResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetSharedFolders(request *IMachinegetSharedFolders) (*IMachinegetSharedFoldersResponse, error) {
	response := new(IMachinegetSharedFoldersResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetClipboardMode(request *IMachinegetClipboardMode) (*IMachinegetClipboardModeResponse, error) {
	response := new(IMachinegetClipboardModeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetClipboardMode(request *IMachinesetClipboardMode) (*IMachinesetClipboardModeResponse, error) {
	response := new(IMachinesetClipboardModeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetGuestPropertyNotificationPatterns(request *IMachinegetGuestPropertyNotificationPatterns) (*IMachinegetGuestPropertyNotificationPatternsResponse, error) {
	response := new(IMachinegetGuestPropertyNotificationPatternsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetGuestPropertyNotificationPatterns(request *IMachinesetGuestPropertyNotificationPatterns) (*IMachinesetGuestPropertyNotificationPatternsResponse, error) {
	response := new(IMachinesetGuestPropertyNotificationPatternsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetTeleporterEnabled(request *IMachinegetTeleporterEnabled) (*IMachinegetTeleporterEnabledResponse, error) {
	response := new(IMachinegetTeleporterEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetTeleporterEnabled(request *IMachinesetTeleporterEnabled) (*IMachinesetTeleporterEnabledResponse, error) {
	response := new(IMachinesetTeleporterEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetTeleporterPort(request *IMachinegetTeleporterPort) (*IMachinegetTeleporterPortResponse, error) {
	response := new(IMachinegetTeleporterPortResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetTeleporterPort(request *IMachinesetTeleporterPort) (*IMachinesetTeleporterPortResponse, error) {
	response := new(IMachinesetTeleporterPortResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetTeleporterAddress(request *IMachinegetTeleporterAddress) (*IMachinegetTeleporterAddressResponse, error) {
	response := new(IMachinegetTeleporterAddressResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetTeleporterAddress(request *IMachinesetTeleporterAddress) (*IMachinesetTeleporterAddressResponse, error) {
	response := new(IMachinesetTeleporterAddressResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetTeleporterPassword(request *IMachinegetTeleporterPassword) (*IMachinegetTeleporterPasswordResponse, error) {
	response := new(IMachinegetTeleporterPasswordResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetTeleporterPassword(request *IMachinesetTeleporterPassword) (*IMachinesetTeleporterPasswordResponse, error) {
	response := new(IMachinesetTeleporterPasswordResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetRTCUseUTC(request *IMachinegetRTCUseUTC) (*IMachinegetRTCUseUTCResponse, error) {
	response := new(IMachinegetRTCUseUTCResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetRTCUseUTC(request *IMachinesetRTCUseUTC) (*IMachinesetRTCUseUTCResponse, error) {
	response := new(IMachinesetRTCUseUTCResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetIoCacheEnabled(request *IMachinegetIoCacheEnabled) (*IMachinegetIoCacheEnabledResponse, error) {
	response := new(IMachinegetIoCacheEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetIoCacheEnabled(request *IMachinesetIoCacheEnabled) (*IMachinesetIoCacheEnabledResponse, error) {
	response := new(IMachinesetIoCacheEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetIoCacheSize(request *IMachinegetIoCacheSize) (*IMachinegetIoCacheSizeResponse, error) {
	response := new(IMachinegetIoCacheSizeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetIoCacheSize(request *IMachinesetIoCacheSize) (*IMachinesetIoCacheSizeResponse, error) {
	response := new(IMachinesetIoCacheSizeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetIoBandwidthMax(request *IMachinegetIoBandwidthMax) (*IMachinegetIoBandwidthMaxResponse, error) {
	response := new(IMachinegetIoBandwidthMaxResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetIoBandwidthMax(request *IMachinesetIoBandwidthMax) (*IMachinesetIoBandwidthMaxResponse, error) {
	response := new(IMachinesetIoBandwidthMaxResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetBootOrder(request *IMachinesetBootOrder) (*IMachinesetBootOrderResponse, error) {
	response := new(IMachinesetBootOrderResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetBootOrder(request *IMachinegetBootOrder) (*IMachinegetBootOrderResponse, error) {
	response := new(IMachinegetBootOrderResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineattachDevice(request *IMachineattachDevice) (*IMachineattachDeviceResponse, error) {
	response := new(IMachineattachDeviceResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinedetachDevice(request *IMachinedetachDevice) (*IMachinedetachDeviceResponse, error) {
	response := new(IMachinedetachDeviceResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinepassthroughDevice(request *IMachinepassthroughDevice) (*IMachinepassthroughDeviceResponse, error) {
	response := new(IMachinepassthroughDeviceResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinemountMedium(request *IMachinemountMedium) (*IMachinemountMediumResponse, error) {
	response := new(IMachinemountMediumResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetMedium(request *IMachinegetMedium) (*IMachinegetMediumResponse, error) {
	response := new(IMachinegetMediumResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetMediumAttachmentsOfController(request *IMachinegetMediumAttachmentsOfController) (*IMachinegetMediumAttachmentsOfControllerResponse, error) {
	response := new(IMachinegetMediumAttachmentsOfControllerResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetMediumAttachment(request *IMachinegetMediumAttachment) (*IMachinegetMediumAttachmentResponse, error) {
	response := new(IMachinegetMediumAttachmentResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetNetworkAdapter(request *IMachinegetNetworkAdapter) (*IMachinegetNetworkAdapterResponse, error) {
	response := new(IMachinegetNetworkAdapterResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineaddStorageController(request *IMachineaddStorageController) (*IMachineaddStorageControllerResponse, error) {
	response := new(IMachineaddStorageControllerResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetStorageControllerByName(request *IMachinegetStorageControllerByName) (*IMachinegetStorageControllerByNameResponse, error) {
	response := new(IMachinegetStorageControllerByNameResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetStorageControllerByInstance(request *IMachinegetStorageControllerByInstance) (*IMachinegetStorageControllerByInstanceResponse, error) {
	response := new(IMachinegetStorageControllerByInstanceResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineremoveStorageController(request *IMachineremoveStorageController) (*IMachineremoveStorageControllerResponse, error) {
	response := new(IMachineremoveStorageControllerResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetSerialPort(request *IMachinegetSerialPort) (*IMachinegetSerialPortResponse, error) {
	response := new(IMachinegetSerialPortResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetParallelPort(request *IMachinegetParallelPort) (*IMachinegetParallelPortResponse, error) {
	response := new(IMachinegetParallelPortResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetExtraDataKeys(request *IMachinegetExtraDataKeys) (*IMachinegetExtraDataKeysResponse, error) {
	response := new(IMachinegetExtraDataKeysResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetExtraData(request *IMachinegetExtraData) (*IMachinegetExtraDataResponse, error) {
	response := new(IMachinegetExtraDataResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetExtraData(request *IMachinesetExtraData) (*IMachinesetExtraDataResponse, error) {
	response := new(IMachinesetExtraDataResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetCPUProperty(request *IMachinegetCPUProperty) (*IMachinegetCPUPropertyResponse, error) {
	response := new(IMachinegetCPUPropertyResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetCPUProperty(request *IMachinesetCPUProperty) (*IMachinesetCPUPropertyResponse, error) {
	response := new(IMachinesetCPUPropertyResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetCPUIDLeaf(request *IMachinegetCPUIDLeaf) (*IMachinegetCPUIDLeafResponse, error) {
	response := new(IMachinegetCPUIDLeafResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetCPUIDLeaf(request *IMachinesetCPUIDLeaf) (*IMachinesetCPUIDLeafResponse, error) {
	response := new(IMachinesetCPUIDLeafResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineremoveCPUIDLeaf(request *IMachineremoveCPUIDLeaf) (*IMachineremoveCPUIDLeafResponse, error) {
	response := new(IMachineremoveCPUIDLeafResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineremoveAllCPUIDLeaves(request *IMachineremoveAllCPUIDLeaves) (*IMachineremoveAllCPUIDLeavesResponse, error) {
	response := new(IMachineremoveAllCPUIDLeavesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetHWVirtExProperty(request *IMachinegetHWVirtExProperty) (*IMachinegetHWVirtExPropertyResponse, error) {
	response := new(IMachinegetHWVirtExPropertyResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetHWVirtExProperty(request *IMachinesetHWVirtExProperty) (*IMachinesetHWVirtExPropertyResponse, error) {
	response := new(IMachinesetHWVirtExPropertyResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesaveSettings(request *IMachinesaveSettings) (*IMachinesaveSettingsResponse, error) {
	response := new(IMachinesaveSettingsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinediscardSettings(request *IMachinediscardSettings) (*IMachinediscardSettingsResponse, error) {
	response := new(IMachinediscardSettingsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinedeleteSettings(request *IMachinedeleteSettings) (*IMachinedeleteSettingsResponse, error) {
	response := new(IMachinedeleteSettingsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineexport(request *IMachineexport) (*IMachineexportResponse, error) {
	response := new(IMachineexportResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetSnapshot(request *IMachinegetSnapshot) (*IMachinegetSnapshotResponse, error) {
	response := new(IMachinegetSnapshotResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinefindSnapshot(request *IMachinefindSnapshot) (*IMachinefindSnapshotResponse, error) {
	response := new(IMachinefindSnapshotResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetCurrentSnapshot(request *IMachinesetCurrentSnapshot) (*IMachinesetCurrentSnapshotResponse, error) {
	response := new(IMachinesetCurrentSnapshotResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinecreateSharedFolder(request *IMachinecreateSharedFolder) (*IMachinecreateSharedFolderResponse, error) {
	response := new(IMachinecreateSharedFolderResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineremoveSharedFolder(request *IMachineremoveSharedFolder) (*IMachineremoveSharedFolderResponse, error) {
	response := new(IMachineremoveSharedFolderResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinecanShowConsoleWindow(request *IMachinecanShowConsoleWindow) (*IMachinecanShowConsoleWindowResponse, error) {
	response := new(IMachinecanShowConsoleWindowResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineshowConsoleWindow(request *IMachineshowConsoleWindow) (*IMachineshowConsoleWindowResponse, error) {
	response := new(IMachineshowConsoleWindowResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetGuestProperty(request *IMachinegetGuestProperty) (*IMachinegetGuestPropertyResponse, error) {
	response := new(IMachinegetGuestPropertyResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetGuestPropertyValue(request *IMachinegetGuestPropertyValue) (*IMachinegetGuestPropertyValueResponse, error) {
	response := new(IMachinegetGuestPropertyValueResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetGuestPropertyTimestamp(request *IMachinegetGuestPropertyTimestamp) (*IMachinegetGuestPropertyTimestampResponse, error) {
	response := new(IMachinegetGuestPropertyTimestampResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetGuestProperty(request *IMachinesetGuestProperty) (*IMachinesetGuestPropertyResponse, error) {
	response := new(IMachinesetGuestPropertyResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetGuestPropertyValue(request *IMachinesetGuestPropertyValue) (*IMachinesetGuestPropertyValueResponse, error) {
	response := new(IMachinesetGuestPropertyValueResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineenumerateGuestProperties(request *IMachineenumerateGuestProperties) (*IMachineenumerateGuestPropertiesResponse, error) {
	response := new(IMachineenumerateGuestPropertiesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinequerySavedThumbnailSize(request *IMachinequerySavedThumbnailSize) (*IMachinequerySavedThumbnailSizeResponse, error) {
	response := new(IMachinequerySavedThumbnailSizeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinereadSavedThumbnailToArray(request *IMachinereadSavedThumbnailToArray) (*IMachinereadSavedThumbnailToArrayResponse, error) {
	response := new(IMachinereadSavedThumbnailToArrayResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinequerySavedScreenshotPNGSize(request *IMachinequerySavedScreenshotPNGSize) (*IMachinequerySavedScreenshotPNGSizeResponse, error) {
	response := new(IMachinequerySavedScreenshotPNGSizeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinereadSavedScreenshotPNGToArray(request *IMachinereadSavedScreenshotPNGToArray) (*IMachinereadSavedScreenshotPNGToArrayResponse, error) {
	response := new(IMachinereadSavedScreenshotPNGToArrayResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinehotPlugCPU(request *IMachinehotPlugCPU) (*IMachinehotPlugCPUResponse, error) {
	response := new(IMachinehotPlugCPUResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinehotUnplugCPU(request *IMachinehotUnplugCPU) (*IMachinehotUnplugCPUResponse, error) {
	response := new(IMachinehotUnplugCPUResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetCPUStatus(request *IMachinegetCPUStatus) (*IMachinegetCPUStatusResponse, error) {
	response := new(IMachinegetCPUStatusResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinequeryLogFilename(request *IMachinequeryLogFilename) (*IMachinequeryLogFilenameResponse, error) {
	response := new(IMachinequeryLogFilenameResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinereadLog(request *IMachinereadLog) (*IMachinereadLogResponse, error) {
	response := new(IMachinereadLogResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IConsolegetMachine(request *IConsolegetMachine) (*IConsolegetMachineResponse, error) {
	response := new(IConsolegetMachineResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IConsolegetState(request *IConsolegetState) (*IConsolegetStateResponse, error) {
	response := new(IConsolegetStateResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IConsolegetGuest(request *IConsolegetGuest) (*IConsolegetGuestResponse, error) {
	response := new(IConsolegetGuestResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IConsolegetKeyboard(request *IConsolegetKeyboard) (*IConsolegetKeyboardResponse, error) {
	response := new(IConsolegetKeyboardResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IConsolegetMouse(request *IConsolegetMouse) (*IConsolegetMouseResponse, error) {
	response := new(IConsolegetMouseResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IConsolegetDisplay(request *IConsolegetDisplay) (*IConsolegetDisplayResponse, error) {
	response := new(IConsolegetDisplayResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IConsolegetUSBDevices(request *IConsolegetUSBDevices) (*IConsolegetUSBDevicesResponse, error) {
	response := new(IConsolegetUSBDevicesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IConsolegetRemoteUSBDevices(request *IConsolegetRemoteUSBDevices) (*IConsolegetRemoteUSBDevicesResponse, error) {
	response := new(IConsolegetRemoteUSBDevicesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IConsolegetSharedFolders(request *IConsolegetSharedFolders) (*IConsolegetSharedFoldersResponse, error) {
	response := new(IConsolegetSharedFoldersResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IConsolegetRemoteDisplayInfo(request *IConsolegetRemoteDisplayInfo) (*IConsolegetRemoteDisplayInfoResponse, error) {
	response := new(IConsolegetRemoteDisplayInfoResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IConsolepowerUp(request *IConsolepowerUp) (*IConsolepowerUpResponse, error) {
	response := new(IConsolepowerUpResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IConsolepowerUpPaused(request *IConsolepowerUpPaused) (*IConsolepowerUpPausedResponse, error) {
	response := new(IConsolepowerUpPausedResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IConsolepowerDown(request *IConsolepowerDown) (*IConsolepowerDownResponse, error) {
	response := new(IConsolepowerDownResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IConsolereset(request *IConsolereset) (*IConsoleresetResponse, error) {
	response := new(IConsoleresetResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IConsolepause(request *IConsolepause) (*IConsolepauseResponse, error) {
	response := new(IConsolepauseResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IConsoleresume(request *IConsoleresume) (*IConsoleresumeResponse, error) {
	response := new(IConsoleresumeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IConsolepowerButton(request *IConsolepowerButton) (*IConsolepowerButtonResponse, error) {
	response := new(IConsolepowerButtonResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IConsolesleepButton(request *IConsolesleepButton) (*IConsolesleepButtonResponse, error) {
	response := new(IConsolesleepButtonResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IConsolegetPowerButtonHandled(request *IConsolegetPowerButtonHandled) (*IConsolegetPowerButtonHandledResponse, error) {
	response := new(IConsolegetPowerButtonHandledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IConsolegetGuestEnteredACPIMode(request *IConsolegetGuestEnteredACPIMode) (*IConsolegetGuestEnteredACPIModeResponse, error) {
	response := new(IConsolegetGuestEnteredACPIModeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IConsolesaveState(request *IConsolesaveState) (*IConsolesaveStateResponse, error) {
	response := new(IConsolesaveStateResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IConsoleadoptSavedState(request *IConsoleadoptSavedState) (*IConsoleadoptSavedStateResponse, error) {
	response := new(IConsoleadoptSavedStateResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IConsoleforgetSavedState(request *IConsoleforgetSavedState) (*IConsoleforgetSavedStateResponse, error) {
	response := new(IConsoleforgetSavedStateResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IConsolegetDeviceActivity(request *IConsolegetDeviceActivity) (*IConsolegetDeviceActivityResponse, error) {
	response := new(IConsolegetDeviceActivityResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IConsoleattachUSBDevice(request *IConsoleattachUSBDevice) (*IConsoleattachUSBDeviceResponse, error) {
	response := new(IConsoleattachUSBDeviceResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IConsoledetachUSBDevice(request *IConsoledetachUSBDevice) (*IConsoledetachUSBDeviceResponse, error) {
	response := new(IConsoledetachUSBDeviceResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IConsolefindUSBDeviceByAddress(request *IConsolefindUSBDeviceByAddress) (*IConsolefindUSBDeviceByAddressResponse, error) {
	response := new(IConsolefindUSBDeviceByAddressResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IConsolefindUSBDeviceById(request *IConsolefindUSBDeviceById) (*IConsolefindUSBDeviceByIdResponse, error) {
	response := new(IConsolefindUSBDeviceByIdResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IConsolecreateSharedFolder(request *IConsolecreateSharedFolder) (*IConsolecreateSharedFolderResponse, error) {
	response := new(IConsolecreateSharedFolderResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IConsoleremoveSharedFolder(request *IConsoleremoveSharedFolder) (*IConsoleremoveSharedFolderResponse, error) {
	response := new(IConsoleremoveSharedFolderResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IConsoletakeSnapshot(request *IConsoletakeSnapshot) (*IConsoletakeSnapshotResponse, error) {
	response := new(IConsoletakeSnapshotResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IConsoledeleteSnapshot(request *IConsoledeleteSnapshot) (*IConsoledeleteSnapshotResponse, error) {
	response := new(IConsoledeleteSnapshotResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IConsolerestoreSnapshot(request *IConsolerestoreSnapshot) (*IConsolerestoreSnapshotResponse, error) {
	response := new(IConsolerestoreSnapshotResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IConsoleteleport(request *IConsoleteleport) (*IConsoleteleportResponse, error) {
	response := new(IConsoleteleportResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostNetworkInterfacegetName(request *IHostNetworkInterfacegetName) (*IHostNetworkInterfacegetNameResponse, error) {
	response := new(IHostNetworkInterfacegetNameResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostNetworkInterfacegetId(request *IHostNetworkInterfacegetId) (*IHostNetworkInterfacegetIdResponse, error) {
	response := new(IHostNetworkInterfacegetIdResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostNetworkInterfacegetNetworkName(request *IHostNetworkInterfacegetNetworkName) (*IHostNetworkInterfacegetNetworkNameResponse, error) {
	response := new(IHostNetworkInterfacegetNetworkNameResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostNetworkInterfacegetDhcpEnabled(request *IHostNetworkInterfacegetDhcpEnabled) (*IHostNetworkInterfacegetDhcpEnabledResponse, error) {
	response := new(IHostNetworkInterfacegetDhcpEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostNetworkInterfacegetIPAddress(request *IHostNetworkInterfacegetIPAddress) (*IHostNetworkInterfacegetIPAddressResponse, error) {
	response := new(IHostNetworkInterfacegetIPAddressResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostNetworkInterfacegetNetworkMask(request *IHostNetworkInterfacegetNetworkMask) (*IHostNetworkInterfacegetNetworkMaskResponse, error) {
	response := new(IHostNetworkInterfacegetNetworkMaskResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostNetworkInterfacegetIPV6Supported(request *IHostNetworkInterfacegetIPV6Supported) (*IHostNetworkInterfacegetIPV6SupportedResponse, error) {
	response := new(IHostNetworkInterfacegetIPV6SupportedResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostNetworkInterfacegetIPV6Address(request *IHostNetworkInterfacegetIPV6Address) (*IHostNetworkInterfacegetIPV6AddressResponse, error) {
	response := new(IHostNetworkInterfacegetIPV6AddressResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostNetworkInterfacegetIPV6NetworkMaskPrefixLength(request *IHostNetworkInterfacegetIPV6NetworkMaskPrefixLength) (*IHostNetworkInterfacegetIPV6NetworkMaskPrefixLengthResponse, error) {
	response := new(IHostNetworkInterfacegetIPV6NetworkMaskPrefixLengthResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostNetworkInterfacegetHardwareAddress(request *IHostNetworkInterfacegetHardwareAddress) (*IHostNetworkInterfacegetHardwareAddressResponse, error) {
	response := new(IHostNetworkInterfacegetHardwareAddressResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostNetworkInterfacegetMediumType(request *IHostNetworkInterfacegetMediumType) (*IHostNetworkInterfacegetMediumTypeResponse, error) {
	response := new(IHostNetworkInterfacegetMediumTypeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostNetworkInterfacegetStatus(request *IHostNetworkInterfacegetStatus) (*IHostNetworkInterfacegetStatusResponse, error) {
	response := new(IHostNetworkInterfacegetStatusResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostNetworkInterfacegetInterfaceType(request *IHostNetworkInterfacegetInterfaceType) (*IHostNetworkInterfacegetInterfaceTypeResponse, error) {
	response := new(IHostNetworkInterfacegetInterfaceTypeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostNetworkInterfaceenableStaticIpConfig(request *IHostNetworkInterfaceenableStaticIpConfig) (*IHostNetworkInterfaceenableStaticIpConfigResponse, error) {
	response := new(IHostNetworkInterfaceenableStaticIpConfigResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostNetworkInterfaceenableStaticIpConfigV6(request *IHostNetworkInterfaceenableStaticIpConfigV6) (*IHostNetworkInterfaceenableStaticIpConfigV6Response, error) {
	response := new(IHostNetworkInterfaceenableStaticIpConfigV6Response)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostNetworkInterfaceenableDynamicIpConfig(request *IHostNetworkInterfaceenableDynamicIpConfig) (*IHostNetworkInterfaceenableDynamicIpConfigResponse, error) {
	response := new(IHostNetworkInterfaceenableDynamicIpConfigResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostNetworkInterfacedhcpRediscover(request *IHostNetworkInterfacedhcpRediscover) (*IHostNetworkInterfacedhcpRediscoverResponse, error) {
	response := new(IHostNetworkInterfacedhcpRediscoverResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostgetDVDDrives(request *IHostgetDVDDrives) (*IHostgetDVDDrivesResponse, error) {
	response := new(IHostgetDVDDrivesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostgetFloppyDrives(request *IHostgetFloppyDrives) (*IHostgetFloppyDrivesResponse, error) {
	response := new(IHostgetFloppyDrivesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostgetUSBDevices(request *IHostgetUSBDevices) (*IHostgetUSBDevicesResponse, error) {
	response := new(IHostgetUSBDevicesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostgetUSBDeviceFilters(request *IHostgetUSBDeviceFilters) (*IHostgetUSBDeviceFiltersResponse, error) {
	response := new(IHostgetUSBDeviceFiltersResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostgetNetworkInterfaces(request *IHostgetNetworkInterfaces) (*IHostgetNetworkInterfacesResponse, error) {
	response := new(IHostgetNetworkInterfacesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostgetProcessorCount(request *IHostgetProcessorCount) (*IHostgetProcessorCountResponse, error) {
	response := new(IHostgetProcessorCountResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostgetProcessorOnlineCount(request *IHostgetProcessorOnlineCount) (*IHostgetProcessorOnlineCountResponse, error) {
	response := new(IHostgetProcessorOnlineCountResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostgetProcessorCoreCount(request *IHostgetProcessorCoreCount) (*IHostgetProcessorCoreCountResponse, error) {
	response := new(IHostgetProcessorCoreCountResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostgetMemorySize(request *IHostgetMemorySize) (*IHostgetMemorySizeResponse, error) {
	response := new(IHostgetMemorySizeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostgetMemoryAvailable(request *IHostgetMemoryAvailable) (*IHostgetMemoryAvailableResponse, error) {
	response := new(IHostgetMemoryAvailableResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostgetOperatingSystem(request *IHostgetOperatingSystem) (*IHostgetOperatingSystemResponse, error) {
	response := new(IHostgetOperatingSystemResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostgetOSVersion(request *IHostgetOSVersion) (*IHostgetOSVersionResponse, error) {
	response := new(IHostgetOSVersionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostgetUTCTime(request *IHostgetUTCTime) (*IHostgetUTCTimeResponse, error) {
	response := new(IHostgetUTCTimeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostgetAcceleration3DAvailable(request *IHostgetAcceleration3DAvailable) (*IHostgetAcceleration3DAvailableResponse, error) {
	response := new(IHostgetAcceleration3DAvailableResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostgetProcessorSpeed(request *IHostgetProcessorSpeed) (*IHostgetProcessorSpeedResponse, error) {
	response := new(IHostgetProcessorSpeedResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostgetProcessorFeature(request *IHostgetProcessorFeature) (*IHostgetProcessorFeatureResponse, error) {
	response := new(IHostgetProcessorFeatureResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostgetProcessorDescription(request *IHostgetProcessorDescription) (*IHostgetProcessorDescriptionResponse, error) {
	response := new(IHostgetProcessorDescriptionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostgetProcessorCPUIDLeaf(request *IHostgetProcessorCPUIDLeaf) (*IHostgetProcessorCPUIDLeafResponse, error) {
	response := new(IHostgetProcessorCPUIDLeafResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostcreateHostOnlyNetworkInterface(request *IHostcreateHostOnlyNetworkInterface) (*IHostcreateHostOnlyNetworkInterfaceResponse, error) {
	response := new(IHostcreateHostOnlyNetworkInterfaceResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostremoveHostOnlyNetworkInterface(request *IHostremoveHostOnlyNetworkInterface) (*IHostremoveHostOnlyNetworkInterfaceResponse, error) {
	response := new(IHostremoveHostOnlyNetworkInterfaceResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostcreateUSBDeviceFilter(request *IHostcreateUSBDeviceFilter) (*IHostcreateUSBDeviceFilterResponse, error) {
	response := new(IHostcreateUSBDeviceFilterResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostinsertUSBDeviceFilter(request *IHostinsertUSBDeviceFilter) (*IHostinsertUSBDeviceFilterResponse, error) {
	response := new(IHostinsertUSBDeviceFilterResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostremoveUSBDeviceFilter(request *IHostremoveUSBDeviceFilter) (*IHostremoveUSBDeviceFilterResponse, error) {
	response := new(IHostremoveUSBDeviceFilterResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostfindHostDVDDrive(request *IHostfindHostDVDDrive) (*IHostfindHostDVDDriveResponse, error) {
	response := new(IHostfindHostDVDDriveResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostfindHostFloppyDrive(request *IHostfindHostFloppyDrive) (*IHostfindHostFloppyDriveResponse, error) {
	response := new(IHostfindHostFloppyDriveResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostfindHostNetworkInterfaceByName(request *IHostfindHostNetworkInterfaceByName) (*IHostfindHostNetworkInterfaceByNameResponse, error) {
	response := new(IHostfindHostNetworkInterfaceByNameResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostfindHostNetworkInterfaceById(request *IHostfindHostNetworkInterfaceById) (*IHostfindHostNetworkInterfaceByIdResponse, error) {
	response := new(IHostfindHostNetworkInterfaceByIdResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostfindHostNetworkInterfacesOfType(request *IHostfindHostNetworkInterfacesOfType) (*IHostfindHostNetworkInterfacesOfTypeResponse, error) {
	response := new(IHostfindHostNetworkInterfacesOfTypeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostfindUSBDeviceById(request *IHostfindUSBDeviceById) (*IHostfindUSBDeviceByIdResponse, error) {
	response := new(IHostfindUSBDeviceByIdResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostfindUSBDeviceByAddress(request *IHostfindUSBDeviceByAddress) (*IHostfindUSBDeviceByAddressResponse, error) {
	response := new(IHostfindUSBDeviceByAddressResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiesgetMinGuestRAM(request *ISystemPropertiesgetMinGuestRAM) (*ISystemPropertiesgetMinGuestRAMResponse, error) {
	response := new(ISystemPropertiesgetMinGuestRAMResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiesgetMaxGuestRAM(request *ISystemPropertiesgetMaxGuestRAM) (*ISystemPropertiesgetMaxGuestRAMResponse, error) {
	response := new(ISystemPropertiesgetMaxGuestRAMResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiesgetMinGuestVRAM(request *ISystemPropertiesgetMinGuestVRAM) (*ISystemPropertiesgetMinGuestVRAMResponse, error) {
	response := new(ISystemPropertiesgetMinGuestVRAMResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiesgetMaxGuestVRAM(request *ISystemPropertiesgetMaxGuestVRAM) (*ISystemPropertiesgetMaxGuestVRAMResponse, error) {
	response := new(ISystemPropertiesgetMaxGuestVRAMResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiesgetMinGuestCPUCount(request *ISystemPropertiesgetMinGuestCPUCount) (*ISystemPropertiesgetMinGuestCPUCountResponse, error) {
	response := new(ISystemPropertiesgetMinGuestCPUCountResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiesgetMaxGuestCPUCount(request *ISystemPropertiesgetMaxGuestCPUCount) (*ISystemPropertiesgetMaxGuestCPUCountResponse, error) {
	response := new(ISystemPropertiesgetMaxGuestCPUCountResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiesgetMaxGuestMonitors(request *ISystemPropertiesgetMaxGuestMonitors) (*ISystemPropertiesgetMaxGuestMonitorsResponse, error) {
	response := new(ISystemPropertiesgetMaxGuestMonitorsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiesgetMaxVDISize(request *ISystemPropertiesgetMaxVDISize) (*ISystemPropertiesgetMaxVDISizeResponse, error) {
	response := new(ISystemPropertiesgetMaxVDISizeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiesgetNetworkAdapterCount(request *ISystemPropertiesgetNetworkAdapterCount) (*ISystemPropertiesgetNetworkAdapterCountResponse, error) {
	response := new(ISystemPropertiesgetNetworkAdapterCountResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiesgetSerialPortCount(request *ISystemPropertiesgetSerialPortCount) (*ISystemPropertiesgetSerialPortCountResponse, error) {
	response := new(ISystemPropertiesgetSerialPortCountResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiesgetParallelPortCount(request *ISystemPropertiesgetParallelPortCount) (*ISystemPropertiesgetParallelPortCountResponse, error) {
	response := new(ISystemPropertiesgetParallelPortCountResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiesgetMaxBootPosition(request *ISystemPropertiesgetMaxBootPosition) (*ISystemPropertiesgetMaxBootPositionResponse, error) {
	response := new(ISystemPropertiesgetMaxBootPositionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiesgetDefaultMachineFolder(request *ISystemPropertiesgetDefaultMachineFolder) (*ISystemPropertiesgetDefaultMachineFolderResponse, error) {
	response := new(ISystemPropertiesgetDefaultMachineFolderResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiessetDefaultMachineFolder(request *ISystemPropertiessetDefaultMachineFolder) (*ISystemPropertiessetDefaultMachineFolderResponse, error) {
	response := new(ISystemPropertiessetDefaultMachineFolderResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiesgetDefaultHardDiskFolder(request *ISystemPropertiesgetDefaultHardDiskFolder) (*ISystemPropertiesgetDefaultHardDiskFolderResponse, error) {
	response := new(ISystemPropertiesgetDefaultHardDiskFolderResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiessetDefaultHardDiskFolder(request *ISystemPropertiessetDefaultHardDiskFolder) (*ISystemPropertiessetDefaultHardDiskFolderResponse, error) {
	response := new(ISystemPropertiessetDefaultHardDiskFolderResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiesgetMediumFormats(request *ISystemPropertiesgetMediumFormats) (*ISystemPropertiesgetMediumFormatsResponse, error) {
	response := new(ISystemPropertiesgetMediumFormatsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiesgetDefaultHardDiskFormat(request *ISystemPropertiesgetDefaultHardDiskFormat) (*ISystemPropertiesgetDefaultHardDiskFormatResponse, error) {
	response := new(ISystemPropertiesgetDefaultHardDiskFormatResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiessetDefaultHardDiskFormat(request *ISystemPropertiessetDefaultHardDiskFormat) (*ISystemPropertiessetDefaultHardDiskFormatResponse, error) {
	response := new(ISystemPropertiessetDefaultHardDiskFormatResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiesgetFreeDiskSpaceWarning(request *ISystemPropertiesgetFreeDiskSpaceWarning) (*ISystemPropertiesgetFreeDiskSpaceWarningResponse, error) {
	response := new(ISystemPropertiesgetFreeDiskSpaceWarningResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiessetFreeDiskSpaceWarning(request *ISystemPropertiessetFreeDiskSpaceWarning) (*ISystemPropertiessetFreeDiskSpaceWarningResponse, error) {
	response := new(ISystemPropertiessetFreeDiskSpaceWarningResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiesgetFreeDiskSpacePercentWarning(request *ISystemPropertiesgetFreeDiskSpacePercentWarning) (*ISystemPropertiesgetFreeDiskSpacePercentWarningResponse, error) {
	response := new(ISystemPropertiesgetFreeDiskSpacePercentWarningResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiessetFreeDiskSpacePercentWarning(request *ISystemPropertiessetFreeDiskSpacePercentWarning) (*ISystemPropertiessetFreeDiskSpacePercentWarningResponse, error) {
	response := new(ISystemPropertiessetFreeDiskSpacePercentWarningResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiesgetFreeDiskSpaceError(request *ISystemPropertiesgetFreeDiskSpaceError) (*ISystemPropertiesgetFreeDiskSpaceErrorResponse, error) {
	response := new(ISystemPropertiesgetFreeDiskSpaceErrorResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiessetFreeDiskSpaceError(request *ISystemPropertiessetFreeDiskSpaceError) (*ISystemPropertiessetFreeDiskSpaceErrorResponse, error) {
	response := new(ISystemPropertiessetFreeDiskSpaceErrorResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiesgetFreeDiskSpacePercentError(request *ISystemPropertiesgetFreeDiskSpacePercentError) (*ISystemPropertiesgetFreeDiskSpacePercentErrorResponse, error) {
	response := new(ISystemPropertiesgetFreeDiskSpacePercentErrorResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiessetFreeDiskSpacePercentError(request *ISystemPropertiessetFreeDiskSpacePercentError) (*ISystemPropertiessetFreeDiskSpacePercentErrorResponse, error) {
	response := new(ISystemPropertiessetFreeDiskSpacePercentErrorResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiesgetRemoteDisplayAuthLibrary(request *ISystemPropertiesgetRemoteDisplayAuthLibrary) (*ISystemPropertiesgetRemoteDisplayAuthLibraryResponse, error) {
	response := new(ISystemPropertiesgetRemoteDisplayAuthLibraryResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiessetRemoteDisplayAuthLibrary(request *ISystemPropertiessetRemoteDisplayAuthLibrary) (*ISystemPropertiessetRemoteDisplayAuthLibraryResponse, error) {
	response := new(ISystemPropertiessetRemoteDisplayAuthLibraryResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiesgetWebServiceAuthLibrary(request *ISystemPropertiesgetWebServiceAuthLibrary) (*ISystemPropertiesgetWebServiceAuthLibraryResponse, error) {
	response := new(ISystemPropertiesgetWebServiceAuthLibraryResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiessetWebServiceAuthLibrary(request *ISystemPropertiessetWebServiceAuthLibrary) (*ISystemPropertiessetWebServiceAuthLibraryResponse, error) {
	response := new(ISystemPropertiessetWebServiceAuthLibraryResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiesgetLogHistoryCount(request *ISystemPropertiesgetLogHistoryCount) (*ISystemPropertiesgetLogHistoryCountResponse, error) {
	response := new(ISystemPropertiesgetLogHistoryCountResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiessetLogHistoryCount(request *ISystemPropertiessetLogHistoryCount) (*ISystemPropertiessetLogHistoryCountResponse, error) {
	response := new(ISystemPropertiessetLogHistoryCountResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiesgetDefaultAudioDriver(request *ISystemPropertiesgetDefaultAudioDriver) (*ISystemPropertiesgetDefaultAudioDriverResponse, error) {
	response := new(ISystemPropertiesgetDefaultAudioDriverResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiesgetMaxDevicesPerPortForStorageBus(request *ISystemPropertiesgetMaxDevicesPerPortForStorageBus) (*ISystemPropertiesgetMaxDevicesPerPortForStorageBusResponse, error) {
	response := new(ISystemPropertiesgetMaxDevicesPerPortForStorageBusResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiesgetMinPortCountForStorageBus(request *ISystemPropertiesgetMinPortCountForStorageBus) (*ISystemPropertiesgetMinPortCountForStorageBusResponse, error) {
	response := new(ISystemPropertiesgetMinPortCountForStorageBusResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiesgetMaxPortCountForStorageBus(request *ISystemPropertiesgetMaxPortCountForStorageBus) (*ISystemPropertiesgetMaxPortCountForStorageBusResponse, error) {
	response := new(ISystemPropertiesgetMaxPortCountForStorageBusResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiesgetMaxInstancesOfStorageBus(request *ISystemPropertiesgetMaxInstancesOfStorageBus) (*ISystemPropertiesgetMaxInstancesOfStorageBusResponse, error) {
	response := new(ISystemPropertiesgetMaxInstancesOfStorageBusResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiesgetDeviceTypesForStorageBus(request *ISystemPropertiesgetDeviceTypesForStorageBus) (*ISystemPropertiesgetDeviceTypesForStorageBusResponse, error) {
	response := new(ISystemPropertiesgetDeviceTypesForStorageBusResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestgetOSTypeId(request *IGuestgetOSTypeId) (*IGuestgetOSTypeIdResponse, error) {
	response := new(IGuestgetOSTypeIdResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestgetAdditionsActive(request *IGuestgetAdditionsActive) (*IGuestgetAdditionsActiveResponse, error) {
	response := new(IGuestgetAdditionsActiveResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestgetAdditionsVersion(request *IGuestgetAdditionsVersion) (*IGuestgetAdditionsVersionResponse, error) {
	response := new(IGuestgetAdditionsVersionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestgetSupportsSeamless(request *IGuestgetSupportsSeamless) (*IGuestgetSupportsSeamlessResponse, error) {
	response := new(IGuestgetSupportsSeamlessResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestgetSupportsGraphics(request *IGuestgetSupportsGraphics) (*IGuestgetSupportsGraphicsResponse, error) {
	response := new(IGuestgetSupportsGraphicsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestgetMemoryBalloonSize(request *IGuestgetMemoryBalloonSize) (*IGuestgetMemoryBalloonSizeResponse, error) {
	response := new(IGuestgetMemoryBalloonSizeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestsetMemoryBalloonSize(request *IGuestsetMemoryBalloonSize) (*IGuestsetMemoryBalloonSizeResponse, error) {
	response := new(IGuestsetMemoryBalloonSizeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestgetPageFusionEnabled(request *IGuestgetPageFusionEnabled) (*IGuestgetPageFusionEnabledResponse, error) {
	response := new(IGuestgetPageFusionEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestsetPageFusionEnabled(request *IGuestsetPageFusionEnabled) (*IGuestsetPageFusionEnabledResponse, error) {
	response := new(IGuestsetPageFusionEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestgetStatisticsUpdateInterval(request *IGuestgetStatisticsUpdateInterval) (*IGuestgetStatisticsUpdateIntervalResponse, error) {
	response := new(IGuestgetStatisticsUpdateIntervalResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestsetStatisticsUpdateInterval(request *IGuestsetStatisticsUpdateInterval) (*IGuestsetStatisticsUpdateIntervalResponse, error) {
	response := new(IGuestsetStatisticsUpdateIntervalResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestinternalGetStatistics(request *IGuestinternalGetStatistics) (*IGuestinternalGetStatisticsResponse, error) {
	response := new(IGuestinternalGetStatisticsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestsetCredentials(request *IGuestsetCredentials) (*IGuestsetCredentialsResponse, error) {
	response := new(IGuestsetCredentialsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestexecuteProcess(request *IGuestexecuteProcess) (*IGuestexecuteProcessResponse, error) {
	response := new(IGuestexecuteProcessResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestgetProcessOutput(request *IGuestgetProcessOutput) (*IGuestgetProcessOutputResponse, error) {
	response := new(IGuestgetProcessOutputResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestgetProcessStatus(request *IGuestgetProcessStatus) (*IGuestgetProcessStatusResponse, error) {
	response := new(IGuestgetProcessStatusResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IProgressgetId(request *IProgressgetId) (*IProgressgetIdResponse, error) {
	response := new(IProgressgetIdResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IProgressgetDescription(request *IProgressgetDescription) (*IProgressgetDescriptionResponse, error) {
	response := new(IProgressgetDescriptionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IProgressgetInitiator(request *IProgressgetInitiator) (*IProgressgetInitiatorResponse, error) {
	response := new(IProgressgetInitiatorResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IProgressgetCancelable(request *IProgressgetCancelable) (*IProgressgetCancelableResponse, error) {
	response := new(IProgressgetCancelableResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IProgressgetPercent(request *IProgressgetPercent) (*IProgressgetPercentResponse, error) {
	response := new(IProgressgetPercentResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IProgressgetTimeRemaining(request *IProgressgetTimeRemaining) (*IProgressgetTimeRemainingResponse, error) {
	response := new(IProgressgetTimeRemainingResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IProgressgetCompleted(request *IProgressgetCompleted) (*IProgressgetCompletedResponse, error) {
	response := new(IProgressgetCompletedResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IProgressgetCanceled(request *IProgressgetCanceled) (*IProgressgetCanceledResponse, error) {
	response := new(IProgressgetCanceledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IProgressgetResultCode(request *IProgressgetResultCode) (*IProgressgetResultCodeResponse, error) {
	response := new(IProgressgetResultCodeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IProgressgetErrorInfo(request *IProgressgetErrorInfo) (*IProgressgetErrorInfoResponse, error) {
	response := new(IProgressgetErrorInfoResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IProgressgetOperationCount(request *IProgressgetOperationCount) (*IProgressgetOperationCountResponse, error) {
	response := new(IProgressgetOperationCountResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IProgressgetOperation(request *IProgressgetOperation) (*IProgressgetOperationResponse, error) {
	response := new(IProgressgetOperationResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IProgressgetOperationDescription(request *IProgressgetOperationDescription) (*IProgressgetOperationDescriptionResponse, error) {
	response := new(IProgressgetOperationDescriptionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IProgressgetOperationPercent(request *IProgressgetOperationPercent) (*IProgressgetOperationPercentResponse, error) {
	response := new(IProgressgetOperationPercentResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IProgressgetTimeout(request *IProgressgetTimeout) (*IProgressgetTimeoutResponse, error) {
	response := new(IProgressgetTimeoutResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IProgresssetTimeout(request *IProgresssetTimeout) (*IProgresssetTimeoutResponse, error) {
	response := new(IProgresssetTimeoutResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IProgresssetCurrentOperationProgress(request *IProgresssetCurrentOperationProgress) (*IProgresssetCurrentOperationProgressResponse, error) {
	response := new(IProgresssetCurrentOperationProgressResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IProgresssetNextOperation(request *IProgresssetNextOperation) (*IProgresssetNextOperationResponse, error) {
	response := new(IProgresssetNextOperationResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IProgresswaitForCompletion(request *IProgresswaitForCompletion) (*IProgresswaitForCompletionResponse, error) {
	response := new(IProgresswaitForCompletionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IProgresswaitForOperationCompletion(request *IProgresswaitForOperationCompletion) (*IProgresswaitForOperationCompletionResponse, error) {
	response := new(IProgresswaitForOperationCompletionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IProgresscancel(request *IProgresscancel) (*IProgresscancelResponse, error) {
	response := new(IProgresscancelResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISnapshotgetId(request *ISnapshotgetId) (*ISnapshotgetIdResponse, error) {
	response := new(ISnapshotgetIdResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISnapshotgetName(request *ISnapshotgetName) (*ISnapshotgetNameResponse, error) {
	response := new(ISnapshotgetNameResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISnapshotsetName(request *ISnapshotsetName) (*ISnapshotsetNameResponse, error) {
	response := new(ISnapshotsetNameResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISnapshotgetDescription(request *ISnapshotgetDescription) (*ISnapshotgetDescriptionResponse, error) {
	response := new(ISnapshotgetDescriptionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISnapshotsetDescription(request *ISnapshotsetDescription) (*ISnapshotsetDescriptionResponse, error) {
	response := new(ISnapshotsetDescriptionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISnapshotgetTimeStamp(request *ISnapshotgetTimeStamp) (*ISnapshotgetTimeStampResponse, error) {
	response := new(ISnapshotgetTimeStampResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISnapshotgetOnline(request *ISnapshotgetOnline) (*ISnapshotgetOnlineResponse, error) {
	response := new(ISnapshotgetOnlineResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISnapshotgetMachine(request *ISnapshotgetMachine) (*ISnapshotgetMachineResponse, error) {
	response := new(ISnapshotgetMachineResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISnapshotgetParent(request *ISnapshotgetParent) (*ISnapshotgetParentResponse, error) {
	response := new(ISnapshotgetParentResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISnapshotgetChildren(request *ISnapshotgetChildren) (*ISnapshotgetChildrenResponse, error) {
	response := new(ISnapshotgetChildrenResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumgetId(request *IMediumgetId) (*IMediumgetIdResponse, error) {
	response := new(IMediumgetIdResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumgetDescription(request *IMediumgetDescription) (*IMediumgetDescriptionResponse, error) {
	response := new(IMediumgetDescriptionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumsetDescription(request *IMediumsetDescription) (*IMediumsetDescriptionResponse, error) {
	response := new(IMediumsetDescriptionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumgetState(request *IMediumgetState) (*IMediumgetStateResponse, error) {
	response := new(IMediumgetStateResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumgetLocation(request *IMediumgetLocation) (*IMediumgetLocationResponse, error) {
	response := new(IMediumgetLocationResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumsetLocation(request *IMediumsetLocation) (*IMediumsetLocationResponse, error) {
	response := new(IMediumsetLocationResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumgetName(request *IMediumgetName) (*IMediumgetNameResponse, error) {
	response := new(IMediumgetNameResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumgetDeviceType(request *IMediumgetDeviceType) (*IMediumgetDeviceTypeResponse, error) {
	response := new(IMediumgetDeviceTypeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumgetHostDrive(request *IMediumgetHostDrive) (*IMediumgetHostDriveResponse, error) {
	response := new(IMediumgetHostDriveResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumgetSize(request *IMediumgetSize) (*IMediumgetSizeResponse, error) {
	response := new(IMediumgetSizeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumgetFormat(request *IMediumgetFormat) (*IMediumgetFormatResponse, error) {
	response := new(IMediumgetFormatResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumgetMediumFormat(request *IMediumgetMediumFormat) (*IMediumgetMediumFormatResponse, error) {
	response := new(IMediumgetMediumFormatResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumgetType(request *IMediumgetType) (*IMediumgetTypeResponse, error) {
	response := new(IMediumgetTypeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumsetType(request *IMediumsetType) (*IMediumsetTypeResponse, error) {
	response := new(IMediumsetTypeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumgetParent(request *IMediumgetParent) (*IMediumgetParentResponse, error) {
	response := new(IMediumgetParentResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumgetChildren(request *IMediumgetChildren) (*IMediumgetChildrenResponse, error) {
	response := new(IMediumgetChildrenResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumgetBase(request *IMediumgetBase) (*IMediumgetBaseResponse, error) {
	response := new(IMediumgetBaseResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumgetReadOnly(request *IMediumgetReadOnly) (*IMediumgetReadOnlyResponse, error) {
	response := new(IMediumgetReadOnlyResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumgetLogicalSize(request *IMediumgetLogicalSize) (*IMediumgetLogicalSizeResponse, error) {
	response := new(IMediumgetLogicalSizeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumgetAutoReset(request *IMediumgetAutoReset) (*IMediumgetAutoResetResponse, error) {
	response := new(IMediumgetAutoResetResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumsetAutoReset(request *IMediumsetAutoReset) (*IMediumsetAutoResetResponse, error) {
	response := new(IMediumsetAutoResetResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumgetLastAccessError(request *IMediumgetLastAccessError) (*IMediumgetLastAccessErrorResponse, error) {
	response := new(IMediumgetLastAccessErrorResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumgetMachineIds(request *IMediumgetMachineIds) (*IMediumgetMachineIdsResponse, error) {
	response := new(IMediumgetMachineIdsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumrefreshState(request *IMediumrefreshState) (*IMediumrefreshStateResponse, error) {
	response := new(IMediumrefreshStateResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumgetSnapshotIds(request *IMediumgetSnapshotIds) (*IMediumgetSnapshotIdsResponse, error) {
	response := new(IMediumgetSnapshotIdsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumlockRead(request *IMediumlockRead) (*IMediumlockReadResponse, error) {
	response := new(IMediumlockReadResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumunlockRead(request *IMediumunlockRead) (*IMediumunlockReadResponse, error) {
	response := new(IMediumunlockReadResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumlockWrite(request *IMediumlockWrite) (*IMediumlockWriteResponse, error) {
	response := new(IMediumlockWriteResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumunlockWrite(request *IMediumunlockWrite) (*IMediumunlockWriteResponse, error) {
	response := new(IMediumunlockWriteResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumclose(request *IMediumclose) (*IMediumcloseResponse, error) {
	response := new(IMediumcloseResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumgetProperty(request *IMediumgetProperty) (*IMediumgetPropertyResponse, error) {
	response := new(IMediumgetPropertyResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumsetProperty(request *IMediumsetProperty) (*IMediumsetPropertyResponse, error) {
	response := new(IMediumsetPropertyResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumgetProperties(request *IMediumgetProperties) (*IMediumgetPropertiesResponse, error) {
	response := new(IMediumgetPropertiesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumsetProperties(request *IMediumsetProperties) (*IMediumsetPropertiesResponse, error) {
	response := new(IMediumsetPropertiesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumcreateBaseStorage(request *IMediumcreateBaseStorage) (*IMediumcreateBaseStorageResponse, error) {
	response := new(IMediumcreateBaseStorageResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumdeleteStorage(request *IMediumdeleteStorage) (*IMediumdeleteStorageResponse, error) {
	response := new(IMediumdeleteStorageResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumcreateDiffStorage(request *IMediumcreateDiffStorage) (*IMediumcreateDiffStorageResponse, error) {
	response := new(IMediumcreateDiffStorageResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediummergeTo(request *IMediummergeTo) (*IMediummergeToResponse, error) {
	response := new(IMediummergeToResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumcloneTo(request *IMediumcloneTo) (*IMediumcloneToResponse, error) {
	response := new(IMediumcloneToResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumcompact(request *IMediumcompact) (*IMediumcompactResponse, error) {
	response := new(IMediumcompactResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumresize(request *IMediumresize) (*IMediumresizeResponse, error) {
	response := new(IMediumresizeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumreset(request *IMediumreset) (*IMediumresetResponse, error) {
	response := new(IMediumresetResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumFormatgetId(request *IMediumFormatgetId) (*IMediumFormatgetIdResponse, error) {
	response := new(IMediumFormatgetIdResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumFormatgetName(request *IMediumFormatgetName) (*IMediumFormatgetNameResponse, error) {
	response := new(IMediumFormatgetNameResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumFormatgetFileExtensions(request *IMediumFormatgetFileExtensions) (*IMediumFormatgetFileExtensionsResponse, error) {
	response := new(IMediumFormatgetFileExtensionsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumFormatgetCapabilities(request *IMediumFormatgetCapabilities) (*IMediumFormatgetCapabilitiesResponse, error) {
	response := new(IMediumFormatgetCapabilitiesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumFormatdescribeProperties(request *IMediumFormatdescribeProperties) (*IMediumFormatdescribePropertiesResponse, error) {
	response := new(IMediumFormatdescribePropertiesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IKeyboardputScancode(request *IKeyboardputScancode) (*IKeyboardputScancodeResponse, error) {
	response := new(IKeyboardputScancodeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IKeyboardputScancodes(request *IKeyboardputScancodes) (*IKeyboardputScancodesResponse, error) {
	response := new(IKeyboardputScancodesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IKeyboardputCAD(request *IKeyboardputCAD) (*IKeyboardputCADResponse, error) {
	response := new(IKeyboardputCADResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMousegetAbsoluteSupported(request *IMousegetAbsoluteSupported) (*IMousegetAbsoluteSupportedResponse, error) {
	response := new(IMousegetAbsoluteSupportedResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMousegetRelativeSupported(request *IMousegetRelativeSupported) (*IMousegetRelativeSupportedResponse, error) {
	response := new(IMousegetRelativeSupportedResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMousegetNeedsHostCursor(request *IMousegetNeedsHostCursor) (*IMousegetNeedsHostCursorResponse, error) {
	response := new(IMousegetNeedsHostCursorResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMouseputMouseEvent(request *IMouseputMouseEvent) (*IMouseputMouseEventResponse, error) {
	response := new(IMouseputMouseEventResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMouseputMouseEventAbsolute(request *IMouseputMouseEventAbsolute) (*IMouseputMouseEventAbsoluteResponse, error) {
	response := new(IMouseputMouseEventAbsoluteResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IDisplaygetScreenResolution(request *IDisplaygetScreenResolution) (*IDisplaygetScreenResolutionResponse, error) {
	response := new(IDisplaygetScreenResolutionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IDisplaysetVideoModeHint(request *IDisplaysetVideoModeHint) (*IDisplaysetVideoModeHintResponse, error) {
	response := new(IDisplaysetVideoModeHintResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IDisplaysetSeamlessMode(request *IDisplaysetSeamlessMode) (*IDisplaysetSeamlessModeResponse, error) {
	response := new(IDisplaysetSeamlessModeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IDisplaytakeScreenShotToArray(request *IDisplaytakeScreenShotToArray) (*IDisplaytakeScreenShotToArrayResponse, error) {
	response := new(IDisplaytakeScreenShotToArrayResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IDisplayinvalidateAndUpdate(request *IDisplayinvalidateAndUpdate) (*IDisplayinvalidateAndUpdateResponse, error) {
	response := new(IDisplayinvalidateAndUpdateResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IDisplayresizeCompleted(request *IDisplayresizeCompleted) (*IDisplayresizeCompletedResponse, error) {
	response := new(IDisplayresizeCompletedResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INetworkAdaptergetAdapterType(request *INetworkAdaptergetAdapterType) (*INetworkAdaptergetAdapterTypeResponse, error) {
	response := new(INetworkAdaptergetAdapterTypeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INetworkAdaptersetAdapterType(request *INetworkAdaptersetAdapterType) (*INetworkAdaptersetAdapterTypeResponse, error) {
	response := new(INetworkAdaptersetAdapterTypeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INetworkAdaptergetSlot(request *INetworkAdaptergetSlot) (*INetworkAdaptergetSlotResponse, error) {
	response := new(INetworkAdaptergetSlotResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INetworkAdaptergetEnabled(request *INetworkAdaptergetEnabled) (*INetworkAdaptergetEnabledResponse, error) {
	response := new(INetworkAdaptergetEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INetworkAdaptersetEnabled(request *INetworkAdaptersetEnabled) (*INetworkAdaptersetEnabledResponse, error) {
	response := new(INetworkAdaptersetEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INetworkAdaptergetMACAddress(request *INetworkAdaptergetMACAddress) (*INetworkAdaptergetMACAddressResponse, error) {
	response := new(INetworkAdaptergetMACAddressResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INetworkAdaptersetMACAddress(request *INetworkAdaptersetMACAddress) (*INetworkAdaptersetMACAddressResponse, error) {
	response := new(INetworkAdaptersetMACAddressResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INetworkAdaptergetAttachmentType(request *INetworkAdaptergetAttachmentType) (*INetworkAdaptergetAttachmentTypeResponse, error) {
	response := new(INetworkAdaptergetAttachmentTypeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INetworkAdaptergetHostInterface(request *INetworkAdaptergetHostInterface) (*INetworkAdaptergetHostInterfaceResponse, error) {
	response := new(INetworkAdaptergetHostInterfaceResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INetworkAdaptersetHostInterface(request *INetworkAdaptersetHostInterface) (*INetworkAdaptersetHostInterfaceResponse, error) {
	response := new(INetworkAdaptersetHostInterfaceResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INetworkAdaptergetInternalNetwork(request *INetworkAdaptergetInternalNetwork) (*INetworkAdaptergetInternalNetworkResponse, error) {
	response := new(INetworkAdaptergetInternalNetworkResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INetworkAdaptersetInternalNetwork(request *INetworkAdaptersetInternalNetwork) (*INetworkAdaptersetInternalNetworkResponse, error) {
	response := new(INetworkAdaptersetInternalNetworkResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INetworkAdaptergetNATNetwork(request *INetworkAdaptergetNATNetwork) (*INetworkAdaptergetNATNetworkResponse, error) {
	response := new(INetworkAdaptergetNATNetworkResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INetworkAdaptersetNATNetwork(request *INetworkAdaptersetNATNetwork) (*INetworkAdaptersetNATNetworkResponse, error) {
	response := new(INetworkAdaptersetNATNetworkResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INetworkAdaptergetVDENetwork(request *INetworkAdaptergetVDENetwork) (*INetworkAdaptergetVDENetworkResponse, error) {
	response := new(INetworkAdaptergetVDENetworkResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INetworkAdaptersetVDENetwork(request *INetworkAdaptersetVDENetwork) (*INetworkAdaptersetVDENetworkResponse, error) {
	response := new(INetworkAdaptersetVDENetworkResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INetworkAdaptergetCableConnected(request *INetworkAdaptergetCableConnected) (*INetworkAdaptergetCableConnectedResponse, error) {
	response := new(INetworkAdaptergetCableConnectedResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INetworkAdaptersetCableConnected(request *INetworkAdaptersetCableConnected) (*INetworkAdaptersetCableConnectedResponse, error) {
	response := new(INetworkAdaptersetCableConnectedResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INetworkAdaptergetLineSpeed(request *INetworkAdaptergetLineSpeed) (*INetworkAdaptergetLineSpeedResponse, error) {
	response := new(INetworkAdaptergetLineSpeedResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INetworkAdaptersetLineSpeed(request *INetworkAdaptersetLineSpeed) (*INetworkAdaptersetLineSpeedResponse, error) {
	response := new(INetworkAdaptersetLineSpeedResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INetworkAdaptergetTraceEnabled(request *INetworkAdaptergetTraceEnabled) (*INetworkAdaptergetTraceEnabledResponse, error) {
	response := new(INetworkAdaptergetTraceEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INetworkAdaptersetTraceEnabled(request *INetworkAdaptersetTraceEnabled) (*INetworkAdaptersetTraceEnabledResponse, error) {
	response := new(INetworkAdaptersetTraceEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INetworkAdaptergetTraceFile(request *INetworkAdaptergetTraceFile) (*INetworkAdaptergetTraceFileResponse, error) {
	response := new(INetworkAdaptergetTraceFileResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INetworkAdaptersetTraceFile(request *INetworkAdaptersetTraceFile) (*INetworkAdaptersetTraceFileResponse, error) {
	response := new(INetworkAdaptersetTraceFileResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INetworkAdaptergetNatDriver(request *INetworkAdaptergetNatDriver) (*INetworkAdaptergetNatDriverResponse, error) {
	response := new(INetworkAdaptergetNatDriverResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INetworkAdaptergetBootPriority(request *INetworkAdaptergetBootPriority) (*INetworkAdaptergetBootPriorityResponse, error) {
	response := new(INetworkAdaptergetBootPriorityResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INetworkAdaptersetBootPriority(request *INetworkAdaptersetBootPriority) (*INetworkAdaptersetBootPriorityResponse, error) {
	response := new(INetworkAdaptersetBootPriorityResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INetworkAdapterattachToNAT(request *INetworkAdapterattachToNAT) (*INetworkAdapterattachToNATResponse, error) {
	response := new(INetworkAdapterattachToNATResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INetworkAdapterattachToBridgedInterface(request *INetworkAdapterattachToBridgedInterface) (*INetworkAdapterattachToBridgedInterfaceResponse, error) {
	response := new(INetworkAdapterattachToBridgedInterfaceResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INetworkAdapterattachToInternalNetwork(request *INetworkAdapterattachToInternalNetwork) (*INetworkAdapterattachToInternalNetworkResponse, error) {
	response := new(INetworkAdapterattachToInternalNetworkResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INetworkAdapterattachToHostOnlyInterface(request *INetworkAdapterattachToHostOnlyInterface) (*INetworkAdapterattachToHostOnlyInterfaceResponse, error) {
	response := new(INetworkAdapterattachToHostOnlyInterfaceResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INetworkAdapterattachToVDE(request *INetworkAdapterattachToVDE) (*INetworkAdapterattachToVDEResponse, error) {
	response := new(INetworkAdapterattachToVDEResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INetworkAdapterdetach(request *INetworkAdapterdetach) (*INetworkAdapterdetachResponse, error) {
	response := new(INetworkAdapterdetachResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISerialPortgetSlot(request *ISerialPortgetSlot) (*ISerialPortgetSlotResponse, error) {
	response := new(ISerialPortgetSlotResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISerialPortgetEnabled(request *ISerialPortgetEnabled) (*ISerialPortgetEnabledResponse, error) {
	response := new(ISerialPortgetEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISerialPortsetEnabled(request *ISerialPortsetEnabled) (*ISerialPortsetEnabledResponse, error) {
	response := new(ISerialPortsetEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISerialPortgetIOBase(request *ISerialPortgetIOBase) (*ISerialPortgetIOBaseResponse, error) {
	response := new(ISerialPortgetIOBaseResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISerialPortsetIOBase(request *ISerialPortsetIOBase) (*ISerialPortsetIOBaseResponse, error) {
	response := new(ISerialPortsetIOBaseResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISerialPortgetIRQ(request *ISerialPortgetIRQ) (*ISerialPortgetIRQResponse, error) {
	response := new(ISerialPortgetIRQResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISerialPortsetIRQ(request *ISerialPortsetIRQ) (*ISerialPortsetIRQResponse, error) {
	response := new(ISerialPortsetIRQResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISerialPortgetHostMode(request *ISerialPortgetHostMode) (*ISerialPortgetHostModeResponse, error) {
	response := new(ISerialPortgetHostModeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISerialPortsetHostMode(request *ISerialPortsetHostMode) (*ISerialPortsetHostModeResponse, error) {
	response := new(ISerialPortsetHostModeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISerialPortgetServer(request *ISerialPortgetServer) (*ISerialPortgetServerResponse, error) {
	response := new(ISerialPortgetServerResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISerialPortsetServer(request *ISerialPortsetServer) (*ISerialPortsetServerResponse, error) {
	response := new(ISerialPortsetServerResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISerialPortgetPath(request *ISerialPortgetPath) (*ISerialPortgetPathResponse, error) {
	response := new(ISerialPortgetPathResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISerialPortsetPath(request *ISerialPortsetPath) (*ISerialPortsetPathResponse, error) {
	response := new(ISerialPortsetPathResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IParallelPortgetSlot(request *IParallelPortgetSlot) (*IParallelPortgetSlotResponse, error) {
	response := new(IParallelPortgetSlotResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IParallelPortgetEnabled(request *IParallelPortgetEnabled) (*IParallelPortgetEnabledResponse, error) {
	response := new(IParallelPortgetEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IParallelPortsetEnabled(request *IParallelPortsetEnabled) (*IParallelPortsetEnabledResponse, error) {
	response := new(IParallelPortsetEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IParallelPortgetIOBase(request *IParallelPortgetIOBase) (*IParallelPortgetIOBaseResponse, error) {
	response := new(IParallelPortgetIOBaseResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IParallelPortsetIOBase(request *IParallelPortsetIOBase) (*IParallelPortsetIOBaseResponse, error) {
	response := new(IParallelPortsetIOBaseResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IParallelPortgetIRQ(request *IParallelPortgetIRQ) (*IParallelPortgetIRQResponse, error) {
	response := new(IParallelPortgetIRQResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IParallelPortsetIRQ(request *IParallelPortsetIRQ) (*IParallelPortsetIRQResponse, error) {
	response := new(IParallelPortsetIRQResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IParallelPortgetPath(request *IParallelPortgetPath) (*IParallelPortgetPathResponse, error) {
	response := new(IParallelPortgetPathResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IParallelPortsetPath(request *IParallelPortsetPath) (*IParallelPortsetPathResponse, error) {
	response := new(IParallelPortsetPathResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IUSBControllergetEnabled(request *IUSBControllergetEnabled) (*IUSBControllergetEnabledResponse, error) {
	response := new(IUSBControllergetEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IUSBControllersetEnabled(request *IUSBControllersetEnabled) (*IUSBControllersetEnabledResponse, error) {
	response := new(IUSBControllersetEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IUSBControllergetEnabledEhci(request *IUSBControllergetEnabledEhci) (*IUSBControllergetEnabledEhciResponse, error) {
	response := new(IUSBControllergetEnabledEhciResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IUSBControllersetEnabledEhci(request *IUSBControllersetEnabledEhci) (*IUSBControllersetEnabledEhciResponse, error) {
	response := new(IUSBControllersetEnabledEhciResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IUSBControllergetProxyAvailable(request *IUSBControllergetProxyAvailable) (*IUSBControllergetProxyAvailableResponse, error) {
	response := new(IUSBControllergetProxyAvailableResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IUSBControllergetUSBStandard(request *IUSBControllergetUSBStandard) (*IUSBControllergetUSBStandardResponse, error) {
	response := new(IUSBControllergetUSBStandardResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IUSBControllergetDeviceFilters(request *IUSBControllergetDeviceFilters) (*IUSBControllergetDeviceFiltersResponse, error) {
	response := new(IUSBControllergetDeviceFiltersResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IUSBControllercreateDeviceFilter(request *IUSBControllercreateDeviceFilter) (*IUSBControllercreateDeviceFilterResponse, error) {
	response := new(IUSBControllercreateDeviceFilterResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IUSBControllerinsertDeviceFilter(request *IUSBControllerinsertDeviceFilter) (*IUSBControllerinsertDeviceFilterResponse, error) {
	response := new(IUSBControllerinsertDeviceFilterResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IUSBControllerremoveDeviceFilter(request *IUSBControllerremoveDeviceFilter) (*IUSBControllerremoveDeviceFilterResponse, error) {
	response := new(IUSBControllerremoveDeviceFilterResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IUSBDevicegetId(request *IUSBDevicegetId) (*IUSBDevicegetIdResponse, error) {
	response := new(IUSBDevicegetIdResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IUSBDevicegetVendorId(request *IUSBDevicegetVendorId) (*IUSBDevicegetVendorIdResponse, error) {
	response := new(IUSBDevicegetVendorIdResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IUSBDevicegetProductId(request *IUSBDevicegetProductId) (*IUSBDevicegetProductIdResponse, error) {
	response := new(IUSBDevicegetProductIdResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IUSBDevicegetRevision(request *IUSBDevicegetRevision) (*IUSBDevicegetRevisionResponse, error) {
	response := new(IUSBDevicegetRevisionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IUSBDevicegetManufacturer(request *IUSBDevicegetManufacturer) (*IUSBDevicegetManufacturerResponse, error) {
	response := new(IUSBDevicegetManufacturerResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IUSBDevicegetProduct(request *IUSBDevicegetProduct) (*IUSBDevicegetProductResponse, error) {
	response := new(IUSBDevicegetProductResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IUSBDevicegetSerialNumber(request *IUSBDevicegetSerialNumber) (*IUSBDevicegetSerialNumberResponse, error) {
	response := new(IUSBDevicegetSerialNumberResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IUSBDevicegetAddress(request *IUSBDevicegetAddress) (*IUSBDevicegetAddressResponse, error) {
	response := new(IUSBDevicegetAddressResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IUSBDevicegetPort(request *IUSBDevicegetPort) (*IUSBDevicegetPortResponse, error) {
	response := new(IUSBDevicegetPortResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IUSBDevicegetVersion(request *IUSBDevicegetVersion) (*IUSBDevicegetVersionResponse, error) {
	response := new(IUSBDevicegetVersionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IUSBDevicegetPortVersion(request *IUSBDevicegetPortVersion) (*IUSBDevicegetPortVersionResponse, error) {
	response := new(IUSBDevicegetPortVersionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IUSBDevicegetRemote(request *IUSBDevicegetRemote) (*IUSBDevicegetRemoteResponse, error) {
	response := new(IUSBDevicegetRemoteResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IUSBDeviceFiltergetName(request *IUSBDeviceFiltergetName) (*IUSBDeviceFiltergetNameResponse, error) {
	response := new(IUSBDeviceFiltergetNameResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IUSBDeviceFiltersetName(request *IUSBDeviceFiltersetName) (*IUSBDeviceFiltersetNameResponse, error) {
	response := new(IUSBDeviceFiltersetNameResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IUSBDeviceFiltergetActive(request *IUSBDeviceFiltergetActive) (*IUSBDeviceFiltergetActiveResponse, error) {
	response := new(IUSBDeviceFiltergetActiveResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IUSBDeviceFiltersetActive(request *IUSBDeviceFiltersetActive) (*IUSBDeviceFiltersetActiveResponse, error) {
	response := new(IUSBDeviceFiltersetActiveResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IUSBDeviceFiltergetVendorId(request *IUSBDeviceFiltergetVendorId) (*IUSBDeviceFiltergetVendorIdResponse, error) {
	response := new(IUSBDeviceFiltergetVendorIdResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IUSBDeviceFiltersetVendorId(request *IUSBDeviceFiltersetVendorId) (*IUSBDeviceFiltersetVendorIdResponse, error) {
	response := new(IUSBDeviceFiltersetVendorIdResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IUSBDeviceFiltergetProductId(request *IUSBDeviceFiltergetProductId) (*IUSBDeviceFiltergetProductIdResponse, error) {
	response := new(IUSBDeviceFiltergetProductIdResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IUSBDeviceFiltersetProductId(request *IUSBDeviceFiltersetProductId) (*IUSBDeviceFiltersetProductIdResponse, error) {
	response := new(IUSBDeviceFiltersetProductIdResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IUSBDeviceFiltergetRevision(request *IUSBDeviceFiltergetRevision) (*IUSBDeviceFiltergetRevisionResponse, error) {
	response := new(IUSBDeviceFiltergetRevisionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IUSBDeviceFiltersetRevision(request *IUSBDeviceFiltersetRevision) (*IUSBDeviceFiltersetRevisionResponse, error) {
	response := new(IUSBDeviceFiltersetRevisionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IUSBDeviceFiltergetManufacturer(request *IUSBDeviceFiltergetManufacturer) (*IUSBDeviceFiltergetManufacturerResponse, error) {
	response := new(IUSBDeviceFiltergetManufacturerResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IUSBDeviceFiltersetManufacturer(request *IUSBDeviceFiltersetManufacturer) (*IUSBDeviceFiltersetManufacturerResponse, error) {
	response := new(IUSBDeviceFiltersetManufacturerResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IUSBDeviceFiltergetProduct(request *IUSBDeviceFiltergetProduct) (*IUSBDeviceFiltergetProductResponse, error) {
	response := new(IUSBDeviceFiltergetProductResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IUSBDeviceFiltersetProduct(request *IUSBDeviceFiltersetProduct) (*IUSBDeviceFiltersetProductResponse, error) {
	response := new(IUSBDeviceFiltersetProductResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IUSBDeviceFiltergetSerialNumber(request *IUSBDeviceFiltergetSerialNumber) (*IUSBDeviceFiltergetSerialNumberResponse, error) {
	response := new(IUSBDeviceFiltergetSerialNumberResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IUSBDeviceFiltersetSerialNumber(request *IUSBDeviceFiltersetSerialNumber) (*IUSBDeviceFiltersetSerialNumberResponse, error) {
	response := new(IUSBDeviceFiltersetSerialNumberResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IUSBDeviceFiltergetPort(request *IUSBDeviceFiltergetPort) (*IUSBDeviceFiltergetPortResponse, error) {
	response := new(IUSBDeviceFiltergetPortResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IUSBDeviceFiltersetPort(request *IUSBDeviceFiltersetPort) (*IUSBDeviceFiltersetPortResponse, error) {
	response := new(IUSBDeviceFiltersetPortResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IUSBDeviceFiltergetRemote(request *IUSBDeviceFiltergetRemote) (*IUSBDeviceFiltergetRemoteResponse, error) {
	response := new(IUSBDeviceFiltergetRemoteResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IUSBDeviceFiltersetRemote(request *IUSBDeviceFiltersetRemote) (*IUSBDeviceFiltersetRemoteResponse, error) {
	response := new(IUSBDeviceFiltersetRemoteResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IUSBDeviceFiltergetMaskedInterfaces(request *IUSBDeviceFiltergetMaskedInterfaces) (*IUSBDeviceFiltergetMaskedInterfacesResponse, error) {
	response := new(IUSBDeviceFiltergetMaskedInterfacesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IUSBDeviceFiltersetMaskedInterfaces(request *IUSBDeviceFiltersetMaskedInterfaces) (*IUSBDeviceFiltersetMaskedInterfacesResponse, error) {
	response := new(IUSBDeviceFiltersetMaskedInterfacesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostUSBDevicegetState(request *IHostUSBDevicegetState) (*IHostUSBDevicegetStateResponse, error) {
	response := new(IHostUSBDevicegetStateResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostUSBDeviceFiltergetAction(request *IHostUSBDeviceFiltergetAction) (*IHostUSBDeviceFiltergetActionResponse, error) {
	response := new(IHostUSBDeviceFiltergetActionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostUSBDeviceFiltersetAction(request *IHostUSBDeviceFiltersetAction) (*IHostUSBDeviceFiltersetActionResponse, error) {
	response := new(IHostUSBDeviceFiltersetActionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IAudioAdaptergetEnabled(request *IAudioAdaptergetEnabled) (*IAudioAdaptergetEnabledResponse, error) {
	response := new(IAudioAdaptergetEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IAudioAdaptersetEnabled(request *IAudioAdaptersetEnabled) (*IAudioAdaptersetEnabledResponse, error) {
	response := new(IAudioAdaptersetEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IAudioAdaptergetAudioController(request *IAudioAdaptergetAudioController) (*IAudioAdaptergetAudioControllerResponse, error) {
	response := new(IAudioAdaptergetAudioControllerResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IAudioAdaptersetAudioController(request *IAudioAdaptersetAudioController) (*IAudioAdaptersetAudioControllerResponse, error) {
	response := new(IAudioAdaptersetAudioControllerResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IAudioAdaptergetAudioDriver(request *IAudioAdaptergetAudioDriver) (*IAudioAdaptergetAudioDriverResponse, error) {
	response := new(IAudioAdaptergetAudioDriverResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IAudioAdaptersetAudioDriver(request *IAudioAdaptersetAudioDriver) (*IAudioAdaptersetAudioDriverResponse, error) {
	response := new(IAudioAdaptersetAudioDriverResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVRDPServergetEnabled(request *IVRDPServergetEnabled) (*IVRDPServergetEnabledResponse, error) {
	response := new(IVRDPServergetEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVRDPServersetEnabled(request *IVRDPServersetEnabled) (*IVRDPServersetEnabledResponse, error) {
	response := new(IVRDPServersetEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVRDPServergetPorts(request *IVRDPServergetPorts) (*IVRDPServergetPortsResponse, error) {
	response := new(IVRDPServergetPortsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVRDPServersetPorts(request *IVRDPServersetPorts) (*IVRDPServersetPortsResponse, error) {
	response := new(IVRDPServersetPortsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVRDPServergetNetAddress(request *IVRDPServergetNetAddress) (*IVRDPServergetNetAddressResponse, error) {
	response := new(IVRDPServergetNetAddressResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVRDPServersetNetAddress(request *IVRDPServersetNetAddress) (*IVRDPServersetNetAddressResponse, error) {
	response := new(IVRDPServersetNetAddressResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVRDPServergetAuthType(request *IVRDPServergetAuthType) (*IVRDPServergetAuthTypeResponse, error) {
	response := new(IVRDPServergetAuthTypeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVRDPServersetAuthType(request *IVRDPServersetAuthType) (*IVRDPServersetAuthTypeResponse, error) {
	response := new(IVRDPServersetAuthTypeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVRDPServergetAuthTimeout(request *IVRDPServergetAuthTimeout) (*IVRDPServergetAuthTimeoutResponse, error) {
	response := new(IVRDPServergetAuthTimeoutResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVRDPServersetAuthTimeout(request *IVRDPServersetAuthTimeout) (*IVRDPServersetAuthTimeoutResponse, error) {
	response := new(IVRDPServersetAuthTimeoutResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVRDPServergetAllowMultiConnection(request *IVRDPServergetAllowMultiConnection) (*IVRDPServergetAllowMultiConnectionResponse, error) {
	response := new(IVRDPServergetAllowMultiConnectionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVRDPServersetAllowMultiConnection(request *IVRDPServersetAllowMultiConnection) (*IVRDPServersetAllowMultiConnectionResponse, error) {
	response := new(IVRDPServersetAllowMultiConnectionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVRDPServergetReuseSingleConnection(request *IVRDPServergetReuseSingleConnection) (*IVRDPServergetReuseSingleConnectionResponse, error) {
	response := new(IVRDPServergetReuseSingleConnectionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVRDPServersetReuseSingleConnection(request *IVRDPServersetReuseSingleConnection) (*IVRDPServersetReuseSingleConnectionResponse, error) {
	response := new(IVRDPServersetReuseSingleConnectionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVRDPServergetVideoChannel(request *IVRDPServergetVideoChannel) (*IVRDPServergetVideoChannelResponse, error) {
	response := new(IVRDPServergetVideoChannelResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVRDPServersetVideoChannel(request *IVRDPServersetVideoChannel) (*IVRDPServersetVideoChannelResponse, error) {
	response := new(IVRDPServersetVideoChannelResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVRDPServergetVideoChannelQuality(request *IVRDPServergetVideoChannelQuality) (*IVRDPServergetVideoChannelQualityResponse, error) {
	response := new(IVRDPServergetVideoChannelQualityResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVRDPServersetVideoChannelQuality(request *IVRDPServersetVideoChannelQuality) (*IVRDPServersetVideoChannelQualityResponse, error) {
	response := new(IVRDPServersetVideoChannelQualityResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISessiongetState(request *ISessiongetState) (*ISessiongetStateResponse, error) {
	response := new(ISessiongetStateResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISessiongetType(request *ISessiongetType) (*ISessiongetTypeResponse, error) {
	response := new(ISessiongetTypeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISessiongetMachine(request *ISessiongetMachine) (*ISessiongetMachineResponse, error) {
	response := new(ISessiongetMachineResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISessiongetConsole(request *ISessiongetConsole) (*ISessiongetConsoleResponse, error) {
	response := new(ISessiongetConsoleResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISessionclose(request *ISessionclose) (*ISessioncloseResponse, error) {
	response := new(ISessioncloseResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IStorageControllergetName(request *IStorageControllergetName) (*IStorageControllergetNameResponse, error) {
	response := new(IStorageControllergetNameResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IStorageControllergetMaxDevicesPerPortCount(request *IStorageControllergetMaxDevicesPerPortCount) (*IStorageControllergetMaxDevicesPerPortCountResponse, error) {
	response := new(IStorageControllergetMaxDevicesPerPortCountResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IStorageControllergetMinPortCount(request *IStorageControllergetMinPortCount) (*IStorageControllergetMinPortCountResponse, error) {
	response := new(IStorageControllergetMinPortCountResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IStorageControllergetMaxPortCount(request *IStorageControllergetMaxPortCount) (*IStorageControllergetMaxPortCountResponse, error) {
	response := new(IStorageControllergetMaxPortCountResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IStorageControllergetInstance(request *IStorageControllergetInstance) (*IStorageControllergetInstanceResponse, error) {
	response := new(IStorageControllergetInstanceResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IStorageControllersetInstance(request *IStorageControllersetInstance) (*IStorageControllersetInstanceResponse, error) {
	response := new(IStorageControllersetInstanceResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IStorageControllergetPortCount(request *IStorageControllergetPortCount) (*IStorageControllergetPortCountResponse, error) {
	response := new(IStorageControllergetPortCountResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IStorageControllersetPortCount(request *IStorageControllersetPortCount) (*IStorageControllersetPortCountResponse, error) {
	response := new(IStorageControllersetPortCountResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IStorageControllergetBus(request *IStorageControllergetBus) (*IStorageControllergetBusResponse, error) {
	response := new(IStorageControllergetBusResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IStorageControllergetControllerType(request *IStorageControllergetControllerType) (*IStorageControllergetControllerTypeResponse, error) {
	response := new(IStorageControllergetControllerTypeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IStorageControllersetControllerType(request *IStorageControllersetControllerType) (*IStorageControllersetControllerTypeResponse, error) {
	response := new(IStorageControllersetControllerTypeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IStorageControllergetUseHostIOCache(request *IStorageControllergetUseHostIOCache) (*IStorageControllergetUseHostIOCacheResponse, error) {
	response := new(IStorageControllergetUseHostIOCacheResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IStorageControllersetUseHostIOCache(request *IStorageControllersetUseHostIOCache) (*IStorageControllersetUseHostIOCacheResponse, error) {
	response := new(IStorageControllersetUseHostIOCacheResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IStorageControllergetIDEEmulationPort(request *IStorageControllergetIDEEmulationPort) (*IStorageControllergetIDEEmulationPortResponse, error) {
	response := new(IStorageControllergetIDEEmulationPortResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IStorageControllersetIDEEmulationPort(request *IStorageControllersetIDEEmulationPort) (*IStorageControllersetIDEEmulationPortResponse, error) {
	response := new(IStorageControllersetIDEEmulationPortResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IManagedObjectRefgetInterfaceName(request *IManagedObjectRefgetInterfaceName) (*IManagedObjectRefgetInterfaceNameResponse, error) {
	response := new(IManagedObjectRefgetInterfaceNameResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IManagedObjectRefrelease(request *IManagedObjectRefrelease) (*IManagedObjectRefreleaseResponse, error) {
	response := new(IManagedObjectRefreleaseResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IWebsessionManagerlogon(request *IWebsessionManagerlogon) (*IWebsessionManagerlogonResponse, error) {
	response := new(IWebsessionManagerlogonResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IWebsessionManagergetSessionObject(request *IWebsessionManagergetSessionObject) (*IWebsessionManagergetSessionObjectResponse, error) {
	response := new(IWebsessionManagergetSessionObjectResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IWebsessionManagerlogoff(request *IWebsessionManagerlogoff) (*IWebsessionManagerlogoffResponse, error) {
	response := new(IWebsessionManagerlogoffResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IPerformanceMetricgetMetricName(request *IPerformanceMetricgetMetricName) (*IPerformanceMetricgetMetricNameResponse, error) {
	response := new(IPerformanceMetricgetMetricNameResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IPerformanceMetricgetObject(request *IPerformanceMetricgetObject) (*IPerformanceMetricgetObjectResponse, error) {
	response := new(IPerformanceMetricgetObjectResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IPerformanceMetricgetDescription(request *IPerformanceMetricgetDescription) (*IPerformanceMetricgetDescriptionResponse, error) {
	response := new(IPerformanceMetricgetDescriptionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IPerformanceMetricgetPeriod(request *IPerformanceMetricgetPeriod) (*IPerformanceMetricgetPeriodResponse, error) {
	response := new(IPerformanceMetricgetPeriodResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IPerformanceMetricgetCount(request *IPerformanceMetricgetCount) (*IPerformanceMetricgetCountResponse, error) {
	response := new(IPerformanceMetricgetCountResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IPerformanceMetricgetUnit(request *IPerformanceMetricgetUnit) (*IPerformanceMetricgetUnitResponse, error) {
	response := new(IPerformanceMetricgetUnitResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IPerformanceMetricgetMinimumValue(request *IPerformanceMetricgetMinimumValue) (*IPerformanceMetricgetMinimumValueResponse, error) {
	response := new(IPerformanceMetricgetMinimumValueResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IPerformanceMetricgetMaximumValue(request *IPerformanceMetricgetMaximumValue) (*IPerformanceMetricgetMaximumValueResponse, error) {
	response := new(IPerformanceMetricgetMaximumValueResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IPerformanceCollectorgetMetricNames(request *IPerformanceCollectorgetMetricNames) (*IPerformanceCollectorgetMetricNamesResponse, error) {
	response := new(IPerformanceCollectorgetMetricNamesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IPerformanceCollectorgetMetrics(request *IPerformanceCollectorgetMetrics) (*IPerformanceCollectorgetMetricsResponse, error) {
	response := new(IPerformanceCollectorgetMetricsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IPerformanceCollectorsetupMetrics(request *IPerformanceCollectorsetupMetrics) (*IPerformanceCollectorsetupMetricsResponse, error) {
	response := new(IPerformanceCollectorsetupMetricsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IPerformanceCollectorenableMetrics(request *IPerformanceCollectorenableMetrics) (*IPerformanceCollectorenableMetricsResponse, error) {
	response := new(IPerformanceCollectorenableMetricsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IPerformanceCollectordisableMetrics(request *IPerformanceCollectordisableMetrics) (*IPerformanceCollectordisableMetricsResponse, error) {
	response := new(IPerformanceCollectordisableMetricsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IPerformanceCollectorqueryMetricsData(request *IPerformanceCollectorqueryMetricsData) (*IPerformanceCollectorqueryMetricsDataResponse, error) {
	response := new(IPerformanceCollectorqueryMetricsDataResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATEnginegetNetwork(request *INATEnginegetNetwork) (*INATEnginegetNetworkResponse, error) {
	response := new(INATEnginegetNetworkResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATEnginesetNetwork(request *INATEnginesetNetwork) (*INATEnginesetNetworkResponse, error) {
	response := new(INATEnginesetNetworkResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATEnginegetHostIP(request *INATEnginegetHostIP) (*INATEnginegetHostIPResponse, error) {
	response := new(INATEnginegetHostIPResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATEnginesetHostIP(request *INATEnginesetHostIP) (*INATEnginesetHostIPResponse, error) {
	response := new(INATEnginesetHostIPResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATEnginegetTftpPrefix(request *INATEnginegetTftpPrefix) (*INATEnginegetTftpPrefixResponse, error) {
	response := new(INATEnginegetTftpPrefixResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATEnginesetTftpPrefix(request *INATEnginesetTftpPrefix) (*INATEnginesetTftpPrefixResponse, error) {
	response := new(INATEnginesetTftpPrefixResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATEnginegetTftpBootFile(request *INATEnginegetTftpBootFile) (*INATEnginegetTftpBootFileResponse, error) {
	response := new(INATEnginegetTftpBootFileResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATEnginesetTftpBootFile(request *INATEnginesetTftpBootFile) (*INATEnginesetTftpBootFileResponse, error) {
	response := new(INATEnginesetTftpBootFileResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATEnginegetTftpNextServer(request *INATEnginegetTftpNextServer) (*INATEnginegetTftpNextServerResponse, error) {
	response := new(INATEnginegetTftpNextServerResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATEnginesetTftpNextServer(request *INATEnginesetTftpNextServer) (*INATEnginesetTftpNextServerResponse, error) {
	response := new(INATEnginesetTftpNextServerResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATEnginegetAliasMode(request *INATEnginegetAliasMode) (*INATEnginegetAliasModeResponse, error) {
	response := new(INATEnginegetAliasModeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATEnginesetAliasMode(request *INATEnginesetAliasMode) (*INATEnginesetAliasModeResponse, error) {
	response := new(INATEnginesetAliasModeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATEnginegetDnsPassDomain(request *INATEnginegetDnsPassDomain) (*INATEnginegetDnsPassDomainResponse, error) {
	response := new(INATEnginegetDnsPassDomainResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATEnginesetDnsPassDomain(request *INATEnginesetDnsPassDomain) (*INATEnginesetDnsPassDomainResponse, error) {
	response := new(INATEnginesetDnsPassDomainResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATEnginegetDnsProxy(request *INATEnginegetDnsProxy) (*INATEnginegetDnsProxyResponse, error) {
	response := new(INATEnginegetDnsProxyResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATEnginesetDnsProxy(request *INATEnginesetDnsProxy) (*INATEnginesetDnsProxyResponse, error) {
	response := new(INATEnginesetDnsProxyResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATEnginegetDnsUseHostResolver(request *INATEnginegetDnsUseHostResolver) (*INATEnginegetDnsUseHostResolverResponse, error) {
	response := new(INATEnginegetDnsUseHostResolverResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATEnginesetDnsUseHostResolver(request *INATEnginesetDnsUseHostResolver) (*INATEnginesetDnsUseHostResolverResponse, error) {
	response := new(INATEnginesetDnsUseHostResolverResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATEnginegetRedirects(request *INATEnginegetRedirects) (*INATEnginegetRedirectsResponse, error) {
	response := new(INATEnginegetRedirectsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATEnginesetNetworkSettings(request *INATEnginesetNetworkSettings) (*INATEnginesetNetworkSettingsResponse, error) {
	response := new(INATEnginesetNetworkSettingsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATEnginegetNetworkSettings(request *INATEnginegetNetworkSettings) (*INATEnginegetNetworkSettingsResponse, error) {
	response := new(INATEnginegetNetworkSettingsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATEngineaddRedirect(request *INATEngineaddRedirect) (*INATEngineaddRedirectResponse, error) {
	response := new(INATEngineaddRedirectResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATEngineremoveRedirect(request *INATEngineremoveRedirect) (*INATEngineremoveRedirectResponse, error) {
	response := new(INATEngineremoveRedirectResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

var timeout = time.Duration(30 * time.Second)

func dialTimeout(network, addr string) (net.Conn, error) {
	return net.DialTimeout(network, addr, timeout)
}

type SOAPEnvelope struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`

	Body SOAPBody
}

type SOAPHeader struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Header"`

	Header interface{}
}

type SOAPBody struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body"`

	Fault   *SOAPFault  `xml:",omitempty"`
	Content interface{} `xml:",omitempty"`
}

type SOAPFault struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault"`

	Code   string `xml:"faultcode,omitempty"`
	String string `xml:"faultstring,omitempty"`
	Actor  string `xml:"faultactor,omitempty"`
	Detail string `xml:"detail,omitempty"`
}

type BasicAuth struct {
	Login    string
	Password string
}

type SOAPClient struct {
	url  string
	tls  bool
	auth *BasicAuth
}

func (b *SOAPBody) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	if b.Content == nil {
		return xml.UnmarshalError("Content must be a pointer to a struct")
	}

	var (
		token    xml.Token
		err      error
		consumed bool
	)

Loop:
	for {
		if token, err = d.Token(); err != nil {
			return err
		}

		if token == nil {
			break
		}

		switch se := token.(type) {
		case xml.StartElement:
			if consumed {
				return xml.UnmarshalError("Found multiple elements inside SOAP body; not wrapped-document/literal WS-I compliant")
			} else if se.Name.Space == "http://schemas.xmlsoap.org/soap/envelope/" && se.Name.Local == "Fault" {
				b.Fault = &SOAPFault{}
				b.Content = nil

				err = d.DecodeElement(b.Fault, &se)
				if err != nil {
					return err
				}

				consumed = true
			} else {
				if err = d.DecodeElement(b.Content, &se); err != nil {
					return err
				}

				consumed = true
			}
		case xml.EndElement:
			break Loop
		}
	}

	return nil
}

func (f *SOAPFault) Error() string {
	return f.String
}

func NewSOAPClient(url string, tls bool, auth *BasicAuth) *SOAPClient {
	return &SOAPClient{
		url:  url,
		tls:  tls,
		auth: auth,
	}
}

func (s *SOAPClient) Call(soapAction string, request, response interface{}) error {
	envelope := SOAPEnvelope{
	//Header:        SoapHeader{},
	}

	envelope.Body.Content = request
	buffer := new(bytes.Buffer)

	encoder := xml.NewEncoder(buffer)
	//encoder.Indent("  ", "    ")

	if err := encoder.Encode(envelope); err != nil {
		return err
	}

	if err := encoder.Flush(); err != nil {
		return err
	}

	// log.Println(buffer.String())

	req, err := http.NewRequest("POST", s.url, buffer)
	if err != nil {
		return err
	}
	if s.auth != nil {
		req.SetBasicAuth(s.auth.Login, s.auth.Password)
	}

	req.Header.Add("Content-Type", "text/xml; charset=\"utf-8\"")
	if soapAction != "" {
		req.Header.Add("SOAPAction", soapAction)
	}

	req.Header.Set("User-Agent", "gowsdl/0.1")
	req.Close = true

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: s.tls,
		},
		Dial: dialTimeout,
	}

	client := &http.Client{Transport: tr}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	rawbody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	if len(rawbody) == 0 {
		log.Println("empty response")
		return nil
	}

	// log.Println(string(rawbody))
	respEnvelope := new(SOAPEnvelope)
	respEnvelope.Body = SOAPBody{Content: response}
	err = xml.Unmarshal(rawbody, respEnvelope)
	if err != nil {
		return err
	}

	fault := respEnvelope.Body.Fault
	if fault != nil {
		return fault
	}

	return nil
}
