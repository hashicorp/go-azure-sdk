package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ScheduleChangeRequestAssignedTo string

const (
	ScheduleChangeRequestAssignedTo_Manager   ScheduleChangeRequestAssignedTo = "manager"
	ScheduleChangeRequestAssignedTo_Recipient ScheduleChangeRequestAssignedTo = "recipient"
	ScheduleChangeRequestAssignedTo_Sender    ScheduleChangeRequestAssignedTo = "sender"
	ScheduleChangeRequestAssignedTo_System    ScheduleChangeRequestAssignedTo = "system"
)

func PossibleValuesForScheduleChangeRequestAssignedTo() []string {
	return []string{
		string(ScheduleChangeRequestAssignedTo_Manager),
		string(ScheduleChangeRequestAssignedTo_Recipient),
		string(ScheduleChangeRequestAssignedTo_Sender),
		string(ScheduleChangeRequestAssignedTo_System),
	}
}

func (s *ScheduleChangeRequestAssignedTo) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseScheduleChangeRequestAssignedTo(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseScheduleChangeRequestAssignedTo(input string) (*ScheduleChangeRequestAssignedTo, error) {
	vals := map[string]ScheduleChangeRequestAssignedTo{
		"manager":   ScheduleChangeRequestAssignedTo_Manager,
		"recipient": ScheduleChangeRequestAssignedTo_Recipient,
		"sender":    ScheduleChangeRequestAssignedTo_Sender,
		"system":    ScheduleChangeRequestAssignedTo_System,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ScheduleChangeRequestAssignedTo(input)
	return &out, nil
}
