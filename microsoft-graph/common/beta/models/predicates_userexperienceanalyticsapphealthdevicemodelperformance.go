package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsAppHealthDeviceModelPerformanceOperationPredicate struct {
	ActiveDeviceCount          *int64
	DeviceManufacturer         *string
	DeviceModel                *string
	Id                         *string
	MeanTimeToFailureInMinutes *int64
	ODataType                  *string
}

func (p UserExperienceAnalyticsAppHealthDeviceModelPerformanceOperationPredicate) Matches(input UserExperienceAnalyticsAppHealthDeviceModelPerformance) bool {

	if p.ActiveDeviceCount != nil && (input.ActiveDeviceCount == nil || *p.ActiveDeviceCount != *input.ActiveDeviceCount) {
		return false
	}

	if p.DeviceManufacturer != nil && (input.DeviceManufacturer == nil || *p.DeviceManufacturer != *input.DeviceManufacturer) {
		return false
	}

	if p.DeviceModel != nil && (input.DeviceModel == nil || *p.DeviceModel != *input.DeviceModel) {
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
