package standardassignments

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StandardAssignmentPropertiesExemptionData struct {
	AssignedAssessment *AssignedAssessmentItem `json:"assignedAssessment,omitempty"`
	ExemptionCategory  *ExemptionCategory      `json:"exemptionCategory,omitempty"`
}
