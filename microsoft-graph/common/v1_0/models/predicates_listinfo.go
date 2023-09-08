package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListInfoOperationPredicate struct {
	ContentTypesEnabled *bool
	Hidden              *bool
	ODataType           *string
	Template            *string
}

func (p ListInfoOperationPredicate) Matches(input ListInfo) bool {

	if p.ContentTypesEnabled != nil && (input.ContentTypesEnabled == nil || *p.ContentTypesEnabled != *input.ContentTypesEnabled) {
		return false
	}

	if p.Hidden != nil && (input.Hidden == nil || *p.Hidden != *input.Hidden) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Template != nil && (input.Template == nil || *p.Template != *input.Template) {
		return false
	}

	return true
}
