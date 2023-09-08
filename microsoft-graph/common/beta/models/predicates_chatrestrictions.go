package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChatRestrictionsOperationPredicate struct {
	AllowTextOnly *bool
	ODataType     *string
}

func (p ChatRestrictionsOperationPredicate) Matches(input ChatRestrictions) bool {

	if p.AllowTextOnly != nil && (input.AllowTextOnly == nil || *p.AllowTextOnly != *input.AllowTextOnly) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
