package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessPackageAssignmentOperationPredicate struct {
	ExpiredDateTime *string
	Id              *string
	ODataType       *string
	Status          *string
}

func (p AccessPackageAssignmentOperationPredicate) Matches(input AccessPackageAssignment) bool {

	if p.ExpiredDateTime != nil && (input.ExpiredDateTime == nil || *p.ExpiredDateTime != *input.ExpiredDateTime) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Status != nil && (input.Status == nil || *p.Status != *input.Status) {
		return false
	}

	return true
}
