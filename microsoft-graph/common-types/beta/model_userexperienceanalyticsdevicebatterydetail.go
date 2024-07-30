package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsDeviceBatteryDetail struct {
	BatteryId             *string `json:"batteryId,omitempty"`
	FullBatteryDrainCount *int64  `json:"fullBatteryDrainCount,omitempty"`
	MaxCapacityPercentage *int64  `json:"maxCapacityPercentage,omitempty"`
	ODataType             *string `json:"@odata.type,omitempty"`
}
