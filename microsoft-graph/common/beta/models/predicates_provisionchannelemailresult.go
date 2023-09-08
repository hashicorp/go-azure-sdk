package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProvisionChannelEmailResultOperationPredicate struct {
	Email     *string
	ODataType *string
}

func (p ProvisionChannelEmailResultOperationPredicate) Matches(input ProvisionChannelEmailResult) bool {

	if p.Email != nil && (input.Email == nil || *p.Email != *input.Email) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
