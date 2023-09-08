package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementCollectionSettingDefinitionValueType string

const (
	DeviceManagementCollectionSettingDefinitionValueTypeabstractComplex DeviceManagementCollectionSettingDefinitionValueType = "AbstractComplex"
	DeviceManagementCollectionSettingDefinitionValueTypeboolean         DeviceManagementCollectionSettingDefinitionValueType = "Boolean"
	DeviceManagementCollectionSettingDefinitionValueTypecollection      DeviceManagementCollectionSettingDefinitionValueType = "Collection"
	DeviceManagementCollectionSettingDefinitionValueTypecomplex         DeviceManagementCollectionSettingDefinitionValueType = "Complex"
	DeviceManagementCollectionSettingDefinitionValueTypeinteger         DeviceManagementCollectionSettingDefinitionValueType = "Integer"
	DeviceManagementCollectionSettingDefinitionValueTypestring          DeviceManagementCollectionSettingDefinitionValueType = "String"
)

func PossibleValuesForDeviceManagementCollectionSettingDefinitionValueType() []string {
	return []string{
		string(DeviceManagementCollectionSettingDefinitionValueTypeabstractComplex),
		string(DeviceManagementCollectionSettingDefinitionValueTypeboolean),
		string(DeviceManagementCollectionSettingDefinitionValueTypecollection),
		string(DeviceManagementCollectionSettingDefinitionValueTypecomplex),
		string(DeviceManagementCollectionSettingDefinitionValueTypeinteger),
		string(DeviceManagementCollectionSettingDefinitionValueTypestring),
	}
}

func parseDeviceManagementCollectionSettingDefinitionValueType(input string) (*DeviceManagementCollectionSettingDefinitionValueType, error) {
	vals := map[string]DeviceManagementCollectionSettingDefinitionValueType{
		"abstractcomplex": DeviceManagementCollectionSettingDefinitionValueTypeabstractComplex,
		"boolean":         DeviceManagementCollectionSettingDefinitionValueTypeboolean,
		"collection":      DeviceManagementCollectionSettingDefinitionValueTypecollection,
		"complex":         DeviceManagementCollectionSettingDefinitionValueTypecomplex,
		"integer":         DeviceManagementCollectionSettingDefinitionValueTypeinteger,
		"string":          DeviceManagementCollectionSettingDefinitionValueTypestring,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementCollectionSettingDefinitionValueType(input)
	return &out, nil
}
