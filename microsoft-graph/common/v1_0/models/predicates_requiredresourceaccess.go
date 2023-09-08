package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RequiredResourceAccessOperationPredicate struct {
	ODataType     *string
	ResourceAppId *string
}

func (p RequiredResourceAccessOperationPredicate) Matches(input RequiredResourceAccess) bool {

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.ResourceAppId != nil && (input.ResourceAppId == nil || *p.ResourceAppId != *input.ResourceAppId) {
		return false
	}

	return true
}
