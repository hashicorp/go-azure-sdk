package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric struct {
	Id                                      *string  `json:"id,omitempty"`
	ODataType                               *string  `json:"@odata.type,omitempty"`
	OsCheckFailedPercentage                 *float64 `json:"osCheckFailedPercentage,omitempty"`
	Processor64BitCheckFailedPercentage     *float64 `json:"processor64BitCheckFailedPercentage,omitempty"`
	ProcessorCoreCountCheckFailedPercentage *float64 `json:"processorCoreCountCheckFailedPercentage,omitempty"`
	ProcessorFamilyCheckFailedPercentage    *float64 `json:"processorFamilyCheckFailedPercentage,omitempty"`
	ProcessorSpeedCheckFailedPercentage     *float64 `json:"processorSpeedCheckFailedPercentage,omitempty"`
	RamCheckFailedPercentage                *float64 `json:"ramCheckFailedPercentage,omitempty"`
	SecureBootCheckFailedPercentage         *float64 `json:"secureBootCheckFailedPercentage,omitempty"`
	StorageCheckFailedPercentage            *float64 `json:"storageCheckFailedPercentage,omitempty"`
	TotalDeviceCount                        *int64   `json:"totalDeviceCount,omitempty"`
	TpmCheckFailedPercentage                *float64 `json:"tpmCheckFailedPercentage,omitempty"`
	UpgradeEligibleDeviceCount              *int64   `json:"upgradeEligibleDeviceCount,omitempty"`
}
