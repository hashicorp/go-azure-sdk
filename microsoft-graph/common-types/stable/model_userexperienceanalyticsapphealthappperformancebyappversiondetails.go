package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetails struct {
	AppCrashCount          *int64  `json:"appCrashCount,omitempty"`
	AppDisplayName         *string `json:"appDisplayName,omitempty"`
	AppName                *string `json:"appName,omitempty"`
	AppPublisher           *string `json:"appPublisher,omitempty"`
	AppVersion             *string `json:"appVersion,omitempty"`
	DeviceCountWithCrashes *int64  `json:"deviceCountWithCrashes,omitempty"`
	Id                     *string `json:"id,omitempty"`
	IsLatestUsedVersion    *bool   `json:"isLatestUsedVersion,omitempty"`
	IsMostUsedVersion      *bool   `json:"isMostUsedVersion,omitempty"`
	ODataType              *string `json:"@odata.type,omitempty"`
}
