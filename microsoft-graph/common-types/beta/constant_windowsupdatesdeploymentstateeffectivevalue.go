package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUpdatesDeploymentStateEffectiveValue string

const (
	WindowsUpdatesDeploymentStateEffectiveValue_Archived  WindowsUpdatesDeploymentStateEffectiveValue = "archived"
	WindowsUpdatesDeploymentStateEffectiveValue_Faulted   WindowsUpdatesDeploymentStateEffectiveValue = "faulted"
	WindowsUpdatesDeploymentStateEffectiveValue_Offering  WindowsUpdatesDeploymentStateEffectiveValue = "offering"
	WindowsUpdatesDeploymentStateEffectiveValue_Paused    WindowsUpdatesDeploymentStateEffectiveValue = "paused"
	WindowsUpdatesDeploymentStateEffectiveValue_Scheduled WindowsUpdatesDeploymentStateEffectiveValue = "scheduled"
)

func PossibleValuesForWindowsUpdatesDeploymentStateEffectiveValue() []string {
	return []string{
		string(WindowsUpdatesDeploymentStateEffectiveValue_Archived),
		string(WindowsUpdatesDeploymentStateEffectiveValue_Faulted),
		string(WindowsUpdatesDeploymentStateEffectiveValue_Offering),
		string(WindowsUpdatesDeploymentStateEffectiveValue_Paused),
		string(WindowsUpdatesDeploymentStateEffectiveValue_Scheduled),
	}
}

func (s *WindowsUpdatesDeploymentStateEffectiveValue) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsUpdatesDeploymentStateEffectiveValue(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsUpdatesDeploymentStateEffectiveValue(input string) (*WindowsUpdatesDeploymentStateEffectiveValue, error) {
	vals := map[string]WindowsUpdatesDeploymentStateEffectiveValue{
		"archived":  WindowsUpdatesDeploymentStateEffectiveValue_Archived,
		"faulted":   WindowsUpdatesDeploymentStateEffectiveValue_Faulted,
		"offering":  WindowsUpdatesDeploymentStateEffectiveValue_Offering,
		"paused":    WindowsUpdatesDeploymentStateEffectiveValue_Paused,
		"scheduled": WindowsUpdatesDeploymentStateEffectiveValue_Scheduled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsUpdatesDeploymentStateEffectiveValue(input)
	return &out, nil
}
