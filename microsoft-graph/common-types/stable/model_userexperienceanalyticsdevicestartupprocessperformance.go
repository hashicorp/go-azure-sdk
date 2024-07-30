package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsDeviceStartupProcessPerformance struct {
	DeviceCount      *int64  `json:"deviceCount,omitempty"`
	Id               *string `json:"id,omitempty"`
	MedianImpactInMs *int64  `json:"medianImpactInMs,omitempty"`
	ODataType        *string `json:"@odata.type,omitempty"`
	ProcessName      *string `json:"processName,omitempty"`
	ProductName      *string `json:"productName,omitempty"`
	Publisher        *string `json:"publisher,omitempty"`
	TotalImpactInMs  *int64  `json:"totalImpactInMs,omitempty"`
}
