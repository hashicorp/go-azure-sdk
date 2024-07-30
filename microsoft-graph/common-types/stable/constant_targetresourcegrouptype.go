package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TargetResourceGroupType string

const (
	TargetResourceGroupType_AzureAD       TargetResourceGroupType = "azureAD"
	TargetResourceGroupType_UnifiedGroups TargetResourceGroupType = "unifiedGroups"
)

func PossibleValuesForTargetResourceGroupType() []string {
	return []string{
		string(TargetResourceGroupType_AzureAD),
		string(TargetResourceGroupType_UnifiedGroups),
	}
}

func (s *TargetResourceGroupType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTargetResourceGroupType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTargetResourceGroupType(input string) (*TargetResourceGroupType, error) {
	vals := map[string]TargetResourceGroupType{
		"azuread":       TargetResourceGroupType_AzureAD,
		"unifiedgroups": TargetResourceGroupType_UnifiedGroups,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TargetResourceGroupType(input)
	return &out, nil
}
