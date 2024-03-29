package assessments

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AssessmentStatus struct {
	Cause       *string              `json:"cause,omitempty"`
	Code        AssessmentStatusCode `json:"code"`
	Description *string              `json:"description,omitempty"`
}
