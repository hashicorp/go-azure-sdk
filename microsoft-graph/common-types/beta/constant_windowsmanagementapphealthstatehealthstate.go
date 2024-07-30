package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsManagementAppHealthStateHealthState string

const (
	WindowsManagementAppHealthStateHealthState_Healthy   WindowsManagementAppHealthStateHealthState = "healthy"
	WindowsManagementAppHealthStateHealthState_Unhealthy WindowsManagementAppHealthStateHealthState = "unhealthy"
	WindowsManagementAppHealthStateHealthState_Unknown   WindowsManagementAppHealthStateHealthState = "unknown"
)

func PossibleValuesForWindowsManagementAppHealthStateHealthState() []string {
	return []string{
		string(WindowsManagementAppHealthStateHealthState_Healthy),
		string(WindowsManagementAppHealthStateHealthState_Unhealthy),
		string(WindowsManagementAppHealthStateHealthState_Unknown),
	}
}

func (s *WindowsManagementAppHealthStateHealthState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsManagementAppHealthStateHealthState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsManagementAppHealthStateHealthState(input string) (*WindowsManagementAppHealthStateHealthState, error) {
	vals := map[string]WindowsManagementAppHealthStateHealthState{
		"healthy":   WindowsManagementAppHealthStateHealthState_Healthy,
		"unhealthy": WindowsManagementAppHealthStateHealthState_Unhealthy,
		"unknown":   WindowsManagementAppHealthStateHealthState_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsManagementAppHealthStateHealthState(input)
	return &out, nil
}
