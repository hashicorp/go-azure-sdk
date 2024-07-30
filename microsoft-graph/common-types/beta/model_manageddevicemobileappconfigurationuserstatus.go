package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedDeviceMobileAppConfigurationUserStatus struct {
	DevicesCount         *int64                                               `json:"devicesCount,omitempty"`
	Id                   *string                                              `json:"id,omitempty"`
	LastReportedDateTime *string                                              `json:"lastReportedDateTime,omitempty"`
	ODataType            *string                                              `json:"@odata.type,omitempty"`
	Status               *ManagedDeviceMobileAppConfigurationUserStatusStatus `json:"status,omitempty"`
	UserDisplayName      *string                                              `json:"userDisplayName,omitempty"`
	UserPrincipalName    *string                                              `json:"userPrincipalName,omitempty"`
}
