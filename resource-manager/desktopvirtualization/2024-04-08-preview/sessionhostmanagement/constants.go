package sessionhostmanagement

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HostPoolUpdateAction string

const (
	HostPoolUpdateActionCancel HostPoolUpdateAction = "Cancel"
	HostPoolUpdateActionPause  HostPoolUpdateAction = "Pause"
	HostPoolUpdateActionResume HostPoolUpdateAction = "Resume"
	HostPoolUpdateActionRetry  HostPoolUpdateAction = "Retry"
	HostPoolUpdateActionStart  HostPoolUpdateAction = "Start"
)

func PossibleValuesForHostPoolUpdateAction() []string {
	return []string{
		string(HostPoolUpdateActionCancel),
		string(HostPoolUpdateActionPause),
		string(HostPoolUpdateActionResume),
		string(HostPoolUpdateActionRetry),
		string(HostPoolUpdateActionStart),
	}
}

func (s *HostPoolUpdateAction) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseHostPoolUpdateAction(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseHostPoolUpdateAction(input string) (*HostPoolUpdateAction, error) {
	vals := map[string]HostPoolUpdateAction{
		"cancel": HostPoolUpdateActionCancel,
		"pause":  HostPoolUpdateActionPause,
		"resume": HostPoolUpdateActionResume,
		"retry":  HostPoolUpdateActionRetry,
		"start":  HostPoolUpdateActionStart,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := HostPoolUpdateAction(input)
	return &out, nil
}
