package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationAppPolicyDetailsOperationPredicate struct {
	ODataType  *string
	PolicyName *string
}

func (p AuthenticationAppPolicyDetailsOperationPredicate) Matches(input AuthenticationAppPolicyDetails) bool {

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.PolicyName != nil && (input.PolicyName == nil || *p.PolicyName != *input.PolicyName) {
		return false
	}

	return true
}
