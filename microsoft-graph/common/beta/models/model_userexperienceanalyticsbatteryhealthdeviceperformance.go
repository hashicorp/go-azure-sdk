package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsBatteryHealthDevicePerformance struct {
	BatteryAgeInDays          *int64                                                             `json:"batteryAgeInDays,omitempty"`
	DeviceBatteryCount        *int64                                                             `json:"deviceBatteryCount,omitempty"`
	DeviceBatteryHealthScore  *int64                                                             `json:"deviceBatteryHealthScore,omitempty"`
	DeviceId                  *string                                                            `json:"deviceId,omitempty"`
	DeviceName                *string                                                            `json:"deviceName,omitempty"`
	EstimatedRuntimeInMinutes *int64                                                             `json:"estimatedRuntimeInMinutes,omitempty"`
	FullBatteryDrainCount     *int64                                                             `json:"fullBatteryDrainCount,omitempty"`
	HealthStatus              *UserExperienceAnalyticsBatteryHealthDevicePerformanceHealthStatus `json:"healthStatus,omitempty"`
	Id                        *string                                                            `json:"id,omitempty"`
	Manufacturer              *string                                                            `json:"manufacturer,omitempty"`
	MaxCapacityPercentage     *int64                                                             `json:"maxCapacityPercentage,omitempty"`
	Model                     *string                                                            `json:"model,omitempty"`
	ODataType                 *string                                                            `json:"@odata.type,omitempty"`
}
