package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SignInPreferencesOperationPredicate struct {
	IsSystemPreferredAuthenticationMethodEnabled *bool
	ODataType                                    *string
}

func (p SignInPreferencesOperationPredicate) Matches(input SignInPreferences) bool {

	if p.IsSystemPreferredAuthenticationMethodEnabled != nil && (input.IsSystemPreferredAuthenticationMethodEnabled == nil || *p.IsSystemPreferredAuthenticationMethodEnabled != *input.IsSystemPreferredAuthenticationMethodEnabled) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
