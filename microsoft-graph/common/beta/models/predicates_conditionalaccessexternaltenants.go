package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConditionalAccessExternalTenantsOperationPredicate struct {
	ODataType *string
}

func (p ConditionalAccessExternalTenantsOperationPredicate) Matches(input ConditionalAccessExternalTenants) bool {

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
