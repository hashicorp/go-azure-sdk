package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceHealthScriptPolicyStateDetectionState string

const (
	DeviceHealthScriptPolicyStateDetectionState_Fail          DeviceHealthScriptPolicyStateDetectionState = "fail"
	DeviceHealthScriptPolicyStateDetectionState_NotApplicable DeviceHealthScriptPolicyStateDetectionState = "notApplicable"
	DeviceHealthScriptPolicyStateDetectionState_Pending       DeviceHealthScriptPolicyStateDetectionState = "pending"
	DeviceHealthScriptPolicyStateDetectionState_ScriptError   DeviceHealthScriptPolicyStateDetectionState = "scriptError"
	DeviceHealthScriptPolicyStateDetectionState_Success       DeviceHealthScriptPolicyStateDetectionState = "success"
	DeviceHealthScriptPolicyStateDetectionState_Unknown       DeviceHealthScriptPolicyStateDetectionState = "unknown"
)

func PossibleValuesForDeviceHealthScriptPolicyStateDetectionState() []string {
	return []string{
		string(DeviceHealthScriptPolicyStateDetectionState_Fail),
		string(DeviceHealthScriptPolicyStateDetectionState_NotApplicable),
		string(DeviceHealthScriptPolicyStateDetectionState_Pending),
		string(DeviceHealthScriptPolicyStateDetectionState_ScriptError),
		string(DeviceHealthScriptPolicyStateDetectionState_Success),
		string(DeviceHealthScriptPolicyStateDetectionState_Unknown),
	}
}

func (s *DeviceHealthScriptPolicyStateDetectionState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceHealthScriptPolicyStateDetectionState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceHealthScriptPolicyStateDetectionState(input string) (*DeviceHealthScriptPolicyStateDetectionState, error) {
	vals := map[string]DeviceHealthScriptPolicyStateDetectionState{
		"fail":          DeviceHealthScriptPolicyStateDetectionState_Fail,
		"notapplicable": DeviceHealthScriptPolicyStateDetectionState_NotApplicable,
		"pending":       DeviceHealthScriptPolicyStateDetectionState_Pending,
		"scripterror":   DeviceHealthScriptPolicyStateDetectionState_ScriptError,
		"success":       DeviceHealthScriptPolicyStateDetectionState_Success,
		"unknown":       DeviceHealthScriptPolicyStateDetectionState_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceHealthScriptPolicyStateDetectionState(input)
	return &out, nil
}
