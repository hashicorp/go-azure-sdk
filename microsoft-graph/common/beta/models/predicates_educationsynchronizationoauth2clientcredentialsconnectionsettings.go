package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EducationSynchronizationOAuth2ClientCredentialsConnectionSettingsOperationPredicate struct {
	ClientId     *string
	ClientSecret *string
	ODataType    *string
	Scope        *string
	TokenUrl     *string
}

func (p EducationSynchronizationOAuth2ClientCredentialsConnectionSettingsOperationPredicate) Matches(input EducationSynchronizationOAuth2ClientCredentialsConnectionSettings) bool {

	if p.ClientId != nil && (input.ClientId == nil || *p.ClientId != *input.ClientId) {
		return false
	}

	if p.ClientSecret != nil && (input.ClientSecret == nil || *p.ClientSecret != *input.ClientSecret) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Scope != nil && (input.Scope == nil || *p.Scope != *input.Scope) {
		return false
	}

	if p.TokenUrl != nil && (input.TokenUrl == nil || *p.TokenUrl != *input.TokenUrl) {
		return false
	}

	return true
}
