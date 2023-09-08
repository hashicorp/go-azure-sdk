package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationOptionDefinitionOperationPredicate struct {
	Description *string
	DisplayName *string
	HelpText    *string
	ItemId      *string
	Name        *string
	ODataType   *string
}

func (p DeviceManagementConfigurationOptionDefinitionOperationPredicate) Matches(input DeviceManagementConfigurationOptionDefinition) bool {

	if p.Description != nil && (input.Description == nil || *p.Description != *input.Description) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.HelpText != nil && (input.HelpText == nil || *p.HelpText != *input.HelpText) {
		return false
	}

	if p.ItemId != nil && (input.ItemId == nil || *p.ItemId != *input.ItemId) {
		return false
	}

	if p.Name != nil && (input.Name == nil || *p.Name != *input.Name) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
