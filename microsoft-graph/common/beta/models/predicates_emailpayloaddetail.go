package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EmailPayloadDetailOperationPredicate struct {
	Content          *string
	FromEmail        *string
	FromName         *string
	IsExternalSender *bool
	ODataType        *string
	PhishingUrl      *string
	Subject          *string
}

func (p EmailPayloadDetailOperationPredicate) Matches(input EmailPayloadDetail) bool {

	if p.Content != nil && (input.Content == nil || *p.Content != *input.Content) {
		return false
	}

	if p.FromEmail != nil && (input.FromEmail == nil || *p.FromEmail != *input.FromEmail) {
		return false
	}

	if p.FromName != nil && (input.FromName == nil || *p.FromName != *input.FromName) {
		return false
	}

	if p.IsExternalSender != nil && (input.IsExternalSender == nil || *p.IsExternalSender != *input.IsExternalSender) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.PhishingUrl != nil && (input.PhishingUrl == nil || *p.PhishingUrl != *input.PhishingUrl) {
		return false
	}

	if p.Subject != nil && (input.Subject == nil || *p.Subject != *input.Subject) {
		return false
	}

	return true
}
