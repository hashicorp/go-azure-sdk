package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUpdatesSafeguardSettingsOperationPredicate struct {
	ODataType *string
}

func (p WindowsUpdatesSafeguardSettingsOperationPredicate) Matches(input WindowsUpdatesSafeguardSettings) bool {

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
