package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RemoteLockActionResultOperationPredicate struct {
	ActionName          *string
	LastUpdatedDateTime *string
	ODataType           *string
	StartDateTime       *string
	UnlockPin           *string
}

func (p RemoteLockActionResultOperationPredicate) Matches(input RemoteLockActionResult) bool {

	if p.ActionName != nil && (input.ActionName == nil || *p.ActionName != *input.ActionName) {
		return false
	}

	if p.LastUpdatedDateTime != nil && (input.LastUpdatedDateTime == nil || *p.LastUpdatedDateTime != *input.LastUpdatedDateTime) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.StartDateTime != nil && (input.StartDateTime == nil || *p.StartDateTime != *input.StartDateTime) {
		return false
	}

	if p.UnlockPin != nil && (input.UnlockPin == nil || *p.UnlockPin != *input.UnlockPin) {
		return false
	}

	return true
}
