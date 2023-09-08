package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsWorkFromAnywhereModelPerformance struct {
	HealthStatus     *UserExperienceAnalyticsWorkFromAnywhereModelPerformanceHealthStatus `json:"healthStatus,omitempty"`
	Id               *string                                                              `json:"id,omitempty"`
	Manufacturer     *string                                                              `json:"manufacturer,omitempty"`
	Model            *string                                                              `json:"model,omitempty"`
	ModelDeviceCount *int64                                                               `json:"modelDeviceCount,omitempty"`
	ODataType        *string                                                              `json:"@odata.type,omitempty"`
}
