package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TimeOffRequestAssignedTo string

const (
	TimeOffRequestAssignedTo_Manager   TimeOffRequestAssignedTo = "manager"
	TimeOffRequestAssignedTo_Recipient TimeOffRequestAssignedTo = "recipient"
	TimeOffRequestAssignedTo_Sender    TimeOffRequestAssignedTo = "sender"
	TimeOffRequestAssignedTo_System    TimeOffRequestAssignedTo = "system"
)

func PossibleValuesForTimeOffRequestAssignedTo() []string {
	return []string{
		string(TimeOffRequestAssignedTo_Manager),
		string(TimeOffRequestAssignedTo_Recipient),
		string(TimeOffRequestAssignedTo_Sender),
		string(TimeOffRequestAssignedTo_System),
	}
}

func (s *TimeOffRequestAssignedTo) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTimeOffRequestAssignedTo(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTimeOffRequestAssignedTo(input string) (*TimeOffRequestAssignedTo, error) {
	vals := map[string]TimeOffRequestAssignedTo{
		"manager":   TimeOffRequestAssignedTo_Manager,
		"recipient": TimeOffRequestAssignedTo_Recipient,
		"sender":    TimeOffRequestAssignedTo_Sender,
		"system":    TimeOffRequestAssignedTo_System,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TimeOffRequestAssignedTo(input)
	return &out, nil
}
