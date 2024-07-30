package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceComplianceScriptDeviceStateDetectionState string

const (
	DeviceComplianceScriptDeviceStateDetectionState_Fail          DeviceComplianceScriptDeviceStateDetectionState = "fail"
	DeviceComplianceScriptDeviceStateDetectionState_NotApplicable DeviceComplianceScriptDeviceStateDetectionState = "notApplicable"
	DeviceComplianceScriptDeviceStateDetectionState_Pending       DeviceComplianceScriptDeviceStateDetectionState = "pending"
	DeviceComplianceScriptDeviceStateDetectionState_ScriptError   DeviceComplianceScriptDeviceStateDetectionState = "scriptError"
	DeviceComplianceScriptDeviceStateDetectionState_Success       DeviceComplianceScriptDeviceStateDetectionState = "success"
	DeviceComplianceScriptDeviceStateDetectionState_Unknown       DeviceComplianceScriptDeviceStateDetectionState = "unknown"
)

func PossibleValuesForDeviceComplianceScriptDeviceStateDetectionState() []string {
	return []string{
		string(DeviceComplianceScriptDeviceStateDetectionState_Fail),
		string(DeviceComplianceScriptDeviceStateDetectionState_NotApplicable),
		string(DeviceComplianceScriptDeviceStateDetectionState_Pending),
		string(DeviceComplianceScriptDeviceStateDetectionState_ScriptError),
		string(DeviceComplianceScriptDeviceStateDetectionState_Success),
		string(DeviceComplianceScriptDeviceStateDetectionState_Unknown),
	}
}

func (s *DeviceComplianceScriptDeviceStateDetectionState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceComplianceScriptDeviceStateDetectionState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceComplianceScriptDeviceStateDetectionState(input string) (*DeviceComplianceScriptDeviceStateDetectionState, error) {
	vals := map[string]DeviceComplianceScriptDeviceStateDetectionState{
		"fail":          DeviceComplianceScriptDeviceStateDetectionState_Fail,
		"notapplicable": DeviceComplianceScriptDeviceStateDetectionState_NotApplicable,
		"pending":       DeviceComplianceScriptDeviceStateDetectionState_Pending,
		"scripterror":   DeviceComplianceScriptDeviceStateDetectionState_ScriptError,
		"success":       DeviceComplianceScriptDeviceStateDetectionState_Success,
		"unknown":       DeviceComplianceScriptDeviceStateDetectionState_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceComplianceScriptDeviceStateDetectionState(input)
	return &out, nil
}
