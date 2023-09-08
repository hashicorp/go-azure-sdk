package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityBuiltInUserFlowAttributeUserFlowAttributeType string

const (
	IdentityBuiltInUserFlowAttributeUserFlowAttributeTypebuiltIn  IdentityBuiltInUserFlowAttributeUserFlowAttributeType = "BuiltIn"
	IdentityBuiltInUserFlowAttributeUserFlowAttributeTypecustom   IdentityBuiltInUserFlowAttributeUserFlowAttributeType = "Custom"
	IdentityBuiltInUserFlowAttributeUserFlowAttributeTyperequired IdentityBuiltInUserFlowAttributeUserFlowAttributeType = "Required"
)

func PossibleValuesForIdentityBuiltInUserFlowAttributeUserFlowAttributeType() []string {
	return []string{
		string(IdentityBuiltInUserFlowAttributeUserFlowAttributeTypebuiltIn),
		string(IdentityBuiltInUserFlowAttributeUserFlowAttributeTypecustom),
		string(IdentityBuiltInUserFlowAttributeUserFlowAttributeTyperequired),
	}
}

func parseIdentityBuiltInUserFlowAttributeUserFlowAttributeType(input string) (*IdentityBuiltInUserFlowAttributeUserFlowAttributeType, error) {
	vals := map[string]IdentityBuiltInUserFlowAttributeUserFlowAttributeType{
		"builtin":  IdentityBuiltInUserFlowAttributeUserFlowAttributeTypebuiltIn,
		"custom":   IdentityBuiltInUserFlowAttributeUserFlowAttributeTypecustom,
		"required": IdentityBuiltInUserFlowAttributeUserFlowAttributeTyperequired,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IdentityBuiltInUserFlowAttributeUserFlowAttributeType(input)
	return &out, nil
}
