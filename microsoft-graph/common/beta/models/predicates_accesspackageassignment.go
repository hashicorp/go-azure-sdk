package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessPackageAssignmentOperationPredicate struct {
	AccessPackageId    *string
	AssignmentPolicyId *string
	AssignmentState    *string
	AssignmentStatus   *string
	CatalogId          *string
	ExpiredDateTime    *string
	Id                 *string
	IsExtended         *bool
	ODataType          *string
	TargetId           *string
}

func (p AccessPackageAssignmentOperationPredicate) Matches(input AccessPackageAssignment) bool {

	if p.AccessPackageId != nil && (input.AccessPackageId == nil || *p.AccessPackageId != *input.AccessPackageId) {
		return false
	}

	if p.AssignmentPolicyId != nil && (input.AssignmentPolicyId == nil || *p.AssignmentPolicyId != *input.AssignmentPolicyId) {
		return false
	}

	if p.AssignmentState != nil && (input.AssignmentState == nil || *p.AssignmentState != *input.AssignmentState) {
		return false
	}

	if p.AssignmentStatus != nil && (input.AssignmentStatus == nil || *p.AssignmentStatus != *input.AssignmentStatus) {
		return false
	}

	if p.CatalogId != nil && (input.CatalogId == nil || *p.CatalogId != *input.CatalogId) {
		return false
	}

	if p.ExpiredDateTime != nil && (input.ExpiredDateTime == nil || *p.ExpiredDateTime != *input.ExpiredDateTime) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.IsExtended != nil && (input.IsExtended == nil || *p.IsExtended != *input.IsExtended) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.TargetId != nil && (input.TargetId == nil || *p.TargetId != *input.TargetId) {
		return false
	}

	return true
}
