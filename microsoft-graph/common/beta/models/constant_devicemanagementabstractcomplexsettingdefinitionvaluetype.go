package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementAbstractComplexSettingDefinitionValueType string

const (
	DeviceManagementAbstractComplexSettingDefinitionValueTypeabstractComplex DeviceManagementAbstractComplexSettingDefinitionValueType = "AbstractComplex"
	DeviceManagementAbstractComplexSettingDefinitionValueTypeboolean         DeviceManagementAbstractComplexSettingDefinitionValueType = "Boolean"
	DeviceManagementAbstractComplexSettingDefinitionValueTypecollection      DeviceManagementAbstractComplexSettingDefinitionValueType = "Collection"
	DeviceManagementAbstractComplexSettingDefinitionValueTypecomplex         DeviceManagementAbstractComplexSettingDefinitionValueType = "Complex"
	DeviceManagementAbstractComplexSettingDefinitionValueTypeinteger         DeviceManagementAbstractComplexSettingDefinitionValueType = "Integer"
	DeviceManagementAbstractComplexSettingDefinitionValueTypestring          DeviceManagementAbstractComplexSettingDefinitionValueType = "String"
)

func PossibleValuesForDeviceManagementAbstractComplexSettingDefinitionValueType() []string {
	return []string{
		string(DeviceManagementAbstractComplexSettingDefinitionValueTypeabstractComplex),
		string(DeviceManagementAbstractComplexSettingDefinitionValueTypeboolean),
		string(DeviceManagementAbstractComplexSettingDefinitionValueTypecollection),
		string(DeviceManagementAbstractComplexSettingDefinitionValueTypecomplex),
		string(DeviceManagementAbstractComplexSettingDefinitionValueTypeinteger),
		string(DeviceManagementAbstractComplexSettingDefinitionValueTypestring),
	}
}

func parseDeviceManagementAbstractComplexSettingDefinitionValueType(input string) (*DeviceManagementAbstractComplexSettingDefinitionValueType, error) {
	vals := map[string]DeviceManagementAbstractComplexSettingDefinitionValueType{
		"abstractcomplex": DeviceManagementAbstractComplexSettingDefinitionValueTypeabstractComplex,
		"boolean":         DeviceManagementAbstractComplexSettingDefinitionValueTypeboolean,
		"collection":      DeviceManagementAbstractComplexSettingDefinitionValueTypecollection,
		"complex":         DeviceManagementAbstractComplexSettingDefinitionValueTypecomplex,
		"integer":         DeviceManagementAbstractComplexSettingDefinitionValueTypeinteger,
		"string":          DeviceManagementAbstractComplexSettingDefinitionValueTypestring,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementAbstractComplexSettingDefinitionValueType(input)
	return &out, nil
}
