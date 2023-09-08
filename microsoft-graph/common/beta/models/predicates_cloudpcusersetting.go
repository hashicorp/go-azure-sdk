package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCUserSettingOperationPredicate struct {
	CreatedDateTime      *string
	DisplayName          *string
	Id                   *string
	LastModifiedDateTime *string
	LocalAdminEnabled    *bool
	ODataType            *string
	ResetEnabled         *bool
	SelfServiceEnabled   *bool
}

func (p CloudPCUserSettingOperationPredicate) Matches(input CloudPCUserSetting) bool {

	if p.CreatedDateTime != nil && (input.CreatedDateTime == nil || *p.CreatedDateTime != *input.CreatedDateTime) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.LastModifiedDateTime != nil && (input.LastModifiedDateTime == nil || *p.LastModifiedDateTime != *input.LastModifiedDateTime) {
		return false
	}

	if p.LocalAdminEnabled != nil && (input.LocalAdminEnabled == nil || *p.LocalAdminEnabled != *input.LocalAdminEnabled) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.ResetEnabled != nil && (input.ResetEnabled == nil || *p.ResetEnabled != *input.ResetEnabled) {
		return false
	}

	if p.SelfServiceEnabled != nil && (input.SelfServiceEnabled == nil || *p.SelfServiceEnabled != *input.SelfServiceEnabled) {
		return false
	}

	return true
}
