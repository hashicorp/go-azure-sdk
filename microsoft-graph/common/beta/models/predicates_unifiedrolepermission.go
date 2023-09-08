package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UnifiedRolePermissionOperationPredicate struct {
	Condition *string
	ODataType *string
}

func (p UnifiedRolePermissionOperationPredicate) Matches(input UnifiedRolePermission) bool {

	if p.Condition != nil && (input.Condition == nil || *p.Condition != *input.Condition) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
