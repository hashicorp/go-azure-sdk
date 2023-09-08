package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsImpactingProcess struct {
	Category    *string `json:"category,omitempty"`
	Description *string `json:"description,omitempty"`
	DeviceId    *string `json:"deviceId,omitempty"`
	Id          *string `json:"id,omitempty"`
	ODataType   *string `json:"@odata.type,omitempty"`
	ProcessName *string `json:"processName,omitempty"`
	Publisher   *string `json:"publisher,omitempty"`
}
