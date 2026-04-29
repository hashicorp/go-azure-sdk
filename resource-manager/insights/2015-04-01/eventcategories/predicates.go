package eventcategories

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MicrosoftCommonLocalizableStringOperationPredicate struct {
	LocalizedValue *string
	Value          *string
}

func (p MicrosoftCommonLocalizableStringOperationPredicate) Matches(input MicrosoftCommonLocalizableString) bool {

	if p.LocalizedValue != nil && (input.LocalizedValue == nil || *p.LocalizedValue != *input.LocalizedValue) {
		return false
	}

	if p.Value != nil && *p.Value != input.Value {
		return false
	}

	return true
}
