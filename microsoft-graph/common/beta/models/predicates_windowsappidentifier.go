package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsAppIdentifierOperationPredicate struct {
	ODataType    *string
	WindowsAppId *string
}

func (p WindowsAppIdentifierOperationPredicate) Matches(input WindowsAppIdentifier) bool {

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.WindowsAppId != nil && (input.WindowsAppId == nil || *p.WindowsAppId != *input.WindowsAppId) {
		return false
	}

	return true
}
