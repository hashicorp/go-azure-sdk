package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityHostPairOperationPredicate struct {
	FirstSeenDateTime *string
	Id                *string
	LastSeenDateTime  *string
	LinkKind          *string
	ODataType         *string
}

func (p SecurityHostPairOperationPredicate) Matches(input SecurityHostPair) bool {

	if p.FirstSeenDateTime != nil && (input.FirstSeenDateTime == nil || *p.FirstSeenDateTime != *input.FirstSeenDateTime) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.LastSeenDateTime != nil && (input.LastSeenDateTime == nil || *p.LastSeenDateTime != *input.LastSeenDateTime) {
		return false
	}

	if p.LinkKind != nil && (input.LinkKind == nil || *p.LinkKind != *input.LinkKind) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
