package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivacyProfileOperationPredicate struct {
	ContactEmail *string
	ODataType    *string
	StatementUrl *string
}

func (p PrivacyProfileOperationPredicate) Matches(input PrivacyProfile) bool {

	if p.ContactEmail != nil && (input.ContactEmail == nil || *p.ContactEmail != *input.ContactEmail) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.StatementUrl != nil && (input.StatementUrl == nil || *p.StatementUrl != *input.StatementUrl) {
		return false
	}

	return true
}
