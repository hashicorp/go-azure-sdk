package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsAnomalyCorrelationGroupOverview struct {
	AnomalyCorrelationGroupCount         *int64                                                                            `json:"anomalyCorrelationGroupCount,omitempty"`
	AnomalyId                            *string                                                                           `json:"anomalyId,omitempty"`
	CorrelationGroupAnomalousDeviceCount *int64                                                                            `json:"correlationGroupAnomalousDeviceCount,omitempty"`
	CorrelationGroupAtRiskDeviceCount    *int64                                                                            `json:"correlationGroupAtRiskDeviceCount,omitempty"`
	CorrelationGroupDeviceCount          *int64                                                                            `json:"correlationGroupDeviceCount,omitempty"`
	CorrelationGroupFeatures             *[]UserExperienceAnalyticsAnomalyCorrelationGroupFeature                          `json:"correlationGroupFeatures,omitempty"`
	CorrelationGroupId                   *string                                                                           `json:"correlationGroupId,omitempty"`
	CorrelationGroupPrevalence           *UserExperienceAnalyticsAnomalyCorrelationGroupOverviewCorrelationGroupPrevalence `json:"correlationGroupPrevalence,omitempty"`
	CorrelationGroupPrevalencePercentage *float64                                                                          `json:"correlationGroupPrevalencePercentage,omitempty"`
	Id                                   *string                                                                           `json:"id,omitempty"`
	ODataType                            *string                                                                           `json:"@odata.type,omitempty"`
	TotalDeviceCount                     *int64                                                                            `json:"totalDeviceCount,omitempty"`
}
