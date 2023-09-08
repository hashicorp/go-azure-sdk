package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationDetailOperationPredicate struct {
	AuthenticationMethod           *string
	AuthenticationMethodDetail     *string
	AuthenticationStepDateTime     *string
	AuthenticationStepRequirement  *string
	AuthenticationStepResultDetail *string
	ODataType                      *string
	Succeeded                      *bool
}

func (p AuthenticationDetailOperationPredicate) Matches(input AuthenticationDetail) bool {

	if p.AuthenticationMethod != nil && (input.AuthenticationMethod == nil || *p.AuthenticationMethod != *input.AuthenticationMethod) {
		return false
	}

	if p.AuthenticationMethodDetail != nil && (input.AuthenticationMethodDetail == nil || *p.AuthenticationMethodDetail != *input.AuthenticationMethodDetail) {
		return false
	}

	if p.AuthenticationStepDateTime != nil && (input.AuthenticationStepDateTime == nil || *p.AuthenticationStepDateTime != *input.AuthenticationStepDateTime) {
		return false
	}

	if p.AuthenticationStepRequirement != nil && (input.AuthenticationStepRequirement == nil || *p.AuthenticationStepRequirement != *input.AuthenticationStepRequirement) {
		return false
	}

	if p.AuthenticationStepResultDetail != nil && (input.AuthenticationStepResultDetail == nil || *p.AuthenticationStepResultDetail != *input.AuthenticationStepResultDetail) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Succeeded != nil && (input.Succeeded == nil || *p.Succeeded != *input.Succeeded) {
		return false
	}

	return true
}
