package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsBatteryHealthCapacityDetails struct {
	ActiveDevices         *int64  `json:"activeDevices,omitempty"`
	BatteryCapacityFair   *int64  `json:"batteryCapacityFair,omitempty"`
	BatteryCapacityGood   *int64  `json:"batteryCapacityGood,omitempty"`
	BatteryCapacityPoor   *int64  `json:"batteryCapacityPoor,omitempty"`
	Id                    *string `json:"id,omitempty"`
	LastRefreshedDateTime *string `json:"lastRefreshedDateTime,omitempty"`
	ODataType             *string `json:"@odata.type,omitempty"`
}
