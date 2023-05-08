package policyexemptions

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PolicyExemptionUpdateProperties struct {
	AssignmentScopeValidation *AssignmentScopeValidation `json:"assignmentScopeValidation,omitempty"`
	ResourceSelectors         *[]ResourceSelector        `json:"resourceSelectors,omitempty"`
}
