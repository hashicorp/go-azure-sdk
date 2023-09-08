package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessPackageLocalizedTextOperationPredicate struct {
	LanguageCode *string
	ODataType    *string
	Text         *string
}

func (p AccessPackageLocalizedTextOperationPredicate) Matches(input AccessPackageLocalizedText) bool {

	if p.LanguageCode != nil && (input.LanguageCode == nil || *p.LanguageCode != *input.LanguageCode) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Text != nil && (input.Text == nil || *p.Text != *input.Text) {
		return false
	}

	return true
}
