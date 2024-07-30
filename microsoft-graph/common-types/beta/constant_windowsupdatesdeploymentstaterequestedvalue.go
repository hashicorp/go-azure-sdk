package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUpdatesDeploymentStateRequestedValue string

const (
	WindowsUpdatesDeploymentStateRequestedValue_Archived WindowsUpdatesDeploymentStateRequestedValue = "archived"
	WindowsUpdatesDeploymentStateRequestedValue_None     WindowsUpdatesDeploymentStateRequestedValue = "none"
	WindowsUpdatesDeploymentStateRequestedValue_Paused   WindowsUpdatesDeploymentStateRequestedValue = "paused"
)

func PossibleValuesForWindowsUpdatesDeploymentStateRequestedValue() []string {
	return []string{
		string(WindowsUpdatesDeploymentStateRequestedValue_Archived),
		string(WindowsUpdatesDeploymentStateRequestedValue_None),
		string(WindowsUpdatesDeploymentStateRequestedValue_Paused),
	}
}

func (s *WindowsUpdatesDeploymentStateRequestedValue) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsUpdatesDeploymentStateRequestedValue(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsUpdatesDeploymentStateRequestedValue(input string) (*WindowsUpdatesDeploymentStateRequestedValue, error) {
	vals := map[string]WindowsUpdatesDeploymentStateRequestedValue{
		"archived": WindowsUpdatesDeploymentStateRequestedValue_Archived,
		"none":     WindowsUpdatesDeploymentStateRequestedValue_None,
		"paused":   WindowsUpdatesDeploymentStateRequestedValue_Paused,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsUpdatesDeploymentStateRequestedValue(input)
	return &out, nil
}
