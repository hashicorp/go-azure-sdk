package iotsecuritysolutionsanalytics

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IoTSecurityAggregatedAlertProperties struct {
	ActionTaken          *string           `json:"actionTaken,omitempty"`
	AggregatedDateUtc    *string           `json:"aggregatedDateUtc,omitempty"`
	AlertDisplayName     *string           `json:"alertDisplayName,omitempty"`
	AlertType            *string           `json:"alertType,omitempty"`
	Count                *int64            `json:"count,omitempty"`
	Description          *string           `json:"description,omitempty"`
	EffectedResourceType *string           `json:"effectedResourceType,omitempty"`
	LogAnalyticsQuery    *string           `json:"logAnalyticsQuery,omitempty"`
	RemediationSteps     *string           `json:"remediationSteps,omitempty"`
	ReportedSeverity     *ReportedSeverity `json:"reportedSeverity,omitempty"`
	SystemSource         *string           `json:"systemSource,omitempty"`
	VendorName           *string           `json:"vendorName,omitempty"`
}
