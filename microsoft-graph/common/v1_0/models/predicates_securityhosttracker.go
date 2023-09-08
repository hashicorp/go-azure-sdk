package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityHostTrackerOperationPredicate struct {
	FirstSeenDateTime *string
	Id                *string
	Kind              *string
	LastSeenDateTime  *string
	ODataType         *string
	Value             *string
}

func (p SecurityHostTrackerOperationPredicate) Matches(input SecurityHostTracker) bool {

	if p.FirstSeenDateTime != nil && (input.FirstSeenDateTime == nil || *p.FirstSeenDateTime != *input.FirstSeenDateTime) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.Kind != nil && (input.Kind == nil || *p.Kind != *input.Kind) {
		return false
	}

	if p.LastSeenDateTime != nil && (input.LastSeenDateTime == nil || *p.LastSeenDateTime != *input.LastSeenDateTime) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Value != nil && (input.Value == nil || *p.Value != *input.Value) {
		return false
	}

	return true
}
