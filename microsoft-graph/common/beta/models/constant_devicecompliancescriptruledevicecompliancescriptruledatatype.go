package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceComplianceScriptRuleDeviceComplianceScriptRuleDataType string

const (
	DeviceComplianceScriptRuleDeviceComplianceScriptRuleDataTypebase64        DeviceComplianceScriptRuleDeviceComplianceScriptRuleDataType = "Base64"
	DeviceComplianceScriptRuleDeviceComplianceScriptRuleDataTypeboolean       DeviceComplianceScriptRuleDeviceComplianceScriptRuleDataType = "Boolean"
	DeviceComplianceScriptRuleDeviceComplianceScriptRuleDataTypebooleanArray  DeviceComplianceScriptRuleDeviceComplianceScriptRuleDataType = "BooleanArray"
	DeviceComplianceScriptRuleDeviceComplianceScriptRuleDataTypedateTime      DeviceComplianceScriptRuleDeviceComplianceScriptRuleDataType = "DateTime"
	DeviceComplianceScriptRuleDeviceComplianceScriptRuleDataTypedateTimeArray DeviceComplianceScriptRuleDeviceComplianceScriptRuleDataType = "DateTimeArray"
	DeviceComplianceScriptRuleDeviceComplianceScriptRuleDataTypedouble        DeviceComplianceScriptRuleDeviceComplianceScriptRuleDataType = "Double"
	DeviceComplianceScriptRuleDeviceComplianceScriptRuleDataTypedoubleArray   DeviceComplianceScriptRuleDeviceComplianceScriptRuleDataType = "DoubleArray"
	DeviceComplianceScriptRuleDeviceComplianceScriptRuleDataTypeint64         DeviceComplianceScriptRuleDeviceComplianceScriptRuleDataType = "Int64"
	DeviceComplianceScriptRuleDeviceComplianceScriptRuleDataTypeint64Array    DeviceComplianceScriptRuleDeviceComplianceScriptRuleDataType = "Int64Array"
	DeviceComplianceScriptRuleDeviceComplianceScriptRuleDataTypenone          DeviceComplianceScriptRuleDeviceComplianceScriptRuleDataType = "None"
	DeviceComplianceScriptRuleDeviceComplianceScriptRuleDataTypestring        DeviceComplianceScriptRuleDeviceComplianceScriptRuleDataType = "String"
	DeviceComplianceScriptRuleDeviceComplianceScriptRuleDataTypestringArray   DeviceComplianceScriptRuleDeviceComplianceScriptRuleDataType = "StringArray"
	DeviceComplianceScriptRuleDeviceComplianceScriptRuleDataTypeversion       DeviceComplianceScriptRuleDeviceComplianceScriptRuleDataType = "Version"
	DeviceComplianceScriptRuleDeviceComplianceScriptRuleDataTypeversionArray  DeviceComplianceScriptRuleDeviceComplianceScriptRuleDataType = "VersionArray"
	DeviceComplianceScriptRuleDeviceComplianceScriptRuleDataTypexml           DeviceComplianceScriptRuleDeviceComplianceScriptRuleDataType = "Xml"
)

func PossibleValuesForDeviceComplianceScriptRuleDeviceComplianceScriptRuleDataType() []string {
	return []string{
		string(DeviceComplianceScriptRuleDeviceComplianceScriptRuleDataTypebase64),
		string(DeviceComplianceScriptRuleDeviceComplianceScriptRuleDataTypeboolean),
		string(DeviceComplianceScriptRuleDeviceComplianceScriptRuleDataTypebooleanArray),
		string(DeviceComplianceScriptRuleDeviceComplianceScriptRuleDataTypedateTime),
		string(DeviceComplianceScriptRuleDeviceComplianceScriptRuleDataTypedateTimeArray),
		string(DeviceComplianceScriptRuleDeviceComplianceScriptRuleDataTypedouble),
		string(DeviceComplianceScriptRuleDeviceComplianceScriptRuleDataTypedoubleArray),
		string(DeviceComplianceScriptRuleDeviceComplianceScriptRuleDataTypeint64),
		string(DeviceComplianceScriptRuleDeviceComplianceScriptRuleDataTypeint64Array),
		string(DeviceComplianceScriptRuleDeviceComplianceScriptRuleDataTypenone),
		string(DeviceComplianceScriptRuleDeviceComplianceScriptRuleDataTypestring),
		string(DeviceComplianceScriptRuleDeviceComplianceScriptRuleDataTypestringArray),
		string(DeviceComplianceScriptRuleDeviceComplianceScriptRuleDataTypeversion),
		string(DeviceComplianceScriptRuleDeviceComplianceScriptRuleDataTypeversionArray),
		string(DeviceComplianceScriptRuleDeviceComplianceScriptRuleDataTypexml),
	}
}

func parseDeviceComplianceScriptRuleDeviceComplianceScriptRuleDataType(input string) (*DeviceComplianceScriptRuleDeviceComplianceScriptRuleDataType, error) {
	vals := map[string]DeviceComplianceScriptRuleDeviceComplianceScriptRuleDataType{
		"base64":        DeviceComplianceScriptRuleDeviceComplianceScriptRuleDataTypebase64,
		"boolean":       DeviceComplianceScriptRuleDeviceComplianceScriptRuleDataTypeboolean,
		"booleanarray":  DeviceComplianceScriptRuleDeviceComplianceScriptRuleDataTypebooleanArray,
		"datetime":      DeviceComplianceScriptRuleDeviceComplianceScriptRuleDataTypedateTime,
		"datetimearray": DeviceComplianceScriptRuleDeviceComplianceScriptRuleDataTypedateTimeArray,
		"double":        DeviceComplianceScriptRuleDeviceComplianceScriptRuleDataTypedouble,
		"doublearray":   DeviceComplianceScriptRuleDeviceComplianceScriptRuleDataTypedoubleArray,
		"int64":         DeviceComplianceScriptRuleDeviceComplianceScriptRuleDataTypeint64,
		"int64array":    DeviceComplianceScriptRuleDeviceComplianceScriptRuleDataTypeint64Array,
		"none":          DeviceComplianceScriptRuleDeviceComplianceScriptRuleDataTypenone,
		"string":        DeviceComplianceScriptRuleDeviceComplianceScriptRuleDataTypestring,
		"stringarray":   DeviceComplianceScriptRuleDeviceComplianceScriptRuleDataTypestringArray,
		"version":       DeviceComplianceScriptRuleDeviceComplianceScriptRuleDataTypeversion,
		"versionarray":  DeviceComplianceScriptRuleDeviceComplianceScriptRuleDataTypeversionArray,
		"xml":           DeviceComplianceScriptRuleDeviceComplianceScriptRuleDataTypexml,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceComplianceScriptRuleDeviceComplianceScriptRuleDataType(input)
	return &out, nil
}
