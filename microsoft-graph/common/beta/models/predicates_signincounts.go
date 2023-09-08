package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SignInCountsOperationPredicate struct {
	NoSignIn   *int64
	ODataType  *string
	WithSignIn *int64
}

func (p SignInCountsOperationPredicate) Matches(input SignInCounts) bool {

	if p.NoSignIn != nil && (input.NoSignIn == nil || *p.NoSignIn != *input.NoSignIn) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.WithSignIn != nil && (input.WithSignIn == nil || *p.WithSignIn != *input.WithSignIn) {
		return false
	}

	return true
}
