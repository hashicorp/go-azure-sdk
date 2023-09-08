package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProtectGroupOperationPredicate struct {
	AllowEmailFromGuestUsers *bool
	AllowGuestUsers          *bool
	Name                     *string
	ODataType                *string
}

func (p ProtectGroupOperationPredicate) Matches(input ProtectGroup) bool {

	if p.AllowEmailFromGuestUsers != nil && (input.AllowEmailFromGuestUsers == nil || *p.AllowEmailFromGuestUsers != *input.AllowEmailFromGuestUsers) {
		return false
	}

	if p.AllowGuestUsers != nil && (input.AllowGuestUsers == nil || *p.AllowGuestUsers != *input.AllowGuestUsers) {
		return false
	}

	if p.Name != nil && (input.Name == nil || *p.Name != *input.Name) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
