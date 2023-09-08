package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SearchAnswerVariantOperationPredicate struct {
	Description *string
	DisplayName *string
	LanguageTag *string
	ODataType   *string
	WebUrl      *string
}

func (p SearchAnswerVariantOperationPredicate) Matches(input SearchAnswerVariant) bool {

	if p.Description != nil && (input.Description == nil || *p.Description != *input.Description) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.LanguageTag != nil && (input.LanguageTag == nil || *p.LanguageTag != *input.LanguageTag) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.WebUrl != nil && (input.WebUrl == nil || *p.WebUrl != *input.WebUrl) {
		return false
	}

	return true
}
