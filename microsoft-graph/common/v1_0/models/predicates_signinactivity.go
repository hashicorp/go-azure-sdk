package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SignInActivityOperationPredicate struct {
	LastNonInteractiveSignInDateTime  *string
	LastNonInteractiveSignInRequestId *string
	LastSignInDateTime                *string
	LastSignInRequestId               *string
	ODataType                         *string
}

func (p SignInActivityOperationPredicate) Matches(input SignInActivity) bool {

	if p.LastNonInteractiveSignInDateTime != nil && (input.LastNonInteractiveSignInDateTime == nil || *p.LastNonInteractiveSignInDateTime != *input.LastNonInteractiveSignInDateTime) {
		return false
	}

	if p.LastNonInteractiveSignInRequestId != nil && (input.LastNonInteractiveSignInRequestId == nil || *p.LastNonInteractiveSignInRequestId != *input.LastNonInteractiveSignInRequestId) {
		return false
	}

	if p.LastSignInDateTime != nil && (input.LastSignInDateTime == nil || *p.LastSignInDateTime != *input.LastSignInDateTime) {
		return false
	}

	if p.LastSignInRequestId != nil && (input.LastSignInRequestId == nil || *p.LastSignInRequestId != *input.LastSignInRequestId) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
