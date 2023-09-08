package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsRemoteConnection struct {
	DeviceCount       *int64  `json:"deviceCount,omitempty"`
	DeviceId          *string `json:"deviceId,omitempty"`
	DeviceName        *string `json:"deviceName,omitempty"`
	Id                *string `json:"id,omitempty"`
	Manufacturer      *string `json:"manufacturer,omitempty"`
	Model             *string `json:"model,omitempty"`
	ODataType         *string `json:"@odata.type,omitempty"`
	UserPrincipalName *string `json:"userPrincipalName,omitempty"`
	VirtualNetwork    *string `json:"virtualNetwork,omitempty"`
}
