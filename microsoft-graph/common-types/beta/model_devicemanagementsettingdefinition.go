package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementSettingDefinition struct {
	Constraints      *[]DeviceManagementConstraint               `json:"constraints,omitempty"`
	Dependencies     *[]DeviceManagementSettingDependency        `json:"dependencies,omitempty"`
	Description      *string                                     `json:"description,omitempty"`
	DisplayName      *string                                     `json:"displayName,omitempty"`
	DocumentationUrl *string                                     `json:"documentationUrl,omitempty"`
	HeaderSubtitle   *string                                     `json:"headerSubtitle,omitempty"`
	HeaderTitle      *string                                     `json:"headerTitle,omitempty"`
	Id               *string                                     `json:"id,omitempty"`
	IsTopLevel       *bool                                       `json:"isTopLevel,omitempty"`
	Keywords         *[]string                                   `json:"keywords,omitempty"`
	ODataType        *string                                     `json:"@odata.type,omitempty"`
	PlaceholderText  *string                                     `json:"placeholderText,omitempty"`
	ValueType        *DeviceManagementSettingDefinitionValueType `json:"valueType,omitempty"`
}
