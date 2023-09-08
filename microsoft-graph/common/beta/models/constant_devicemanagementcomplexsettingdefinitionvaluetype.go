package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementComplexSettingDefinitionValueType string

const (
	DeviceManagementComplexSettingDefinitionValueTypeabstractComplex DeviceManagementComplexSettingDefinitionValueType = "AbstractComplex"
	DeviceManagementComplexSettingDefinitionValueTypeboolean         DeviceManagementComplexSettingDefinitionValueType = "Boolean"
	DeviceManagementComplexSettingDefinitionValueTypecollection      DeviceManagementComplexSettingDefinitionValueType = "Collection"
	DeviceManagementComplexSettingDefinitionValueTypecomplex         DeviceManagementComplexSettingDefinitionValueType = "Complex"
	DeviceManagementComplexSettingDefinitionValueTypeinteger         DeviceManagementComplexSettingDefinitionValueType = "Integer"
	DeviceManagementComplexSettingDefinitionValueTypestring          DeviceManagementComplexSettingDefinitionValueType = "String"
)

func PossibleValuesForDeviceManagementComplexSettingDefinitionValueType() []string {
	return []string{
		string(DeviceManagementComplexSettingDefinitionValueTypeabstractComplex),
		string(DeviceManagementComplexSettingDefinitionValueTypeboolean),
		string(DeviceManagementComplexSettingDefinitionValueTypecollection),
		string(DeviceManagementComplexSettingDefinitionValueTypecomplex),
		string(DeviceManagementComplexSettingDefinitionValueTypeinteger),
		string(DeviceManagementComplexSettingDefinitionValueTypestring),
	}
}

func parseDeviceManagementComplexSettingDefinitionValueType(input string) (*DeviceManagementComplexSettingDefinitionValueType, error) {
	vals := map[string]DeviceManagementComplexSettingDefinitionValueType{
		"abstractcomplex": DeviceManagementComplexSettingDefinitionValueTypeabstractComplex,
		"boolean":         DeviceManagementComplexSettingDefinitionValueTypeboolean,
		"collection":      DeviceManagementComplexSettingDefinitionValueTypecollection,
		"complex":         DeviceManagementComplexSettingDefinitionValueTypecomplex,
		"integer":         DeviceManagementComplexSettingDefinitionValueTypeinteger,
		"string":          DeviceManagementComplexSettingDefinitionValueTypestring,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementComplexSettingDefinitionValueType(input)
	return &out, nil
}
