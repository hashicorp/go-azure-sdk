package iotsecuritysolutionsanalytics

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IoTSecurityAggregatedRecommendationProperties struct {
	Description               *string           `json:"description,omitempty"`
	DetectedBy                *string           `json:"detectedBy,omitempty"`
	HealthyDevices            *int64            `json:"healthyDevices,omitempty"`
	LogAnalyticsQuery         *string           `json:"logAnalyticsQuery,omitempty"`
	RecommendationDisplayName *string           `json:"recommendationDisplayName,omitempty"`
	RecommendationName        *string           `json:"recommendationName,omitempty"`
	RecommendationTypeId      *string           `json:"recommendationTypeId,omitempty"`
	RemediationSteps          *string           `json:"remediationSteps,omitempty"`
	ReportedSeverity          *ReportedSeverity `json:"reportedSeverity,omitempty"`
	UnhealthyDeviceCount      *int64            `json:"unhealthyDeviceCount,omitempty"`
}
