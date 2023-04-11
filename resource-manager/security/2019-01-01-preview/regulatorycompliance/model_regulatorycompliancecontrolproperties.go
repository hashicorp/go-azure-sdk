package regulatorycompliance

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RegulatoryComplianceControlProperties struct {
	Description        *string `json:"description,omitempty"`
	FailedAssessments  *int64  `json:"failedAssessments,omitempty"`
	PassedAssessments  *int64  `json:"passedAssessments,omitempty"`
	SkippedAssessments *int64  `json:"skippedAssessments,omitempty"`
	State              *State  `json:"state,omitempty"`
}
