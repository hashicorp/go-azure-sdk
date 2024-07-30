package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementAutopilotEventEnrollmentState string

const (
	DeviceManagementAutopilotEventEnrollmentState_Blocked      DeviceManagementAutopilotEventEnrollmentState = "blocked"
	DeviceManagementAutopilotEventEnrollmentState_Enrolled     DeviceManagementAutopilotEventEnrollmentState = "enrolled"
	DeviceManagementAutopilotEventEnrollmentState_Failed       DeviceManagementAutopilotEventEnrollmentState = "failed"
	DeviceManagementAutopilotEventEnrollmentState_NotContacted DeviceManagementAutopilotEventEnrollmentState = "notContacted"
	DeviceManagementAutopilotEventEnrollmentState_PendingReset DeviceManagementAutopilotEventEnrollmentState = "pendingReset"
	DeviceManagementAutopilotEventEnrollmentState_Unknown      DeviceManagementAutopilotEventEnrollmentState = "unknown"
)

func PossibleValuesForDeviceManagementAutopilotEventEnrollmentState() []string {
	return []string{
		string(DeviceManagementAutopilotEventEnrollmentState_Blocked),
		string(DeviceManagementAutopilotEventEnrollmentState_Enrolled),
		string(DeviceManagementAutopilotEventEnrollmentState_Failed),
		string(DeviceManagementAutopilotEventEnrollmentState_NotContacted),
		string(DeviceManagementAutopilotEventEnrollmentState_PendingReset),
		string(DeviceManagementAutopilotEventEnrollmentState_Unknown),
	}
}

func (s *DeviceManagementAutopilotEventEnrollmentState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementAutopilotEventEnrollmentState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementAutopilotEventEnrollmentState(input string) (*DeviceManagementAutopilotEventEnrollmentState, error) {
	vals := map[string]DeviceManagementAutopilotEventEnrollmentState{
		"blocked":      DeviceManagementAutopilotEventEnrollmentState_Blocked,
		"enrolled":     DeviceManagementAutopilotEventEnrollmentState_Enrolled,
		"failed":       DeviceManagementAutopilotEventEnrollmentState_Failed,
		"notcontacted": DeviceManagementAutopilotEventEnrollmentState_NotContacted,
		"pendingreset": DeviceManagementAutopilotEventEnrollmentState_PendingReset,
		"unknown":      DeviceManagementAutopilotEventEnrollmentState_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementAutopilotEventEnrollmentState(input)
	return &out, nil
}
