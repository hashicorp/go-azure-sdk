package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessPackageAssignmentRequestRequirementsOperationPredicate struct {
	IsApprovalRequired                *bool
	IsApprovalRequiredForExtension    *bool
	IsCustomAssignmentScheduleAllowed *bool
	IsRequestorJustificationRequired  *bool
	ODataType                         *string
	PolicyDescription                 *string
	PolicyDisplayName                 *string
	PolicyId                          *string
}

func (p AccessPackageAssignmentRequestRequirementsOperationPredicate) Matches(input AccessPackageAssignmentRequestRequirements) bool {

	if p.IsApprovalRequired != nil && (input.IsApprovalRequired == nil || *p.IsApprovalRequired != *input.IsApprovalRequired) {
		return false
	}

	if p.IsApprovalRequiredForExtension != nil && (input.IsApprovalRequiredForExtension == nil || *p.IsApprovalRequiredForExtension != *input.IsApprovalRequiredForExtension) {
		return false
	}

	if p.IsCustomAssignmentScheduleAllowed != nil && (input.IsCustomAssignmentScheduleAllowed == nil || *p.IsCustomAssignmentScheduleAllowed != *input.IsCustomAssignmentScheduleAllowed) {
		return false
	}

	if p.IsRequestorJustificationRequired != nil && (input.IsRequestorJustificationRequired == nil || *p.IsRequestorJustificationRequired != *input.IsRequestorJustificationRequired) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.PolicyDescription != nil && (input.PolicyDescription == nil || *p.PolicyDescription != *input.PolicyDescription) {
		return false
	}

	if p.PolicyDisplayName != nil && (input.PolicyDisplayName == nil || *p.PolicyDisplayName != *input.PolicyDisplayName) {
		return false
	}

	if p.PolicyId != nil && (input.PolicyId == nil || *p.PolicyId != *input.PolicyId) {
		return false
	}

	return true
}
