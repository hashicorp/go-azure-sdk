package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ScheduleChangeRequestState string

const (
	ScheduleChangeRequestState_Approved ScheduleChangeRequestState = "approved"
	ScheduleChangeRequestState_Declined ScheduleChangeRequestState = "declined"
	ScheduleChangeRequestState_Pending  ScheduleChangeRequestState = "pending"
)

func PossibleValuesForScheduleChangeRequestState() []string {
	return []string{
		string(ScheduleChangeRequestState_Approved),
		string(ScheduleChangeRequestState_Declined),
		string(ScheduleChangeRequestState_Pending),
	}
}

func (s *ScheduleChangeRequestState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseScheduleChangeRequestState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseScheduleChangeRequestState(input string) (*ScheduleChangeRequestState, error) {
	vals := map[string]ScheduleChangeRequestState{
		"approved": ScheduleChangeRequestState_Approved,
		"declined": ScheduleChangeRequestState_Declined,
		"pending":  ScheduleChangeRequestState_Pending,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ScheduleChangeRequestState(input)
	return &out, nil
}
