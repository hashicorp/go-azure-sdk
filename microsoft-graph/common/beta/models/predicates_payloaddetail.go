package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PayloadDetailOperationPredicate struct {
	Content     *string
	ODataType   *string
	PhishingUrl *string
}

func (p PayloadDetailOperationPredicate) Matches(input PayloadDetail) bool {

	if p.Content != nil && (input.Content == nil || *p.Content != *input.Content) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.PhishingUrl != nil && (input.PhishingUrl == nil || *p.PhishingUrl != *input.PhishingUrl) {
		return false
	}

	return true
}
