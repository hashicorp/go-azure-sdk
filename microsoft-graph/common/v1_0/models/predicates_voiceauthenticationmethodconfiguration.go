package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VoiceAuthenticationMethodConfigurationOperationPredicate struct {
	Id                   *string
	IsOfficePhoneAllowed *bool
	ODataType            *string
}

func (p VoiceAuthenticationMethodConfigurationOperationPredicate) Matches(input VoiceAuthenticationMethodConfiguration) bool {

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.IsOfficePhoneAllowed != nil && (input.IsOfficePhoneAllowed == nil || *p.IsOfficePhoneAllowed != *input.IsOfficePhoneAllowed) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
