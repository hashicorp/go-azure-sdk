package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdOperationPredicate struct {
	AppCrashCount     *int64
	AppDisplayName    *string
	AppName           *string
	AppPublisher      *string
	AppVersion        *string
	DeviceDisplayName *string
	DeviceId          *string
	Id                *string
	ODataType         *string
	ProcessedDateTime *string
}

func (p UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdOperationPredicate) Matches(input UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceId) bool {

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

	if p.DeviceDisplayName != nil && (input.DeviceDisplayName == nil || *p.DeviceDisplayName != *input.DeviceDisplayName) {
		return false
	}

	if p.DeviceId != nil && (input.DeviceId == nil || *p.DeviceId != *input.DeviceId) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.ProcessedDateTime != nil && (input.ProcessedDateTime == nil || *p.ProcessedDateTime != *input.ProcessedDateTime) {
		return false
	}

	return true
}
