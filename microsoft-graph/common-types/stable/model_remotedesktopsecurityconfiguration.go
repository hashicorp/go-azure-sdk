package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RemoteDesktopSecurityConfiguration struct {
	Id                             *string              `json:"id,omitempty"`
	IsRemoteDesktopProtocolEnabled *bool                `json:"isRemoteDesktopProtocolEnabled,omitempty"`
	ODataType                      *string              `json:"@odata.type,omitempty"`
	TargetDeviceGroups             *[]TargetDeviceGroup `json:"targetDeviceGroups,omitempty"`
}
