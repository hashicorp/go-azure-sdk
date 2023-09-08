package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EmployeeOrgDataOperationPredicate struct {
	CostCenter *string
	Division   *string
	ODataType  *string
}

func (p EmployeeOrgDataOperationPredicate) Matches(input EmployeeOrgData) bool {

	if p.CostCenter != nil && (input.CostCenter == nil || *p.CostCenter != *input.CostCenter) {
		return false
	}

	if p.Division != nil && (input.Division == nil || *p.Division != *input.Division) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
