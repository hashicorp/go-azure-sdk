package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MicrosoftAuthenticatorAuthenticationMethodConfigurationOperationPredicate struct {
	Id                    *string
	IsSoftwareOathEnabled *bool
	ODataType             *string
}

func (p MicrosoftAuthenticatorAuthenticationMethodConfigurationOperationPredicate) Matches(input MicrosoftAuthenticatorAuthenticationMethodConfiguration) bool {

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.IsSoftwareOathEnabled != nil && (input.IsSoftwareOathEnabled == nil || *p.IsSoftwareOathEnabled != *input.IsSoftwareOathEnabled) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
