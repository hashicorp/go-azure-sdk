package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsAnomalyCorrelationGroupFeature struct {
	DeviceFeatureType *UserExperienceAnalyticsAnomalyCorrelationGroupFeatureDeviceFeatureType `json:"deviceFeatureType,omitempty"`
	ODataType         *string                                                                 `json:"@odata.type,omitempty"`
	Values            *[]string                                                               `json:"values,omitempty"`
}
