package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsKioskActiveDirectoryGroupOperationPredicate struct {
	GroupName *string
	ODataType *string
}

func (p WindowsKioskActiveDirectoryGroupOperationPredicate) Matches(input WindowsKioskActiveDirectoryGroup) bool {

	if p.GroupName != nil && (input.GroupName == nil || *p.GroupName != *input.GroupName) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
