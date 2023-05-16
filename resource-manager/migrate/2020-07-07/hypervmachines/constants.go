package hypervmachines

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HighlyAvailable string

const (
	HighlyAvailableNo      HighlyAvailable = "No"
	HighlyAvailableUnknown HighlyAvailable = "Unknown"
	HighlyAvailableYes     HighlyAvailable = "Yes"
)

func PossibleValuesForHighlyAvailable() []string {
	return []string{
		string(HighlyAvailableNo),
		string(HighlyAvailableUnknown),
		string(HighlyAvailableYes),
	}
}

func (s *HighlyAvailable) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseHighlyAvailable(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseHighlyAvailable(input string) (*HighlyAvailable, error) {
	vals := map[string]HighlyAvailable{
		"no":      HighlyAvailableNo,
		"unknown": HighlyAvailableUnknown,
		"yes":     HighlyAvailableYes,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := HighlyAvailable(input)
	return &out, nil
}
