package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementSettingDefinitionValueType string

const (
	DeviceManagementSettingDefinitionValueTypeabstractComplex DeviceManagementSettingDefinitionValueType = "AbstractComplex"
	DeviceManagementSettingDefinitionValueTypeboolean         DeviceManagementSettingDefinitionValueType = "Boolean"
	DeviceManagementSettingDefinitionValueTypecollection      DeviceManagementSettingDefinitionValueType = "Collection"
	DeviceManagementSettingDefinitionValueTypecomplex         DeviceManagementSettingDefinitionValueType = "Complex"
	DeviceManagementSettingDefinitionValueTypeinteger         DeviceManagementSettingDefinitionValueType = "Integer"
	DeviceManagementSettingDefinitionValueTypestring          DeviceManagementSettingDefinitionValueType = "String"
)

func PossibleValuesForDeviceManagementSettingDefinitionValueType() []string {
	return []string{
		string(DeviceManagementSettingDefinitionValueTypeabstractComplex),
		string(DeviceManagementSettingDefinitionValueTypeboolean),
		string(DeviceManagementSettingDefinitionValueTypecollection),
		string(DeviceManagementSettingDefinitionValueTypecomplex),
		string(DeviceManagementSettingDefinitionValueTypeinteger),
		string(DeviceManagementSettingDefinitionValueTypestring),
	}
}

func parseDeviceManagementSettingDefinitionValueType(input string) (*DeviceManagementSettingDefinitionValueType, error) {
	vals := map[string]DeviceManagementSettingDefinitionValueType{
		"abstractcomplex": DeviceManagementSettingDefinitionValueTypeabstractComplex,
		"boolean":         DeviceManagementSettingDefinitionValueTypeboolean,
		"collection":      DeviceManagementSettingDefinitionValueTypecollection,
		"complex":         DeviceManagementSettingDefinitionValueTypecomplex,
		"integer":         DeviceManagementSettingDefinitionValueTypeinteger,
		"string":          DeviceManagementSettingDefinitionValueTypestring,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementSettingDefinitionValueType(input)
	return &out, nil
}
