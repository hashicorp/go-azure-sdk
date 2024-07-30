package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityGovernanceWorkflowCategory string

const (
	IdentityGovernanceWorkflowCategory_Joiner IdentityGovernanceWorkflowCategory = "joiner"
	IdentityGovernanceWorkflowCategory_Leaver IdentityGovernanceWorkflowCategory = "leaver"
	IdentityGovernanceWorkflowCategory_Mover  IdentityGovernanceWorkflowCategory = "mover"
)

func PossibleValuesForIdentityGovernanceWorkflowCategory() []string {
	return []string{
		string(IdentityGovernanceWorkflowCategory_Joiner),
		string(IdentityGovernanceWorkflowCategory_Leaver),
		string(IdentityGovernanceWorkflowCategory_Mover),
	}
}

func (s *IdentityGovernanceWorkflowCategory) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIdentityGovernanceWorkflowCategory(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIdentityGovernanceWorkflowCategory(input string) (*IdentityGovernanceWorkflowCategory, error) {
	vals := map[string]IdentityGovernanceWorkflowCategory{
		"joiner": IdentityGovernanceWorkflowCategory_Joiner,
		"leaver": IdentityGovernanceWorkflowCategory_Leaver,
		"mover":  IdentityGovernanceWorkflowCategory_Mover,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IdentityGovernanceWorkflowCategory(input)
	return &out, nil
}
