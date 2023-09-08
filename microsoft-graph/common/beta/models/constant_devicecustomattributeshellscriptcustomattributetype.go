package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceCustomAttributeShellScriptCustomAttributeType string

const (
	DeviceCustomAttributeShellScriptCustomAttributeTypedateTime DeviceCustomAttributeShellScriptCustomAttributeType = "DateTime"
	DeviceCustomAttributeShellScriptCustomAttributeTypeinteger  DeviceCustomAttributeShellScriptCustomAttributeType = "Integer"
	DeviceCustomAttributeShellScriptCustomAttributeTypestring   DeviceCustomAttributeShellScriptCustomAttributeType = "String"
)

func PossibleValuesForDeviceCustomAttributeShellScriptCustomAttributeType() []string {
	return []string{
		string(DeviceCustomAttributeShellScriptCustomAttributeTypedateTime),
		string(DeviceCustomAttributeShellScriptCustomAttributeTypeinteger),
		string(DeviceCustomAttributeShellScriptCustomAttributeTypestring),
	}
}

func parseDeviceCustomAttributeShellScriptCustomAttributeType(input string) (*DeviceCustomAttributeShellScriptCustomAttributeType, error) {
	vals := map[string]DeviceCustomAttributeShellScriptCustomAttributeType{
		"datetime": DeviceCustomAttributeShellScriptCustomAttributeTypedateTime,
		"integer":  DeviceCustomAttributeShellScriptCustomAttributeTypeinteger,
		"string":   DeviceCustomAttributeShellScriptCustomAttributeTypestring,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceCustomAttributeShellScriptCustomAttributeType(input)
	return &out, nil
}
