package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityUserFlowAttributeUserFlowAttributeType string

const (
	IdentityUserFlowAttributeUserFlowAttributeType_BuiltIn  IdentityUserFlowAttributeUserFlowAttributeType = "builtIn"
	IdentityUserFlowAttributeUserFlowAttributeType_Custom   IdentityUserFlowAttributeUserFlowAttributeType = "custom"
	IdentityUserFlowAttributeUserFlowAttributeType_Required IdentityUserFlowAttributeUserFlowAttributeType = "required"
)

func PossibleValuesForIdentityUserFlowAttributeUserFlowAttributeType() []string {
	return []string{
		string(IdentityUserFlowAttributeUserFlowAttributeType_BuiltIn),
		string(IdentityUserFlowAttributeUserFlowAttributeType_Custom),
		string(IdentityUserFlowAttributeUserFlowAttributeType_Required),
	}
}

func (s *IdentityUserFlowAttributeUserFlowAttributeType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIdentityUserFlowAttributeUserFlowAttributeType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIdentityUserFlowAttributeUserFlowAttributeType(input string) (*IdentityUserFlowAttributeUserFlowAttributeType, error) {
	vals := map[string]IdentityUserFlowAttributeUserFlowAttributeType{
		"builtin":  IdentityUserFlowAttributeUserFlowAttributeType_BuiltIn,
		"custom":   IdentityUserFlowAttributeUserFlowAttributeType_Custom,
		"required": IdentityUserFlowAttributeUserFlowAttributeType_Required,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IdentityUserFlowAttributeUserFlowAttributeType(input)
	return &out, nil
}
