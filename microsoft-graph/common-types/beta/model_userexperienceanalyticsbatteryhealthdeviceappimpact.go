package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsBatteryHealthDeviceAppImpact struct {
	AppDisplayName         *string  `json:"appDisplayName,omitempty"`
	AppName                *string  `json:"appName,omitempty"`
	AppPublisher           *string  `json:"appPublisher,omitempty"`
	BatteryUsagePercentage *float64 `json:"batteryUsagePercentage,omitempty"`
	DeviceId               *string  `json:"deviceId,omitempty"`
	Id                     *string  `json:"id,omitempty"`
	IsForegroundApp        *bool    `json:"isForegroundApp,omitempty"`
	ODataType              *string  `json:"@odata.type,omitempty"`
}
