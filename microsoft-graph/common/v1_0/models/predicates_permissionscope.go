package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PermissionScopeOperationPredicate struct {
	AdminConsentDescription *string
	AdminConsentDisplayName *string
	Id                      *string
	IsEnabled               *bool
	ODataType               *string
	Origin                  *string
	Type                    *string
	UserConsentDescription  *string
	UserConsentDisplayName  *string
	Value                   *string
}

func (p PermissionScopeOperationPredicate) Matches(input PermissionScope) bool {

	if p.AdminConsentDescription != nil && (input.AdminConsentDescription == nil || *p.AdminConsentDescription != *input.AdminConsentDescription) {
		return false
	}

	if p.AdminConsentDisplayName != nil && (input.AdminConsentDisplayName == nil || *p.AdminConsentDisplayName != *input.AdminConsentDisplayName) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.IsEnabled != nil && (input.IsEnabled == nil || *p.IsEnabled != *input.IsEnabled) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Origin != nil && (input.Origin == nil || *p.Origin != *input.Origin) {
		return false
	}

	if p.Type != nil && (input.Type == nil || *p.Type != *input.Type) {
		return false
	}

	if p.UserConsentDescription != nil && (input.UserConsentDescription == nil || *p.UserConsentDescription != *input.UserConsentDescription) {
		return false
	}

	if p.UserConsentDisplayName != nil && (input.UserConsentDisplayName == nil || *p.UserConsentDisplayName != *input.UserConsentDisplayName) {
		return false
	}

	if p.Value != nil && (input.Value == nil || *p.Value != *input.Value) {
		return false
	}

	return true
}
