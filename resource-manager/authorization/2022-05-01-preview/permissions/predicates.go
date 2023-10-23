package permissions

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PermissionOperationPredicate struct {
	Condition        *string
	ConditionVersion *string
}

func (p PermissionOperationPredicate) Matches(input Permission) bool {

	if p.Condition != nil && (input.Condition == nil || *p.Condition != *input.Condition) {
		return false
	}

	if p.ConditionVersion != nil && (input.ConditionVersion == nil || *p.ConditionVersion != *input.ConditionVersion) {
		return false
	}

	return true
}
