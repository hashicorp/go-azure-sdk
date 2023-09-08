package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsApplicationOperationPredicate struct {
	ODataType  *string
	PackageSid *string
}

func (p WindowsApplicationOperationPredicate) Matches(input WindowsApplication) bool {

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.PackageSid != nil && (input.PackageSid == nil || *p.PackageSid != *input.PackageSid) {
		return false
	}

	return true
}
