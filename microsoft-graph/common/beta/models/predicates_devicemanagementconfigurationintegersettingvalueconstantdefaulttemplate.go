package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationIntegerSettingValueConstantDefaultTemplateOperationPredicate struct {
	ConstantValue *int64
	ODataType     *string
}

func (p DeviceManagementConfigurationIntegerSettingValueConstantDefaultTemplateOperationPredicate) Matches(input DeviceManagementConfigurationIntegerSettingValueConstantDefaultTemplate) bool {

	if p.ConstantValue != nil && (input.ConstantValue == nil || *p.ConstantValue != *input.ConstantValue) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
