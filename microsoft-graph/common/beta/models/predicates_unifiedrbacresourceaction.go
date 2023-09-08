package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UnifiedRbacResourceActionOperationPredicate struct {
	ActionVerb                      *string
	AuthenticationContextId         *string
	Description                     *string
	Id                              *string
	IsAuthenticationContextSettable *bool
	IsPrivileged                    *bool
	Name                            *string
	ODataType                       *string
	ResourceScopeId                 *string
}

func (p UnifiedRbacResourceActionOperationPredicate) Matches(input UnifiedRbacResourceAction) bool {

	if p.ActionVerb != nil && (input.ActionVerb == nil || *p.ActionVerb != *input.ActionVerb) {
		return false
	}

	if p.AuthenticationContextId != nil && (input.AuthenticationContextId == nil || *p.AuthenticationContextId != *input.AuthenticationContextId) {
		return false
	}

	if p.Description != nil && (input.Description == nil || *p.Description != *input.Description) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.IsAuthenticationContextSettable != nil && (input.IsAuthenticationContextSettable == nil || *p.IsAuthenticationContextSettable != *input.IsAuthenticationContextSettable) {
		return false
	}

	if p.IsPrivileged != nil && (input.IsPrivileged == nil || *p.IsPrivileged != *input.IsPrivileged) {
		return false
	}

	if p.Name != nil && (input.Name == nil || *p.Name != *input.Name) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.ResourceScopeId != nil && (input.ResourceScopeId == nil || *p.ResourceScopeId != *input.ResourceScopeId) {
		return false
	}

	return true
}
