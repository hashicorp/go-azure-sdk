package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TimeOffRequestState string

const (
	TimeOffRequestState_Approved TimeOffRequestState = "approved"
	TimeOffRequestState_Declined TimeOffRequestState = "declined"
	TimeOffRequestState_Pending  TimeOffRequestState = "pending"
)

func PossibleValuesForTimeOffRequestState() []string {
	return []string{
		string(TimeOffRequestState_Approved),
		string(TimeOffRequestState_Declined),
		string(TimeOffRequestState_Pending),
	}
}

func (s *TimeOffRequestState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTimeOffRequestState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTimeOffRequestState(input string) (*TimeOffRequestState, error) {
	vals := map[string]TimeOffRequestState{
		"approved": TimeOffRequestState_Approved,
		"declined": TimeOffRequestState_Declined,
		"pending":  TimeOffRequestState_Pending,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TimeOffRequestState(input)
	return &out, nil
}
