package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedTenantsDeviceHealthStatus struct {
	BlueScreenCount     *int64  `json:"blueScreenCount,omitempty"`
	DeviceId            *string `json:"deviceId,omitempty"`
	DeviceMake          *string `json:"deviceMake,omitempty"`
	DeviceModel         *string `json:"deviceModel,omitempty"`
	DeviceName          *string `json:"deviceName,omitempty"`
	HealthStatus        *string `json:"healthStatus,omitempty"`
	Id                  *string `json:"id,omitempty"`
	LastUpdatedDateTime *string `json:"lastUpdatedDateTime,omitempty"`
	ODataType           *string `json:"@odata.type,omitempty"`
	OsVersion           *string `json:"osVersion,omitempty"`
	PrimaryDiskType     *string `json:"primaryDiskType,omitempty"`
	RestartCount        *int64  `json:"restartCount,omitempty"`
	TenantDisplayName   *string `json:"tenantDisplayName,omitempty"`
	TenantId            *string `json:"tenantId,omitempty"`
	TopProcesses        *string `json:"topProcesses,omitempty"`
}
