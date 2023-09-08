package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityBaselineDeviceState struct {
	DeviceDisplayName    *string                           `json:"deviceDisplayName,omitempty"`
	Id                   *string                           `json:"id,omitempty"`
	LastReportedDateTime *string                           `json:"lastReportedDateTime,omitempty"`
	ManagedDeviceId      *string                           `json:"managedDeviceId,omitempty"`
	ODataType            *string                           `json:"@odata.type,omitempty"`
	State                *SecurityBaselineDeviceStateState `json:"state,omitempty"`
	UserPrincipalName    *string                           `json:"userPrincipalName,omitempty"`
}
