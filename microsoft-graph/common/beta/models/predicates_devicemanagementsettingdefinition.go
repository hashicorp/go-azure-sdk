package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementSettingDefinitionOperationPredicate struct {
	Description      *string
	DisplayName      *string
	DocumentationUrl *string
	HeaderSubtitle   *string
	HeaderTitle      *string
	Id               *string
	IsTopLevel       *bool
	ODataType        *string
	PlaceholderText  *string
}

func (p DeviceManagementSettingDefinitionOperationPredicate) Matches(input DeviceManagementSettingDefinition) bool {

	if p.Description != nil && (input.Description == nil || *p.Description != *input.Description) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.DocumentationUrl != nil && (input.DocumentationUrl == nil || *p.DocumentationUrl != *input.DocumentationUrl) {
		return false
	}

	if p.HeaderSubtitle != nil && (input.HeaderSubtitle == nil || *p.HeaderSubtitle != *input.HeaderSubtitle) {
		return false
	}

	if p.HeaderTitle != nil && (input.HeaderTitle == nil || *p.HeaderTitle != *input.HeaderTitle) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.IsTopLevel != nil && (input.IsTopLevel == nil || *p.IsTopLevel != *input.IsTopLevel) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.PlaceholderText != nil && (input.PlaceholderText == nil || *p.PlaceholderText != *input.PlaceholderText) {
		return false
	}

	return true
}
