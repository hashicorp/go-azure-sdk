package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceComplianceScriptRuleDataType string

const (
	DeviceComplianceScriptRuleDataTypebase64        DeviceComplianceScriptRuleDataType = "Base64"
	DeviceComplianceScriptRuleDataTypeboolean       DeviceComplianceScriptRuleDataType = "Boolean"
	DeviceComplianceScriptRuleDataTypebooleanArray  DeviceComplianceScriptRuleDataType = "BooleanArray"
	DeviceComplianceScriptRuleDataTypedateTime      DeviceComplianceScriptRuleDataType = "DateTime"
	DeviceComplianceScriptRuleDataTypedateTimeArray DeviceComplianceScriptRuleDataType = "DateTimeArray"
	DeviceComplianceScriptRuleDataTypedouble        DeviceComplianceScriptRuleDataType = "Double"
	DeviceComplianceScriptRuleDataTypedoubleArray   DeviceComplianceScriptRuleDataType = "DoubleArray"
	DeviceComplianceScriptRuleDataTypeint64         DeviceComplianceScriptRuleDataType = "Int64"
	DeviceComplianceScriptRuleDataTypeint64Array    DeviceComplianceScriptRuleDataType = "Int64Array"
	DeviceComplianceScriptRuleDataTypenone          DeviceComplianceScriptRuleDataType = "None"
	DeviceComplianceScriptRuleDataTypestring        DeviceComplianceScriptRuleDataType = "String"
	DeviceComplianceScriptRuleDataTypestringArray   DeviceComplianceScriptRuleDataType = "StringArray"
	DeviceComplianceScriptRuleDataTypeversion       DeviceComplianceScriptRuleDataType = "Version"
	DeviceComplianceScriptRuleDataTypeversionArray  DeviceComplianceScriptRuleDataType = "VersionArray"
	DeviceComplianceScriptRuleDataTypexml           DeviceComplianceScriptRuleDataType = "Xml"
)

func PossibleValuesForDeviceComplianceScriptRuleDataType() []string {
	return []string{
		string(DeviceComplianceScriptRuleDataTypebase64),
		string(DeviceComplianceScriptRuleDataTypeboolean),
		string(DeviceComplianceScriptRuleDataTypebooleanArray),
		string(DeviceComplianceScriptRuleDataTypedateTime),
		string(DeviceComplianceScriptRuleDataTypedateTimeArray),
		string(DeviceComplianceScriptRuleDataTypedouble),
		string(DeviceComplianceScriptRuleDataTypedoubleArray),
		string(DeviceComplianceScriptRuleDataTypeint64),
		string(DeviceComplianceScriptRuleDataTypeint64Array),
		string(DeviceComplianceScriptRuleDataTypenone),
		string(DeviceComplianceScriptRuleDataTypestring),
		string(DeviceComplianceScriptRuleDataTypestringArray),
		string(DeviceComplianceScriptRuleDataTypeversion),
		string(DeviceComplianceScriptRuleDataTypeversionArray),
		string(DeviceComplianceScriptRuleDataTypexml),
	}
}

func parseDeviceComplianceScriptRuleDataType(input string) (*DeviceComplianceScriptRuleDataType, error) {
	vals := map[string]DeviceComplianceScriptRuleDataType{
		"base64":        DeviceComplianceScriptRuleDataTypebase64,
		"boolean":       DeviceComplianceScriptRuleDataTypeboolean,
		"booleanarray":  DeviceComplianceScriptRuleDataTypebooleanArray,
		"datetime":      DeviceComplianceScriptRuleDataTypedateTime,
		"datetimearray": DeviceComplianceScriptRuleDataTypedateTimeArray,
		"double":        DeviceComplianceScriptRuleDataTypedouble,
		"doublearray":   DeviceComplianceScriptRuleDataTypedoubleArray,
		"int64":         DeviceComplianceScriptRuleDataTypeint64,
		"int64array":    DeviceComplianceScriptRuleDataTypeint64Array,
		"none":          DeviceComplianceScriptRuleDataTypenone,
		"string":        DeviceComplianceScriptRuleDataTypestring,
		"stringarray":   DeviceComplianceScriptRuleDataTypestringArray,
		"version":       DeviceComplianceScriptRuleDataTypeversion,
		"versionarray":  DeviceComplianceScriptRuleDataTypeversionArray,
		"xml":           DeviceComplianceScriptRuleDataTypexml,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceComplianceScriptRuleDataType(input)
	return &out, nil
}
