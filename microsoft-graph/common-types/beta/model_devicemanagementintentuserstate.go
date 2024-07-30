package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementIntentUserState struct {
	DeviceCount          *int64                                `json:"deviceCount,omitempty"`
	Id                   *string                               `json:"id,omitempty"`
	LastReportedDateTime *string                               `json:"lastReportedDateTime,omitempty"`
	ODataType            *string                               `json:"@odata.type,omitempty"`
	State                *DeviceManagementIntentUserStateState `json:"state,omitempty"`
	UserName             *string                               `json:"userName,omitempty"`
	UserPrincipalName    *string                               `json:"userPrincipalName,omitempty"`
}
