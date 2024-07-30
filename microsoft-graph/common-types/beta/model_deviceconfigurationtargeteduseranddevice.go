package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceConfigurationTargetedUserAndDevice struct {
	DeviceId            *string `json:"deviceId,omitempty"`
	DeviceName          *string `json:"deviceName,omitempty"`
	LastCheckinDateTime *string `json:"lastCheckinDateTime,omitempty"`
	ODataType           *string `json:"@odata.type,omitempty"`
	UserDisplayName     *string `json:"userDisplayName,omitempty"`
	UserId              *string `json:"userId,omitempty"`
	UserPrincipalName   *string `json:"userPrincipalName,omitempty"`
}
