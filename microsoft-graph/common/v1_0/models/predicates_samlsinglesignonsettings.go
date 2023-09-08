package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SamlSingleSignOnSettingsOperationPredicate struct {
	ODataType  *string
	RelayState *string
}

func (p SamlSingleSignOnSettingsOperationPredicate) Matches(input SamlSingleSignOnSettings) bool {

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.RelayState != nil && (input.RelayState == nil || *p.RelayState != *input.RelayState) {
		return false
	}

	return true
}
