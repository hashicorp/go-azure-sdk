package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserRegistrationCountOperationPredicate struct {
	ODataType         *string
	RegistrationCount *int64
}

func (p UserRegistrationCountOperationPredicate) Matches(input UserRegistrationCount) bool {

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.RegistrationCount != nil && (input.RegistrationCount == nil || *p.RegistrationCount != *input.RegistrationCount) {
		return false
	}

	return true
}
