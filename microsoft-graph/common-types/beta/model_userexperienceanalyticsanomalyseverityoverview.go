package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsAnomalySeverityOverview struct {
	HighSeverityAnomalyCount          *int64  `json:"highSeverityAnomalyCount,omitempty"`
	InformationalSeverityAnomalyCount *int64  `json:"informationalSeverityAnomalyCount,omitempty"`
	LowSeverityAnomalyCount           *int64  `json:"lowSeverityAnomalyCount,omitempty"`
	MediumSeverityAnomalyCount        *int64  `json:"mediumSeverityAnomalyCount,omitempty"`
	ODataType                         *string `json:"@odata.type,omitempty"`
}
