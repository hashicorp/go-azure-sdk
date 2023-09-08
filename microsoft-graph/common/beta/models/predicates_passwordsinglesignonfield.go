package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PasswordSingleSignOnFieldOperationPredicate struct {
	CustomizedLabel *string
	DefaultLabel    *string
	FieldId         *string
	ODataType       *string
	Type            *string
}

func (p PasswordSingleSignOnFieldOperationPredicate) Matches(input PasswordSingleSignOnField) bool {

	if p.CustomizedLabel != nil && (input.CustomizedLabel == nil || *p.CustomizedLabel != *input.CustomizedLabel) {
		return false
	}

	if p.DefaultLabel != nil && (input.DefaultLabel == nil || *p.DefaultLabel != *input.DefaultLabel) {
		return false
	}

	if p.FieldId != nil && (input.FieldId == nil || *p.FieldId != *input.FieldId) {
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
