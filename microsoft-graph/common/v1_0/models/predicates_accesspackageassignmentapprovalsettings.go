package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessPackageAssignmentApprovalSettingsOperationPredicate struct {
	IsApprovalRequiredForAdd    *bool
	IsApprovalRequiredForUpdate *bool
	ODataType                   *string
}

func (p AccessPackageAssignmentApprovalSettingsOperationPredicate) Matches(input AccessPackageAssignmentApprovalSettings) bool {

	if p.IsApprovalRequiredForAdd != nil && (input.IsApprovalRequiredForAdd == nil || *p.IsApprovalRequiredForAdd != *input.IsApprovalRequiredForAdd) {
		return false
	}

	if p.IsApprovalRequiredForUpdate != nil && (input.IsApprovalRequiredForUpdate == nil || *p.IsApprovalRequiredForUpdate != *input.IsApprovalRequiredForUpdate) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
