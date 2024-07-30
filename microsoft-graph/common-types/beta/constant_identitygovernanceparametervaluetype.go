package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityGovernanceParameterValueType string

const (
	IdentityGovernanceParameterValueType_Bool   IdentityGovernanceParameterValueType = "bool"
	IdentityGovernanceParameterValueType_Enum   IdentityGovernanceParameterValueType = "enum"
	IdentityGovernanceParameterValueType_Int    IdentityGovernanceParameterValueType = "int"
	IdentityGovernanceParameterValueType_String IdentityGovernanceParameterValueType = "string"
)

func PossibleValuesForIdentityGovernanceParameterValueType() []string {
	return []string{
		string(IdentityGovernanceParameterValueType_Bool),
		string(IdentityGovernanceParameterValueType_Enum),
		string(IdentityGovernanceParameterValueType_Int),
		string(IdentityGovernanceParameterValueType_String),
	}
}

func (s *IdentityGovernanceParameterValueType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIdentityGovernanceParameterValueType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIdentityGovernanceParameterValueType(input string) (*IdentityGovernanceParameterValueType, error) {
	vals := map[string]IdentityGovernanceParameterValueType{
		"bool":   IdentityGovernanceParameterValueType_Bool,
		"enum":   IdentityGovernanceParameterValueType_Enum,
		"int":    IdentityGovernanceParameterValueType_Int,
		"string": IdentityGovernanceParameterValueType_String,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IdentityGovernanceParameterValueType(input)
	return &out, nil
}
