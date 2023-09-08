package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VoiceAuthenticationMethodTargetOperationPredicate struct {
	Id                     *string
	IsRegistrationRequired *bool
	ODataType              *string
}

func (p VoiceAuthenticationMethodTargetOperationPredicate) Matches(input VoiceAuthenticationMethodTarget) bool {

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.IsRegistrationRequired != nil && (input.IsRegistrationRequired == nil || *p.IsRegistrationRequired != *input.IsRegistrationRequired) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
