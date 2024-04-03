package sharedgalleries

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SharedToValues string

const (
	SharedToValuesTenant SharedToValues = "tenant"
)

func PossibleValuesForSharedToValues() []string {
	return []string{
		string(SharedToValuesTenant),
	}
}

func (s *SharedToValues) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSharedToValues(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSharedToValues(input string) (*SharedToValues, error) {
	vals := map[string]SharedToValues{
		"tenant": SharedToValuesTenant,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SharedToValues(input)
	return &out, nil
}
