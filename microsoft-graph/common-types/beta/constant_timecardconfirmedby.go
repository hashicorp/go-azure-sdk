package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TimeCardConfirmedBy string

const (
	TimeCardConfirmedBy_Manager TimeCardConfirmedBy = "manager"
	TimeCardConfirmedBy_None    TimeCardConfirmedBy = "none"
	TimeCardConfirmedBy_User    TimeCardConfirmedBy = "user"
)

func PossibleValuesForTimeCardConfirmedBy() []string {
	return []string{
		string(TimeCardConfirmedBy_Manager),
		string(TimeCardConfirmedBy_None),
		string(TimeCardConfirmedBy_User),
	}
}

func (s *TimeCardConfirmedBy) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTimeCardConfirmedBy(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTimeCardConfirmedBy(input string) (*TimeCardConfirmedBy, error) {
	vals := map[string]TimeCardConfirmedBy{
		"manager": TimeCardConfirmedBy_Manager,
		"none":    TimeCardConfirmedBy_None,
		"user":    TimeCardConfirmedBy_User,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TimeCardConfirmedBy(input)
	return &out, nil
}
