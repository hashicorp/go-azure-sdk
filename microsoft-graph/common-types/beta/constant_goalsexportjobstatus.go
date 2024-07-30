package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GoalsExportJobStatus string

const (
	GoalsExportJobStatus_Failed     GoalsExportJobStatus = "failed"
	GoalsExportJobStatus_NotStarted GoalsExportJobStatus = "notStarted"
	GoalsExportJobStatus_Running    GoalsExportJobStatus = "running"
	GoalsExportJobStatus_Succeeded  GoalsExportJobStatus = "succeeded"
)

func PossibleValuesForGoalsExportJobStatus() []string {
	return []string{
		string(GoalsExportJobStatus_Failed),
		string(GoalsExportJobStatus_NotStarted),
		string(GoalsExportJobStatus_Running),
		string(GoalsExportJobStatus_Succeeded),
	}
}

func (s *GoalsExportJobStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseGoalsExportJobStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseGoalsExportJobStatus(input string) (*GoalsExportJobStatus, error) {
	vals := map[string]GoalsExportJobStatus{
		"failed":     GoalsExportJobStatus_Failed,
		"notstarted": GoalsExportJobStatus_NotStarted,
		"running":    GoalsExportJobStatus_Running,
		"succeeded":  GoalsExportJobStatus_Succeeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := GoalsExportJobStatus(input)
	return &out, nil
}
