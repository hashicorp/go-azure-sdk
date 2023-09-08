package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationOptionDefinitionTemplateOperationPredicate struct {
	ItemId    *string
	ODataType *string
}

func (p DeviceManagementConfigurationOptionDefinitionTemplateOperationPredicate) Matches(input DeviceManagementConfigurationOptionDefinitionTemplate) bool {

	if p.ItemId != nil && (input.ItemId == nil || *p.ItemId != *input.ItemId) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
