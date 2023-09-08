package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserRegistrationMethodCountOperationPredicate struct {
	AuthenticationMethod *string
	ODataType            *string
	UserCount            *int64
}

func (p UserRegistrationMethodCountOperationPredicate) Matches(input UserRegistrationMethodCount) bool {

	if p.AuthenticationMethod != nil && (input.AuthenticationMethod == nil || *p.AuthenticationMethod != *input.AuthenticationMethod) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.UserCount != nil && (input.UserCount == nil || *p.UserCount != *input.UserCount) {
		return false
	}

	return true
}
