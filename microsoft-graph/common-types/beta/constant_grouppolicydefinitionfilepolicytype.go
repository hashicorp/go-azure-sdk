package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupPolicyDefinitionFilePolicyType string

const (
	GroupPolicyDefinitionFilePolicyType_AdmxBacked   GroupPolicyDefinitionFilePolicyType = "admxBacked"
	GroupPolicyDefinitionFilePolicyType_AdmxIngested GroupPolicyDefinitionFilePolicyType = "admxIngested"
)

func PossibleValuesForGroupPolicyDefinitionFilePolicyType() []string {
	return []string{
		string(GroupPolicyDefinitionFilePolicyType_AdmxBacked),
		string(GroupPolicyDefinitionFilePolicyType_AdmxIngested),
	}
}

func (s *GroupPolicyDefinitionFilePolicyType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseGroupPolicyDefinitionFilePolicyType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseGroupPolicyDefinitionFilePolicyType(input string) (*GroupPolicyDefinitionFilePolicyType, error) {
	vals := map[string]GroupPolicyDefinitionFilePolicyType{
		"admxbacked":   GroupPolicyDefinitionFilePolicyType_AdmxBacked,
		"admxingested": GroupPolicyDefinitionFilePolicyType_AdmxIngested,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := GroupPolicyDefinitionFilePolicyType(input)
	return &out, nil
}
