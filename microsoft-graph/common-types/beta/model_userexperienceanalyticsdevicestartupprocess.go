package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsDeviceStartupProcess struct {
	Id                *string `json:"id,omitempty"`
	ManagedDeviceId   *string `json:"managedDeviceId,omitempty"`
	ODataType         *string `json:"@odata.type,omitempty"`
	ProcessName       *string `json:"processName,omitempty"`
	ProductName       *string `json:"productName,omitempty"`
	Publisher         *string `json:"publisher,omitempty"`
	StartupImpactInMs *int64  `json:"startupImpactInMs,omitempty"`
}
