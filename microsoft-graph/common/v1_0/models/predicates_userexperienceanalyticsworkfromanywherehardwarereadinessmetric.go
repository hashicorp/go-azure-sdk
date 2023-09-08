package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricOperationPredicate struct {
	Id                         *string
	ODataType                  *string
	TotalDeviceCount           *int64
	UpgradeEligibleDeviceCount *int64
}

func (p UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricOperationPredicate) Matches(input UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric) bool {

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.TotalDeviceCount != nil && (input.TotalDeviceCount == nil || *p.TotalDeviceCount != *input.TotalDeviceCount) {
		return false
	}

	if p.UpgradeEligibleDeviceCount != nil && (input.UpgradeEligibleDeviceCount == nil || *p.UpgradeEligibleDeviceCount != *input.UpgradeEligibleDeviceCount) {
		return false
	}

	return true
}
