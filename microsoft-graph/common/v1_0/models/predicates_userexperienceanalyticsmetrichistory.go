package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsMetricHistoryOperationPredicate struct {
	DeviceId       *string
	Id             *string
	MetricDateTime *string
	MetricType     *string
	ODataType      *string
}

func (p UserExperienceAnalyticsMetricHistoryOperationPredicate) Matches(input UserExperienceAnalyticsMetricHistory) bool {

	if p.DeviceId != nil && (input.DeviceId == nil || *p.DeviceId != *input.DeviceId) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.MetricDateTime != nil && (input.MetricDateTime == nil || *p.MetricDateTime != *input.MetricDateTime) {
		return false
	}

	if p.MetricType != nil && (input.MetricType == nil || *p.MetricType != *input.MetricType) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
