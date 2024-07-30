package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedIdentityMsiType string

const (
	ManagedIdentityMsiType_None           ManagedIdentityMsiType = "none"
	ManagedIdentityMsiType_SystemAssigned ManagedIdentityMsiType = "systemAssigned"
	ManagedIdentityMsiType_UserAssigned   ManagedIdentityMsiType = "userAssigned"
)

func PossibleValuesForManagedIdentityMsiType() []string {
	return []string{
		string(ManagedIdentityMsiType_None),
		string(ManagedIdentityMsiType_SystemAssigned),
		string(ManagedIdentityMsiType_UserAssigned),
	}
}

func (s *ManagedIdentityMsiType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedIdentityMsiType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedIdentityMsiType(input string) (*ManagedIdentityMsiType, error) {
	vals := map[string]ManagedIdentityMsiType{
		"none":           ManagedIdentityMsiType_None,
		"systemassigned": ManagedIdentityMsiType_SystemAssigned,
		"userassigned":   ManagedIdentityMsiType_UserAssigned,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedIdentityMsiType(input)
	return &out, nil
}
