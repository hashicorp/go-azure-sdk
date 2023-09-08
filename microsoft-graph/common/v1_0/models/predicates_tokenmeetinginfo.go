package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TokenMeetingInfoOperationPredicate struct {
	ODataType *string
	Token     *string
}

func (p TokenMeetingInfoOperationPredicate) Matches(input TokenMeetingInfo) bool {

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Token != nil && (input.Token == nil || *p.Token != *input.Token) {
		return false
	}

	return true
}
