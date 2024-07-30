package virtualendpointprovisioningpolicy

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsInformationProtectionDeviceRegistration struct {
	DeviceMacAddress     *string `json:"deviceMacAddress,omitempty"`
	DeviceName           *string `json:"deviceName,omitempty"`
	DeviceRegistrationId *string `json:"deviceRegistrationId,omitempty"`
	DeviceType           *string `json:"deviceType,omitempty"`
	Id                   *string `json:"id,omitempty"`
	LastCheckInDateTime  *string `json:"lastCheckInDateTime,omitempty"`
	ODataType            *string `json:"@odata.type,omitempty"`
	UserId               *string `json:"userId,omitempty"`
}
