package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityUserFlowAttributeUserFlowAttributeType string

const (
	IdentityUserFlowAttributeUserFlowAttributeTypebuiltIn  IdentityUserFlowAttributeUserFlowAttributeType = "BuiltIn"
	IdentityUserFlowAttributeUserFlowAttributeTypecustom   IdentityUserFlowAttributeUserFlowAttributeType = "Custom"
	IdentityUserFlowAttributeUserFlowAttributeTyperequired IdentityUserFlowAttributeUserFlowAttributeType = "Required"
)

func PossibleValuesForIdentityUserFlowAttributeUserFlowAttributeType() []string {
	return []string{
		string(IdentityUserFlowAttributeUserFlowAttributeTypebuiltIn),
		string(IdentityUserFlowAttributeUserFlowAttributeTypecustom),
		string(IdentityUserFlowAttributeUserFlowAttributeTyperequired),
	}
}

func parseIdentityUserFlowAttributeUserFlowAttributeType(input string) (*IdentityUserFlowAttributeUserFlowAttributeType, error) {
	vals := map[string]IdentityUserFlowAttributeUserFlowAttributeType{
		"builtin":  IdentityUserFlowAttributeUserFlowAttributeTypebuiltIn,
		"custom":   IdentityUserFlowAttributeUserFlowAttributeTypecustom,
		"required": IdentityUserFlowAttributeUserFlowAttributeTyperequired,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IdentityUserFlowAttributeUserFlowAttributeType(input)
	return &out, nil
}
