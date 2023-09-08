package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SessionLifetimePolicyOperationPredicate struct {
	Detail    *string
	ODataType *string
}

func (p SessionLifetimePolicyOperationPredicate) Matches(input SessionLifetimePolicy) bool {

	if p.Detail != nil && (input.Detail == nil || *p.Detail != *input.Detail) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
