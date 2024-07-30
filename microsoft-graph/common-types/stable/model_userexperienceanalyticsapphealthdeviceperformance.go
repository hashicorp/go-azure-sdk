package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsAppHealthDevicePerformance struct {
	AppCrashCount              *int64                                                         `json:"appCrashCount,omitempty"`
	AppHangCount               *int64                                                         `json:"appHangCount,omitempty"`
	CrashedAppCount            *int64                                                         `json:"crashedAppCount,omitempty"`
	DeviceAppHealthScore       *float64                                                       `json:"deviceAppHealthScore,omitempty"`
	DeviceDisplayName          *string                                                        `json:"deviceDisplayName,omitempty"`
	DeviceId                   *string                                                        `json:"deviceId,omitempty"`
	DeviceManufacturer         *string                                                        `json:"deviceManufacturer,omitempty"`
	DeviceModel                *string                                                        `json:"deviceModel,omitempty"`
	HealthStatus               *UserExperienceAnalyticsAppHealthDevicePerformanceHealthStatus `json:"healthStatus,omitempty"`
	Id                         *string                                                        `json:"id,omitempty"`
	MeanTimeToFailureInMinutes *int64                                                         `json:"meanTimeToFailureInMinutes,omitempty"`
	ODataType                  *string                                                        `json:"@odata.type,omitempty"`
	ProcessedDateTime          *string                                                        `json:"processedDateTime,omitempty"`
}
