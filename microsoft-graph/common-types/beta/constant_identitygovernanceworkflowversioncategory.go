package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityGovernanceWorkflowVersionCategory string

const (
	IdentityGovernanceWorkflowVersionCategory_Joiner IdentityGovernanceWorkflowVersionCategory = "joiner"
	IdentityGovernanceWorkflowVersionCategory_Leaver IdentityGovernanceWorkflowVersionCategory = "leaver"
	IdentityGovernanceWorkflowVersionCategory_Mover  IdentityGovernanceWorkflowVersionCategory = "mover"
)

func PossibleValuesForIdentityGovernanceWorkflowVersionCategory() []string {
	return []string{
		string(IdentityGovernanceWorkflowVersionCategory_Joiner),
		string(IdentityGovernanceWorkflowVersionCategory_Leaver),
		string(IdentityGovernanceWorkflowVersionCategory_Mover),
	}
}

func (s *IdentityGovernanceWorkflowVersionCategory) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIdentityGovernanceWorkflowVersionCategory(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIdentityGovernanceWorkflowVersionCategory(input string) (*IdentityGovernanceWorkflowVersionCategory, error) {
	vals := map[string]IdentityGovernanceWorkflowVersionCategory{
		"joiner": IdentityGovernanceWorkflowVersionCategory_Joiner,
		"leaver": IdentityGovernanceWorkflowVersionCategory_Leaver,
		"mover":  IdentityGovernanceWorkflowVersionCategory_Mover,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IdentityGovernanceWorkflowVersionCategory(input)
	return &out, nil
}
