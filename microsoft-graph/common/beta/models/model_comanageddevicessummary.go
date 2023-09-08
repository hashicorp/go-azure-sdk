package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ComanagedDevicesSummary struct {
	CompliancePolicyCount         *int64  `json:"compliancePolicyCount,omitempty"`
	ConfigurationSettingsCount    *int64  `json:"configurationSettingsCount,omitempty"`
	EndpointProtectionCount       *int64  `json:"endpointProtectionCount,omitempty"`
	InventoryCount                *int64  `json:"inventoryCount,omitempty"`
	ModernAppsCount               *int64  `json:"modernAppsCount,omitempty"`
	ODataType                     *string `json:"@odata.type,omitempty"`
	OfficeAppsCount               *int64  `json:"officeAppsCount,omitempty"`
	ResourceAccessCount           *int64  `json:"resourceAccessCount,omitempty"`
	TotalComanagedCount           *int64  `json:"totalComanagedCount,omitempty"`
	WindowsUpdateForBusinessCount *int64  `json:"windowsUpdateForBusinessCount,omitempty"`
}
