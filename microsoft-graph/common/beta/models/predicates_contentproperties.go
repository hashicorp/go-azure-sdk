package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ContentPropertiesOperationPredicate struct {
	LastModifiedBy       *string
	LastModifiedDateTime *string
	ODataType            *string
}

func (p ContentPropertiesOperationPredicate) Matches(input ContentProperties) bool {

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
