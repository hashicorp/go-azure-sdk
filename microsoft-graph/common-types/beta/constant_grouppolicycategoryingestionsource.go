package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupPolicyCategoryIngestionSource string

const (
	GroupPolicyCategoryIngestionSource_BuiltIn GroupPolicyCategoryIngestionSource = "builtIn"
	GroupPolicyCategoryIngestionSource_Custom  GroupPolicyCategoryIngestionSource = "custom"
	GroupPolicyCategoryIngestionSource_Unknown GroupPolicyCategoryIngestionSource = "unknown"
)

func PossibleValuesForGroupPolicyCategoryIngestionSource() []string {
	return []string{
		string(GroupPolicyCategoryIngestionSource_BuiltIn),
		string(GroupPolicyCategoryIngestionSource_Custom),
		string(GroupPolicyCategoryIngestionSource_Unknown),
	}
}

func (s *GroupPolicyCategoryIngestionSource) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseGroupPolicyCategoryIngestionSource(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseGroupPolicyCategoryIngestionSource(input string) (*GroupPolicyCategoryIngestionSource, error) {
	vals := map[string]GroupPolicyCategoryIngestionSource{
		"builtin": GroupPolicyCategoryIngestionSource_BuiltIn,
		"custom":  GroupPolicyCategoryIngestionSource_Custom,
		"unknown": GroupPolicyCategoryIngestionSource_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := GroupPolicyCategoryIngestionSource(input)
	return &out, nil
}
