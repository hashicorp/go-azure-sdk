package loganalytics

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IntervalInMins string

const (
	IntervalInMinsFiveMins   IntervalInMins = "FiveMins"
	IntervalInMinsSixtyMins  IntervalInMins = "SixtyMins"
	IntervalInMinsThirtyMins IntervalInMins = "ThirtyMins"
	IntervalInMinsThreeMins  IntervalInMins = "ThreeMins"
)

func PossibleValuesForIntervalInMins() []string {
	return []string{
		string(IntervalInMinsFiveMins),
		string(IntervalInMinsSixtyMins),
		string(IntervalInMinsThirtyMins),
		string(IntervalInMinsThreeMins),
	}
}

func (s *IntervalInMins) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIntervalInMins(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIntervalInMins(input string) (*IntervalInMins, error) {
	vals := map[string]IntervalInMins{
		"fivemins":   IntervalInMinsFiveMins,
		"sixtymins":  IntervalInMinsSixtyMins,
		"thirtymins": IntervalInMinsThirtyMins,
		"threemins":  IntervalInMinsThreeMins,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IntervalInMins(input)
	return &out, nil
}
