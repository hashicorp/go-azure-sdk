package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsCloudManagementDevicesSummaryOperationPredicate struct {
	CoManagedDeviceCount    *int64
	IntuneDeviceCount       *int64
	ODataType               *string
	TenantAttachDeviceCount *int64
}

func (p UserExperienceAnalyticsCloudManagementDevicesSummaryOperationPredicate) Matches(input UserExperienceAnalyticsCloudManagementDevicesSummary) bool {

	if p.CoManagedDeviceCount != nil && (input.CoManagedDeviceCount == nil || *p.CoManagedDeviceCount != *input.CoManagedDeviceCount) {
		return false
	}

	if p.IntuneDeviceCount != nil && (input.IntuneDeviceCount == nil || *p.IntuneDeviceCount != *input.IntuneDeviceCount) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.TenantAttachDeviceCount != nil && (input.TenantAttachDeviceCount == nil || *p.TenantAttachDeviceCount != *input.TenantAttachDeviceCount) {
		return false
	}

	return true
}
