package entitytypes

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IoTDeviceEntityProperties struct {
	AdditionalData     *interface{}          `json:"additionalData,omitempty"`
	DeviceId           *string               `json:"deviceId,omitempty"`
	DeviceName         *string               `json:"deviceName,omitempty"`
	DeviceType         *string               `json:"deviceType,omitempty"`
	EdgeId             *string               `json:"edgeId,omitempty"`
	FirmwareVersion    *string               `json:"firmwareVersion,omitempty"`
	FriendlyName       *string               `json:"friendlyName,omitempty"`
	HostEntityId       *string               `json:"hostEntityId,omitempty"`
	IPAddressEntityId  *string               `json:"ipAddressEntityId,omitempty"`
	IotHubEntityId     *string               `json:"iotHubEntityId,omitempty"`
	IotSecurityAgentId *string               `json:"iotSecurityAgentId,omitempty"`
	MacAddress         *string               `json:"macAddress,omitempty"`
	Model              *string               `json:"model,omitempty"`
	OperatingSystem    *string               `json:"operatingSystem,omitempty"`
	Protocols          *[]string             `json:"protocols,omitempty"`
	SerialNumber       *string               `json:"serialNumber,omitempty"`
	Source             *string               `json:"source,omitempty"`
	ThreatIntelligence *[]ThreatIntelligence `json:"threatIntelligence,omitempty"`
	Vendor             *string               `json:"vendor,omitempty"`
}
