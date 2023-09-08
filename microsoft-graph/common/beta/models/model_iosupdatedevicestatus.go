package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosUpdateDeviceStatus struct {
	ComplianceGracePeriodExpirationDateTime *string                             `json:"complianceGracePeriodExpirationDateTime,omitempty"`
	DeviceDisplayName                       *string                             `json:"deviceDisplayName,omitempty"`
	DeviceId                                *string                             `json:"deviceId,omitempty"`
	DeviceModel                             *string                             `json:"deviceModel,omitempty"`
	Id                                      *string                             `json:"id,omitempty"`
	InstallStatus                           *IosUpdateDeviceStatusInstallStatus `json:"installStatus,omitempty"`
	LastReportedDateTime                    *string                             `json:"lastReportedDateTime,omitempty"`
	ODataType                               *string                             `json:"@odata.type,omitempty"`
	OsVersion                               *string                             `json:"osVersion,omitempty"`
	Platform                                *int64                              `json:"platform,omitempty"`
	Status                                  *IosUpdateDeviceStatusStatus        `json:"status,omitempty"`
	UserId                                  *string                             `json:"userId,omitempty"`
	UserName                                *string                             `json:"userName,omitempty"`
	UserPrincipalName                       *string                             `json:"userPrincipalName,omitempty"`
}
