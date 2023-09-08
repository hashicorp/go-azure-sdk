package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SelfServiceSignUpAuthenticationFlowConfigurationOperationPredicate struct {
	IsEnabled *bool
	ODataType *string
}

func (p SelfServiceSignUpAuthenticationFlowConfigurationOperationPredicate) Matches(input SelfServiceSignUpAuthenticationFlowConfiguration) bool {

	if p.IsEnabled != nil && (input.IsEnabled == nil || *p.IsEnabled != *input.IsEnabled) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
