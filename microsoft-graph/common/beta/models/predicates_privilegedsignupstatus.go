package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivilegedSignupStatusOperationPredicate struct {
	Id           *string
	IsRegistered *bool
	ODataType    *string
}

func (p PrivilegedSignupStatusOperationPredicate) Matches(input PrivilegedSignupStatus) bool {

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.IsRegistered != nil && (input.IsRegistered == nil || *p.IsRegistered != *input.IsRegistered) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
