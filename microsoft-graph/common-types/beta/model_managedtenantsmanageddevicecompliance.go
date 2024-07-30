package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedTenantsManagedDeviceCompliance struct {
	ComplianceStatus           *string `json:"complianceStatus,omitempty"`
	DeviceType                 *string `json:"deviceType,omitempty"`
	Id                         *string `json:"id,omitempty"`
	InGracePeriodUntilDateTime *string `json:"inGracePeriodUntilDateTime,omitempty"`
	LastRefreshedDateTime      *string `json:"lastRefreshedDateTime,omitempty"`
	LastSyncDateTime           *string `json:"lastSyncDateTime,omitempty"`
	ManagedDeviceId            *string `json:"managedDeviceId,omitempty"`
	ManagedDeviceName          *string `json:"managedDeviceName,omitempty"`
	Manufacturer               *string `json:"manufacturer,omitempty"`
	Model                      *string `json:"model,omitempty"`
	ODataType                  *string `json:"@odata.type,omitempty"`
	OsDescription              *string `json:"osDescription,omitempty"`
	OsVersion                  *string `json:"osVersion,omitempty"`
	OwnerType                  *string `json:"ownerType,omitempty"`
	TenantDisplayName          *string `json:"tenantDisplayName,omitempty"`
	TenantId                   *string `json:"tenantId,omitempty"`
}
