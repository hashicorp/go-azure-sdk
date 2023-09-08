package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationStrengthOperationPredicate struct {
	AuthenticationStrengthId *string
	DisplayName              *string
	ODataType                *string
}

func (p AuthenticationStrengthOperationPredicate) Matches(input AuthenticationStrength) bool {

	if p.AuthenticationStrengthId != nil && (input.AuthenticationStrengthId == nil || *p.AuthenticationStrengthId != *input.AuthenticationStrengthId) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
