package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationIntegerSettingValueOperationPredicate struct {
	ODataType *string
	Value     *int64
}

func (p DeviceManagementConfigurationIntegerSettingValueOperationPredicate) Matches(input DeviceManagementConfigurationIntegerSettingValue) bool {

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Value != nil && (input.Value == nil || *p.Value != *input.Value) {
		return false
	}

	return true
}
