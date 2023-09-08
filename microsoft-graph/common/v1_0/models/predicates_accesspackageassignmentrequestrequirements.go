package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessPackageAssignmentRequestRequirementsOperationPredicate struct {
	AllowCustomAssignmentSchedule *bool
	IsApprovalRequiredForAdd      *bool
	IsApprovalRequiredForUpdate   *bool
	ODataType                     *string
	PolicyDescription             *string
	PolicyDisplayName             *string
	PolicyId                      *string
}

func (p AccessPackageAssignmentRequestRequirementsOperationPredicate) Matches(input AccessPackageAssignmentRequestRequirements) bool {

	if p.AllowCustomAssignmentSchedule != nil && (input.AllowCustomAssignmentSchedule == nil || *p.AllowCustomAssignmentSchedule != *input.AllowCustomAssignmentSchedule) {
		return false
	}

	if p.IsApprovalRequiredForAdd != nil && (input.IsApprovalRequiredForAdd == nil || *p.IsApprovalRequiredForAdd != *input.IsApprovalRequiredForAdd) {
		return false
	}

	if p.IsApprovalRequiredForUpdate != nil && (input.IsApprovalRequiredForUpdate == nil || *p.IsApprovalRequiredForUpdate != *input.IsApprovalRequiredForUpdate) {
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
