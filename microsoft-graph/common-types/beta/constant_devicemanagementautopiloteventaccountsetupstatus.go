package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementAutopilotEventAccountSetupStatus string

const (
	DeviceManagementAutopilotEventAccountSetupStatus_Disabled           DeviceManagementAutopilotEventAccountSetupStatus = "disabled"
	DeviceManagementAutopilotEventAccountSetupStatus_Failure            DeviceManagementAutopilotEventAccountSetupStatus = "failure"
	DeviceManagementAutopilotEventAccountSetupStatus_InProgress         DeviceManagementAutopilotEventAccountSetupStatus = "inProgress"
	DeviceManagementAutopilotEventAccountSetupStatus_NotAttempted       DeviceManagementAutopilotEventAccountSetupStatus = "notAttempted"
	DeviceManagementAutopilotEventAccountSetupStatus_Success            DeviceManagementAutopilotEventAccountSetupStatus = "success"
	DeviceManagementAutopilotEventAccountSetupStatus_SuccessOnRetry     DeviceManagementAutopilotEventAccountSetupStatus = "successOnRetry"
	DeviceManagementAutopilotEventAccountSetupStatus_SuccessWithTimeout DeviceManagementAutopilotEventAccountSetupStatus = "successWithTimeout"
	DeviceManagementAutopilotEventAccountSetupStatus_Unknown            DeviceManagementAutopilotEventAccountSetupStatus = "unknown"
)

func PossibleValuesForDeviceManagementAutopilotEventAccountSetupStatus() []string {
	return []string{
		string(DeviceManagementAutopilotEventAccountSetupStatus_Disabled),
		string(DeviceManagementAutopilotEventAccountSetupStatus_Failure),
		string(DeviceManagementAutopilotEventAccountSetupStatus_InProgress),
		string(DeviceManagementAutopilotEventAccountSetupStatus_NotAttempted),
		string(DeviceManagementAutopilotEventAccountSetupStatus_Success),
		string(DeviceManagementAutopilotEventAccountSetupStatus_SuccessOnRetry),
		string(DeviceManagementAutopilotEventAccountSetupStatus_SuccessWithTimeout),
		string(DeviceManagementAutopilotEventAccountSetupStatus_Unknown),
	}
}

func (s *DeviceManagementAutopilotEventAccountSetupStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementAutopilotEventAccountSetupStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementAutopilotEventAccountSetupStatus(input string) (*DeviceManagementAutopilotEventAccountSetupStatus, error) {
	vals := map[string]DeviceManagementAutopilotEventAccountSetupStatus{
		"disabled":           DeviceManagementAutopilotEventAccountSetupStatus_Disabled,
		"failure":            DeviceManagementAutopilotEventAccountSetupStatus_Failure,
		"inprogress":         DeviceManagementAutopilotEventAccountSetupStatus_InProgress,
		"notattempted":       DeviceManagementAutopilotEventAccountSetupStatus_NotAttempted,
		"success":            DeviceManagementAutopilotEventAccountSetupStatus_Success,
		"successonretry":     DeviceManagementAutopilotEventAccountSetupStatus_SuccessOnRetry,
		"successwithtimeout": DeviceManagementAutopilotEventAccountSetupStatus_SuccessWithTimeout,
		"unknown":            DeviceManagementAutopilotEventAccountSetupStatus_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementAutopilotEventAccountSetupStatus(input)
	return &out, nil
}
