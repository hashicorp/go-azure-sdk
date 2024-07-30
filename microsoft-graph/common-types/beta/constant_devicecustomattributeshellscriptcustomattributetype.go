package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceCustomAttributeShellScriptCustomAttributeType string

const (
	DeviceCustomAttributeShellScriptCustomAttributeType_DateTime DeviceCustomAttributeShellScriptCustomAttributeType = "dateTime"
	DeviceCustomAttributeShellScriptCustomAttributeType_Integer  DeviceCustomAttributeShellScriptCustomAttributeType = "integer"
	DeviceCustomAttributeShellScriptCustomAttributeType_String   DeviceCustomAttributeShellScriptCustomAttributeType = "string"
)

func PossibleValuesForDeviceCustomAttributeShellScriptCustomAttributeType() []string {
	return []string{
		string(DeviceCustomAttributeShellScriptCustomAttributeType_DateTime),
		string(DeviceCustomAttributeShellScriptCustomAttributeType_Integer),
		string(DeviceCustomAttributeShellScriptCustomAttributeType_String),
	}
}

func (s *DeviceCustomAttributeShellScriptCustomAttributeType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceCustomAttributeShellScriptCustomAttributeType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceCustomAttributeShellScriptCustomAttributeType(input string) (*DeviceCustomAttributeShellScriptCustomAttributeType, error) {
	vals := map[string]DeviceCustomAttributeShellScriptCustomAttributeType{
		"datetime": DeviceCustomAttributeShellScriptCustomAttributeType_DateTime,
		"integer":  DeviceCustomAttributeShellScriptCustomAttributeType_Integer,
		"string":   DeviceCustomAttributeShellScriptCustomAttributeType_String,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceCustomAttributeShellScriptCustomAttributeType(input)
	return &out, nil
}
