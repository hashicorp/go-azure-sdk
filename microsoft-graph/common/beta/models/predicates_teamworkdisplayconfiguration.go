package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamworkDisplayConfigurationOperationPredicate struct {
	DisplayCount                *int64
	IsContentDuplicationAllowed *bool
	IsDualDisplayModeEnabled    *bool
	ODataType                   *string
}

func (p TeamworkDisplayConfigurationOperationPredicate) Matches(input TeamworkDisplayConfiguration) bool {

	if p.DisplayCount != nil && (input.DisplayCount == nil || *p.DisplayCount != *input.DisplayCount) {
		return false
	}

	if p.IsContentDuplicationAllowed != nil && (input.IsContentDuplicationAllowed == nil || *p.IsContentDuplicationAllowed != *input.IsContentDuplicationAllowed) {
		return false
	}

	if p.IsDualDisplayModeEnabled != nil && (input.IsDualDisplayModeEnabled == nil || *p.IsDualDisplayModeEnabled != *input.IsDualDisplayModeEnabled) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
