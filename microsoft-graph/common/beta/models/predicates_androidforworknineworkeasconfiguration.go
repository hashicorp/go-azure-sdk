package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidForWorkNineWorkEasConfigurationOperationPredicate struct {
	CreatedDateTime      *string
	Description          *string
	DisplayName          *string
	HostName             *string
	Id                   *string
	LastModifiedDateTime *string
	ODataType            *string
	RequireSsl           *bool
	SupportsScopeTags    *bool
	SyncCalendar         *bool
	SyncContacts         *bool
	SyncTasks            *bool
	Version              *int64
}

func (p AndroidForWorkNineWorkEasConfigurationOperationPredicate) Matches(input AndroidForWorkNineWorkEasConfiguration) bool {

	if p.CreatedDateTime != nil && (input.CreatedDateTime == nil || *p.CreatedDateTime != *input.CreatedDateTime) {
		return false
	}

	if p.Description != nil && (input.Description == nil || *p.Description != *input.Description) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.HostName != nil && (input.HostName == nil || *p.HostName != *input.HostName) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.LastModifiedDateTime != nil && (input.LastModifiedDateTime == nil || *p.LastModifiedDateTime != *input.LastModifiedDateTime) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.RequireSsl != nil && (input.RequireSsl == nil || *p.RequireSsl != *input.RequireSsl) {
		return false
	}

	if p.SupportsScopeTags != nil && (input.SupportsScopeTags == nil || *p.SupportsScopeTags != *input.SupportsScopeTags) {
		return false
	}

	if p.SyncCalendar != nil && (input.SyncCalendar == nil || *p.SyncCalendar != *input.SyncCalendar) {
		return false
	}

	if p.SyncContacts != nil && (input.SyncContacts == nil || *p.SyncContacts != *input.SyncContacts) {
		return false
	}

	if p.SyncTasks != nil && (input.SyncTasks == nil || *p.SyncTasks != *input.SyncTasks) {
		return false
	}

	if p.Version != nil && (input.Version == nil || *p.Version != *input.Version) {
		return false
	}

	return true
}
