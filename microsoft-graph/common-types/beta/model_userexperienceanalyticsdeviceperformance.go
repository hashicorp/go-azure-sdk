package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsDevicePerformance struct {
	AverageBlueScreens           *float64                                              `json:"averageBlueScreens,omitempty"`
	AverageRestarts              *float64                                              `json:"averageRestarts,omitempty"`
	BlueScreenCount              *int64                                                `json:"blueScreenCount,omitempty"`
	BootScore                    *int64                                                `json:"bootScore,omitempty"`
	CoreBootTimeInMs             *int64                                                `json:"coreBootTimeInMs,omitempty"`
	CoreLoginTimeInMs            *int64                                                `json:"coreLoginTimeInMs,omitempty"`
	DeviceCount                  *int64                                                `json:"deviceCount,omitempty"`
	DeviceName                   *string                                               `json:"deviceName,omitempty"`
	DiskType                     *UserExperienceAnalyticsDevicePerformanceDiskType     `json:"diskType,omitempty"`
	GroupPolicyBootTimeInMs      *int64                                                `json:"groupPolicyBootTimeInMs,omitempty"`
	GroupPolicyLoginTimeInMs     *int64                                                `json:"groupPolicyLoginTimeInMs,omitempty"`
	HealthStatus                 *UserExperienceAnalyticsDevicePerformanceHealthStatus `json:"healthStatus,omitempty"`
	Id                           *string                                               `json:"id,omitempty"`
	LoginScore                   *int64                                                `json:"loginScore,omitempty"`
	Manufacturer                 *string                                               `json:"manufacturer,omitempty"`
	Model                        *string                                               `json:"model,omitempty"`
	ModelStartupPerformanceScore *float64                                              `json:"modelStartupPerformanceScore,omitempty"`
	ODataType                    *string                                               `json:"@odata.type,omitempty"`
	OperatingSystemVersion       *string                                               `json:"operatingSystemVersion,omitempty"`
	ResponsiveDesktopTimeInMs    *int64                                                `json:"responsiveDesktopTimeInMs,omitempty"`
	RestartCount                 *int64                                                `json:"restartCount,omitempty"`
	StartupPerformanceScore      *float64                                              `json:"startupPerformanceScore,omitempty"`
}
