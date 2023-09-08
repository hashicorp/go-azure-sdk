package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityGovernanceWorkflowBaseCategory string

const (
	IdentityGovernanceWorkflowBaseCategoryjoiner IdentityGovernanceWorkflowBaseCategory = "Joiner"
	IdentityGovernanceWorkflowBaseCategoryleaver IdentityGovernanceWorkflowBaseCategory = "Leaver"
	IdentityGovernanceWorkflowBaseCategorymover  IdentityGovernanceWorkflowBaseCategory = "Mover"
)

func PossibleValuesForIdentityGovernanceWorkflowBaseCategory() []string {
	return []string{
		string(IdentityGovernanceWorkflowBaseCategoryjoiner),
		string(IdentityGovernanceWorkflowBaseCategoryleaver),
		string(IdentityGovernanceWorkflowBaseCategorymover),
	}
}

func parseIdentityGovernanceWorkflowBaseCategory(input string) (*IdentityGovernanceWorkflowBaseCategory, error) {
	vals := map[string]IdentityGovernanceWorkflowBaseCategory{
		"joiner": IdentityGovernanceWorkflowBaseCategoryjoiner,
		"leaver": IdentityGovernanceWorkflowBaseCategoryleaver,
		"mover":  IdentityGovernanceWorkflowBaseCategorymover,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IdentityGovernanceWorkflowBaseCategory(input)
	return &out, nil
}
