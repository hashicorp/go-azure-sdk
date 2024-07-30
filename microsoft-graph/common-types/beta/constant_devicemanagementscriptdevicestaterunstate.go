package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementScriptDeviceStateRunState string

const (
	DeviceManagementScriptDeviceStateRunState_Fail          DeviceManagementScriptDeviceStateRunState = "fail"
	DeviceManagementScriptDeviceStateRunState_NotApplicable DeviceManagementScriptDeviceStateRunState = "notApplicable"
	DeviceManagementScriptDeviceStateRunState_Pending       DeviceManagementScriptDeviceStateRunState = "pending"
	DeviceManagementScriptDeviceStateRunState_ScriptError   DeviceManagementScriptDeviceStateRunState = "scriptError"
	DeviceManagementScriptDeviceStateRunState_Success       DeviceManagementScriptDeviceStateRunState = "success"
	DeviceManagementScriptDeviceStateRunState_Unknown       DeviceManagementScriptDeviceStateRunState = "unknown"
)

func PossibleValuesForDeviceManagementScriptDeviceStateRunState() []string {
	return []string{
		string(DeviceManagementScriptDeviceStateRunState_Fail),
		string(DeviceManagementScriptDeviceStateRunState_NotApplicable),
		string(DeviceManagementScriptDeviceStateRunState_Pending),
		string(DeviceManagementScriptDeviceStateRunState_ScriptError),
		string(DeviceManagementScriptDeviceStateRunState_Success),
		string(DeviceManagementScriptDeviceStateRunState_Unknown),
	}
}

func (s *DeviceManagementScriptDeviceStateRunState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementScriptDeviceStateRunState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementScriptDeviceStateRunState(input string) (*DeviceManagementScriptDeviceStateRunState, error) {
	vals := map[string]DeviceManagementScriptDeviceStateRunState{
		"fail":          DeviceManagementScriptDeviceStateRunState_Fail,
		"notapplicable": DeviceManagementScriptDeviceStateRunState_NotApplicable,
		"pending":       DeviceManagementScriptDeviceStateRunState_Pending,
		"scripterror":   DeviceManagementScriptDeviceStateRunState_ScriptError,
		"success":       DeviceManagementScriptDeviceStateRunState_Success,
		"unknown":       DeviceManagementScriptDeviceStateRunState_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementScriptDeviceStateRunState(input)
	return &out, nil
}
