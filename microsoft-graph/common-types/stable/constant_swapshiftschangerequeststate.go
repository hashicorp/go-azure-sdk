package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SwapShiftsChangeRequestState string

const (
	SwapShiftsChangeRequestState_Approved SwapShiftsChangeRequestState = "approved"
	SwapShiftsChangeRequestState_Declined SwapShiftsChangeRequestState = "declined"
	SwapShiftsChangeRequestState_Pending  SwapShiftsChangeRequestState = "pending"
)

func PossibleValuesForSwapShiftsChangeRequestState() []string {
	return []string{
		string(SwapShiftsChangeRequestState_Approved),
		string(SwapShiftsChangeRequestState_Declined),
		string(SwapShiftsChangeRequestState_Pending),
	}
}

func (s *SwapShiftsChangeRequestState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSwapShiftsChangeRequestState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSwapShiftsChangeRequestState(input string) (*SwapShiftsChangeRequestState, error) {
	vals := map[string]SwapShiftsChangeRequestState{
		"approved": SwapShiftsChangeRequestState_Approved,
		"declined": SwapShiftsChangeRequestState_Declined,
		"pending":  SwapShiftsChangeRequestState_Pending,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SwapShiftsChangeRequestState(input)
	return &out, nil
}
