package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type B2cAuthenticationMethodsPolicyOperationPredicate struct {
	Id                                          *string
	IsEmailPasswordAuthenticationEnabled        *bool
	IsPhoneOneTimePasswordAuthenticationEnabled *bool
	IsUserNameAuthenticationEnabled             *bool
	ODataType                                   *string
}

func (p B2cAuthenticationMethodsPolicyOperationPredicate) Matches(input B2cAuthenticationMethodsPolicy) bool {

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.IsEmailPasswordAuthenticationEnabled != nil && (input.IsEmailPasswordAuthenticationEnabled == nil || *p.IsEmailPasswordAuthenticationEnabled != *input.IsEmailPasswordAuthenticationEnabled) {
		return false
	}

	if p.IsPhoneOneTimePasswordAuthenticationEnabled != nil && (input.IsPhoneOneTimePasswordAuthenticationEnabled == nil || *p.IsPhoneOneTimePasswordAuthenticationEnabled != *input.IsPhoneOneTimePasswordAuthenticationEnabled) {
		return false
	}

	if p.IsUserNameAuthenticationEnabled != nil && (input.IsUserNameAuthenticationEnabled == nil || *p.IsUserNameAuthenticationEnabled != *input.IsUserNameAuthenticationEnabled) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
