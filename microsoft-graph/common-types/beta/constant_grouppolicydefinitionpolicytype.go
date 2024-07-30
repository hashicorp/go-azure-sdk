package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupPolicyDefinitionPolicyType string

const (
	GroupPolicyDefinitionPolicyType_AdmxBacked   GroupPolicyDefinitionPolicyType = "admxBacked"
	GroupPolicyDefinitionPolicyType_AdmxIngested GroupPolicyDefinitionPolicyType = "admxIngested"
)

func PossibleValuesForGroupPolicyDefinitionPolicyType() []string {
	return []string{
		string(GroupPolicyDefinitionPolicyType_AdmxBacked),
		string(GroupPolicyDefinitionPolicyType_AdmxIngested),
	}
}

func (s *GroupPolicyDefinitionPolicyType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseGroupPolicyDefinitionPolicyType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseGroupPolicyDefinitionPolicyType(input string) (*GroupPolicyDefinitionPolicyType, error) {
	vals := map[string]GroupPolicyDefinitionPolicyType{
		"admxbacked":   GroupPolicyDefinitionPolicyType_AdmxBacked,
		"admxingested": GroupPolicyDefinitionPolicyType_AdmxIngested,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := GroupPolicyDefinitionPolicyType(input)
	return &out, nil
}
