package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExpirationPatternOperationPredicate struct {
	Duration    *string
	EndDateTime *string
	ODataType   *string
}

func (p ExpirationPatternOperationPredicate) Matches(input ExpirationPattern) bool {

	if p.Duration != nil && (input.Duration == nil || *p.Duration != *input.Duration) {
		return false
	}

	if p.EndDateTime != nil && (input.EndDateTime == nil || *p.EndDateTime != *input.EndDateTime) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
