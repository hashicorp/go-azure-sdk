package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityIoTDeviceEvidence struct {
	CreatedDateTime          *string                                     `json:"createdDateTime,omitempty"`
	DetailedRoles            *[]string                                   `json:"detailedRoles,omitempty"`
	DeviceId                 *string                                     `json:"deviceId,omitempty"`
	DeviceName               *string                                     `json:"deviceName,omitempty"`
	DevicePageLink           *string                                     `json:"devicePageLink,omitempty"`
	DeviceSubType            *string                                     `json:"deviceSubType,omitempty"`
	DeviceType               *string                                     `json:"deviceType,omitempty"`
	Importance               *SecurityIoTDeviceEvidenceImportance        `json:"importance,omitempty"`
	IoTHub                   *SecurityAzureResourceEvidence              `json:"ioTHub,omitempty"`
	IoTSecurityAgentId       *string                                     `json:"ioTSecurityAgentId,omitempty"`
	IpAddress                *SecurityIpEvidence                         `json:"ipAddress,omitempty"`
	IsAuthorized             *bool                                       `json:"isAuthorized,omitempty"`
	IsProgramming            *bool                                       `json:"isProgramming,omitempty"`
	IsScanner                *bool                                       `json:"isScanner,omitempty"`
	MacAddress               *string                                     `json:"macAddress,omitempty"`
	Manufacturer             *string                                     `json:"manufacturer,omitempty"`
	Model                    *string                                     `json:"model,omitempty"`
	Nics                     *SecurityNicEvidence                        `json:"nics,omitempty"`
	ODataType                *string                                     `json:"@odata.type,omitempty"`
	OperatingSystem          *string                                     `json:"operatingSystem,omitempty"`
	Owners                   *[]string                                   `json:"owners,omitempty"`
	Protocols                *[]string                                   `json:"protocols,omitempty"`
	PurdueLayer              *string                                     `json:"purdueLayer,omitempty"`
	RemediationStatus        *SecurityIoTDeviceEvidenceRemediationStatus `json:"remediationStatus,omitempty"`
	RemediationStatusDetails *string                                     `json:"remediationStatusDetails,omitempty"`
	Roles                    *SecurityIoTDeviceEvidenceRoles             `json:"roles,omitempty"`
	Sensor                   *string                                     `json:"sensor,omitempty"`
	SerialNumber             *string                                     `json:"serialNumber,omitempty"`
	Site                     *string                                     `json:"site,omitempty"`
	Source                   *string                                     `json:"source,omitempty"`
	SourceRef                *SecurityUrlEvidence                        `json:"sourceRef,omitempty"`
	Tags                     *[]string                                   `json:"tags,omitempty"`
	Verdict                  *SecurityIoTDeviceEvidenceVerdict           `json:"verdict,omitempty"`
	Zone                     *string                                     `json:"zone,omitempty"`
}
