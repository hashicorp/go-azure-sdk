package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupPolicyConfigurationPolicyConfigurationIngestionType string

const (
	GroupPolicyConfigurationPolicyConfigurationIngestionType_BuiltIn GroupPolicyConfigurationPolicyConfigurationIngestionType = "builtIn"
	GroupPolicyConfigurationPolicyConfigurationIngestionType_Custom  GroupPolicyConfigurationPolicyConfigurationIngestionType = "custom"
	GroupPolicyConfigurationPolicyConfigurationIngestionType_Mixed   GroupPolicyConfigurationPolicyConfigurationIngestionType = "mixed"
	GroupPolicyConfigurationPolicyConfigurationIngestionType_Unknown GroupPolicyConfigurationPolicyConfigurationIngestionType = "unknown"
)

func PossibleValuesForGroupPolicyConfigurationPolicyConfigurationIngestionType() []string {
	return []string{
		string(GroupPolicyConfigurationPolicyConfigurationIngestionType_BuiltIn),
		string(GroupPolicyConfigurationPolicyConfigurationIngestionType_Custom),
		string(GroupPolicyConfigurationPolicyConfigurationIngestionType_Mixed),
		string(GroupPolicyConfigurationPolicyConfigurationIngestionType_Unknown),
	}
}

func (s *GroupPolicyConfigurationPolicyConfigurationIngestionType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseGroupPolicyConfigurationPolicyConfigurationIngestionType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseGroupPolicyConfigurationPolicyConfigurationIngestionType(input string) (*GroupPolicyConfigurationPolicyConfigurationIngestionType, error) {
	vals := map[string]GroupPolicyConfigurationPolicyConfigurationIngestionType{
		"builtin": GroupPolicyConfigurationPolicyConfigurationIngestionType_BuiltIn,
		"custom":  GroupPolicyConfigurationPolicyConfigurationIngestionType_Custom,
		"mixed":   GroupPolicyConfigurationPolicyConfigurationIngestionType_Mixed,
		"unknown": GroupPolicyConfigurationPolicyConfigurationIngestionType_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := GroupPolicyConfigurationPolicyConfigurationIngestionType(input)
	return &out, nil
}
