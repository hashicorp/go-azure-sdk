package advisorscore

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Aggregated string

const (
	AggregatedDay   Aggregated = "day"
	AggregatedMonth Aggregated = "month"
	AggregatedWeek  Aggregated = "week"
)

func PossibleValuesForAggregated() []string {
	return []string{
		string(AggregatedDay),
		string(AggregatedMonth),
		string(AggregatedWeek),
	}
}

func (s *Aggregated) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAggregated(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAggregated(input string) (*Aggregated, error) {
	vals := map[string]Aggregated{
		"day":   AggregatedDay,
		"month": AggregatedMonth,
		"week":  AggregatedWeek,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Aggregated(input)
	return &out, nil
}
