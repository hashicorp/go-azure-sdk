package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsCloudManagementDevicesSummary struct {
	CoManagedDeviceCount    *int64  `json:"coManagedDeviceCount,omitempty"`
	IntuneDeviceCount       *int64  `json:"intuneDeviceCount,omitempty"`
	ODataType               *string `json:"@odata.type,omitempty"`
	TenantAttachDeviceCount *int64  `json:"tenantAttachDeviceCount,omitempty"`
}
