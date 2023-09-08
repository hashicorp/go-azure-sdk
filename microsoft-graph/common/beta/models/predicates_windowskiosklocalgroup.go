package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsKioskLocalGroupOperationPredicate struct {
	GroupName *string
	ODataType *string
}

func (p WindowsKioskLocalGroupOperationPredicate) Matches(input WindowsKioskLocalGroup) bool {

	if p.GroupName != nil && (input.GroupName == nil || *p.GroupName != *input.GroupName) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
