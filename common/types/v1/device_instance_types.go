package v1

import ()

// Twin provides a logical representation of control properties (writable properties in the
// device model). The properties can have a Desired state and a Reported state. The cloud configures
// the `Desired`state of a device property and this configuration update is pushed to the edge node.
// The mapper sends a command to the device to change this property value as per the desired state .
// It receives the `Reported` state of the property once the previous operation is complete and sends
// the reported state to the cloud. Offline device interaction in the edge is possible via twin
// properties for control/command operations.
type Twin struct {
	Service string `json:"Service,omitempty"`
	// Required: The property name for which the desired/reported values are specified.
	// This property should be present in the device model.
	PropertyName string `json:"propertyName,omitempty"`
	// Required: the desired property value.
	Desired TwinProperty `json:"desired,omitempty"`
	// Required: the reported property value.
	Reported TwinProperty `json:"reported,omitempty"`
}

// TwinProperty represents the device property for which an Expected/Actual state can be defined.
type TwinProperty struct {
	// Required: The value for this property.
	Value string `json:"value,"`
	// Additional metadata like timestamp when the value was reported etc.
	// +optional
	Metadata map[string]string `json:"metadata,omitempty"`
}

// DeviceSpec represents a single device instance. It is an instantation of a device model.
type DeviceSpec struct {
	// Required: DeviceModelRef is reference to the device model used as a template
	// to create the device instance.
	//this should be the ID of device model.
	DeviceModelRef string `json:"deviceModelRef,omitempty"`

	// Required: The protocol configuration used to connect to the device.
	//this should be a json string
	Protocol string `json:"protocol,omitempty"`

	// ExtensionConfig which describe how to access the device properties,command, and events.
	// +optional
	ExtensionConfig ExtensionConfig `json:"extensionConfig,omitempty"`
}

type ExtensionConfig struct {
	// Required: List of device services.
	Services []*DeviceServiceSpec `json:"Services,omitempty"`
}

// DeviceServiceSpec is the  an instantation of a DeviceServiceModel.
type DeviceServiceSpec struct {
	Name       string                `json:"name"`
	Properties []*DevicePropertySpec `json:"properties,omitempty"`
	Events     []*DeviceEventSpec    `json:"events,omitempty"`
	Commands   []*DeviceCommandSpec  `json:"commands,omitempty"`
}

// DevicePropertySpec is an instantation of a DevicePropertyModel.
type DevicePropertySpec struct {
	*DevicePropertyModel `json:",inline"`
	// List of AccessConfig which describe how to access the device properties,command, and events.
	// AccessConfig must unique by AccessConfig.propertyName.
	// +optional
	//this should be a json string
	// AccessConfig must unique by AccessConfig.propertyName.
	AccessConfig string `json:"accessConfig,omitempty"`
}

// DeviceEventSpec is an instantation of a DeviceEventModel.
type DeviceEventSpec struct {
	*DeviceEventModel `json:",inline"`
	// List of AccessConfig which describe how to access the device properties,command, and events.
	// AccessConfig must unique by AccessConfig.propertyName.
	// +optional
	//this should be a json string
	// AccessConfig must unique by AccessConfig.propertyName.
	AccessConfig string `json:"accessConfig,omitempty"`
}

// DeviceCommandSpec is an instantation of a DeviceCommandModel.
type DeviceCommandSpec struct {
	*DeviceCommandModel `json:",inline"`
	// List of AccessConfig which describe how to access the device properties,command, and events.
	// AccessConfig must unique by AccessConfig.propertyName.
	// +optional
	//this should be a json string
	// AccessConfig must unique by AccessConfig.propertyName.
	AccessConfig string `json:"accessConfig,omitempty"`
}

// DeviceStatus reports the device state and the desired/reported values of twin attributes.
type DeviceStatus struct {
	//device status
	// online, offline, error etc.
	DeviceStatus string `json:"deviceStatus,omitempty"`
	//start/stop collecting flag.
	Collecting bool `json:"collecting,omitempty"`
	// A list of device twins containing desired/reported desired/reported values of twin properties..
	// Optional: A passive device won't have twin properties and this list could be empty.
	// +optional
	Twins []Twin `json:"twins,omitempty"`
}

//Based device info.
type DeviceBase struct {
	//Required: device name, changed by user
	DeviceName string `json:"DeviceName"`
	//Required: unique ID in global.
	DeviceID                 string `json:"DeviceID"`
	DeviceOS                 string `json:"deviceOS,omitempty"`
	DeviceCatagory           string `json:"deviceCatagory,omitempty"`
	DeviceVersion            int    `json:"deviceVersion,omitempty"`
	DeviceIdentificationCode string `json:"deviceIdentificationCode,omitempty"`

	//group.
	//TODO: reserved in future
	GroupID   string `json:"groupId,omitempty"`
	GroupName string `json:"groupName,omitempty"`
	//who create the device by ID.
	Creator         string `json:"creator,omitempty"`
	CreateTimeStamp int64  `json:"createTimeStamp,omitempty"`
	UpdateTimeStamp int64  `json:"updateTimeStamp,omitempty"`
	// +optional
	//TODO: reserved in future
	DeviceAuthType string `json:"deviceAuthType,omitempty"`
	Secret         string `json:"secret,omitempty"`

	//Device Type
	//Kind: Direct, GateWay, SubDevcie.
	DeviceType  string `json:"deviceType,omitempty"`
	GatewayID   string `json:"gatewayId,omitempty"`
	GatewayName string `json:"gatewayName,omitempty"`
	// Additional metadata like tags.
	// +optional
	Tags map[string]string `json:"tags,omitempty"`
}

// Device is the Schema for the devices API
type Device struct {
	DeviceBase `json:"base"`
	Spec       DeviceSpec   `json:"spec,omitempty"`
	Status     DeviceStatus `json:"status,omitempty"`
}

// DeviceList contains a list of Device
type DeviceList struct {
	Items []Device `json:"items"`
}
