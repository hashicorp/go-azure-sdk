package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric struct {
	Id                         *string `json:"id,omitempty"`
	ODataType                  *string `json:"@odata.type,omitempty"`
	TotalDeviceCount           *int64  `json:"totalDeviceCount,omitempty"`
	UpgradeEligibleDeviceCount *int64  `json:"upgradeEligibleDeviceCount,omitempty"`
}
