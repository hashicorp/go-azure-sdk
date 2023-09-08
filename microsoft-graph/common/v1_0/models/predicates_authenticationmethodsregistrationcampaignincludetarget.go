package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationMethodsRegistrationCampaignIncludeTargetOperationPredicate struct {
	Id                           *string
	ODataType                    *string
	TargetedAuthenticationMethod *string
}

func (p AuthenticationMethodsRegistrationCampaignIncludeTargetOperationPredicate) Matches(input AuthenticationMethodsRegistrationCampaignIncludeTarget) bool {

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.TargetedAuthenticationMethod != nil && (input.TargetedAuthenticationMethod == nil || *p.TargetedAuthenticationMethod != *input.TargetedAuthenticationMethod) {
		return false
	}

	return true
}
