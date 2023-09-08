package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityGovernanceParameterValueType string

const (
	IdentityGovernanceParameterValueTypebool   IdentityGovernanceParameterValueType = "Bool"
	IdentityGovernanceParameterValueTypeenum   IdentityGovernanceParameterValueType = "Enum"
	IdentityGovernanceParameterValueTypeint    IdentityGovernanceParameterValueType = "Int"
	IdentityGovernanceParameterValueTypestring IdentityGovernanceParameterValueType = "String"
)

func PossibleValuesForIdentityGovernanceParameterValueType() []string {
	return []string{
		string(IdentityGovernanceParameterValueTypebool),
		string(IdentityGovernanceParameterValueTypeenum),
		string(IdentityGovernanceParameterValueTypeint),
		string(IdentityGovernanceParameterValueTypestring),
	}
}

func parseIdentityGovernanceParameterValueType(input string) (*IdentityGovernanceParameterValueType, error) {
	vals := map[string]IdentityGovernanceParameterValueType{
		"bool":   IdentityGovernanceParameterValueTypebool,
		"enum":   IdentityGovernanceParameterValueTypeenum,
		"int":    IdentityGovernanceParameterValueTypeint,
		"string": IdentityGovernanceParameterValueTypestring,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IdentityGovernanceParameterValueType(input)
	return &out, nil
}
