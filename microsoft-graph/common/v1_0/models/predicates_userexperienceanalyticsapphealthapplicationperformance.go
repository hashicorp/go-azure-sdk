package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsAppHealthApplicationPerformanceOperationPredicate struct {
	ActiveDeviceCount          *int64
	AppCrashCount              *int64
	AppDisplayName             *string
	AppHangCount               *int64
	AppName                    *string
	AppPublisher               *string
	AppUsageDuration           *int64
	Id                         *string
	MeanTimeToFailureInMinutes *int64
	ODataType                  *string
}

func (p UserExperienceAnalyticsAppHealthApplicationPerformanceOperationPredicate) Matches(input UserExperienceAnalyticsAppHealthApplicationPerformance) bool {

	if p.ActiveDeviceCount != nil && (input.ActiveDeviceCount == nil || *p.ActiveDeviceCount != *input.ActiveDeviceCount) {
		return false
	}

	if p.AppCrashCount != nil && (input.AppCrashCount == nil || *p.AppCrashCount != *input.AppCrashCount) {
		return false
	}

	if p.AppDisplayName != nil && (input.AppDisplayName == nil || *p.AppDisplayName != *input.AppDisplayName) {
		return false
	}

	if p.AppHangCount != nil && (input.AppHangCount == nil || *p.AppHangCount != *input.AppHangCount) {
		return false
	}

	if p.AppName != nil && (input.AppName == nil || *p.AppName != *input.AppName) {
		return false
	}

	if p.AppPublisher != nil && (input.AppPublisher == nil || *p.AppPublisher != *input.AppPublisher) {
		return false
	}

	if p.AppUsageDuration != nil && (input.AppUsageDuration == nil || *p.AppUsageDuration != *input.AppUsageDuration) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.MeanTimeToFailureInMinutes != nil && (input.MeanTimeToFailureInMinutes == nil || *p.MeanTimeToFailureInMinutes != *input.MeanTimeToFailureInMinutes) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
