package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CallStartedEventMessageDetailOperationPredicate struct {
	CallId    *string
	ODataType *string
}

func (p CallStartedEventMessageDetailOperationPredicate) Matches(input CallStartedEventMessageDetail) bool {

	if p.CallId != nil && (input.CallId == nil || *p.CallId != *input.CallId) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
