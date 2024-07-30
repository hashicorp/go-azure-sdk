package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsAutopilotDeviceIdentityEnrollmentState string

const (
	WindowsAutopilotDeviceIdentityEnrollmentState_Enrolled     WindowsAutopilotDeviceIdentityEnrollmentState = "enrolled"
	WindowsAutopilotDeviceIdentityEnrollmentState_Failed       WindowsAutopilotDeviceIdentityEnrollmentState = "failed"
	WindowsAutopilotDeviceIdentityEnrollmentState_NotContacted WindowsAutopilotDeviceIdentityEnrollmentState = "notContacted"
	WindowsAutopilotDeviceIdentityEnrollmentState_PendingReset WindowsAutopilotDeviceIdentityEnrollmentState = "pendingReset"
	WindowsAutopilotDeviceIdentityEnrollmentState_Unknown      WindowsAutopilotDeviceIdentityEnrollmentState = "unknown"
)

func PossibleValuesForWindowsAutopilotDeviceIdentityEnrollmentState() []string {
	return []string{
		string(WindowsAutopilotDeviceIdentityEnrollmentState_Enrolled),
		string(WindowsAutopilotDeviceIdentityEnrollmentState_Failed),
		string(WindowsAutopilotDeviceIdentityEnrollmentState_NotContacted),
		string(WindowsAutopilotDeviceIdentityEnrollmentState_PendingReset),
		string(WindowsAutopilotDeviceIdentityEnrollmentState_Unknown),
	}
}

func (s *WindowsAutopilotDeviceIdentityEnrollmentState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsAutopilotDeviceIdentityEnrollmentState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsAutopilotDeviceIdentityEnrollmentState(input string) (*WindowsAutopilotDeviceIdentityEnrollmentState, error) {
	vals := map[string]WindowsAutopilotDeviceIdentityEnrollmentState{
		"enrolled":     WindowsAutopilotDeviceIdentityEnrollmentState_Enrolled,
		"failed":       WindowsAutopilotDeviceIdentityEnrollmentState_Failed,
		"notcontacted": WindowsAutopilotDeviceIdentityEnrollmentState_NotContacted,
		"pendingreset": WindowsAutopilotDeviceIdentityEnrollmentState_PendingReset,
		"unknown":      WindowsAutopilotDeviceIdentityEnrollmentState_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsAutopilotDeviceIdentityEnrollmentState(input)
	return &out, nil
}
