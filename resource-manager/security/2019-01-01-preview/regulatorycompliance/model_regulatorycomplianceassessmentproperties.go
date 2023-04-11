package regulatorycompliance

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RegulatoryComplianceAssessmentProperties struct {
	AssessmentDetailsLink *string `json:"assessmentDetailsLink,omitempty"`
	AssessmentType        *string `json:"assessmentType,omitempty"`
	Description           *string `json:"description,omitempty"`
	FailedResources       *int64  `json:"failedResources,omitempty"`
	PassedResources       *int64  `json:"passedResources,omitempty"`
	SkippedResources      *int64  `json:"skippedResources,omitempty"`
	State                 *State  `json:"state,omitempty"`
	UnsupportedResources  *int64  `json:"unsupportedResources,omitempty"`
}
