package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityGovernanceWorkflowTemplateCategory string

const (
	IdentityGovernanceWorkflowTemplateCategoryjoiner IdentityGovernanceWorkflowTemplateCategory = "Joiner"
	IdentityGovernanceWorkflowTemplateCategoryleaver IdentityGovernanceWorkflowTemplateCategory = "Leaver"
	IdentityGovernanceWorkflowTemplateCategorymover  IdentityGovernanceWorkflowTemplateCategory = "Mover"
)

func PossibleValuesForIdentityGovernanceWorkflowTemplateCategory() []string {
	return []string{
		string(IdentityGovernanceWorkflowTemplateCategoryjoiner),
		string(IdentityGovernanceWorkflowTemplateCategoryleaver),
		string(IdentityGovernanceWorkflowTemplateCategorymover),
	}
}

func parseIdentityGovernanceWorkflowTemplateCategory(input string) (*IdentityGovernanceWorkflowTemplateCategory, error) {
	vals := map[string]IdentityGovernanceWorkflowTemplateCategory{
		"joiner": IdentityGovernanceWorkflowTemplateCategoryjoiner,
		"leaver": IdentityGovernanceWorkflowTemplateCategoryleaver,
		"mover":  IdentityGovernanceWorkflowTemplateCategorymover,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IdentityGovernanceWorkflowTemplateCategory(input)
	return &out, nil
}
