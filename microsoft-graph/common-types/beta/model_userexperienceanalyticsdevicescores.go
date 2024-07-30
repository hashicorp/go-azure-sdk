package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsDeviceScores struct {
	AppReliabilityScore     *float64                                         `json:"appReliabilityScore,omitempty"`
	BatteryHealthScore      *float64                                         `json:"batteryHealthScore,omitempty"`
	DeviceName              *string                                          `json:"deviceName,omitempty"`
	EndpointAnalyticsScore  *float64                                         `json:"endpointAnalyticsScore,omitempty"`
	HealthStatus            *UserExperienceAnalyticsDeviceScoresHealthStatus `json:"healthStatus,omitempty"`
	Id                      *string                                          `json:"id,omitempty"`
	Manufacturer            *string                                          `json:"manufacturer,omitempty"`
	Model                   *string                                          `json:"model,omitempty"`
	ODataType               *string                                          `json:"@odata.type,omitempty"`
	StartupPerformanceScore *float64                                         `json:"startupPerformanceScore,omitempty"`
	WorkFromAnywhereScore   *float64                                         `json:"workFromAnywhereScore,omitempty"`
}
