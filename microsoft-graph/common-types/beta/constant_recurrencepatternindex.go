package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RecurrencePatternIndex string

const (
	RecurrencePatternIndex_First  RecurrencePatternIndex = "first"
	RecurrencePatternIndex_Fourth RecurrencePatternIndex = "fourth"
	RecurrencePatternIndex_Last   RecurrencePatternIndex = "last"
	RecurrencePatternIndex_Second RecurrencePatternIndex = "second"
	RecurrencePatternIndex_Third  RecurrencePatternIndex = "third"
)

func PossibleValuesForRecurrencePatternIndex() []string {
	return []string{
		string(RecurrencePatternIndex_First),
		string(RecurrencePatternIndex_Fourth),
		string(RecurrencePatternIndex_Last),
		string(RecurrencePatternIndex_Second),
		string(RecurrencePatternIndex_Third),
	}
}

func (s *RecurrencePatternIndex) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRecurrencePatternIndex(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRecurrencePatternIndex(input string) (*RecurrencePatternIndex, error) {
	vals := map[string]RecurrencePatternIndex{
		"first":  RecurrencePatternIndex_First,
		"fourth": RecurrencePatternIndex_Fourth,
		"last":   RecurrencePatternIndex_Last,
		"second": RecurrencePatternIndex_Second,
		"third":  RecurrencePatternIndex_Third,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RecurrencePatternIndex(input)
	return &out, nil
}
