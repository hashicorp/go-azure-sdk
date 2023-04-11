package subassessments

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SubAssessmentStatus struct {
	Cause       *string                  `json:"cause,omitempty"`
	Code        *SubAssessmentStatusCode `json:"code,omitempty"`
	Description *string                  `json:"description,omitempty"`
	Severity    *Severity                `json:"severity,omitempty"`
}
