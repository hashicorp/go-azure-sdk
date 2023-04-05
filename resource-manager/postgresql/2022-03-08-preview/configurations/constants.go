package configurations

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConfigurationDataType string

const (
	ConfigurationDataTypeBoolean     ConfigurationDataType = "Boolean"
	ConfigurationDataTypeEnumeration ConfigurationDataType = "Enumeration"
	ConfigurationDataTypeInteger     ConfigurationDataType = "Integer"
	ConfigurationDataTypeNumeric     ConfigurationDataType = "Numeric"
)

func PossibleValuesForConfigurationDataType() []string {
	return []string{
		string(ConfigurationDataTypeBoolean),
		string(ConfigurationDataTypeEnumeration),
		string(ConfigurationDataTypeInteger),
		string(ConfigurationDataTypeNumeric),
	}
}
