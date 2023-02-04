package hypervsites

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteHealthSummary struct {
	AffectedObjectsCount *int64    `json:"affectedObjectsCount,omitempty"`
	AffectedResourceType *string   `json:"affectedResourceType,omitempty"`
	AffectedResources    *[]string `json:"affectedResources,omitempty"`
	ApplianceName        *string   `json:"applianceName,omitempty"`
	ErrorCode            *string   `json:"errorCode,omitempty"`
	ErrorId              *int64    `json:"errorId,omitempty"`
	ErrorMessage         *string   `json:"errorMessage,omitempty"`
	HitCount             *int64    `json:"hitCount,omitempty"`
	RemediationGuidance  *string   `json:"remediationGuidance,omitempty"`
	Severity             *string   `json:"severity,omitempty"`
	SummaryMessage       *string   `json:"summaryMessage,omitempty"`
}
