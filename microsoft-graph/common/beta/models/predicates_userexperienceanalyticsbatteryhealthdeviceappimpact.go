package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsBatteryHealthDeviceAppImpactOperationPredicate struct {
	AppDisplayName  *string
	AppName         *string
	AppPublisher    *string
	DeviceId        *string
	Id              *string
	IsForegroundApp *bool
	ODataType       *string
}

func (p UserExperienceAnalyticsBatteryHealthDeviceAppImpactOperationPredicate) Matches(input UserExperienceAnalyticsBatteryHealthDeviceAppImpact) bool {

	if p.AppDisplayName != nil && (input.AppDisplayName == nil || *p.AppDisplayName != *input.AppDisplayName) {
		return false
	}

	if p.AppName != nil && (input.AppName == nil || *p.AppName != *input.AppName) {
		return false
	}

	if p.AppPublisher != nil && (input.AppPublisher == nil || *p.AppPublisher != *input.AppPublisher) {
		return false
	}

	if p.DeviceId != nil && (input.DeviceId == nil || *p.DeviceId != *input.DeviceId) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.IsForegroundApp != nil && (input.IsForegroundApp == nil || *p.IsForegroundApp != *input.IsForegroundApp) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
