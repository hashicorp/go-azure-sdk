package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupPolicyDefinitionValueConfigurationType string

const (
	GroupPolicyDefinitionValueConfigurationType_Policy     GroupPolicyDefinitionValueConfigurationType = "policy"
	GroupPolicyDefinitionValueConfigurationType_Preference GroupPolicyDefinitionValueConfigurationType = "preference"
)

func PossibleValuesForGroupPolicyDefinitionValueConfigurationType() []string {
	return []string{
		string(GroupPolicyDefinitionValueConfigurationType_Policy),
		string(GroupPolicyDefinitionValueConfigurationType_Preference),
	}
}

func (s *GroupPolicyDefinitionValueConfigurationType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseGroupPolicyDefinitionValueConfigurationType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseGroupPolicyDefinitionValueConfigurationType(input string) (*GroupPolicyDefinitionValueConfigurationType, error) {
	vals := map[string]GroupPolicyDefinitionValueConfigurationType{
		"policy":     GroupPolicyDefinitionValueConfigurationType_Policy,
		"preference": GroupPolicyDefinitionValueConfigurationType_Preference,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := GroupPolicyDefinitionValueConfigurationType(input)
	return &out, nil
}
