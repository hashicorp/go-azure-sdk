package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OfferShiftRequestAssignedTo string

const (
	OfferShiftRequestAssignedTo_Manager   OfferShiftRequestAssignedTo = "manager"
	OfferShiftRequestAssignedTo_Recipient OfferShiftRequestAssignedTo = "recipient"
	OfferShiftRequestAssignedTo_Sender    OfferShiftRequestAssignedTo = "sender"
	OfferShiftRequestAssignedTo_System    OfferShiftRequestAssignedTo = "system"
)

func PossibleValuesForOfferShiftRequestAssignedTo() []string {
	return []string{
		string(OfferShiftRequestAssignedTo_Manager),
		string(OfferShiftRequestAssignedTo_Recipient),
		string(OfferShiftRequestAssignedTo_Sender),
		string(OfferShiftRequestAssignedTo_System),
	}
}

func (s *OfferShiftRequestAssignedTo) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseOfferShiftRequestAssignedTo(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseOfferShiftRequestAssignedTo(input string) (*OfferShiftRequestAssignedTo, error) {
	vals := map[string]OfferShiftRequestAssignedTo{
		"manager":   OfferShiftRequestAssignedTo_Manager,
		"recipient": OfferShiftRequestAssignedTo_Recipient,
		"sender":    OfferShiftRequestAssignedTo_Sender,
		"system":    OfferShiftRequestAssignedTo_System,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OfferShiftRequestAssignedTo(input)
	return &out, nil
}
