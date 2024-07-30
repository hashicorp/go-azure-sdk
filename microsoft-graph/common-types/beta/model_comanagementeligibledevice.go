package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ComanagementEligibleDevice struct {
	ClientRegistrationStatus *ComanagementEligibleDeviceClientRegistrationStatus `json:"clientRegistrationStatus,omitempty"`
	DeviceName               *string                                             `json:"deviceName,omitempty"`
	DeviceType               *ComanagementEligibleDeviceDeviceType               `json:"deviceType,omitempty"`
	EntitySource             *int64                                              `json:"entitySource,omitempty"`
	Id                       *string                                             `json:"id,omitempty"`
	ManagementAgents         *ComanagementEligibleDeviceManagementAgents         `json:"managementAgents,omitempty"`
	ManagementState          *ComanagementEligibleDeviceManagementState          `json:"managementState,omitempty"`
	Manufacturer             *string                                             `json:"manufacturer,omitempty"`
	MdmStatus                *string                                             `json:"mdmStatus,omitempty"`
	Model                    *string                                             `json:"model,omitempty"`
	ODataType                *string                                             `json:"@odata.type,omitempty"`
	OsDescription            *string                                             `json:"osDescription,omitempty"`
	OsVersion                *string                                             `json:"osVersion,omitempty"`
	OwnerType                *ComanagementEligibleDeviceOwnerType                `json:"ownerType,omitempty"`
	ReferenceId              *string                                             `json:"referenceId,omitempty"`
	SerialNumber             *string                                             `json:"serialNumber,omitempty"`
	Status                   *ComanagementEligibleDeviceStatus                   `json:"status,omitempty"`
	Upn                      *string                                             `json:"upn,omitempty"`
	UserEmail                *string                                             `json:"userEmail,omitempty"`
	UserId                   *string                                             `json:"userId,omitempty"`
	UserName                 *string                                             `json:"userName,omitempty"`
}
