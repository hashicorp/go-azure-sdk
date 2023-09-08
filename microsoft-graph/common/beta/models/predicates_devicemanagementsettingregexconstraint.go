package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementSettingRegexConstraintOperationPredicate struct {
	ODataType *string
	Regex     *string
}

func (p DeviceManagementSettingRegexConstraintOperationPredicate) Matches(input DeviceManagementSettingRegexConstraint) bool {

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Regex != nil && (input.Regex == nil || *p.Regex != *input.Regex) {
		return false
	}

	return true
}
