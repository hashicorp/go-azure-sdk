package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataSharingConsentOperationPredicate struct {
	GrantDateTime      *string
	Granted            *bool
	GrantedByUpn       *string
	GrantedByUserId    *string
	Id                 *string
	ODataType          *string
	ServiceDisplayName *string
	TermsUrl           *string
}

func (p DataSharingConsentOperationPredicate) Matches(input DataSharingConsent) bool {

	if p.GrantDateTime != nil && (input.GrantDateTime == nil || *p.GrantDateTime != *input.GrantDateTime) {
		return false
	}

	if p.Granted != nil && (input.Granted == nil || *p.Granted != *input.Granted) {
		return false
	}

	if p.GrantedByUpn != nil && (input.GrantedByUpn == nil || *p.GrantedByUpn != *input.GrantedByUpn) {
		return false
	}

	if p.GrantedByUserId != nil && (input.GrantedByUserId == nil || *p.GrantedByUserId != *input.GrantedByUserId) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.ServiceDisplayName != nil && (input.ServiceDisplayName == nil || *p.ServiceDisplayName != *input.ServiceDisplayName) {
		return false
	}

	if p.TermsUrl != nil && (input.TermsUrl == nil || *p.TermsUrl != *input.TermsUrl) {
		return false
	}

	return true
}
