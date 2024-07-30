package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceHealthScriptPolicyStateRemediationState string

const (
	DeviceHealthScriptPolicyStateRemediationState_RemediationFailed DeviceHealthScriptPolicyStateRemediationState = "remediationFailed"
	DeviceHealthScriptPolicyStateRemediationState_ScriptError       DeviceHealthScriptPolicyStateRemediationState = "scriptError"
	DeviceHealthScriptPolicyStateRemediationState_Skipped           DeviceHealthScriptPolicyStateRemediationState = "skipped"
	DeviceHealthScriptPolicyStateRemediationState_Success           DeviceHealthScriptPolicyStateRemediationState = "success"
	DeviceHealthScriptPolicyStateRemediationState_Unknown           DeviceHealthScriptPolicyStateRemediationState = "unknown"
)

func PossibleValuesForDeviceHealthScriptPolicyStateRemediationState() []string {
	return []string{
		string(DeviceHealthScriptPolicyStateRemediationState_RemediationFailed),
		string(DeviceHealthScriptPolicyStateRemediationState_ScriptError),
		string(DeviceHealthScriptPolicyStateRemediationState_Skipped),
		string(DeviceHealthScriptPolicyStateRemediationState_Success),
		string(DeviceHealthScriptPolicyStateRemediationState_Unknown),
	}
}

func (s *DeviceHealthScriptPolicyStateRemediationState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceHealthScriptPolicyStateRemediationState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceHealthScriptPolicyStateRemediationState(input string) (*DeviceHealthScriptPolicyStateRemediationState, error) {
	vals := map[string]DeviceHealthScriptPolicyStateRemediationState{
		"remediationfailed": DeviceHealthScriptPolicyStateRemediationState_RemediationFailed,
		"scripterror":       DeviceHealthScriptPolicyStateRemediationState_ScriptError,
		"skipped":           DeviceHealthScriptPolicyStateRemediationState_Skipped,
		"success":           DeviceHealthScriptPolicyStateRemediationState_Success,
		"unknown":           DeviceHealthScriptPolicyStateRemediationState_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceHealthScriptPolicyStateRemediationState(input)
	return &out, nil
}
