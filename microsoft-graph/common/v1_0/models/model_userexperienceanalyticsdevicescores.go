package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsDeviceScores struct {
	DeviceName   *string                                          `json:"deviceName,omitempty"`
	HealthStatus *UserExperienceAnalyticsDeviceScoresHealthStatus `json:"healthStatus,omitempty"`
	Id           *string                                          `json:"id,omitempty"`
	Manufacturer *string                                          `json:"manufacturer,omitempty"`
	Model        *string                                          `json:"model,omitempty"`
	ODataType    *string                                          `json:"@odata.type,omitempty"`
}
