package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExactMatchUploadAgentOperationPredicate struct {
	CreationDateTime *string
	Description      *string
	Id               *string
	ODataType        *string
}

func (p ExactMatchUploadAgentOperationPredicate) Matches(input ExactMatchUploadAgent) bool {

	if p.CreationDateTime != nil && (input.CreationDateTime == nil || *p.CreationDateTime != *input.CreationDateTime) {
		return false
	}

	if p.Description != nil && (input.Description == nil || *p.Description != *input.Description) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
