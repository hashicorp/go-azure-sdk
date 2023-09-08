package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VppTokenOperationPredicate struct {
	AppleId                 *string
	AutomaticallyUpdateApps *bool
	CountryOrRegion         *string
	ExpirationDateTime      *string
	Id                      *string
	LastModifiedDateTime    *string
	LastSyncDateTime        *string
	ODataType               *string
	OrganizationName        *string
	Token                   *string
}

func (p VppTokenOperationPredicate) Matches(input VppToken) bool {

	if p.AppleId != nil && (input.AppleId == nil || *p.AppleId != *input.AppleId) {
		return false
	}

	if p.AutomaticallyUpdateApps != nil && (input.AutomaticallyUpdateApps == nil || *p.AutomaticallyUpdateApps != *input.AutomaticallyUpdateApps) {
		return false
	}

	if p.CountryOrRegion != nil && (input.CountryOrRegion == nil || *p.CountryOrRegion != *input.CountryOrRegion) {
		return false
	}

	if p.ExpirationDateTime != nil && (input.ExpirationDateTime == nil || *p.ExpirationDateTime != *input.ExpirationDateTime) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.LastModifiedDateTime != nil && (input.LastModifiedDateTime == nil || *p.LastModifiedDateTime != *input.LastModifiedDateTime) {
		return false
	}

	if p.LastSyncDateTime != nil && (input.LastSyncDateTime == nil || *p.LastSyncDateTime != *input.LastSyncDateTime) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.OrganizationName != nil && (input.OrganizationName == nil || *p.OrganizationName != *input.OrganizationName) {
		return false
	}

	if p.Token != nil && (input.Token == nil || *p.Token != *input.Token) {
		return false
	}

	return true
}
