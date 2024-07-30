package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsAutopilotDeviceIdentityUserlessEnrollmentStatus string

const (
	WindowsAutopilotDeviceIdentityUserlessEnrollmentStatus_Allowed WindowsAutopilotDeviceIdentityUserlessEnrollmentStatus = "allowed"
	WindowsAutopilotDeviceIdentityUserlessEnrollmentStatus_Blocked WindowsAutopilotDeviceIdentityUserlessEnrollmentStatus = "blocked"
	WindowsAutopilotDeviceIdentityUserlessEnrollmentStatus_Unknown WindowsAutopilotDeviceIdentityUserlessEnrollmentStatus = "unknown"
)

func PossibleValuesForWindowsAutopilotDeviceIdentityUserlessEnrollmentStatus() []string {
	return []string{
		string(WindowsAutopilotDeviceIdentityUserlessEnrollmentStatus_Allowed),
		string(WindowsAutopilotDeviceIdentityUserlessEnrollmentStatus_Blocked),
		string(WindowsAutopilotDeviceIdentityUserlessEnrollmentStatus_Unknown),
	}
}

func (s *WindowsAutopilotDeviceIdentityUserlessEnrollmentStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsAutopilotDeviceIdentityUserlessEnrollmentStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsAutopilotDeviceIdentityUserlessEnrollmentStatus(input string) (*WindowsAutopilotDeviceIdentityUserlessEnrollmentStatus, error) {
	vals := map[string]WindowsAutopilotDeviceIdentityUserlessEnrollmentStatus{
		"allowed": WindowsAutopilotDeviceIdentityUserlessEnrollmentStatus_Allowed,
		"blocked": WindowsAutopilotDeviceIdentityUserlessEnrollmentStatus_Blocked,
		"unknown": WindowsAutopilotDeviceIdentityUserlessEnrollmentStatus_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsAutopilotDeviceIdentityUserlessEnrollmentStatus(input)
	return &out, nil
}
