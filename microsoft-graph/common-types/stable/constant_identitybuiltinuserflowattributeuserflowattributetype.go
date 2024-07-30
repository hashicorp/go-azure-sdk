package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityBuiltInUserFlowAttributeUserFlowAttributeType string

const (
	IdentityBuiltInUserFlowAttributeUserFlowAttributeType_BuiltIn  IdentityBuiltInUserFlowAttributeUserFlowAttributeType = "builtIn"
	IdentityBuiltInUserFlowAttributeUserFlowAttributeType_Custom   IdentityBuiltInUserFlowAttributeUserFlowAttributeType = "custom"
	IdentityBuiltInUserFlowAttributeUserFlowAttributeType_Required IdentityBuiltInUserFlowAttributeUserFlowAttributeType = "required"
)

func PossibleValuesForIdentityBuiltInUserFlowAttributeUserFlowAttributeType() []string {
	return []string{
		string(IdentityBuiltInUserFlowAttributeUserFlowAttributeType_BuiltIn),
		string(IdentityBuiltInUserFlowAttributeUserFlowAttributeType_Custom),
		string(IdentityBuiltInUserFlowAttributeUserFlowAttributeType_Required),
	}
}

func (s *IdentityBuiltInUserFlowAttributeUserFlowAttributeType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIdentityBuiltInUserFlowAttributeUserFlowAttributeType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIdentityBuiltInUserFlowAttributeUserFlowAttributeType(input string) (*IdentityBuiltInUserFlowAttributeUserFlowAttributeType, error) {
	vals := map[string]IdentityBuiltInUserFlowAttributeUserFlowAttributeType{
		"builtin":  IdentityBuiltInUserFlowAttributeUserFlowAttributeType_BuiltIn,
		"custom":   IdentityBuiltInUserFlowAttributeUserFlowAttributeType_Custom,
		"required": IdentityBuiltInUserFlowAttributeUserFlowAttributeType_Required,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IdentityBuiltInUserFlowAttributeUserFlowAttributeType(input)
	return &out, nil
}
