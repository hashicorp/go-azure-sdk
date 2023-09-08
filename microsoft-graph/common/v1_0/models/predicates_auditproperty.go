package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuditPropertyOperationPredicate struct {
	DisplayName *string
	NewValue    *string
	ODataType   *string
	OldValue    *string
}

func (p AuditPropertyOperationPredicate) Matches(input AuditProperty) bool {

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.NewValue != nil && (input.NewValue == nil || *p.NewValue != *input.NewValue) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.OldValue != nil && (input.OldValue == nil || *p.OldValue != *input.OldValue) {
		return false
	}

	return true
}
