package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CoachmarkLocationOperationPredicate struct {
	Length    *int64
	ODataType *string
	Offset    *int64
}

func (p CoachmarkLocationOperationPredicate) Matches(input CoachmarkLocation) bool {

	if p.Length != nil && (input.Length == nil || *p.Length != *input.Length) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Offset != nil && (input.Offset == nil || *p.Offset != *input.Offset) {
		return false
	}

	return true
}
