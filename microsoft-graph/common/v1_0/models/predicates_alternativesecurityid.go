package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AlternativeSecurityIdOperationPredicate struct {
	IdentityProvider *string
	Key              *string
	ODataType        *string
	Type             *int64
}

func (p AlternativeSecurityIdOperationPredicate) Matches(input AlternativeSecurityId) bool {

	if p.IdentityProvider != nil && (input.IdentityProvider == nil || *p.IdentityProvider != *input.IdentityProvider) {
		return false
	}

	if p.Key != nil && (input.Key == nil || *p.Key != *input.Key) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Type != nil && (input.Type == nil || *p.Type != *input.Type) {
		return false
	}

	return true
}
