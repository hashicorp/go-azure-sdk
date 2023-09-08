package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppliedAuthenticationEventListenerOperationPredicate struct {
	ExecutedListenerId *string
	ODataType          *string
}

func (p AppliedAuthenticationEventListenerOperationPredicate) Matches(input AppliedAuthenticationEventListener) bool {

	if p.ExecutedListenerId != nil && (input.ExecutedListenerId == nil || *p.ExecutedListenerId != *input.ExecutedListenerId) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
