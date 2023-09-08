package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementSettingBooleanConstraintOperationPredicate struct {
	ODataType *string
	Value     *bool
}

func (p DeviceManagementSettingBooleanConstraintOperationPredicate) Matches(input DeviceManagementSettingBooleanConstraint) bool {

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Value != nil && (input.Value == nil || *p.Value != *input.Value) {
		return false
	}

	return true
}
