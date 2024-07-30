package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupPolicyUploadedDefinitionFilePolicyType string

const (
	GroupPolicyUploadedDefinitionFilePolicyType_AdmxBacked   GroupPolicyUploadedDefinitionFilePolicyType = "admxBacked"
	GroupPolicyUploadedDefinitionFilePolicyType_AdmxIngested GroupPolicyUploadedDefinitionFilePolicyType = "admxIngested"
)

func PossibleValuesForGroupPolicyUploadedDefinitionFilePolicyType() []string {
	return []string{
		string(GroupPolicyUploadedDefinitionFilePolicyType_AdmxBacked),
		string(GroupPolicyUploadedDefinitionFilePolicyType_AdmxIngested),
	}
}

func (s *GroupPolicyUploadedDefinitionFilePolicyType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseGroupPolicyUploadedDefinitionFilePolicyType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseGroupPolicyUploadedDefinitionFilePolicyType(input string) (*GroupPolicyUploadedDefinitionFilePolicyType, error) {
	vals := map[string]GroupPolicyUploadedDefinitionFilePolicyType{
		"admxbacked":   GroupPolicyUploadedDefinitionFilePolicyType_AdmxBacked,
		"admxingested": GroupPolicyUploadedDefinitionFilePolicyType_AdmxIngested,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := GroupPolicyUploadedDefinitionFilePolicyType(input)
	return &out, nil
}
