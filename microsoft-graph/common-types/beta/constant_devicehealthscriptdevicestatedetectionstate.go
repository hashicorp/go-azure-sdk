package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceHealthScriptDeviceStateDetectionState string

const (
	DeviceHealthScriptDeviceStateDetectionState_Fail          DeviceHealthScriptDeviceStateDetectionState = "fail"
	DeviceHealthScriptDeviceStateDetectionState_NotApplicable DeviceHealthScriptDeviceStateDetectionState = "notApplicable"
	DeviceHealthScriptDeviceStateDetectionState_Pending       DeviceHealthScriptDeviceStateDetectionState = "pending"
	DeviceHealthScriptDeviceStateDetectionState_ScriptError   DeviceHealthScriptDeviceStateDetectionState = "scriptError"
	DeviceHealthScriptDeviceStateDetectionState_Success       DeviceHealthScriptDeviceStateDetectionState = "success"
	DeviceHealthScriptDeviceStateDetectionState_Unknown       DeviceHealthScriptDeviceStateDetectionState = "unknown"
)

func PossibleValuesForDeviceHealthScriptDeviceStateDetectionState() []string {
	return []string{
		string(DeviceHealthScriptDeviceStateDetectionState_Fail),
		string(DeviceHealthScriptDeviceStateDetectionState_NotApplicable),
		string(DeviceHealthScriptDeviceStateDetectionState_Pending),
		string(DeviceHealthScriptDeviceStateDetectionState_ScriptError),
		string(DeviceHealthScriptDeviceStateDetectionState_Success),
		string(DeviceHealthScriptDeviceStateDetectionState_Unknown),
	}
}

func (s *DeviceHealthScriptDeviceStateDetectionState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceHealthScriptDeviceStateDetectionState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceHealthScriptDeviceStateDetectionState(input string) (*DeviceHealthScriptDeviceStateDetectionState, error) {
	vals := map[string]DeviceHealthScriptDeviceStateDetectionState{
		"fail":          DeviceHealthScriptDeviceStateDetectionState_Fail,
		"notapplicable": DeviceHealthScriptDeviceStateDetectionState_NotApplicable,
		"pending":       DeviceHealthScriptDeviceStateDetectionState_Pending,
		"scripterror":   DeviceHealthScriptDeviceStateDetectionState_ScriptError,
		"success":       DeviceHealthScriptDeviceStateDetectionState_Success,
		"unknown":       DeviceHealthScriptDeviceStateDetectionState_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceHealthScriptDeviceStateDetectionState(input)
	return &out, nil
}
