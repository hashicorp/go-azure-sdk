package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityCustomUserFlowAttributeUserFlowAttributeType string

const (
	IdentityCustomUserFlowAttributeUserFlowAttributeType_BuiltIn  IdentityCustomUserFlowAttributeUserFlowAttributeType = "builtIn"
	IdentityCustomUserFlowAttributeUserFlowAttributeType_Custom   IdentityCustomUserFlowAttributeUserFlowAttributeType = "custom"
	IdentityCustomUserFlowAttributeUserFlowAttributeType_Required IdentityCustomUserFlowAttributeUserFlowAttributeType = "required"
)

func PossibleValuesForIdentityCustomUserFlowAttributeUserFlowAttributeType() []string {
	return []string{
		string(IdentityCustomUserFlowAttributeUserFlowAttributeType_BuiltIn),
		string(IdentityCustomUserFlowAttributeUserFlowAttributeType_Custom),
		string(IdentityCustomUserFlowAttributeUserFlowAttributeType_Required),
	}
}

func (s *IdentityCustomUserFlowAttributeUserFlowAttributeType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIdentityCustomUserFlowAttributeUserFlowAttributeType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIdentityCustomUserFlowAttributeUserFlowAttributeType(input string) (*IdentityCustomUserFlowAttributeUserFlowAttributeType, error) {
	vals := map[string]IdentityCustomUserFlowAttributeUserFlowAttributeType{
		"builtin":  IdentityCustomUserFlowAttributeUserFlowAttributeType_BuiltIn,
		"custom":   IdentityCustomUserFlowAttributeUserFlowAttributeType_Custom,
		"required": IdentityCustomUserFlowAttributeUserFlowAttributeType_Required,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IdentityCustomUserFlowAttributeUserFlowAttributeType(input)
	return &out, nil
}
