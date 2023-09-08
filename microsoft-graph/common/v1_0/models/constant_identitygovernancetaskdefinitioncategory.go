package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityGovernanceTaskDefinitionCategory string

const (
	IdentityGovernanceTaskDefinitionCategoryjoiner IdentityGovernanceTaskDefinitionCategory = "Joiner"
	IdentityGovernanceTaskDefinitionCategoryleaver IdentityGovernanceTaskDefinitionCategory = "Leaver"
	IdentityGovernanceTaskDefinitionCategorymover  IdentityGovernanceTaskDefinitionCategory = "Mover"
)

func PossibleValuesForIdentityGovernanceTaskDefinitionCategory() []string {
	return []string{
		string(IdentityGovernanceTaskDefinitionCategoryjoiner),
		string(IdentityGovernanceTaskDefinitionCategoryleaver),
		string(IdentityGovernanceTaskDefinitionCategorymover),
	}
}

func parseIdentityGovernanceTaskDefinitionCategory(input string) (*IdentityGovernanceTaskDefinitionCategory, error) {
	vals := map[string]IdentityGovernanceTaskDefinitionCategory{
		"joiner": IdentityGovernanceTaskDefinitionCategoryjoiner,
		"leaver": IdentityGovernanceTaskDefinitionCategoryleaver,
		"mover":  IdentityGovernanceTaskDefinitionCategorymover,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IdentityGovernanceTaskDefinitionCategory(input)
	return &out, nil
}
