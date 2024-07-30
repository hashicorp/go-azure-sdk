package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceComplianceScriptRuleDeviceComplianceScriptRuleDataType string

const (
	DeviceComplianceScriptRuleDeviceComplianceScriptRuleDataType_Base64        DeviceComplianceScriptRuleDeviceComplianceScriptRuleDataType = "base64"
	DeviceComplianceScriptRuleDeviceComplianceScriptRuleDataType_Boolean       DeviceComplianceScriptRuleDeviceComplianceScriptRuleDataType = "boolean"
	DeviceComplianceScriptRuleDeviceComplianceScriptRuleDataType_BooleanArray  DeviceComplianceScriptRuleDeviceComplianceScriptRuleDataType = "booleanArray"
	DeviceComplianceScriptRuleDeviceComplianceScriptRuleDataType_DateTime      DeviceComplianceScriptRuleDeviceComplianceScriptRuleDataType = "dateTime"
	DeviceComplianceScriptRuleDeviceComplianceScriptRuleDataType_DateTimeArray DeviceComplianceScriptRuleDeviceComplianceScriptRuleDataType = "dateTimeArray"
	DeviceComplianceScriptRuleDeviceComplianceScriptRuleDataType_Double        DeviceComplianceScriptRuleDeviceComplianceScriptRuleDataType = "double"
	DeviceComplianceScriptRuleDeviceComplianceScriptRuleDataType_DoubleArray   DeviceComplianceScriptRuleDeviceComplianceScriptRuleDataType = "doubleArray"
	DeviceComplianceScriptRuleDeviceComplianceScriptRuleDataType_Int64         DeviceComplianceScriptRuleDeviceComplianceScriptRuleDataType = "int64"
	DeviceComplianceScriptRuleDeviceComplianceScriptRuleDataType_Int64Array    DeviceComplianceScriptRuleDeviceComplianceScriptRuleDataType = "int64Array"
	DeviceComplianceScriptRuleDeviceComplianceScriptRuleDataType_None          DeviceComplianceScriptRuleDeviceComplianceScriptRuleDataType = "none"
	DeviceComplianceScriptRuleDeviceComplianceScriptRuleDataType_String        DeviceComplianceScriptRuleDeviceComplianceScriptRuleDataType = "string"
	DeviceComplianceScriptRuleDeviceComplianceScriptRuleDataType_StringArray   DeviceComplianceScriptRuleDeviceComplianceScriptRuleDataType = "stringArray"
	DeviceComplianceScriptRuleDeviceComplianceScriptRuleDataType_Version       DeviceComplianceScriptRuleDeviceComplianceScriptRuleDataType = "version"
	DeviceComplianceScriptRuleDeviceComplianceScriptRuleDataType_VersionArray  DeviceComplianceScriptRuleDeviceComplianceScriptRuleDataType = "versionArray"
	DeviceComplianceScriptRuleDeviceComplianceScriptRuleDataType_Xml           DeviceComplianceScriptRuleDeviceComplianceScriptRuleDataType = "xml"
)

func PossibleValuesForDeviceComplianceScriptRuleDeviceComplianceScriptRuleDataType() []string {
	return []string{
		string(DeviceComplianceScriptRuleDeviceComplianceScriptRuleDataType_Base64),
		string(DeviceComplianceScriptRuleDeviceComplianceScriptRuleDataType_Boolean),
		string(DeviceComplianceScriptRuleDeviceComplianceScriptRuleDataType_BooleanArray),
		string(DeviceComplianceScriptRuleDeviceComplianceScriptRuleDataType_DateTime),
		string(DeviceComplianceScriptRuleDeviceComplianceScriptRuleDataType_DateTimeArray),
		string(DeviceComplianceScriptRuleDeviceComplianceScriptRuleDataType_Double),
		string(DeviceComplianceScriptRuleDeviceComplianceScriptRuleDataType_DoubleArray),
		string(DeviceComplianceScriptRuleDeviceComplianceScriptRuleDataType_Int64),
		string(DeviceComplianceScriptRuleDeviceComplianceScriptRuleDataType_Int64Array),
		string(DeviceComplianceScriptRuleDeviceComplianceScriptRuleDataType_None),
		string(DeviceComplianceScriptRuleDeviceComplianceScriptRuleDataType_String),
		string(DeviceComplianceScriptRuleDeviceComplianceScriptRuleDataType_StringArray),
		string(DeviceComplianceScriptRuleDeviceComplianceScriptRuleDataType_Version),
		string(DeviceComplianceScriptRuleDeviceComplianceScriptRuleDataType_VersionArray),
		string(DeviceComplianceScriptRuleDeviceComplianceScriptRuleDataType_Xml),
	}
}

func (s *DeviceComplianceScriptRuleDeviceComplianceScriptRuleDataType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceComplianceScriptRuleDeviceComplianceScriptRuleDataType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceComplianceScriptRuleDeviceComplianceScriptRuleDataType(input string) (*DeviceComplianceScriptRuleDeviceComplianceScriptRuleDataType, error) {
	vals := map[string]DeviceComplianceScriptRuleDeviceComplianceScriptRuleDataType{
		"base64":        DeviceComplianceScriptRuleDeviceComplianceScriptRuleDataType_Base64,
		"boolean":       DeviceComplianceScriptRuleDeviceComplianceScriptRuleDataType_Boolean,
		"booleanarray":  DeviceComplianceScriptRuleDeviceComplianceScriptRuleDataType_BooleanArray,
		"datetime":      DeviceComplianceScriptRuleDeviceComplianceScriptRuleDataType_DateTime,
		"datetimearray": DeviceComplianceScriptRuleDeviceComplianceScriptRuleDataType_DateTimeArray,
		"double":        DeviceComplianceScriptRuleDeviceComplianceScriptRuleDataType_Double,
		"doublearray":   DeviceComplianceScriptRuleDeviceComplianceScriptRuleDataType_DoubleArray,
		"int64":         DeviceComplianceScriptRuleDeviceComplianceScriptRuleDataType_Int64,
		"int64array":    DeviceComplianceScriptRuleDeviceComplianceScriptRuleDataType_Int64Array,
		"none":          DeviceComplianceScriptRuleDeviceComplianceScriptRuleDataType_None,
		"string":        DeviceComplianceScriptRuleDeviceComplianceScriptRuleDataType_String,
		"stringarray":   DeviceComplianceScriptRuleDeviceComplianceScriptRuleDataType_StringArray,
		"version":       DeviceComplianceScriptRuleDeviceComplianceScriptRuleDataType_Version,
		"versionarray":  DeviceComplianceScriptRuleDeviceComplianceScriptRuleDataType_VersionArray,
		"xml":           DeviceComplianceScriptRuleDeviceComplianceScriptRuleDataType_Xml,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceComplianceScriptRuleDeviceComplianceScriptRuleDataType(input)
	return &out, nil
}
