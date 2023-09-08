package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessPackageAssignmentWorkflowExtensionOperationPredicate struct {
	CreatedBy            *string
	CreatedDateTime      *string
	Description          *string
	DisplayName          *string
	Id                   *string
	LastModifiedBy       *string
	LastModifiedDateTime *string
	ODataType            *string
}

func (p AccessPackageAssignmentWorkflowExtensionOperationPredicate) Matches(input AccessPackageAssignmentWorkflowExtension) bool {

	if p.CreatedBy != nil && (input.CreatedBy == nil || *p.CreatedBy != *input.CreatedBy) {
		return false
	}

	if p.CreatedDateTime != nil && (input.CreatedDateTime == nil || *p.CreatedDateTime != *input.CreatedDateTime) {
		return false
	}

	if p.Description != nil && (input.Description == nil || *p.Description != *input.Description) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.LastModifiedBy != nil && (input.LastModifiedBy == nil || *p.LastModifiedBy != *input.LastModifiedBy) {
		return false
	}

	if p.LastModifiedDateTime != nil && (input.LastModifiedDateTime == nil || *p.LastModifiedDateTime != *input.LastModifiedDateTime) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
