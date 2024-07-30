package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsDeviceStartupHistory struct {
	CoreBootTimeInMs          *int64                                                      `json:"coreBootTimeInMs,omitempty"`
	CoreLoginTimeInMs         *int64                                                      `json:"coreLoginTimeInMs,omitempty"`
	DeviceId                  *string                                                     `json:"deviceId,omitempty"`
	FeatureUpdateBootTimeInMs *int64                                                      `json:"featureUpdateBootTimeInMs,omitempty"`
	GroupPolicyBootTimeInMs   *int64                                                      `json:"groupPolicyBootTimeInMs,omitempty"`
	GroupPolicyLoginTimeInMs  *int64                                                      `json:"groupPolicyLoginTimeInMs,omitempty"`
	Id                        *string                                                     `json:"id,omitempty"`
	IsFeatureUpdate           *bool                                                       `json:"isFeatureUpdate,omitempty"`
	IsFirstLogin              *bool                                                       `json:"isFirstLogin,omitempty"`
	ODataType                 *string                                                     `json:"@odata.type,omitempty"`
	OperatingSystemVersion    *string                                                     `json:"operatingSystemVersion,omitempty"`
	ResponsiveDesktopTimeInMs *int64                                                      `json:"responsiveDesktopTimeInMs,omitempty"`
	RestartCategory           *UserExperienceAnalyticsDeviceStartupHistoryRestartCategory `json:"restartCategory,omitempty"`
	RestartFaultBucket        *string                                                     `json:"restartFaultBucket,omitempty"`
	RestartStopCode           *string                                                     `json:"restartStopCode,omitempty"`
	StartTime                 *string                                                     `json:"startTime,omitempty"`
	TotalBootTimeInMs         *int64                                                      `json:"totalBootTimeInMs,omitempty"`
	TotalLoginTimeInMs        *int64                                                      `json:"totalLoginTimeInMs,omitempty"`
}
