package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsCategory struct {
	Id           *string                           `json:"id,omitempty"`
	Insights     *[]UserExperienceAnalyticsInsight `json:"insights,omitempty"`
	MetricValues *[]UserExperienceAnalyticsMetric  `json:"metricValues,omitempty"`
	ODataType    *string                           `json:"@odata.type,omitempty"`
}
