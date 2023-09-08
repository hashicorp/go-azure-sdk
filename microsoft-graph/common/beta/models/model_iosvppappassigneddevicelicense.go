package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosVppAppAssignedDeviceLicense struct {
	DeviceName        *string `json:"deviceName,omitempty"`
	Id                *string `json:"id,omitempty"`
	ManagedDeviceId   *string `json:"managedDeviceId,omitempty"`
	ODataType         *string `json:"@odata.type,omitempty"`
	UserEmailAddress  *string `json:"userEmailAddress,omitempty"`
	UserId            *string `json:"userId,omitempty"`
	UserName          *string `json:"userName,omitempty"`
	UserPrincipalName *string `json:"userPrincipalName,omitempty"`
}
