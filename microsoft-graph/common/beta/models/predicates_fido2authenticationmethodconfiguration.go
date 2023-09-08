package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Fido2AuthenticationMethodConfigurationOperationPredicate struct {
	Id                               *string
	IsAttestationEnforced            *bool
	IsSelfServiceRegistrationAllowed *bool
	ODataType                        *string
}

func (p Fido2AuthenticationMethodConfigurationOperationPredicate) Matches(input Fido2AuthenticationMethodConfiguration) bool {

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.IsAttestationEnforced != nil && (input.IsAttestationEnforced == nil || *p.IsAttestationEnforced != *input.IsAttestationEnforced) {
		return false
	}

	if p.IsSelfServiceRegistrationAllowed != nil && (input.IsSelfServiceRegistrationAllowed == nil || *p.IsSelfServiceRegistrationAllowed != *input.IsSelfServiceRegistrationAllowed) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
