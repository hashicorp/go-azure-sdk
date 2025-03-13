package incidentbookmarks

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IoTDeviceEntityProperties struct {
	AdditionalData     *interface{}          `json:"additionalData,omitempty"`
	DeviceId           *string               `json:"deviceId,omitempty"`
	DeviceName         *string               `json:"deviceName,omitempty"`
	DeviceSubType      *string               `json:"deviceSubType,omitempty"`
	DeviceType         *string               `json:"deviceType,omitempty"`
	EdgeId             *string               `json:"edgeId,omitempty"`
	FirmwareVersion    *string               `json:"firmwareVersion,omitempty"`
	FriendlyName       *string               `json:"friendlyName,omitempty"`
	HostEntityId       *string               `json:"hostEntityId,omitempty"`
	IPAddressEntityId  *string               `json:"ipAddressEntityId,omitempty"`
	Importance         *DeviceImportance     `json:"importance,omitempty"`
	IotHubEntityId     *string               `json:"iotHubEntityId,omitempty"`
	IotSecurityAgentId *string               `json:"iotSecurityAgentId,omitempty"`
	IsAuthorized       *bool                 `json:"isAuthorized,omitempty"`
	IsProgramming      *bool                 `json:"isProgramming,omitempty"`
	IsScanner          *bool                 `json:"isScanner,omitempty"`
	MacAddress         *string               `json:"macAddress,omitempty"`
	Model              *string               `json:"model,omitempty"`
	NicEntityIds       *[]string             `json:"nicEntityIds,omitempty"`
	OperatingSystem    *string               `json:"operatingSystem,omitempty"`
	Owners             *[]string             `json:"owners,omitempty"`
	Protocols          *[]string             `json:"protocols,omitempty"`
	PurdueLayer        *string               `json:"purdueLayer,omitempty"`
	Sensor             *string               `json:"sensor,omitempty"`
	SerialNumber       *string               `json:"serialNumber,omitempty"`
	Site               *string               `json:"site,omitempty"`
	Source             *string               `json:"source,omitempty"`
	ThreatIntelligence *[]ThreatIntelligence `json:"threatIntelligence,omitempty"`
	Vendor             *string               `json:"vendor,omitempty"`
	Zone               *string               `json:"zone,omitempty"`
}
