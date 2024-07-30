package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsAppHealthAppPerformanceByOSVersion struct {
	ActiveDeviceCount          *int64  `json:"activeDeviceCount,omitempty"`
	AppCrashCount              *int64  `json:"appCrashCount,omitempty"`
	AppDisplayName             *string `json:"appDisplayName,omitempty"`
	AppName                    *string `json:"appName,omitempty"`
	AppPublisher               *string `json:"appPublisher,omitempty"`
	AppUsageDuration           *int64  `json:"appUsageDuration,omitempty"`
	Id                         *string `json:"id,omitempty"`
	MeanTimeToFailureInMinutes *int64  `json:"meanTimeToFailureInMinutes,omitempty"`
	ODataType                  *string `json:"@odata.type,omitempty"`
	OsBuildNumber              *string `json:"osBuildNumber,omitempty"`
	OsVersion                  *string `json:"osVersion,omitempty"`
}
