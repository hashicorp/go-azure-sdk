package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityGovernanceTaskCategory string

const (
	IdentityGovernanceTaskCategory_Joiner IdentityGovernanceTaskCategory = "joiner"
	IdentityGovernanceTaskCategory_Leaver IdentityGovernanceTaskCategory = "leaver"
	IdentityGovernanceTaskCategory_Mover  IdentityGovernanceTaskCategory = "mover"
)

func PossibleValuesForIdentityGovernanceTaskCategory() []string {
	return []string{
		string(IdentityGovernanceTaskCategory_Joiner),
		string(IdentityGovernanceTaskCategory_Leaver),
		string(IdentityGovernanceTaskCategory_Mover),
	}
}

func (s *IdentityGovernanceTaskCategory) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIdentityGovernanceTaskCategory(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIdentityGovernanceTaskCategory(input string) (*IdentityGovernanceTaskCategory, error) {
	vals := map[string]IdentityGovernanceTaskCategory{
		"joiner": IdentityGovernanceTaskCategory_Joiner,
		"leaver": IdentityGovernanceTaskCategory_Leaver,
		"mover":  IdentityGovernanceTaskCategory_Mover,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IdentityGovernanceTaskCategory(input)
	return &out, nil
}
