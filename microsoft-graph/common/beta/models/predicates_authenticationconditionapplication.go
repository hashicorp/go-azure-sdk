package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationConditionApplicationOperationPredicate struct {
	AppId     *string
	ODataType *string
}

func (p AuthenticationConditionApplicationOperationPredicate) Matches(input AuthenticationConditionApplication) bool {

	if p.AppId != nil && (input.AppId == nil || *p.AppId != *input.AppId) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
