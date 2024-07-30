package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SwapShiftsChangeRequestAssignedTo string

const (
	SwapShiftsChangeRequestAssignedTo_Manager   SwapShiftsChangeRequestAssignedTo = "manager"
	SwapShiftsChangeRequestAssignedTo_Recipient SwapShiftsChangeRequestAssignedTo = "recipient"
	SwapShiftsChangeRequestAssignedTo_Sender    SwapShiftsChangeRequestAssignedTo = "sender"
	SwapShiftsChangeRequestAssignedTo_System    SwapShiftsChangeRequestAssignedTo = "system"
)

func PossibleValuesForSwapShiftsChangeRequestAssignedTo() []string {
	return []string{
		string(SwapShiftsChangeRequestAssignedTo_Manager),
		string(SwapShiftsChangeRequestAssignedTo_Recipient),
		string(SwapShiftsChangeRequestAssignedTo_Sender),
		string(SwapShiftsChangeRequestAssignedTo_System),
	}
}

func (s *SwapShiftsChangeRequestAssignedTo) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSwapShiftsChangeRequestAssignedTo(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSwapShiftsChangeRequestAssignedTo(input string) (*SwapShiftsChangeRequestAssignedTo, error) {
	vals := map[string]SwapShiftsChangeRequestAssignedTo{
		"manager":   SwapShiftsChangeRequestAssignedTo_Manager,
		"recipient": SwapShiftsChangeRequestAssignedTo_Recipient,
		"sender":    SwapShiftsChangeRequestAssignedTo_Sender,
		"system":    SwapShiftsChangeRequestAssignedTo_System,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SwapShiftsChangeRequestAssignedTo(input)
	return &out, nil
}
