package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementAutopilotEventDeploymentState string

const (
	DeviceManagementAutopilotEventDeploymentState_Disabled           DeviceManagementAutopilotEventDeploymentState = "disabled"
	DeviceManagementAutopilotEventDeploymentState_Failure            DeviceManagementAutopilotEventDeploymentState = "failure"
	DeviceManagementAutopilotEventDeploymentState_InProgress         DeviceManagementAutopilotEventDeploymentState = "inProgress"
	DeviceManagementAutopilotEventDeploymentState_NotAttempted       DeviceManagementAutopilotEventDeploymentState = "notAttempted"
	DeviceManagementAutopilotEventDeploymentState_Success            DeviceManagementAutopilotEventDeploymentState = "success"
	DeviceManagementAutopilotEventDeploymentState_SuccessOnRetry     DeviceManagementAutopilotEventDeploymentState = "successOnRetry"
	DeviceManagementAutopilotEventDeploymentState_SuccessWithTimeout DeviceManagementAutopilotEventDeploymentState = "successWithTimeout"
	DeviceManagementAutopilotEventDeploymentState_Unknown            DeviceManagementAutopilotEventDeploymentState = "unknown"
)

func PossibleValuesForDeviceManagementAutopilotEventDeploymentState() []string {
	return []string{
		string(DeviceManagementAutopilotEventDeploymentState_Disabled),
		string(DeviceManagementAutopilotEventDeploymentState_Failure),
		string(DeviceManagementAutopilotEventDeploymentState_InProgress),
		string(DeviceManagementAutopilotEventDeploymentState_NotAttempted),
		string(DeviceManagementAutopilotEventDeploymentState_Success),
		string(DeviceManagementAutopilotEventDeploymentState_SuccessOnRetry),
		string(DeviceManagementAutopilotEventDeploymentState_SuccessWithTimeout),
		string(DeviceManagementAutopilotEventDeploymentState_Unknown),
	}
}

func (s *DeviceManagementAutopilotEventDeploymentState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementAutopilotEventDeploymentState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementAutopilotEventDeploymentState(input string) (*DeviceManagementAutopilotEventDeploymentState, error) {
	vals := map[string]DeviceManagementAutopilotEventDeploymentState{
		"disabled":           DeviceManagementAutopilotEventDeploymentState_Disabled,
		"failure":            DeviceManagementAutopilotEventDeploymentState_Failure,
		"inprogress":         DeviceManagementAutopilotEventDeploymentState_InProgress,
		"notattempted":       DeviceManagementAutopilotEventDeploymentState_NotAttempted,
		"success":            DeviceManagementAutopilotEventDeploymentState_Success,
		"successonretry":     DeviceManagementAutopilotEventDeploymentState_SuccessOnRetry,
		"successwithtimeout": DeviceManagementAutopilotEventDeploymentState_SuccessWithTimeout,
		"unknown":            DeviceManagementAutopilotEventDeploymentState_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementAutopilotEventDeploymentState(input)
	return &out, nil
}
