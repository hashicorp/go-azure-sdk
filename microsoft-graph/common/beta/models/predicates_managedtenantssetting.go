package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedTenantsSettingOperationPredicate struct {
	DisplayName      *string
	JsonValue        *string
	ODataType        *string
	OverwriteAllowed *bool
	SettingId        *string
}

func (p ManagedTenantsSettingOperationPredicate) Matches(input ManagedTenantsSetting) bool {

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.JsonValue != nil && (input.JsonValue == nil || *p.JsonValue != *input.JsonValue) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.OverwriteAllowed != nil && (input.OverwriteAllowed == nil || *p.OverwriteAllowed != *input.OverwriteAllowed) {
		return false
	}

	if p.SettingId != nil && (input.SettingId == nil || *p.SettingId != *input.SettingId) {
		return false
	}

	return true
}
