package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceHealthScriptDeviceStateRemediationState string

const (
	DeviceHealthScriptDeviceStateRemediationState_RemediationFailed DeviceHealthScriptDeviceStateRemediationState = "remediationFailed"
	DeviceHealthScriptDeviceStateRemediationState_ScriptError       DeviceHealthScriptDeviceStateRemediationState = "scriptError"
	DeviceHealthScriptDeviceStateRemediationState_Skipped           DeviceHealthScriptDeviceStateRemediationState = "skipped"
	DeviceHealthScriptDeviceStateRemediationState_Success           DeviceHealthScriptDeviceStateRemediationState = "success"
	DeviceHealthScriptDeviceStateRemediationState_Unknown           DeviceHealthScriptDeviceStateRemediationState = "unknown"
)

func PossibleValuesForDeviceHealthScriptDeviceStateRemediationState() []string {
	return []string{
		string(DeviceHealthScriptDeviceStateRemediationState_RemediationFailed),
		string(DeviceHealthScriptDeviceStateRemediationState_ScriptError),
		string(DeviceHealthScriptDeviceStateRemediationState_Skipped),
		string(DeviceHealthScriptDeviceStateRemediationState_Success),
		string(DeviceHealthScriptDeviceStateRemediationState_Unknown),
	}
}

func (s *DeviceHealthScriptDeviceStateRemediationState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceHealthScriptDeviceStateRemediationState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceHealthScriptDeviceStateRemediationState(input string) (*DeviceHealthScriptDeviceStateRemediationState, error) {
	vals := map[string]DeviceHealthScriptDeviceStateRemediationState{
		"remediationfailed": DeviceHealthScriptDeviceStateRemediationState_RemediationFailed,
		"scripterror":       DeviceHealthScriptDeviceStateRemediationState_ScriptError,
		"skipped":           DeviceHealthScriptDeviceStateRemediationState_Skipped,
		"success":           DeviceHealthScriptDeviceStateRemediationState_Success,
		"unknown":           DeviceHealthScriptDeviceStateRemediationState_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceHealthScriptDeviceStateRemediationState(input)
	return &out, nil
}
