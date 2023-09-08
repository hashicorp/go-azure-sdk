package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidForWorkVpnConfigurationOperationPredicate struct {
	ConnectionName       *string
	CreatedDateTime      *string
	Description          *string
	DisplayName          *string
	Fingerprint          *string
	Id                   *string
	LastModifiedDateTime *string
	ODataType            *string
	Realm                *string
	Role                 *string
	SupportsScopeTags    *bool
	Version              *int64
}

func (p AndroidForWorkVpnConfigurationOperationPredicate) Matches(input AndroidForWorkVpnConfiguration) bool {

	if p.ConnectionName != nil && (input.ConnectionName == nil || *p.ConnectionName != *input.ConnectionName) {
		return false
	}

	if p.CreatedDateTime != nil && (input.CreatedDateTime == nil || *p.CreatedDateTime != *input.CreatedDateTime) {
		return false
	}

	if p.Description != nil && (input.Description == nil || *p.Description != *input.Description) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.Fingerprint != nil && (input.Fingerprint == nil || *p.Fingerprint != *input.Fingerprint) {
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

	if p.Realm != nil && (input.Realm == nil || *p.Realm != *input.Realm) {
		return false
	}

	if p.Role != nil && (input.Role == nil || *p.Role != *input.Role) {
		return false
	}

	if p.SupportsScopeTags != nil && (input.SupportsScopeTags == nil || *p.SupportsScopeTags != *input.SupportsScopeTags) {
		return false
	}

	if p.Version != nil && (input.Version == nil || *p.Version != *input.Version) {
		return false
	}

	return true
}
