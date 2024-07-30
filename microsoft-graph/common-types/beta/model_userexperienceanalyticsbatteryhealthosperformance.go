package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsBatteryHealthOsPerformance struct {
	ActiveDevices                    *int64                                                           `json:"activeDevices,omitempty"`
	AverageBatteryAgeInDays          *int64                                                           `json:"averageBatteryAgeInDays,omitempty"`
	AverageEstimatedRuntimeInMinutes *int64                                                           `json:"averageEstimatedRuntimeInMinutes,omitempty"`
	AverageMaxCapacityPercentage     *int64                                                           `json:"averageMaxCapacityPercentage,omitempty"`
	Id                               *string                                                          `json:"id,omitempty"`
	MeanFullBatteryDrainCount        *int64                                                           `json:"meanFullBatteryDrainCount,omitempty"`
	MedianEstimatedRuntimeInMinutes  *int64                                                           `json:"medianEstimatedRuntimeInMinutes,omitempty"`
	MedianFullBatteryDrainCount      *int64                                                           `json:"medianFullBatteryDrainCount,omitempty"`
	MedianMaxCapacityPercentage      *int64                                                           `json:"medianMaxCapacityPercentage,omitempty"`
	ODataType                        *string                                                          `json:"@odata.type,omitempty"`
	OsBatteryHealthScore             *int64                                                           `json:"osBatteryHealthScore,omitempty"`
	OsBuildNumber                    *string                                                          `json:"osBuildNumber,omitempty"`
	OsHealthStatus                   *UserExperienceAnalyticsBatteryHealthOsPerformanceOsHealthStatus `json:"osHealthStatus,omitempty"`
	OsVersion                        *string                                                          `json:"osVersion,omitempty"`
}
