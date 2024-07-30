package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ScheduleProvisionStatus string

const (
	ScheduleProvisionStatus_Completed  ScheduleProvisionStatus = "Completed"
	ScheduleProvisionStatus_Failed     ScheduleProvisionStatus = "Failed"
	ScheduleProvisionStatus_NotStarted ScheduleProvisionStatus = "NotStarted"
	ScheduleProvisionStatus_Running    ScheduleProvisionStatus = "Running"
)

func PossibleValuesForScheduleProvisionStatus() []string {
	return []string{
		string(ScheduleProvisionStatus_Completed),
		string(ScheduleProvisionStatus_Failed),
		string(ScheduleProvisionStatus_NotStarted),
		string(ScheduleProvisionStatus_Running),
	}
}

func (s *ScheduleProvisionStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseScheduleProvisionStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseScheduleProvisionStatus(input string) (*ScheduleProvisionStatus, error) {
	vals := map[string]ScheduleProvisionStatus{
		"completed":  ScheduleProvisionStatus_Completed,
		"failed":     ScheduleProvisionStatus_Failed,
		"notstarted": ScheduleProvisionStatus_NotStarted,
		"running":    ScheduleProvisionStatus_Running,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ScheduleProvisionStatus(input)
	return &out, nil
}
