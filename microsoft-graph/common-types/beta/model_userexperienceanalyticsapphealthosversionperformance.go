package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsAppHealthOSVersionPerformance struct {
	ActiveDeviceCount          *int64   `json:"activeDeviceCount,omitempty"`
	Id                         *string  `json:"id,omitempty"`
	MeanTimeToFailureInMinutes *int64   `json:"meanTimeToFailureInMinutes,omitempty"`
	ODataType                  *string  `json:"@odata.type,omitempty"`
	OsBuildNumber              *string  `json:"osBuildNumber,omitempty"`
	OsVersion                  *string  `json:"osVersion,omitempty"`
	OsVersionAppHealthScore    *float64 `json:"osVersionAppHealthScore,omitempty"`
}
