package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsAppHealthDevicePerformanceDetailsOperationPredicate struct {
	AppDisplayName    *string
	AppPublisher      *string
	AppVersion        *string
	DeviceDisplayName *string
	DeviceId          *string
	EventDateTime     *string
	EventType         *string
	Id                *string
	ODataType         *string
}

func (p UserExperienceAnalyticsAppHealthDevicePerformanceDetailsOperationPredicate) Matches(input UserExperienceAnalyticsAppHealthDevicePerformanceDetails) bool {

	if p.AppDisplayName != nil && (input.AppDisplayName == nil || *p.AppDisplayName != *input.AppDisplayName) {
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

	if p.EventDateTime != nil && (input.EventDateTime == nil || *p.EventDateTime != *input.EventDateTime) {
		return false
	}

	if p.EventType != nil && (input.EventType == nil || *p.EventType != *input.EventType) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
