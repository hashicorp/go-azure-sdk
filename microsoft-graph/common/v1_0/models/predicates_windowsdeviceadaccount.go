package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsDeviceADAccountOperationPredicate struct {
	DomainName *string
	ODataType  *string
	Password   *string
	UserName   *string
}

func (p WindowsDeviceADAccountOperationPredicate) Matches(input WindowsDeviceADAccount) bool {

	if p.DomainName != nil && (input.DomainName == nil || *p.DomainName != *input.DomainName) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Password != nil && (input.Password == nil || *p.Password != *input.Password) {
		return false
	}

	if p.UserName != nil && (input.UserName == nil || *p.UserName != *input.UserName) {
		return false
	}

	return true
}
