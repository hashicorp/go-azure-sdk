package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LocaleInfoOperationPredicate struct {
	DisplayName *string
	Locale      *string
	ODataType   *string
}

func (p LocaleInfoOperationPredicate) Matches(input LocaleInfo) bool {

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.Locale != nil && (input.Locale == nil || *p.Locale != *input.Locale) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
