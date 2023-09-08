package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityGovernanceTaskCategory string

const (
	IdentityGovernanceTaskCategoryjoiner IdentityGovernanceTaskCategory = "Joiner"
	IdentityGovernanceTaskCategoryleaver IdentityGovernanceTaskCategory = "Leaver"
	IdentityGovernanceTaskCategorymover  IdentityGovernanceTaskCategory = "Mover"
)

func PossibleValuesForIdentityGovernanceTaskCategory() []string {
	return []string{
		string(IdentityGovernanceTaskCategoryjoiner),
		string(IdentityGovernanceTaskCategoryleaver),
		string(IdentityGovernanceTaskCategorymover),
	}
}

func parseIdentityGovernanceTaskCategory(input string) (*IdentityGovernanceTaskCategory, error) {
	vals := map[string]IdentityGovernanceTaskCategory{
		"joiner": IdentityGovernanceTaskCategoryjoiner,
		"leaver": IdentityGovernanceTaskCategoryleaver,
		"mover":  IdentityGovernanceTaskCategorymover,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IdentityGovernanceTaskCategory(input)
	return &out, nil
}
