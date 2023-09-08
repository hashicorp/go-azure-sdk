package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsBatteryHealthRuntimeDetailsOperationPredicate struct {
	ActiveDevices         *int64
	BatteryRuntimeFair    *int64
	BatteryRuntimeGood    *int64
	BatteryRuntimePoor    *int64
	Id                    *string
	LastRefreshedDateTime *string
	ODataType             *string
}

func (p UserExperienceAnalyticsBatteryHealthRuntimeDetailsOperationPredicate) Matches(input UserExperienceAnalyticsBatteryHealthRuntimeDetails) bool {

	if p.ActiveDevices != nil && (input.ActiveDevices == nil || *p.ActiveDevices != *input.ActiveDevices) {
		return false
	}

	if p.BatteryRuntimeFair != nil && (input.BatteryRuntimeFair == nil || *p.BatteryRuntimeFair != *input.BatteryRuntimeFair) {
		return false
	}

	if p.BatteryRuntimeGood != nil && (input.BatteryRuntimeGood == nil || *p.BatteryRuntimeGood != *input.BatteryRuntimeGood) {
		return false
	}

	if p.BatteryRuntimePoor != nil && (input.BatteryRuntimePoor == nil || *p.BatteryRuntimePoor != *input.BatteryRuntimePoor) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.LastRefreshedDateTime != nil && (input.LastRefreshedDateTime == nil || *p.LastRefreshedDateTime != *input.LastRefreshedDateTime) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
