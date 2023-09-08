package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsOperationPredicate struct {
	AppCrashCount          *int64
	AppDisplayName         *string
	AppName                *string
	AppPublisher           *string
	AppVersion             *string
	DeviceCountWithCrashes *int64
	Id                     *string
	IsLatestUsedVersion    *bool
	IsMostUsedVersion      *bool
	ODataType              *string
}

func (p UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsOperationPredicate) Matches(input UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetails) bool {

	if p.AppCrashCount != nil && (input.AppCrashCount == nil || *p.AppCrashCount != *input.AppCrashCount) {
		return false
	}

	if p.AppDisplayName != nil && (input.AppDisplayName == nil || *p.AppDisplayName != *input.AppDisplayName) {
		return false
	}

	if p.AppName != nil && (input.AppName == nil || *p.AppName != *input.AppName) {
		return false
	}

	if p.AppPublisher != nil && (input.AppPublisher == nil || *p.AppPublisher != *input.AppPublisher) {
		return false
	}

	if p.AppVersion != nil && (input.AppVersion == nil || *p.AppVersion != *input.AppVersion) {
		return false
	}

	if p.DeviceCountWithCrashes != nil && (input.DeviceCountWithCrashes == nil || *p.DeviceCountWithCrashes != *input.DeviceCountWithCrashes) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.IsLatestUsedVersion != nil && (input.IsLatestUsedVersion == nil || *p.IsLatestUsedVersion != *input.IsLatestUsedVersion) {
		return false
	}

	if p.IsMostUsedVersion != nil && (input.IsMostUsedVersion == nil || *p.IsMostUsedVersion != *input.IsMostUsedVersion) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
