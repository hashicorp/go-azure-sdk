package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsInformationProtectionWipeAction struct {
	Id                           *string                                       `json:"id,omitempty"`
	LastCheckInDateTime          *string                                       `json:"lastCheckInDateTime,omitempty"`
	ODataType                    *string                                       `json:"@odata.type,omitempty"`
	Status                       *WindowsInformationProtectionWipeActionStatus `json:"status,omitempty"`
	TargetedDeviceMacAddress     *string                                       `json:"targetedDeviceMacAddress,omitempty"`
	TargetedDeviceName           *string                                       `json:"targetedDeviceName,omitempty"`
	TargetedDeviceRegistrationId *string                                       `json:"targetedDeviceRegistrationId,omitempty"`
	TargetedUserId               *string                                       `json:"targetedUserId,omitempty"`
}
