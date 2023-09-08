package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessPackageAssignmentResourceRoleOperationPredicate struct {
	Id           *string
	ODataType    *string
	OriginId     *string
	OriginSystem *string
	Status       *string
}

func (p AccessPackageAssignmentResourceRoleOperationPredicate) Matches(input AccessPackageAssignmentResourceRole) bool {

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.OriginId != nil && (input.OriginId == nil || *p.OriginId != *input.OriginId) {
		return false
	}

	if p.OriginSystem != nil && (input.OriginSystem == nil || *p.OriginSystem != *input.OriginSystem) {
		return false
	}

	if p.Status != nil && (input.Status == nil || *p.Status != *input.Status) {
		return false
	}

	return true
}
