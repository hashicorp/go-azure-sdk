package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UnifiedRoleDefinitionAllowedPrincipalTypes string

const (
	UnifiedRoleDefinitionAllowedPrincipalTypesgroup            UnifiedRoleDefinitionAllowedPrincipalTypes = "Group"
	UnifiedRoleDefinitionAllowedPrincipalTypesservicePrincipal UnifiedRoleDefinitionAllowedPrincipalTypes = "ServicePrincipal"
	UnifiedRoleDefinitionAllowedPrincipalTypesuser             UnifiedRoleDefinitionAllowedPrincipalTypes = "User"
)

func PossibleValuesForUnifiedRoleDefinitionAllowedPrincipalTypes() []string {
	return []string{
		string(UnifiedRoleDefinitionAllowedPrincipalTypesgroup),
		string(UnifiedRoleDefinitionAllowedPrincipalTypesservicePrincipal),
		string(UnifiedRoleDefinitionAllowedPrincipalTypesuser),
	}
}

func parseUnifiedRoleDefinitionAllowedPrincipalTypes(input string) (*UnifiedRoleDefinitionAllowedPrincipalTypes, error) {
	vals := map[string]UnifiedRoleDefinitionAllowedPrincipalTypes{
		"group":            UnifiedRoleDefinitionAllowedPrincipalTypesgroup,
		"serviceprincipal": UnifiedRoleDefinitionAllowedPrincipalTypesservicePrincipal,
		"user":             UnifiedRoleDefinitionAllowedPrincipalTypesuser,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UnifiedRoleDefinitionAllowedPrincipalTypes(input)
	return &out, nil
}
