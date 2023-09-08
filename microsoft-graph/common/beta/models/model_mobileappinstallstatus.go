package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MobileAppInstallStatus struct {
	App                         *MobileApp                                         `json:"app,omitempty"`
	DeviceId                    *string                                            `json:"deviceId,omitempty"`
	DeviceName                  *string                                            `json:"deviceName,omitempty"`
	DisplayVersion              *string                                            `json:"displayVersion,omitempty"`
	ErrorCode                   *int64                                             `json:"errorCode,omitempty"`
	Id                          *string                                            `json:"id,omitempty"`
	InstallState                *MobileAppInstallStatusInstallState                `json:"installState,omitempty"`
	InstallStateDetail          *MobileAppInstallStatusInstallStateDetail          `json:"installStateDetail,omitempty"`
	LastSyncDateTime            *string                                            `json:"lastSyncDateTime,omitempty"`
	MobileAppInstallStatusValue *MobileAppInstallStatusMobileAppInstallStatusValue `json:"mobileAppInstallStatusValue,omitempty"`
	ODataType                   *string                                            `json:"@odata.type,omitempty"`
	OsDescription               *string                                            `json:"osDescription,omitempty"`
	OsVersion                   *string                                            `json:"osVersion,omitempty"`
	UserName                    *string                                            `json:"userName,omitempty"`
	UserPrincipalName           *string                                            `json:"userPrincipalName,omitempty"`
}
