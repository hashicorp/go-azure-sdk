package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FollowupFlagOperationPredicate struct {
	ODataType *string
}

func (p FollowupFlagOperationPredicate) Matches(input FollowupFlag) bool {

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
