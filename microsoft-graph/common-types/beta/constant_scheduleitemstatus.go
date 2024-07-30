package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ScheduleItemStatus string

const (
	ScheduleItemStatus_Busy             ScheduleItemStatus = "busy"
	ScheduleItemStatus_Free             ScheduleItemStatus = "free"
	ScheduleItemStatus_Oof              ScheduleItemStatus = "oof"
	ScheduleItemStatus_Tentative        ScheduleItemStatus = "tentative"
	ScheduleItemStatus_Unknown          ScheduleItemStatus = "unknown"
	ScheduleItemStatus_WorkingElsewhere ScheduleItemStatus = "workingElsewhere"
)

func PossibleValuesForScheduleItemStatus() []string {
	return []string{
		string(ScheduleItemStatus_Busy),
		string(ScheduleItemStatus_Free),
		string(ScheduleItemStatus_Oof),
		string(ScheduleItemStatus_Tentative),
		string(ScheduleItemStatus_Unknown),
		string(ScheduleItemStatus_WorkingElsewhere),
	}
}

func (s *ScheduleItemStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseScheduleItemStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseScheduleItemStatus(input string) (*ScheduleItemStatus, error) {
	vals := map[string]ScheduleItemStatus{
		"busy":             ScheduleItemStatus_Busy,
		"free":             ScheduleItemStatus_Free,
		"oof":              ScheduleItemStatus_Oof,
		"tentative":        ScheduleItemStatus_Tentative,
		"unknown":          ScheduleItemStatus_Unknown,
		"workingelsewhere": ScheduleItemStatus_WorkingElsewhere,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ScheduleItemStatus(input)
	return &out, nil
}
