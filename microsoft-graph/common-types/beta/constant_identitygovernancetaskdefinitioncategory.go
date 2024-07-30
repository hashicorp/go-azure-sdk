package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityGovernanceTaskDefinitionCategory string

const (
	IdentityGovernanceTaskDefinitionCategory_Joiner IdentityGovernanceTaskDefinitionCategory = "joiner"
	IdentityGovernanceTaskDefinitionCategory_Leaver IdentityGovernanceTaskDefinitionCategory = "leaver"
	IdentityGovernanceTaskDefinitionCategory_Mover  IdentityGovernanceTaskDefinitionCategory = "mover"
)

func PossibleValuesForIdentityGovernanceTaskDefinitionCategory() []string {
	return []string{
		string(IdentityGovernanceTaskDefinitionCategory_Joiner),
		string(IdentityGovernanceTaskDefinitionCategory_Leaver),
		string(IdentityGovernanceTaskDefinitionCategory_Mover),
	}
}

func (s *IdentityGovernanceTaskDefinitionCategory) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIdentityGovernanceTaskDefinitionCategory(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIdentityGovernanceTaskDefinitionCategory(input string) (*IdentityGovernanceTaskDefinitionCategory, error) {
	vals := map[string]IdentityGovernanceTaskDefinitionCategory{
		"joiner": IdentityGovernanceTaskDefinitionCategory_Joiner,
		"leaver": IdentityGovernanceTaskDefinitionCategory_Leaver,
		"mover":  IdentityGovernanceTaskDefinitionCategory_Mover,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IdentityGovernanceTaskDefinitionCategory(input)
	return &out, nil
}
