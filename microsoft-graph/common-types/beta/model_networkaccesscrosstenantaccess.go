package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessCrossTenantAccess struct {
	DeviceCount                 *int64                                     `json:"deviceCount,omitempty"`
	LastAccessDateTime          *string                                    `json:"lastAccessDateTime,omitempty"`
	ODataType                   *string                                    `json:"@odata.type,omitempty"`
	ResourceTenantId            *string                                    `json:"resourceTenantId,omitempty"`
	ResourceTenantName          *string                                    `json:"resourceTenantName,omitempty"`
	ResourceTenantPrimaryDomain *string                                    `json:"resourceTenantPrimaryDomain,omitempty"`
	UsageStatus                 *NetworkaccessCrossTenantAccessUsageStatus `json:"usageStatus,omitempty"`
	UserCount                   *int64                                     `json:"userCount,omitempty"`
}
