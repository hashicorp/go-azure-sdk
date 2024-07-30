package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessCrossTenantSummary struct {
	AuthTransactionCount  *int64  `json:"authTransactionCount,omitempty"`
	DeviceCount           *int64  `json:"deviceCount,omitempty"`
	NewTenantCount        *int64  `json:"newTenantCount,omitempty"`
	ODataType             *string `json:"@odata.type,omitempty"`
	RarelyUsedTenantCount *int64  `json:"rarelyUsedTenantCount,omitempty"`
	TenantCount           *int64  `json:"tenantCount,omitempty"`
	UserCount             *int64  `json:"userCount,omitempty"`
}
