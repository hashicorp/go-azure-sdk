package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedTenantsCloudPCOverview struct {
	LastRefreshedDateTime                            *string `json:"lastRefreshedDateTime,omitempty"`
	NumberOfCloudPCConnectionStatusFailed            *int64  `json:"numberOfCloudPcConnectionStatusFailed,omitempty"`
	NumberOfCloudPCConnectionStatusPassed            *int64  `json:"numberOfCloudPcConnectionStatusPassed,omitempty"`
	NumberOfCloudPCConnectionStatusPending           *int64  `json:"numberOfCloudPcConnectionStatusPending,omitempty"`
	NumberOfCloudPCConnectionStatusRunning           *int64  `json:"numberOfCloudPcConnectionStatusRunning,omitempty"`
	NumberOfCloudPCConnectionStatusUnkownFutureValue *int64  `json:"numberOfCloudPcConnectionStatusUnkownFutureValue,omitempty"`
	NumberOfCloudPCStatusDeprovisioning              *int64  `json:"numberOfCloudPcStatusDeprovisioning,omitempty"`
	NumberOfCloudPCStatusFailed                      *int64  `json:"numberOfCloudPcStatusFailed,omitempty"`
	NumberOfCloudPCStatusInGracePeriod               *int64  `json:"numberOfCloudPcStatusInGracePeriod,omitempty"`
	NumberOfCloudPCStatusNotProvisioned              *int64  `json:"numberOfCloudPcStatusNotProvisioned,omitempty"`
	NumberOfCloudPCStatusProvisioned                 *int64  `json:"numberOfCloudPcStatusProvisioned,omitempty"`
	NumberOfCloudPCStatusProvisioning                *int64  `json:"numberOfCloudPcStatusProvisioning,omitempty"`
	NumberOfCloudPCStatusUnknown                     *int64  `json:"numberOfCloudPcStatusUnknown,omitempty"`
	NumberOfCloudPCStatusUpgrading                   *int64  `json:"numberOfCloudPcStatusUpgrading,omitempty"`
	ODataType                                        *string `json:"@odata.type,omitempty"`
	TenantDisplayName                                *string `json:"tenantDisplayName,omitempty"`
	TenantId                                         *string `json:"tenantId,omitempty"`
	TotalBusinessLicenses                            *int64  `json:"totalBusinessLicenses,omitempty"`
	TotalCloudPCConnectionStatus                     *int64  `json:"totalCloudPcConnectionStatus,omitempty"`
	TotalCloudPCStatus                               *int64  `json:"totalCloudPcStatus,omitempty"`
	TotalEnterpriseLicenses                          *int64  `json:"totalEnterpriseLicenses,omitempty"`
}
