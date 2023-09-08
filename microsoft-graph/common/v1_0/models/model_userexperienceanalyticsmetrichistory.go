package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsMetricHistory struct {
	DeviceId       *string `json:"deviceId,omitempty"`
	Id             *string `json:"id,omitempty"`
	MetricDateTime *string `json:"metricDateTime,omitempty"`
	MetricType     *string `json:"metricType,omitempty"`
	ODataType      *string `json:"@odata.type,omitempty"`
}
