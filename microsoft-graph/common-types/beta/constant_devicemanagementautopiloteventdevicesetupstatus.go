package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementAutopilotEventDeviceSetupStatus string

const (
	DeviceManagementAutopilotEventDeviceSetupStatus_Disabled           DeviceManagementAutopilotEventDeviceSetupStatus = "disabled"
	DeviceManagementAutopilotEventDeviceSetupStatus_Failure            DeviceManagementAutopilotEventDeviceSetupStatus = "failure"
	DeviceManagementAutopilotEventDeviceSetupStatus_InProgress         DeviceManagementAutopilotEventDeviceSetupStatus = "inProgress"
	DeviceManagementAutopilotEventDeviceSetupStatus_NotAttempted       DeviceManagementAutopilotEventDeviceSetupStatus = "notAttempted"
	DeviceManagementAutopilotEventDeviceSetupStatus_Success            DeviceManagementAutopilotEventDeviceSetupStatus = "success"
	DeviceManagementAutopilotEventDeviceSetupStatus_SuccessOnRetry     DeviceManagementAutopilotEventDeviceSetupStatus = "successOnRetry"
	DeviceManagementAutopilotEventDeviceSetupStatus_SuccessWithTimeout DeviceManagementAutopilotEventDeviceSetupStatus = "successWithTimeout"
	DeviceManagementAutopilotEventDeviceSetupStatus_Unknown            DeviceManagementAutopilotEventDeviceSetupStatus = "unknown"
)

func PossibleValuesForDeviceManagementAutopilotEventDeviceSetupStatus() []string {
	return []string{
		string(DeviceManagementAutopilotEventDeviceSetupStatus_Disabled),
		string(DeviceManagementAutopilotEventDeviceSetupStatus_Failure),
		string(DeviceManagementAutopilotEventDeviceSetupStatus_InProgress),
		string(DeviceManagementAutopilotEventDeviceSetupStatus_NotAttempted),
		string(DeviceManagementAutopilotEventDeviceSetupStatus_Success),
		string(DeviceManagementAutopilotEventDeviceSetupStatus_SuccessOnRetry),
		string(DeviceManagementAutopilotEventDeviceSetupStatus_SuccessWithTimeout),
		string(DeviceManagementAutopilotEventDeviceSetupStatus_Unknown),
	}
}

func (s *DeviceManagementAutopilotEventDeviceSetupStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementAutopilotEventDeviceSetupStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementAutopilotEventDeviceSetupStatus(input string) (*DeviceManagementAutopilotEventDeviceSetupStatus, error) {
	vals := map[string]DeviceManagementAutopilotEventDeviceSetupStatus{
		"disabled":           DeviceManagementAutopilotEventDeviceSetupStatus_Disabled,
		"failure":            DeviceManagementAutopilotEventDeviceSetupStatus_Failure,
		"inprogress":         DeviceManagementAutopilotEventDeviceSetupStatus_InProgress,
		"notattempted":       DeviceManagementAutopilotEventDeviceSetupStatus_NotAttempted,
		"success":            DeviceManagementAutopilotEventDeviceSetupStatus_Success,
		"successonretry":     DeviceManagementAutopilotEventDeviceSetupStatus_SuccessOnRetry,
		"successwithtimeout": DeviceManagementAutopilotEventDeviceSetupStatus_SuccessWithTimeout,
		"unknown":            DeviceManagementAutopilotEventDeviceSetupStatus_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementAutopilotEventDeviceSetupStatus(input)
	return &out, nil
}
