package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SharedOperationPredicate struct {
	ODataType      *string
	Scope          *string
	SharedDateTime *string
}

func (p SharedOperationPredicate) Matches(input Shared) bool {

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Scope != nil && (input.Scope == nil || *p.Scope != *input.Scope) {
		return false
	}

	if p.SharedDateTime != nil && (input.SharedDateTime == nil || *p.SharedDateTime != *input.SharedDateTime) {
		return false
	}

	return true
}
