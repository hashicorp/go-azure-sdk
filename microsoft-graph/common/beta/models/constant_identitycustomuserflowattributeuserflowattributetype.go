package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityCustomUserFlowAttributeUserFlowAttributeType string

const (
	IdentityCustomUserFlowAttributeUserFlowAttributeTypebuiltIn  IdentityCustomUserFlowAttributeUserFlowAttributeType = "BuiltIn"
	IdentityCustomUserFlowAttributeUserFlowAttributeTypecustom   IdentityCustomUserFlowAttributeUserFlowAttributeType = "Custom"
	IdentityCustomUserFlowAttributeUserFlowAttributeTyperequired IdentityCustomUserFlowAttributeUserFlowAttributeType = "Required"
)

func PossibleValuesForIdentityCustomUserFlowAttributeUserFlowAttributeType() []string {
	return []string{
		string(IdentityCustomUserFlowAttributeUserFlowAttributeTypebuiltIn),
		string(IdentityCustomUserFlowAttributeUserFlowAttributeTypecustom),
		string(IdentityCustomUserFlowAttributeUserFlowAttributeTyperequired),
	}
}

func parseIdentityCustomUserFlowAttributeUserFlowAttributeType(input string) (*IdentityCustomUserFlowAttributeUserFlowAttributeType, error) {
	vals := map[string]IdentityCustomUserFlowAttributeUserFlowAttributeType{
		"builtin":  IdentityCustomUserFlowAttributeUserFlowAttributeTypebuiltIn,
		"custom":   IdentityCustomUserFlowAttributeUserFlowAttributeTypecustom,
		"required": IdentityCustomUserFlowAttributeUserFlowAttributeTyperequired,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IdentityCustomUserFlowAttributeUserFlowAttributeType(input)
	return &out, nil
}
