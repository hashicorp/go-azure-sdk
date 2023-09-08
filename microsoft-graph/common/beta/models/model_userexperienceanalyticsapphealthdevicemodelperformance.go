package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsAppHealthDeviceModelPerformance struct {
	ActiveDeviceCount          *int64                                                              `json:"activeDeviceCount,omitempty"`
	DeviceManufacturer         *string                                                             `json:"deviceManufacturer,omitempty"`
	DeviceModel                *string                                                             `json:"deviceModel,omitempty"`
	HealthStatus               *UserExperienceAnalyticsAppHealthDeviceModelPerformanceHealthStatus `json:"healthStatus,omitempty"`
	Id                         *string                                                             `json:"id,omitempty"`
	MeanTimeToFailureInMinutes *int64                                                              `json:"meanTimeToFailureInMinutes,omitempty"`
	ODataType                  *string                                                             `json:"@odata.type,omitempty"`
}
