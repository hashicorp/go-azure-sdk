package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationSettingDependedOnByOperationPredicate struct {
	DependedOnBy *string
	ODataType    *string
	Required     *bool
}

func (p DeviceManagementConfigurationSettingDependedOnByOperationPredicate) Matches(input DeviceManagementConfigurationSettingDependedOnBy) bool {

	if p.DependedOnBy != nil && (input.DependedOnBy == nil || *p.DependedOnBy != *input.DependedOnBy) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Required != nil && (input.Required == nil || *p.Required != *input.Required) {
		return false
	}

	return true
}
