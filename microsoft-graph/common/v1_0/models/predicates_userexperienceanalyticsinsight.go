package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsInsightOperationPredicate struct {
	InsightId                       *string
	ODataType                       *string
	UserExperienceAnalyticsMetricId *string
}

func (p UserExperienceAnalyticsInsightOperationPredicate) Matches(input UserExperienceAnalyticsInsight) bool {

	if p.InsightId != nil && (input.InsightId == nil || *p.InsightId != *input.InsightId) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.UserExperienceAnalyticsMetricId != nil && (input.UserExperienceAnalyticsMetricId == nil || *p.UserExperienceAnalyticsMetricId != *input.UserExperienceAnalyticsMetricId) {
		return false
	}

	return true
}
