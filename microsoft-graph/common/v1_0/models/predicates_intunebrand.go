package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IntuneBrandOperationPredicate struct {
	ContactITEmailAddress     *string
	ContactITName             *string
	ContactITNotes            *string
	ContactITPhoneNumber      *string
	DisplayName               *string
	ODataType                 *string
	OnlineSupportSiteName     *string
	OnlineSupportSiteUrl      *string
	PrivacyUrl                *string
	ShowDisplayNameNextToLogo *bool
	ShowLogo                  *bool
	ShowNameNextToLogo        *bool
}

func (p IntuneBrandOperationPredicate) Matches(input IntuneBrand) bool {

	if p.ContactITEmailAddress != nil && (input.ContactITEmailAddress == nil || *p.ContactITEmailAddress != *input.ContactITEmailAddress) {
		return false
	}

	if p.ContactITName != nil && (input.ContactITName == nil || *p.ContactITName != *input.ContactITName) {
		return false
	}

	if p.ContactITNotes != nil && (input.ContactITNotes == nil || *p.ContactITNotes != *input.ContactITNotes) {
		return false
	}

	if p.ContactITPhoneNumber != nil && (input.ContactITPhoneNumber == nil || *p.ContactITPhoneNumber != *input.ContactITPhoneNumber) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.OnlineSupportSiteName != nil && (input.OnlineSupportSiteName == nil || *p.OnlineSupportSiteName != *input.OnlineSupportSiteName) {
		return false
	}

	if p.OnlineSupportSiteUrl != nil && (input.OnlineSupportSiteUrl == nil || *p.OnlineSupportSiteUrl != *input.OnlineSupportSiteUrl) {
		return false
	}

	if p.PrivacyUrl != nil && (input.PrivacyUrl == nil || *p.PrivacyUrl != *input.PrivacyUrl) {
		return false
	}

	if p.ShowDisplayNameNextToLogo != nil && (input.ShowDisplayNameNextToLogo == nil || *p.ShowDisplayNameNextToLogo != *input.ShowDisplayNameNextToLogo) {
		return false
	}

	if p.ShowLogo != nil && (input.ShowLogo == nil || *p.ShowLogo != *input.ShowLogo) {
		return false
	}

	if p.ShowNameNextToLogo != nil && (input.ShowNameNextToLogo == nil || *p.ShowNameNextToLogo != *input.ShowNameNextToLogo) {
		return false
	}

	return true
}
