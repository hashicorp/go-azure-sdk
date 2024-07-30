package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityGovernanceWorkflowBaseCategory string

const (
	IdentityGovernanceWorkflowBaseCategory_Joiner IdentityGovernanceWorkflowBaseCategory = "joiner"
	IdentityGovernanceWorkflowBaseCategory_Leaver IdentityGovernanceWorkflowBaseCategory = "leaver"
	IdentityGovernanceWorkflowBaseCategory_Mover  IdentityGovernanceWorkflowBaseCategory = "mover"
)

func PossibleValuesForIdentityGovernanceWorkflowBaseCategory() []string {
	return []string{
		string(IdentityGovernanceWorkflowBaseCategory_Joiner),
		string(IdentityGovernanceWorkflowBaseCategory_Leaver),
		string(IdentityGovernanceWorkflowBaseCategory_Mover),
	}
}

func (s *IdentityGovernanceWorkflowBaseCategory) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIdentityGovernanceWorkflowBaseCategory(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIdentityGovernanceWorkflowBaseCategory(input string) (*IdentityGovernanceWorkflowBaseCategory, error) {
	vals := map[string]IdentityGovernanceWorkflowBaseCategory{
		"joiner": IdentityGovernanceWorkflowBaseCategory_Joiner,
		"leaver": IdentityGovernanceWorkflowBaseCategory_Leaver,
		"mover":  IdentityGovernanceWorkflowBaseCategory_Mover,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IdentityGovernanceWorkflowBaseCategory(input)
	return &out, nil
}
