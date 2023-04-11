package iotsecuritysolutionsanalytics

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IoTSecurityDeviceRecommendation struct {
	DevicesCount              *int64            `json:"devicesCount,omitempty"`
	RecommendationDisplayName *string           `json:"recommendationDisplayName,omitempty"`
	ReportedSeverity          *ReportedSeverity `json:"reportedSeverity,omitempty"`
}
