package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsBatteryHealthModelPerformance struct {
	ActiveDevices                    *int64                                                                 `json:"activeDevices,omitempty"`
	AverageBatteryAgeInDays          *int64                                                                 `json:"averageBatteryAgeInDays,omitempty"`
	AverageEstimatedRuntimeInMinutes *int64                                                                 `json:"averageEstimatedRuntimeInMinutes,omitempty"`
	AverageMaxCapacityPercentage     *int64                                                                 `json:"averageMaxCapacityPercentage,omitempty"`
	Id                               *string                                                                `json:"id,omitempty"`
	Manufacturer                     *string                                                                `json:"manufacturer,omitempty"`
	MeanFullBatteryDrainCount        *int64                                                                 `json:"meanFullBatteryDrainCount,omitempty"`
	MedianEstimatedRuntimeInMinutes  *int64                                                                 `json:"medianEstimatedRuntimeInMinutes,omitempty"`
	MedianFullBatteryDrainCount      *int64                                                                 `json:"medianFullBatteryDrainCount,omitempty"`
	MedianMaxCapacityPercentage      *int64                                                                 `json:"medianMaxCapacityPercentage,omitempty"`
	Model                            *string                                                                `json:"model,omitempty"`
	ModelBatteryHealthScore          *int64                                                                 `json:"modelBatteryHealthScore,omitempty"`
	ModelHealthStatus                *UserExperienceAnalyticsBatteryHealthModelPerformanceModelHealthStatus `json:"modelHealthStatus,omitempty"`
	ODataType                        *string                                                                `json:"@odata.type,omitempty"`
}
