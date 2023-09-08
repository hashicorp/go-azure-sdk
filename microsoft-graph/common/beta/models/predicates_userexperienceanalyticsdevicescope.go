package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsDeviceScopeOperationPredicate struct {
	CreatedDateTime      *string
	DeviceScopeName      *string
	Enabled              *bool
	Id                   *string
	IsBuiltIn            *bool
	LastModifiedDateTime *string
	ODataType            *string
	OwnerId              *string
	Value                *string
	ValueObjectId        *string
}

func (p UserExperienceAnalyticsDeviceScopeOperationPredicate) Matches(input UserExperienceAnalyticsDeviceScope) bool {

	if p.CreatedDateTime != nil && (input.CreatedDateTime == nil || *p.CreatedDateTime != *input.CreatedDateTime) {
		return false
	}

	if p.DeviceScopeName != nil && (input.DeviceScopeName == nil || *p.DeviceScopeName != *input.DeviceScopeName) {
		return false
	}

	if p.Enabled != nil && (input.Enabled == nil || *p.Enabled != *input.Enabled) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.IsBuiltIn != nil && (input.IsBuiltIn == nil || *p.IsBuiltIn != *input.IsBuiltIn) {
		return false
	}

	if p.LastModifiedDateTime != nil && (input.LastModifiedDateTime == nil || *p.LastModifiedDateTime != *input.LastModifiedDateTime) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.OwnerId != nil && (input.OwnerId == nil || *p.OwnerId != *input.OwnerId) {
		return false
	}

	if p.Value != nil && (input.Value == nil || *p.Value != *input.Value) {
		return false
	}

	if p.ValueObjectId != nil && (input.ValueObjectId == nil || *p.ValueObjectId != *input.ValueObjectId) {
		return false
	}

	return true
}
