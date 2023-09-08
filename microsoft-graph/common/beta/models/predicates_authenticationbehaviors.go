package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationBehaviorsOperationPredicate struct {
	ODataType                     *string
	RemoveUnverifiedEmailClaim    *bool
	RequireClientServicePrincipal *bool
}

func (p AuthenticationBehaviorsOperationPredicate) Matches(input AuthenticationBehaviors) bool {

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.RemoveUnverifiedEmailClaim != nil && (input.RemoveUnverifiedEmailClaim == nil || *p.RemoveUnverifiedEmailClaim != *input.RemoveUnverifiedEmailClaim) {
		return false
	}

	if p.RequireClientServicePrincipal != nil && (input.RequireClientServicePrincipal == nil || *p.RequireClientServicePrincipal != *input.RequireClientServicePrincipal) {
		return false
	}

	return true
}
