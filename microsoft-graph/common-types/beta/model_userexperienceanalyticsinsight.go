package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsInsight struct {
	InsightId                       *string                                 `json:"insightId,omitempty"`
	ODataType                       *string                                 `json:"@odata.type,omitempty"`
	Severity                        *UserExperienceAnalyticsInsightSeverity `json:"severity,omitempty"`
	UserExperienceAnalyticsMetricId *string                                 `json:"userExperienceAnalyticsMetricId,omitempty"`
	Values                          *[]UserExperienceAnalyticsInsightValue  `json:"values,omitempty"`
}
