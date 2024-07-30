package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceHealthScriptDeviceHealthScriptType string

const (
	DeviceHealthScriptDeviceHealthScriptType_DeviceHealthScript     DeviceHealthScriptDeviceHealthScriptType = "deviceHealthScript"
	DeviceHealthScriptDeviceHealthScriptType_ManagedInstallerScript DeviceHealthScriptDeviceHealthScriptType = "managedInstallerScript"
)

func PossibleValuesForDeviceHealthScriptDeviceHealthScriptType() []string {
	return []string{
		string(DeviceHealthScriptDeviceHealthScriptType_DeviceHealthScript),
		string(DeviceHealthScriptDeviceHealthScriptType_ManagedInstallerScript),
	}
}

func (s *DeviceHealthScriptDeviceHealthScriptType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceHealthScriptDeviceHealthScriptType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceHealthScriptDeviceHealthScriptType(input string) (*DeviceHealthScriptDeviceHealthScriptType, error) {
	vals := map[string]DeviceHealthScriptDeviceHealthScriptType{
		"devicehealthscript":     DeviceHealthScriptDeviceHealthScriptType_DeviceHealthScript,
		"managedinstallerscript": DeviceHealthScriptDeviceHealthScriptType_ManagedInstallerScript,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceHealthScriptDeviceHealthScriptType(input)
	return &out, nil
}
