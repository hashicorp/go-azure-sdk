package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementIntentDeviceState struct {
	DeviceDisplayName    *string                                 `json:"deviceDisplayName,omitempty"`
	DeviceId             *string                                 `json:"deviceId,omitempty"`
	Id                   *string                                 `json:"id,omitempty"`
	LastReportedDateTime *string                                 `json:"lastReportedDateTime,omitempty"`
	ODataType            *string                                 `json:"@odata.type,omitempty"`
	State                *DeviceManagementIntentDeviceStateState `json:"state,omitempty"`
	UserName             *string                                 `json:"userName,omitempty"`
	UserPrincipalName    *string                                 `json:"userPrincipalName,omitempty"`
}
