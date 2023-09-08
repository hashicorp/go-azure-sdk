package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InformationalUrlOperationPredicate struct {
	LogoUrl             *string
	MarketingUrl        *string
	ODataType           *string
	PrivacyStatementUrl *string
	SupportUrl          *string
	TermsOfServiceUrl   *string
}

func (p InformationalUrlOperationPredicate) Matches(input InformationalUrl) bool {

	if p.LogoUrl != nil && (input.LogoUrl == nil || *p.LogoUrl != *input.LogoUrl) {
		return false
	}

	if p.MarketingUrl != nil && (input.MarketingUrl == nil || *p.MarketingUrl != *input.MarketingUrl) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.PrivacyStatementUrl != nil && (input.PrivacyStatementUrl == nil || *p.PrivacyStatementUrl != *input.PrivacyStatementUrl) {
		return false
	}

	if p.SupportUrl != nil && (input.SupportUrl == nil || *p.SupportUrl != *input.SupportUrl) {
		return false
	}

	if p.TermsOfServiceUrl != nil && (input.TermsOfServiceUrl == nil || *p.TermsOfServiceUrl != *input.TermsOfServiceUrl) {
		return false
	}

	return true
}
