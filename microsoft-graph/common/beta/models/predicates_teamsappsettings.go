package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamsAppSettingsOperationPredicate struct {
	AllowUserRequestsForAppAccess                     *bool
	Id                                                *string
	IsChatResourceSpecificConsentEnabled              *bool
	IsUserPersonalScopeResourceSpecificConsentEnabled *bool
	ODataType                                         *string
}

func (p TeamsAppSettingsOperationPredicate) Matches(input TeamsAppSettings) bool {

	if p.AllowUserRequestsForAppAccess != nil && (input.AllowUserRequestsForAppAccess == nil || *p.AllowUserRequestsForAppAccess != *input.AllowUserRequestsForAppAccess) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.IsChatResourceSpecificConsentEnabled != nil && (input.IsChatResourceSpecificConsentEnabled == nil || *p.IsChatResourceSpecificConsentEnabled != *input.IsChatResourceSpecificConsentEnabled) {
		return false
	}

	if p.IsUserPersonalScopeResourceSpecificConsentEnabled != nil && (input.IsUserPersonalScopeResourceSpecificConsentEnabled == nil || *p.IsUserPersonalScopeResourceSpecificConsentEnabled != *input.IsUserPersonalScopeResourceSpecificConsentEnabled) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
