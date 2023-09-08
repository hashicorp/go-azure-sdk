package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SingleServicePrincipalOperationPredicate struct {
	Description        *string
	ODataType          *string
	ServicePrincipalId *string
}

func (p SingleServicePrincipalOperationPredicate) Matches(input SingleServicePrincipal) bool {

	if p.Description != nil && (input.Description == nil || *p.Description != *input.Description) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.ServicePrincipalId != nil && (input.ServicePrincipalId == nil || *p.ServicePrincipalId != *input.ServicePrincipalId) {
		return false
	}

	return true
}
