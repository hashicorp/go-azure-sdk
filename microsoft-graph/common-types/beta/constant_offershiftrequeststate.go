package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OfferShiftRequestState string

const (
	OfferShiftRequestState_Approved OfferShiftRequestState = "approved"
	OfferShiftRequestState_Declined OfferShiftRequestState = "declined"
	OfferShiftRequestState_Pending  OfferShiftRequestState = "pending"
)

func PossibleValuesForOfferShiftRequestState() []string {
	return []string{
		string(OfferShiftRequestState_Approved),
		string(OfferShiftRequestState_Declined),
		string(OfferShiftRequestState_Pending),
	}
}

func (s *OfferShiftRequestState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseOfferShiftRequestState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseOfferShiftRequestState(input string) (*OfferShiftRequestState, error) {
	vals := map[string]OfferShiftRequestState{
		"approved": OfferShiftRequestState_Approved,
		"declined": OfferShiftRequestState_Declined,
		"pending":  OfferShiftRequestState_Pending,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OfferShiftRequestState(input)
	return &out, nil
}
