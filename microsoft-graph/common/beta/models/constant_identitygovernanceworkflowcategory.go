package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityGovernanceWorkflowCategory string

const (
	IdentityGovernanceWorkflowCategoryjoiner IdentityGovernanceWorkflowCategory = "Joiner"
	IdentityGovernanceWorkflowCategoryleaver IdentityGovernanceWorkflowCategory = "Leaver"
	IdentityGovernanceWorkflowCategorymover  IdentityGovernanceWorkflowCategory = "Mover"
)

func PossibleValuesForIdentityGovernanceWorkflowCategory() []string {
	return []string{
		string(IdentityGovernanceWorkflowCategoryjoiner),
		string(IdentityGovernanceWorkflowCategoryleaver),
		string(IdentityGovernanceWorkflowCategorymover),
	}
}

func parseIdentityGovernanceWorkflowCategory(input string) (*IdentityGovernanceWorkflowCategory, error) {
	vals := map[string]IdentityGovernanceWorkflowCategory{
		"joiner": IdentityGovernanceWorkflowCategoryjoiner,
		"leaver": IdentityGovernanceWorkflowCategoryleaver,
		"mover":  IdentityGovernanceWorkflowCategorymover,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IdentityGovernanceWorkflowCategory(input)
	return &out, nil
}
