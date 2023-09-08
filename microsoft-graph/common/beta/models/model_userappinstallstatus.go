package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserAppInstallStatus struct {
	App                     *MobileApp                `json:"app,omitempty"`
	DeviceStatuses          *[]MobileAppInstallStatus `json:"deviceStatuses,omitempty"`
	FailedDeviceCount       *int64                    `json:"failedDeviceCount,omitempty"`
	Id                      *string                   `json:"id,omitempty"`
	InstalledDeviceCount    *int64                    `json:"installedDeviceCount,omitempty"`
	NotInstalledDeviceCount *int64                    `json:"notInstalledDeviceCount,omitempty"`
	ODataType               *string                   `json:"@odata.type,omitempty"`
	UserName                *string                   `json:"userName,omitempty"`
	UserPrincipalName       *string                   `json:"userPrincipalName,omitempty"`
}
