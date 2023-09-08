package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TranslationLanguageOverrideOperationPredicate struct {
	LanguageTag *string
	ODataType   *string
}

func (p TranslationLanguageOverrideOperationPredicate) Matches(input TranslationLanguageOverride) bool {

	if p.LanguageTag != nil && (input.LanguageTag == nil || *p.LanguageTag != *input.LanguageTag) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
