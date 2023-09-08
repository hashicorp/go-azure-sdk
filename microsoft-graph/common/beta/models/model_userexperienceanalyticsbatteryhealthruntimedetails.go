package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsBatteryHealthRuntimeDetails struct {
	ActiveDevices         *int64  `json:"activeDevices,omitempty"`
	BatteryRuntimeFair    *int64  `json:"batteryRuntimeFair,omitempty"`
	BatteryRuntimeGood    *int64  `json:"batteryRuntimeGood,omitempty"`
	BatteryRuntimePoor    *int64  `json:"batteryRuntimePoor,omitempty"`
	Id                    *string `json:"id,omitempty"`
	LastRefreshedDateTime *string `json:"lastRefreshedDateTime,omitempty"`
	ODataType             *string `json:"@odata.type,omitempty"`
}
