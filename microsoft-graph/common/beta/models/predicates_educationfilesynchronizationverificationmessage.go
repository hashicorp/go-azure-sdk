package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EducationFileSynchronizationVerificationMessageOperationPredicate struct {
	Description *string
	FileName    *string
	ODataType   *string
	Type        *string
}

func (p EducationFileSynchronizationVerificationMessageOperationPredicate) Matches(input EducationFileSynchronizationVerificationMessage) bool {

	if p.Description != nil && (input.Description == nil || *p.Description != *input.Description) {
		return false
	}

	if p.FileName != nil && (input.FileName == nil || *p.FileName != *input.FileName) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Type != nil && (input.Type == nil || *p.Type != *input.Type) {
		return false
	}

	return true
}
