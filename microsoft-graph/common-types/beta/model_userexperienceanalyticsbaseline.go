package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsBaseline struct {
	AppHealthMetrics             *UserExperienceAnalyticsCategory `json:"appHealthMetrics,omitempty"`
	BatteryHealthMetrics         *UserExperienceAnalyticsCategory `json:"batteryHealthMetrics,omitempty"`
	BestPracticesMetrics         *UserExperienceAnalyticsCategory `json:"bestPracticesMetrics,omitempty"`
	CreatedDateTime              *string                          `json:"createdDateTime,omitempty"`
	DeviceBootPerformanceMetrics *UserExperienceAnalyticsCategory `json:"deviceBootPerformanceMetrics,omitempty"`
	DisplayName                  *string                          `json:"displayName,omitempty"`
	Id                           *string                          `json:"id,omitempty"`
	IsBuiltIn                    *bool                            `json:"isBuiltIn,omitempty"`
	ODataType                    *string                          `json:"@odata.type,omitempty"`
	RebootAnalyticsMetrics       *UserExperienceAnalyticsCategory `json:"rebootAnalyticsMetrics,omitempty"`
	ResourcePerformanceMetrics   *UserExperienceAnalyticsCategory `json:"resourcePerformanceMetrics,omitempty"`
	WorkFromAnywhereMetrics      *UserExperienceAnalyticsCategory `json:"workFromAnywhereMetrics,omitempty"`
}
