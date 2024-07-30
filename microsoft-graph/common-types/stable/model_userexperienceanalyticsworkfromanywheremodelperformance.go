package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsWorkFromAnywhereModelPerformance struct {
	CloudIdentityScore     *float64                                                             `json:"cloudIdentityScore,omitempty"`
	CloudManagementScore   *float64                                                             `json:"cloudManagementScore,omitempty"`
	CloudProvisioningScore *float64                                                             `json:"cloudProvisioningScore,omitempty"`
	HealthStatus           *UserExperienceAnalyticsWorkFromAnywhereModelPerformanceHealthStatus `json:"healthStatus,omitempty"`
	Id                     *string                                                              `json:"id,omitempty"`
	Manufacturer           *string                                                              `json:"manufacturer,omitempty"`
	Model                  *string                                                              `json:"model,omitempty"`
	ModelDeviceCount       *int64                                                               `json:"modelDeviceCount,omitempty"`
	ODataType              *string                                                              `json:"@odata.type,omitempty"`
	WindowsScore           *float64                                                             `json:"windowsScore,omitempty"`
	WorkFromAnywhereScore  *float64                                                             `json:"workFromAnywhereScore,omitempty"`
}
