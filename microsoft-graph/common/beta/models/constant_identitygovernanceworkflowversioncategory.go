package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityGovernanceWorkflowVersionCategory string

const (
	IdentityGovernanceWorkflowVersionCategoryjoiner IdentityGovernanceWorkflowVersionCategory = "Joiner"
	IdentityGovernanceWorkflowVersionCategoryleaver IdentityGovernanceWorkflowVersionCategory = "Leaver"
	IdentityGovernanceWorkflowVersionCategorymover  IdentityGovernanceWorkflowVersionCategory = "Mover"
)

func PossibleValuesForIdentityGovernanceWorkflowVersionCategory() []string {
	return []string{
		string(IdentityGovernanceWorkflowVersionCategoryjoiner),
		string(IdentityGovernanceWorkflowVersionCategoryleaver),
		string(IdentityGovernanceWorkflowVersionCategorymover),
	}
}

func parseIdentityGovernanceWorkflowVersionCategory(input string) (*IdentityGovernanceWorkflowVersionCategory, error) {
	vals := map[string]IdentityGovernanceWorkflowVersionCategory{
		"joiner": IdentityGovernanceWorkflowVersionCategoryjoiner,
		"leaver": IdentityGovernanceWorkflowVersionCategoryleaver,
		"mover":  IdentityGovernanceWorkflowVersionCategorymover,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IdentityGovernanceWorkflowVersionCategory(input)
	return &out, nil
}
