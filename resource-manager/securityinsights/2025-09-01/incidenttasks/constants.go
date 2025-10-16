package incidenttasks

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IncidentTaskStatus string

const (
	IncidentTaskStatusCompleted IncidentTaskStatus = "Completed"
	IncidentTaskStatusNew       IncidentTaskStatus = "New"
)

func PossibleValuesForIncidentTaskStatus() []string {
	return []string{
		string(IncidentTaskStatusCompleted),
		string(IncidentTaskStatusNew),
	}
}

func (s *IncidentTaskStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIncidentTaskStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIncidentTaskStatus(input string) (*IncidentTaskStatus, error) {
	vals := map[string]IncidentTaskStatus{
		"completed": IncidentTaskStatusCompleted,
		"new":       IncidentTaskStatusNew,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IncidentTaskStatus(input)
	return &out, nil
}
