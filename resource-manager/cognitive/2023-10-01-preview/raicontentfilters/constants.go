package raicontentfilters

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RaiContentFilterType string

const (
	RaiContentFilterTypeMultiLevel RaiContentFilterType = "MultiLevel"
	RaiContentFilterTypeSwitch     RaiContentFilterType = "Switch"
)

func PossibleValuesForRaiContentFilterType() []string {
	return []string{
		string(RaiContentFilterTypeMultiLevel),
		string(RaiContentFilterTypeSwitch),
	}
}

func (s *RaiContentFilterType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRaiContentFilterType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRaiContentFilterType(input string) (*RaiContentFilterType, error) {
	vals := map[string]RaiContentFilterType{
		"multilevel": RaiContentFilterTypeMultiLevel,
		"switch":     RaiContentFilterTypeSwitch,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RaiContentFilterType(input)
	return &out, nil
}
