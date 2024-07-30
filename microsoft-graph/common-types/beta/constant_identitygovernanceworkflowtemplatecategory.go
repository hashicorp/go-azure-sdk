package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityGovernanceWorkflowTemplateCategory string

const (
	IdentityGovernanceWorkflowTemplateCategory_Joiner IdentityGovernanceWorkflowTemplateCategory = "joiner"
	IdentityGovernanceWorkflowTemplateCategory_Leaver IdentityGovernanceWorkflowTemplateCategory = "leaver"
	IdentityGovernanceWorkflowTemplateCategory_Mover  IdentityGovernanceWorkflowTemplateCategory = "mover"
)

func PossibleValuesForIdentityGovernanceWorkflowTemplateCategory() []string {
	return []string{
		string(IdentityGovernanceWorkflowTemplateCategory_Joiner),
		string(IdentityGovernanceWorkflowTemplateCategory_Leaver),
		string(IdentityGovernanceWorkflowTemplateCategory_Mover),
	}
}

func (s *IdentityGovernanceWorkflowTemplateCategory) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIdentityGovernanceWorkflowTemplateCategory(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIdentityGovernanceWorkflowTemplateCategory(input string) (*IdentityGovernanceWorkflowTemplateCategory, error) {
	vals := map[string]IdentityGovernanceWorkflowTemplateCategory{
		"joiner": IdentityGovernanceWorkflowTemplateCategory_Joiner,
		"leaver": IdentityGovernanceWorkflowTemplateCategory_Leaver,
		"mover":  IdentityGovernanceWorkflowTemplateCategory_Mover,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IdentityGovernanceWorkflowTemplateCategory(input)
	return &out, nil
}
