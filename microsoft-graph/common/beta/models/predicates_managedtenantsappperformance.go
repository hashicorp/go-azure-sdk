package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedTenantsAppPerformanceOperationPredicate struct {
	AppFriendlyName            *string
	AppName                    *string
	AppPublisher               *string
	Id                         *string
	LastUpdatedDateTime        *string
	MeanTimeToFailureInMinutes *int64
	ODataType                  *string
	TenantDisplayName          *string
	TenantId                   *string
	TotalActiveDeviceCount     *int64
	TotalAppCrashCount         *int64
	TotalAppFreezeCount        *int64
}

func (p ManagedTenantsAppPerformanceOperationPredicate) Matches(input ManagedTenantsAppPerformance) bool {

	if p.AppFriendlyName != nil && (input.AppFriendlyName == nil || *p.AppFriendlyName != *input.AppFriendlyName) {
		return false
	}

	if p.AppName != nil && (input.AppName == nil || *p.AppName != *input.AppName) {
		return false
	}

	if p.AppPublisher != nil && (input.AppPublisher == nil || *p.AppPublisher != *input.AppPublisher) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.LastUpdatedDateTime != nil && (input.LastUpdatedDateTime == nil || *p.LastUpdatedDateTime != *input.LastUpdatedDateTime) {
		return false
	}

	if p.MeanTimeToFailureInMinutes != nil && (input.MeanTimeToFailureInMinutes == nil || *p.MeanTimeToFailureInMinutes != *input.MeanTimeToFailureInMinutes) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.TenantDisplayName != nil && (input.TenantDisplayName == nil || *p.TenantDisplayName != *input.TenantDisplayName) {
		return false
	}

	if p.TenantId != nil && (input.TenantId == nil || *p.TenantId != *input.TenantId) {
		return false
	}

	if p.TotalActiveDeviceCount != nil && (input.TotalActiveDeviceCount == nil || *p.TotalActiveDeviceCount != *input.TotalActiveDeviceCount) {
		return false
	}

	if p.TotalAppCrashCount != nil && (input.TotalAppCrashCount == nil || *p.TotalAppCrashCount != *input.TotalAppCrashCount) {
		return false
	}

	if p.TotalAppFreezeCount != nil && (input.TotalAppFreezeCount == nil || *p.TotalAppFreezeCount != *input.TotalAppFreezeCount) {
		return false
	}

	return true
}
