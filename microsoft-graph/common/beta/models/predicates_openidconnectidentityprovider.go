package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OpenIdConnectIdentityProviderOperationPredicate struct {
	ClientId     *string
	ClientSecret *string
	DisplayName  *string
	DomainHint   *string
	Id           *string
	MetadataUrl  *string
	ODataType    *string
	Scope        *string
}

func (p OpenIdConnectIdentityProviderOperationPredicate) Matches(input OpenIdConnectIdentityProvider) bool {

	if p.ClientId != nil && (input.ClientId == nil || *p.ClientId != *input.ClientId) {
		return false
	}

	if p.ClientSecret != nil && (input.ClientSecret == nil || *p.ClientSecret != *input.ClientSecret) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.DomainHint != nil && (input.DomainHint == nil || *p.DomainHint != *input.DomainHint) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.MetadataUrl != nil && (input.MetadataUrl == nil || *p.MetadataUrl != *input.MetadataUrl) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Scope != nil && (input.Scope == nil || *p.Scope != *input.Scope) {
		return false
	}

	return true
}
