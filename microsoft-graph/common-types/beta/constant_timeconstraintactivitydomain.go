package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TimeConstraintActivityDomain string

const (
	TimeConstraintActivityDomain_Personal     TimeConstraintActivityDomain = "personal"
	TimeConstraintActivityDomain_Unknown      TimeConstraintActivityDomain = "unknown"
	TimeConstraintActivityDomain_Unrestricted TimeConstraintActivityDomain = "unrestricted"
	TimeConstraintActivityDomain_Work         TimeConstraintActivityDomain = "work"
)

func PossibleValuesForTimeConstraintActivityDomain() []string {
	return []string{
		string(TimeConstraintActivityDomain_Personal),
		string(TimeConstraintActivityDomain_Unknown),
		string(TimeConstraintActivityDomain_Unrestricted),
		string(TimeConstraintActivityDomain_Work),
	}
}

func (s *TimeConstraintActivityDomain) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTimeConstraintActivityDomain(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTimeConstraintActivityDomain(input string) (*TimeConstraintActivityDomain, error) {
	vals := map[string]TimeConstraintActivityDomain{
		"personal":     TimeConstraintActivityDomain_Personal,
		"unknown":      TimeConstraintActivityDomain_Unknown,
		"unrestricted": TimeConstraintActivityDomain_Unrestricted,
		"work":         TimeConstraintActivityDomain_Work,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TimeConstraintActivityDomain(input)
	return &out, nil
}
