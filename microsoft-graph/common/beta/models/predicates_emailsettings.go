package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EmailSettingsOperationPredicate struct {
	ODataType          *string
	SenderDomain       *string
	UseCompanyBranding *bool
}

func (p EmailSettingsOperationPredicate) Matches(input EmailSettings) bool {

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.SenderDomain != nil && (input.SenderDomain == nil || *p.SenderDomain != *input.SenderDomain) {
		return false
	}

	if p.UseCompanyBranding != nil && (input.UseCompanyBranding == nil || *p.UseCompanyBranding != *input.UseCompanyBranding) {
		return false
	}

	return true
}
